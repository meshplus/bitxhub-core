package tss

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync"

	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/cache"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

// ProcessInboundMessages process the different types of p2p msgs received==============================================
func (t *TssInstance) ProcessInboundMessages(wg *sync.WaitGroup) {
	t.logger.Debug("start processing inbound messages")
	defer wg.Done()
	defer t.logger.Debug("stop processing inbound messages")

	for {
		select {
		case <-t.inMsgHandleStopChan:
			t.logger.Debugf("process inbound StopChan")
			return
		case m, ok := <-t.TssMsgChan:
			if !ok {
				t.logger.Debugf("process inbound msg not ok")
				return
			}
			if err := t.ProcessOneMessage(m); err != nil {
				t.logger.Errorf("fail to process the received message: %v", err)
			}
		}
	}
}

// ProcessOneMessage process one p2p msgs received======================================================================
func (t *TssInstance) ProcessOneMessage(msg *pb.Message) error {
	if msg == nil {
		return fmt.Errorf("invalid message")
	}

	wireMsg := &message.WireMessage{}
	if err := json.Unmarshal(msg.Data, wireMsg); err != nil {
		return fmt.Errorf("wire msg unmarshal error: %v", err)
	}
	t.logger.WithFields(logrus.Fields{"msgTyp": wireMsg.MsgType, "msgID": wireMsg.MsgID}).Debug("start processing one message")
	defer t.logger.WithFields(logrus.Fields{"msgTyp": wireMsg.MsgType, "msgID": wireMsg.MsgID}).Debug("finish processing one message")

	switch wireMsg.MsgType {
	case message.TSSKeyGenMsg, message.TSSKeySignMsg:
		// The message returned by the library method.
		// These needs to be processed for party updates once received and acknowledged (broadcast messages need to be acknowledged).
		var taskMsg message.TaskMessage
		if err := json.Unmarshal(wireMsg.MsgData, &taskMsg); nil != err {
			return fmt.Errorf("fail to unmarshal task message: %w", err)
		}

		t.logger.WithFields(logrus.Fields{
			"from":      taskMsg.Routing.From.Id,
			"roundInfo": taskMsg.RoundInfo,
			"msgType":   wireMsg.MsgType,
		}).Debug("process wire msg")

		return t.processTSSMsg(&taskMsg, wireMsg.MsgType, false)
	case message.TSSKeyGenVerMsg, message.TSSKeySignVerMsg:
		// The message is a confirmation for a broadcast message.
		// We need to update the confirmation of the broadcast message after receiving it。
		var bMsg message.BroadcastConfirmMessage
		if err := json.Unmarshal(wireMsg.MsgData, &bMsg); nil != err {
			return fmt.Errorf("fail to unmarshal broadcast confirm message")
		}

		t.logger.WithFields(logrus.Fields{
			"from":      bMsg.FromID,
			"roundInfo": bMsg.Key,
			"msgType":   wireMsg.MsgType,
		}).Debug("process ver msg")

		// we check whether this peer has already send us the VerMsg before update
		ret := t.checkDupAndUpdateVerMsg(&bMsg, bMsg.FromID)
		if ret {
			return t.processVerMsg(&bMsg, wireMsg.MsgType)
		}
	case message.TSSControlMsg:
		// The message is a share request message about a broadcast message from another participant.
		var controlMsg message.TssControl
		if err := json.Unmarshal(wireMsg.MsgData, &controlMsg); nil != err {
			return fmt.Errorf("fail to unmarshal control message: %w", err)
		}

		t.logger.WithFields(logrus.Fields{
			"from":    controlMsg.FromID,
			"reqKey":  controlMsg.ReqKey,
			"reqType": wireMsg.MsgType,
		}).Debug("process control msg")

		if controlMsg.Msg == nil {
			return t.processRequestMsgFromParty([]string{controlMsg.FromID}, &controlMsg, false)
		}
		exist := t.blameMgr.ShareMgr.QueryAndDelete(controlMsg.ReqHash)
		if !exist {
			t.logger.Debug("this request does not exit, maybe already processed")
			return nil
		}
		t.logger.Debug("we got the missing share from the peer")
		return t.processTSSMsg(controlMsg.Msg, controlMsg.RequestType, true)

	case message.TSSTaskDone:
		// If the message is received as the source node task has ended, check whether the task has been marked. If the message is marked, an error log is displayed.
		// If not marked, check whether all other nodes are marked end, if so, close its own end signal channel to return.
		// If not, that is, there is an end that has not received a signal, then we need to wait for their signal, so do not close the end signal channel, directly return.
		var doneMsg message.TssTaskNotifier
		if err := json.Unmarshal(wireMsg.MsgData, &doneMsg); err != nil {
			t.logger.Errorf("fail to unmarshal the notify message")
			return nil
		}

		t.logger.WithFields(logrus.Fields{
			"from":    doneMsg.FromID,
			"reqType": wireMsg.MsgType,
		}).Debug("process task down msg")

		if doneMsg.TaskDone {
			// if we have already logged this node, we return to avoid close of a close channel
			if t.finishedParties[doneMsg.FromID] {
				return fmt.Errorf("duplicated notification from peer %s ignored", doneMsg.FromID)
			}
			t.finishedParties[doneMsg.FromID] = true
			if len(t.finishedParties) == len(t.partyInfo.PartyIDMap)-1 {
				t.logger.WithFields(logrus.Fields{"finishedParties": t.finishedParties}).Infof("we get the confirm of the nodes that generate the signature")
				close(t.taskDoneChan)
			}
			return nil
		}
	}

	return nil
}

