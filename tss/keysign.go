package tss

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	bcommon "github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/meshplus/bitxhub-core/tss/keysign"
	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/meshplus/bitxhub-core/tss/storage"
	"go.uber.org/atomic"
)

var (
	ErrJoinPartyTimeout = errors.New("fail to join party, timeout")
	ErrLeaderNotReady   = errors.New("leader not reachable")
	ErrSignReceived     = errors.New("signature received")
	ErrNotActiveSigner  = errors.New("not active signer")
	ErrSigGenerated     = errors.New("signature generated")
)

func (t *TssInstance) Keysign(req keysign.Request) (*keysign.Response, error) {
	t.logger.Infof("Received keysign request, signers: %v, msg: %v", req.SignerPubKeys, req.Messages)

	// 1. analysis req
	// 1.1 get msgsToSign
	var msgsToSign [][]byte
	for _, val := range req.Messages {
		msgToSign, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("fail to decode message(%s): %w", strings.Join(req.Messages, ","), err)
		}
		msgsToSign = append(msgsToSign, msgToSign)
	}
	sort.SliceStable(msgsToSign, func(i, j int) bool {
		ma, err := conversion.MsgToHashInt(msgsToSign[i])
		if err != nil {
			t.logger.Errorf("fail to convert the hash value")
		}
		mb, err := conversion.MsgToHashInt(msgsToSign[j])
		if err != nil {
			t.logger.Errorf("fail to convert the hash value")
		}
		if ma.Cmp(mb) == -1 {
			return false
		}
		return true
	})

	// 1.2 check signers num
	if len(req.SignerPubKeys) <= t.threshold {
		return nil, fmt.Errorf("at least t+1 signers are required: t-%d", t.threshold)
	}

	var generatedSig *keysign.Response
	var errGen error

	// 2 The first coroutine: generate the signature ourselves
	generatedSig, errGen = t.generateSignature(msgsToSign, req, t.keygenLocalState)

	return generatedSig, errGen
}

func (t *TssInstance) generateSignature(msgsToSign [][]byte, req keysign.Request, localStateItem *storage.KeygenLocalState) (*keysign.Response, error) {
	// 1. Determine if you are one of the participants in the signature request
	allParticipantKeys := req.SignerPubKeys
	localKey := t.localPubK
	pid, _ := conversion.GetPIDFromPubKey(localKey)
	t.logger.Infof("local pid: %s", pid)
	isSignMember := false
	for _, el := range allParticipantKeys {
		elData, _ := el.Bytes()
		localData, _ := localKey.Bytes()
		if string(localData) == string(elData) {
			isSignMember = true
			break
		}
	}

	if !isSignMember {
		t.logger.Infof("we(%s) are not the active signer", pid)
		return nil, ErrNotActiveSigner
	}

	// 2. start sign
	signatureData, err := t.SignMessage(msgsToSign, localStateItem, allParticipantKeys)
	var status conversion.Status
	if err != nil {
		t.logger.Errorf("err in keysign: %v", err)
		status = conversion.Fail
	} else {
		status = conversion.Success
	}

	return keysign.NewResponse(
		conversion.BatchSignatures(signatureData, msgsToSign),
		status,
		t.blameMgr.Blame,
	), err
}