// processTSSMsg =======================================================================================================
// processTSSMsg processes the library message. If the library message is received properly, the process needs to continue
func (t *TssInstance) processTSSMsg(wireMsg *message.TaskMessage, msgType message.TssMsgType, forward bool) error {
	t.logger.WithFields(logrus.Fields{
		"from":        wireMsg.Routing.From.Id,
		"to":          wireMsg.Routing.To,
		"roundInfo":   wireMsg.RoundInfo,
		"isBroadCast": wireMsg.Routing.IsBroadcast,
		"type":        msgType,
	}).Debugf("process wire message")
	defer t.logger.Debugf("finish process wire message")

	// 1. verify msg not empty
	if wireMsg == nil || wireMsg.Routing == nil || wireMsg.Routing.From == nil {
		t.logger.Warnf("received msg invalid")
		return fmt.Errorf("invalid wireMsg")
	}

	// 2. verify sign in msg which is signed by sender
	partyIDMap := t.getPartyInfo().PartyIDMap
	dataOwner, ok := partyIDMap[wireMsg.Routing.From.Id]
	if !ok {
		t.logger.WithFields(logrus.Fields{
			"fromId":     wireMsg.Routing.From.Id,
			"partyIDMap": partyIDMap,
		}).Errorf("error in find the data owner")
		return fmt.Errorf("error in find the data owner")
	}
	pubkey, err := conversion.GetPubKeyFromPartyID(dataOwner)
	if err != nil {
		return fmt.Errorf("get pubkey from party id error: %v", err)
	}
	ok, err = conversion.VerifySignature(pubkey, wireMsg.Message, wireMsg.Sig, t.msgID)
	if err != nil {
		pid, _ := conversion.GetPIDFromPartyID(dataOwner)
		return fmt.Errorf("verify signature error: %v, dataOwnerId: %s, dataOwnerPid: %s", err, dataOwner.Id, pid.String())
	}
	if !ok {
		t.logger.Errorf("fail to verify the signature")
		return fmt.Errorf("signature verify failed")
	}

	// 3 process msg
	// 3.1 for the unicast message, we only update it local party to advance rounds
	if !wireMsg.Routing.IsBroadcast {
		t.logger.Debugf("msg from %s to %+v", wireMsg.Routing.From.Id, wireMsg.Routing.To[0].Id)
		return t.updateLocal(wireMsg)
	}

	// 3.2 broadcast message , we save a copy locally , and then tell all others what we got
	// 3.2.1 tell all others what we got
	if !forward {
		err := t.receiverBroadcastHashToPeers(wireMsg, msgType)
		if err != nil {
			t.logger.Errorf("fail to broadcast msg to peers")
		}
	}

	// 3.2.2 save a copy locally (fromN+roundN -> msgHashN)
	// 3.2.2.1 get cache k/v which will be save
	key := wireMsg.GetCacheKey() // 本地缓存信息，来源地址+轮次信息
	msgHash, err := conversion.BytesToHashString(wireMsg.Message)
	if err != nil {
		return fmt.Errorf("fail to calculate hash of the wire message: %w", err)
	}
	// 3.2.2.2 get local cache to set
	localCacheItem := t.getLocalCacheItem(key)
	// If don't have one, or if have one but it's msg is nil, should save it.
	if nil == localCacheItem {
		t.logger.Debugf("%s doesn't exist yet,add a new one", key)
		localCacheItem = cache.NewLocalCacheItem(wireMsg, msgHash)
		t.setLocalUnconfirmedMessages(key, localCacheItem)
	} else {
		// this means we received the broadcast confirm message from other party first
		t.logger.Debugf("%s exist", key)
		if localCacheItem.Msg == nil {
			t.logger.WithFields(logrus.Fields{"msgHash": msgHash}).Debugf("%s exist, set message", key)
			localCacheItem.Msg = wireMsg
			localCacheItem.Hash = msgHash
		}
	}
	// 3.2.2.3 update confirm list, that is, add localpartyID to the list
	localCacheItem.UpdateConfirmList(t.localPartyID, msgHash)
	t.logger.WithFields(logrus.Fields{"key": localCacheItem.Msg.GetCacheKey()}).Debugf("total confirmed parties:%+v", localCacheItem.ConfirmedList)

	// 3.2.3 handle the sharing confirmation of broadcast messages.
	// check whether the number of caches is sufficient, if sufficient, can be stored to blame and update local
	return t.applyShare(localCacheItem, key, msgType)
}

// The final processing mode of messages. ==============================================================================
// - Unicast messages can be updated directly；
// - Broadcast messages can be updated after confirmation.
// updateLocal will apply the wireMsg to local keygen/keysign party
func (t *TssInstance) updateLocal(taskMsg *message.TaskMessage) error {
	// 1 verify msg not empty
	if taskMsg == nil || taskMsg.Routing == nil || taskMsg.Routing.From == nil {
		t.logger.Warn("wire msg is nil")
		return fmt.Errorf("invalid wireMsg")
	}

	// 2 verify msg sender is known
	partyInfo := t.getPartyInfo()
	if partyInfo == nil {
		return nil
	}
	dataOwnerPartyID, ok := partyInfo.PartyIDMap[taskMsg.Routing.From.Id]
	if !ok {
		return fmt.Errorf("get message from unknown party %s", dataOwnerPartyID.Id)
	}

	// 3 here we log down this peer as the latest unicast peer
	if !taskMsg.Routing.IsBroadcast {
		t.blameMgr.SetLastUnicastParty(dataOwnerPartyID.Id, taskMsg.RoundInfo)
	}

	// 4 get the TaskMessage body content
	// it maybe multiple msg in keysign, but one keygen
	var bulkMsg []BulkWireMsg
	if err := json.Unmarshal(taskMsg.Message, &bulkMsg); err != nil {
		t.logger.Errorf("error to unmarshal the BulkMsg")
		return err
	}

	// 5 update localparty for each msg
	worker := runtime.NumCPU()
	tssJobChan := make(chan *tssJob, len(bulkMsg))

	// 5.1 Each message takes a coroutine and gets the update task from tssJobChan to update the party information of the message
	jobWg := sync.WaitGroup{}
	for i := 0; i < worker; i++ {
		jobWg.Add(1)
		go t.doTssJob(tssJobChan, &jobWg)
	}

	// 5.2 The current main process puts the update task into tssJobChan for each message
	for _, msg := range bulkMsg {
		// 5.2.1 get msg localparty
		data, ok := t.partyInfo.PartyMap.Load(msg.MsgIdentifier)
		if !ok {
			t.logger.Errorf("cannot find the party to this wired msg")
			return fmt.Errorf("cannot find the party")
		}
		localMsgParty := data.(btss.Party)

		// 5.2.2 get msg from partyID
		partyID, ok := t.partyInfo.PartyIDMap[msg.Routing.From.Id]
		if !ok {
			t.logger.Errorf("error in find the partyID")
			return fmt.Errorf("cannot find the party to handle the message")
		}

		// 5.2.3 get msg round
		round, err := conversion.GetMsgRound(msg.WiredBulkMsgs, partyID, msg.Routing.IsBroadcast)
		if err != nil {
			t.logger.Errorf("broken tss share")
			return err
		}
		// we only allow a message be updated only once.
		// here we use round + msgIdentifier as the key for the acceptedShares
		round.MsgIdentifier = msg.MsgIdentifier

		// 5.2.4 check if update duplicately
		// if this share is duplicated, we skip this share
		if t.blameMgr.CheckMsgDuplication(round, partyID.Id) {
			t.logger.Debugf("we received the duplicated message from party %s", partyID.Id)
			continue
		}

		partyInlist := func(el *btss.PartyID, l []*btss.PartyID) bool {
			for _, each := range l {
				if el == each {
					return true
				}
			}
			return false
		}

		// 5.2.5 check whether the source party is a malicious actor
		// If the message sender is already a malicious participant,  return with error
		t.culpritsLock.RLock()
		if len(t.culprits) != 0 && partyInlist(partyID, t.culprits) {
			t.logger.Errorf("the malicious party (party ID:%s) try to send incorrect message to me (party ID:%s)", partyID.Id, localMsgParty.PartyID().Id)
			t.culpritsLock.RUnlock()
			return fmt.Errorf("tss share verification failed")
		}
		t.culpritsLock.RUnlock()

		// 5.2.6 the info is collected and checked, put the task into the channel
		job := newJob(localMsgParty, msg.WiredBulkMsgs, round.MsgIdentifier, partyID, msg.Routing.IsBroadcast)
		tssJobChan <- job
	}

	close(tssJobChan)
	jobWg.Wait()
	return nil
}