func (t *TssInstance) SignMessage(msgsToSign [][]byte, localStateItem *storage.KeygenLocalState, signers []crypto.PubKey) ([]*bcommon.SignatureData, error) {
	// 1. get parties info
	partiesID, localPartyID, err := conversion.GetParties(signers, t.localPubK, t.p2pComm.Peers())
	if err != nil {
		return nil, fmt.Errorf("fail to form key sign party: %w", err)
	}
	t.logger.Debugf("----parties len: %d", len(partiesID))

	// 2. make channel
	outCh := make(chan btss.Message, 2*len(partiesID)*len(msgsToSign))
	endCh := make(chan bcommon.SignatureData, len(partiesID)*len(msgsToSign))
	errCh := make(chan struct{})

	// sign multiple messages and construct a map mapping from messages to localParty
	keySignPartyMap := new(sync.Map)
	// 3. constructing the party for each message
	for i, val := range msgsToSign {
		m, err := conversion.MsgToHashInt(val)
		if err != nil {
			return nil, fmt.Errorf("fail to convert msg to hash int: %w", err)
		}
		moniker := m.String() + ":" + strconv.Itoa(i)

		// 3.1 Construct keysign params
		// - ctx
		ctx := btss.NewPeerContext(partiesID)
		// - localparty
		localPartyID.Moniker = moniker
		// - parties num
		// - threshold
		params := btss.NewParameters(ctx, localPartyID, len(partiesID), t.threshold)

		// 3.2 construct local party
		keySignParty := signing.NewLocalParty(t.logger, m, params, localStateItem.LocalData, outCh, endCh)
		t.localPartyID = keySignParty.PartyID().Id

		// 3.3 record Message-to-party mapping
		keySignPartyMap.Store(moniker, keySignParty)
	}

	// 4. set party info
	// partyID.id到partyID/pid映射关系
	partyIDMap, err := conversion.GetPatyIDInfoMap(partiesID)
	if err != nil {
		t.logger.Errorf("fail to get partyID info map: %v", err)
		return nil, err
	}
	partyInfo := &conversion.PartyInfo{
		PartyMap:   keySignPartyMap,
		PartyIDMap: partyIDMap,
	}
	t.setPartyInfo(partyInfo)
	t.blameMgr.SetPartyInfo(partyInfo)

	// 5 start the key sign
	var keySignWg sync.WaitGroup
	keySignWg.Add(2)

	// 5.1 The first coroutine: batch signature
	go func() {
		defer keySignWg.Done()
		ret := t.startBatchSigning(keySignPartyMap, len(msgsToSign))
		if !ret {
			close(errCh)
		}
	}()

	// 5.2 Second thread: process received p2p messages - in this case, received messages are those sent by 5.3
	go t.ProcessInboundMessages(&keySignWg)

	// 5.3 Current main process: advance the execution of keySign process - send out the pending p2p messages given in the library method as required
	results, err := t.processKeySign(len(msgsToSign), errCh, outCh, endCh)
	if err != nil {
		close(t.inMsgHandleStopChan)
		return nil, fmt.Errorf("fail to process key sign: %w", err)
	}

	kenGenTicker := time.NewTicker(t.conf.KeySignTimeout)
	defer kenGenTicker.Stop()
	select {
	case <-kenGenTicker.C:
		close(t.inMsgHandleStopChan)
	case <-t.taskDoneChan:
		close(t.inMsgHandleStopChan)
	}
	keySignWg.Wait()

	// 6 successfully
	t.logger.Infof("%s successfully sign the message", t.localPartyID)
	sort.SliceStable(results, func(i, j int) bool {
		a := new(big.Int).SetBytes(results[i].M)
		b := new(big.Int).SetBytes(results[j].M)

		return a.Cmp(b) != -1
	})

	return results, nil
}

func (t *TssInstance) startBatchSigning(keySignPartyMap *sync.Map, msgNum int) bool {
	// start the batch sign
	var keySignWg sync.WaitGroup
	ret := atomic.NewBool(true)
	keySignWg.Add(msgNum)
	keySignPartyMap.Range(func(key, value interface{}) bool {
		eachParty := value.(btss.Party)
		go func(eachParty btss.Party) {
			defer keySignWg.Done()
			if err := eachParty.Start(); err != nil {
				t.logger.Errorf("fail to start key sign party: %v", err)
				ret.Store(false)
			}
			t.logger.Infof("local party(%s) %s is ready", eachParty.PartyID().Id, eachParty.PartyID().Moniker)
		}(eachParty)
		return true
	})
	keySignWg.Wait()
	return ret.Load()
}