type tssJob struct {
	wireBytes     []byte
	msgIdentifier string
	partyID       *btss.PartyID
	isBroadcast   bool
	localParty    btss.Party
	// acceptedShares map[conversion.RoundInfo][]string
}

func newJob(party btss.Party, wireBytes []byte, msgIdentifier string, from *btss.PartyID, isBroadcast bool) *tssJob {
	return &tssJob{
		wireBytes:     wireBytes,
		msgIdentifier: msgIdentifier,
		partyID:       from,
		isBroadcast:   isBroadcast,
		localParty:    party,
	}
}

// doTssJob process a job in tssJobChan, essentially call update to advance to the next round
func (t *TssInstance) doTssJob(tssJobChan chan *tssJob, jobWg *sync.WaitGroup) {
	defer func() {
		jobWg.Done()
	}()

	for tssjob := range tssJobChan {
		party := tssjob.localParty
		wireBytes := tssjob.wireBytes
		partyID := tssjob.partyID
		isBroadcast := tssjob.isBroadcast

		// 1. get round
		round, err := conversion.GetMsgRound(wireBytes, partyID, isBroadcast)
		if err != nil {
			t.logger.Errorf("broken tss share")
			continue
		}
		round.MsgIdentifier = tssjob.msgIdentifier

		// 2. Update party with library methods
		_, errUp := party.UpdateFromBytes(wireBytes, partyID, isBroadcast)
		if errUp != nil {
			// set blame with culprits
			err := t.processInvalidMsgBlame(round.RoundMsg, round, errUp)
			t.logger.Errorf("fail to apply the share to tss: %v", err)
			continue
		}
		// we need to retrieve the partylist again as others may update it once we process apply tss share
		t.blameMgr.UpdateAcceptShare(round, partyID.Id)
	}
}

func (t *TssInstance) processInvalidMsgBlame(roundInfo string, round conversion.RoundInfo, err *btss.Error) error {
	// now we get the culprits ID, invalid message and signature the culprits sent
	var culpritsID []string
	var invalidMsgs []*message.TaskMessage
	unicast := conversion.CheckUnicast(round)

	// 1 record culprits
	t.culpritsLock.Lock()
	t.culprits = append(t.culprits, err.Culprits()...)
	t.culpritsLock.Unlock()

	for _, el := range err.Culprits() {
		culpritsID = append(culpritsID, el.Id)
		key := fmt.Sprintf("%s-%s", el.Id, roundInfo)
		storedMsg := t.blameMgr.RoundMgr.Get(key)
		invalidMsgs = append(invalidMsgs, storedMsg)
	}

	// 2 package the blame node and signature information together to return
	var blameNodes []blame.Node
	var msgBody, sig []byte
	for i, partyID := range culpritsID {
		invalidMsg := invalidMsgs[i]
		if invalidMsg == nil {
			t.logger.Error("we cannot find the record of this curlprit, set it as blank")
			msgBody = []byte{}
			sig = []byte{}
		} else {
			msgBody = invalidMsg.Message
			sig = invalidMsg.Sig
		}
		blameNodes = append(blameNodes, blame.NewNode(partyID, msgBody, sig))
	}
	t.blameMgr.Blame.SetBlame(blame.TssBrokenMsg, blameNodes, unicast)
	return fmt.Errorf("fail to set bytes to local party: %w", err)
}

// Broadcast message processing1 =======================================================================================
// receiverBroadcastHashToPeers broadcasts the hash of the received msg to all nodes except the source and myself
func (t *TssInstance) receiverBroadcastHashToPeers(wireMsg *message.TaskMessage, msgType message.TssMsgType) error {
	var ids []uint64
	dataOwnerPartyID := wireMsg.Routing.From.Id
	for id := range t.p2pComm.Peers() {
		if id == dataOwnerPartyID || id == t.localPartyID {
			continue
		}
		tmpId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return fmt.Errorf("receiverBroadcastHashToPeers parse int error: %v", err)
		}
		ids = append(ids, tmpId)
	}

	t.logger.WithFields(logrus.Fields{"typ": msgType}).Debugf("receive broadcast msg")
	msgVerType := conversion.GetBroadcastMessageType(msgType)
	key := wireMsg.GetCacheKey()
	msgHash, err := conversion.BytesToHashString(wireMsg.Message)
	if err != nil {
		return fmt.Errorf("fail to calculate hash of the wire message: %w", err)
	}

	err = t.broadcastHashToPeers(key, msgHash, msgVerType, ids)
	if err != nil {
		t.logger.Errorf("fail to broadcast the hash to peers: %v", err)
		return err
	}
	return nil
}

func (t *TssInstance) broadcastHashToPeers(key, msgHash string, msgType message.TssMsgType, parties []uint64) error {
	broadcastConfirmMsg := &message.BroadcastConfirmMessage{
		FromID: t.localPartyID,
		Key:    key,
		Hash:   msgHash,
	}
	buf, err := json.Marshal(broadcastConfirmMsg)
	if err != nil {
		return fmt.Errorf("fail to marshal borad cast confirm message: %w", err)
	}
	t.logger.Debug("broadcast VerMsg to all other parties")

	wireMsg := &message.WireMessage{
		MsgID:   t.msgID,
		MsgType: msgType,
		MsgData: buf,
	}
	t.renderToP2P(&message.SendMsgChan{
		WireMsg:   wireMsg,
		PartiesID: parties,
	})

	return nil
}

// Broadcast message processing2 =======================================================================================
// applyShare handle the sharing confirmation of broadcast messages.
// check whether the number of caches is sufficient, if sufficient, can be stored to blame and update local
func (t *TssInstance) applyShare(localCacheItem *cache.LocalCacheItem, key string, msgType message.TssMsgType) error {
	unicast := true
	if localCacheItem.Msg.Routing.IsBroadcast {
		unicast = false
	}
	// 1. check for consistent validation of thresholds
	errHashCheck := t.hashCheck(localCacheItem, t.threshold)
	if errHashCheck != nil {
		// If there are not enough participants confirming, return and wait until there are enough participants confirming
		if errors.Is(errHashCheck, blame.ErrNotEnoughPeer) {
			return nil
		}
		// If there are enough participants, but the message we receive is inconsistent with the majority, we need to request the message again
		if errors.Is(errHashCheck, blame.ErrNotMajority) {
			t.logger.Errorf("we send request to get the message match with majority: %v", errHashCheck)
			// Null the local message because it is inconsistent with the majority and may be wrong
			localCacheItem.Msg = nil
			// Request the correct message from other participants
			return t.requestShareFromPeer(localCacheItem, t.threshold, key, msgType)
		}

		// Other errors
		blamePk, err := t.blameMgr.TssWrongShareBlame(localCacheItem.Msg)
		if err != nil {
			t.logger.Errorf("error in get the blame nodes： %v", err)
			t.blameMgr.Blame.SetBlame(blame.HashCheckFail, nil, unicast)
			return fmt.Errorf("error in getting the blame nodes %w, %s", blame.ErrHashCheck, errHashCheck.Error())
		}
		blameNode := blame.NewNode(blamePk, localCacheItem.Msg.Message, localCacheItem.Msg.Sig)
		t.blameMgr.Blame.SetBlame(blame.HashCheckFail, []blame.Node{blameNode}, unicast)
		return fmt.Errorf("%w, %s", blame.ErrHashCheck, errHashCheck)
	}

	// 2. check ok
	// - Save the original message to roundMgr
	// - updateLocal to advance rounds
	// - delete cache
	t.blameMgr.RoundMgr.Set(key, localCacheItem.Msg)
	if err := t.updateLocal(localCacheItem.Msg); nil != err {
		return fmt.Errorf("fail to update the message to local party: %w", err)
	}
	t.logger.Debugf("remove key: %s", key)
	// the information had been confirmed by all party , we don't need it anymore
	t.removeKey(key)
	return nil
}