// Handles messages returned by KeySign library methods that need to be sent
// - outCh: the messages that need to be sent for each turn
// - endCh：the message that ultimately needs to be stored
// - reqNum： number of messages to be signed
func (t *TssInstance) processKeySign(reqNum int,
	errChan chan struct{},
	outCh <-chan btss.Message,
	endCh <-chan bcommon.SignatureData) ([]*bcommon.SignatureData, error) {
	defer t.logger.Debug("finished keysign processd")
	t.logger.Debug("start to read messages from local party")

	var signatures []*bcommon.SignatureData
	kenSignTicker := time.NewTicker(t.conf.KeySignTimeout)
	defer kenSignTicker.Stop()
	for {
		select {
		case <-t.stopChan: // when TSS processor receive signal to quit
			return nil, errors.New("received exit signal")
		case <-errChan: // when key sign return
			t.logger.Error("key sign failed")
			return nil, errors.New("error channel closed fail to start local party")
		case <-kenSignTicker.C:
			// we bail out after KeySignTimeoutSeconds
			t.logger.Errorf("fail to sign message with %s", t.conf.KeySignTimeout.String())

			// 1. get fail reason
			lastMsg := t.blameMgr.GetLastMsg()
			failReason := t.blameMgr.Blame.FailReason
			if failReason == "" {
				failReason = blame.TssTimeout
			}

			// 2. blame problem node in unicast
			if !lastMsg.IsBroadcast() {
				// 2.1 If the last message was unicast, the current message is processed
				blamePeersUnicast, err := t.blameMgr.GetUnicastBlame(lastMsg.Type())
				if err != nil {
					t.logger.Error("error in get unicast blame")
				}
				if len(blamePeersUnicast) > 0 && len(blamePeersUnicast) <= t.threshold {
					t.blameMgr.Blame.SetBlame(failReason, blamePeersUnicast, true)
				}
			} else {
				// 2.2 If the last message was broadcast, the last unicast message is processed
				blamePeersUnicast, err := t.blameMgr.GetUnicastBlame(conversion.GetPreviousKeySignUicast(lastMsg.Type()))
				if err != nil {
					t.logger.Error("error in get unicast blame")
				}
				if len(blamePeersUnicast) > 0 && len(blamePeersUnicast) <= t.threshold {
					t.blameMgr.Blame.SetBlame(failReason, blamePeersUnicast, true)
				}
			}

			// 3. blame problem node in broadcast
			blameNodesBroadcast, err := t.blameMgr.GetBroadcastBlame(lastMsg.Type())
			if err != nil {
				t.logger.Error("error in get broadcast blame")
			}
			t.blameMgr.Blame.AddBlameNodes(blameNodesBroadcast...)

			// 4. blame the node fail to send the shares to the node with batch signing
			// if we cannot find the blame node, we check whether everyone send me the share
			if len(t.blameMgr.Blame.BlameNodes) == 0 {
				blameNodesMisingShare, isUnicast, err := t.blameMgr.TssMissingShareBlame(message.TSSKEYSIGNROUNDS)
				if err != nil {
					t.logger.Error("fail to get the node of missing share ")
				}

				if len(blameNodesMisingShare) > 0 && len(blameNodesMisingShare) <= t.threshold {
					t.blameMgr.Blame.AddBlameNodes(blameNodesMisingShare...)
					t.blameMgr.Blame.IsUnicast = isUnicast
				}
			}

			return nil, blame.ErrTssTimeOut

		case msg := <-outCh:
			t.logger.Debugf(">>>>>>>>>>key sign msg: %s", msg.String())
			t.blameMgr.SetLastMsg(msg)
			err := t.ProcessOutCh(msg, message.TSSKeySignMsg)
			if err != nil {
				t.logger.Errorf("fail to process the message")
				return nil, err
			}

		case msg := <-endCh:
			signatures = append(signatures, &msg)
			if len(signatures) == reqNum {
				t.logger.Debug("we have done the key sign")
				err := t.NotifyTaskDone()
				if err != nil {
					t.logger.Error("fail to broadcast the keysign done")
				}
				return signatures, nil
			}
		}
	}
}