// hashCheck checkes that the hashes confirmed by most people are consistent and meet the threshold, besides, the sender of the message should not be in the confirmation list.
// - 1 Checks whether the number reaches the threshold. If not, the verification fails
// - 2 Check whether there is a sender in the list. If so, delete the sender and return with an error
// - 3 Check whether the hash is consistent, if not, then return with validation failure
func (t *TssInstance) hashCheck(localCacheItem *cache.LocalCacheItem, threshold int) error {
	// 1
	if localCacheItem.TotalConfirmParty() < threshold {
		t.logger.Debug("not enough nodes to evaluate the hash")
		return blame.ErrNotEnoughPeer
	}

	targetHashValue := localCacheItem.Hash
	dataOwner := localCacheItem.Msg.Routing.From
	localCacheItem.Lock.Lock()
	defer localCacheItem.Lock.Unlock()
	// 2
	for partyID := range localCacheItem.ConfirmedList {
		if partyID == dataOwner.Id {
			t.logger.Warnf("we detect that the data owner try to send the hash for his own message\n")
			delete(localCacheItem.ConfirmedList, partyID)
			return blame.ErrHashFromOwner
		}
	}

	// 3
	hash, err := t.getMsgHash(localCacheItem, threshold)
	if err != nil {
		return err
	}
	if targetHashValue == hash {
		t.logger.Debugf("hash check complete for messageID: %v", t.msgID)
		return nil
	}
	return blame.ErrNotMajority
}

// getMsgHash query the hash of the message that most people agree to confirm. The number of confirmed messages must be t-1 or higher
// t-1 was chosen because plus the sender of the message itself is t
func (t *TssInstance) getMsgHash(localCacheItem *cache.LocalCacheItem, threshold int) (string, error) {
	hash, freq, err := getHighestFreq(localCacheItem.ConfirmedList)
	if err != nil {
		t.logger.Errorf("fail to get the hash freq: %v", err)
		return "", blame.ErrHashCheck
	}
	if freq < threshold-1 {
		t.logger.Debugf("fail to have more than 2/3 peers agree on the received message threshold(%d)--total confirmed(%d)\n", threshold, freq)
		return "", blame.ErrHashInconsistency
	}
	return hash, nil
}

func getHighestFreq(confirmedList map[string]string) (string, int, error) {
	if len(confirmedList) == 0 {
		return "", 0, errors.New("empty input")
	}
	freq := make(map[string]int, len(confirmedList))
	for _, n := range confirmedList {
		freq[n]++
	}
	maxFreq := -1
	var data string
	for key, counter := range freq {
		if counter > maxFreq {
			maxFreq = counter
			data = key
		}
	}
	return data, maxFreq, nil
}

// requestShareFromPeer requests the correct message from other participants
// Call condition: Local broadcast confirmation that MSG in localCacheItem is nil
// - The local has not received this message, but has received an acknowledgement hash from another participant
// - Local received the message, but the message hash is not consistent with the majority of people is empty, need to request the correct MSG from other participants
// Construct a message of type Contraol, where from is itself and to is the other participants unanimously identified in localCacheItem
func (t *TssInstance) requestShareFromPeer(localCacheItem *cache.LocalCacheItem, threshold int, key string, msgType message.TssMsgType) error {
	// query the hash of the message that most people agree to confirm
	targetHash, err := t.getMsgHash(localCacheItem, threshold)
	if err != nil {
		// If not found, then we don't know what to ask for (or we can't guarantee that the message is correct if it doesn't reach t) and just return
		t.logger.Debug("we do not know which message to request, so we quit")
		return nil
	}

	// get all participants for which the message has been confirmed consistently
	var partysIDs []string
	for partyID, hash := range localCacheItem.ConfirmedList {
		if hash == targetHash {
			partysIDs = append(partysIDs, partyID)
		}
	}

	msg := &message.TssControl{
		FromID:      t.localPartyID,
		ReqHash:     targetHash,
		ReqKey:      key,
		RequestType: 0,
		Msg:         nil,
	}
	t.blameMgr.ShareMgr.Set(targetHash)

	t.logger.WithFields(logrus.Fields{"key": key, "localID": t.localPartyID, "others": partysIDs}).
		Debugf("we have no msg, request from others")
	switch msgType {
	case message.TSSKeyGenVerMsg:
		msg.RequestType = message.TSSKeyGenMsg
		return t.processRequestMsgFromParty(partysIDs, msg, true)
	case message.TSSKeySignVerMsg:
		msg.RequestType = message.TSSKeySignMsg
		return t.processRequestMsgFromParty(partysIDs, msg, true)
	case message.TSSKeySignMsg, message.TSSKeyGenMsg:
		msg.RequestType = msgType
		return t.processRequestMsgFromParty(partysIDs, msg, true)
	default:
		t.logger.Debug("unknown message type")
		return nil
	}
}

func (t *TssInstance) processRequestMsgFromParty(partiesIDStr []string, msg *message.TssControl, requester bool) error {
	// we need to send msg to the peer
	if !requester {
		if msg == nil {
			return errors.New("empty message")
		}
		reqKey := msg.ReqKey
		storedMsg := t.blameMgr.RoundMgr.Get(reqKey)
		if storedMsg == nil {
			t.logger.Debug("we do not have this message either")
			return nil
		}
		msg.Msg = storedMsg
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("fail to marshal the request body %w", err)
	}

	wireMsg := &message.WireMessage{
		MsgID:   t.msgID,
		MsgType: message.TSSControlMsg,
		MsgData: data,
	}

	partiesID := []uint64{}
	for _, id := range partiesIDStr {
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return fmt.Errorf("parse uint error: %v", err)
		}
		partiesID = append(partiesID, idUint)
	}
	t.renderToP2P(&message.SendMsgChan{
		WireMsg:   wireMsg,
		PartiesID: partiesID,
	})
	return nil
}

func (t *TssInstance) removeKey(key string) {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	delete(t.unConfirmedMessages, key)
}

// TssVer message processing1 ==========================================================================================
// checkDupAndUpdateVerMsg checkes whether this peer has already send us the VerMsg before update
func (t *TssInstance) checkDupAndUpdateVerMsg(bMsg *message.BroadcastConfirmMessage, partyID string) bool {
	localCacheItem := t.getLocalCacheItem(bMsg.Key)
	// we check whether this node has already sent the VerMsg message to avoid eclipse of others VerMsg
	if localCacheItem == nil {
		bMsg.FromID = partyID
		return true
	}

	localCacheItem.Lock.Lock()
	defer localCacheItem.Lock.Unlock()
	if _, ok := localCacheItem.ConfirmedList[partyID]; ok {
		return false
	}
	bMsg.FromID = partyID
	return true
}

// TssVer message processing2 ==========================================================================================
func (t *TssInstance) processVerMsg(broadcastConfirmMsg *message.BroadcastConfirmMessage, msgType message.TssMsgType) error {
	t.logger.WithFields(logrus.Fields{
		"from": broadcastConfirmMsg.FromID,
		"key":  broadcastConfirmMsg.Key,
		"type": msgType,
	}).Debug("process ver msg")
	defer t.logger.Debug("finish process ver msg")
	if nil == broadcastConfirmMsg {
		return nil
	}
	key := broadcastConfirmMsg.Key
	localCacheItem := t.getLocalCacheItem(key)
	if nil == localCacheItem {
		// we didn't receive the TSS Message yet
		t.logger.Debugf("%s doesn't exist yet,add a new one", key)
		// Note that the msg field is nil because the validation they broadcast only has hash and no MSG text, which needs to be filled in when they receive it
		localCacheItem = cache.NewLocalCacheItem(nil, broadcastConfirmMsg.Hash)
		t.setLocalUnconfirmedMessages(key, localCacheItem)
	}

	// update confirm list
	localCacheItem.UpdateConfirmList(broadcastConfirmMsg.FromID, broadcastConfirmMsg.Hash)
	t.logger.WithFields(logrus.Fields{"key": broadcastConfirmMsg.Key}).Debugf("total confirmed parties:%+v", localCacheItem.ConfirmedList)

	// if we do not have the msg, we try to request from peer otherwise, we apply this share
	if localCacheItem.Msg == nil {
		return t.requestShareFromPeer(localCacheItem, t.threshold, key, msgType)
	}
	return t.applyShare(localCacheItem, key, msgType)
}
