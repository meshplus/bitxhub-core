package tss

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/meshplus/bitxhub-core/tss/message"

	"github.com/meshplus/bitxhub-core/tss/blame"

	"github.com/meshplus/bitxhub-core/tss/conversion"

	bkg "github.com/binance-chain/tss-lib/ecdsa/keygen"
	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/meshplus/bitxhub-core/tss/keygen"
	"github.com/meshplus/bitxhub-core/tss/storage"

	//"github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

// Keygen generates the key share of the participants of the threshold signature
func (t *TssManager) Keygen(req keygen.Request) (*keygen.Response, error) {
	t.logger.WithFields(logrus.Fields{}).Info("Received keygen request")
	t.tssKeyGenLocker.Lock()
	defer t.tssKeyGenLocker.Unlock()
	status := conversion.Success

	// 1 Set msg info
	msgID, err := req.RequestToMsgId()
	if err != nil {
		return nil, fmt.Errorf("fail to convert request to msgID: %w", err)
	}
	t.setTssMsgInfo(msgID, 1)

	// 2 Get parties info
	partiesID, localPartyID, err := conversion.GetParties(req.Pubkeys, t.localPubK, t.p2pComm.Peers())
	if err != nil {
		return nil, fmt.Errorf("fail to get keygen parties: %w", err)
	}

	// 3 Construct persistent data information whitch KeyGen required
	partyPksDataMap, err := conversion.GetPubKeyDatasMapFromPartyIDs(partiesID)
	if err != nil {
		return nil, fmt.Errorf("fail to get pids from party ids: %w", err)
	}
	localPkData, err := conversion.GetPubKeyDataFromPartyID(localPartyID)
	if err != nil {
		return nil, fmt.Errorf("fail to get pid from party id: %w", err)
	}
	keyGenLocalStateItem := &storage.KeygenLocalState{
		ParticipantPksMap: partyPksDataMap,
		LocalPartyPk:      localPkData,
	}

	// 4 Construct keygen params
	// - ctx
	ctx := btss.NewPeerContext(partiesID)
	// - localPartyID
	// - parties num
	// - threshold
	params := btss.NewParameters(ctx, localPartyID, len(partiesID), t.threshold)
	// - msg to be sent channel
	outCh := make(chan btss.Message, len(partiesID))
	// - final channel for storing information
	endCh := make(chan bkg.LocalPartySaveData, len(partiesID))
	// - err channel
	errChan := make(chan struct{})
	// - preparams
	if t.keygenPreParams == nil {
		t.logger.Error("keygen: empty pre-parameters")
		return nil, fmt.Errorf("empty keygen pre-parameters")
	}

	// 5 Construct local party
	keyGenParty := bkg.NewLocalParty(t.logger, params, outCh, endCh, *t.keygenPreParams)
	t.localPartyID = keyGenParty.PartyID().Id

	// 6 Set parties info
	// 6.1 Message-to-party mapping, the Keygen only needs one
	keyGenPartyMap := new(sync.Map)
	keyGenPartyMap.Store("", keyGenParty)
	// 6.2 partyID.id-to-partyID mapping
	partyIDMap, err := conversion.GetPatyIDInfoMap(partiesID)
	if err != nil {
		t.logger.Errorf("faile to get partyID info map: %v", err)
		return nil, err
	}
	partyInfo := &conversion.PartyInfo{
		PartyMap:   keyGenPartyMap,
		PartyIDMap: partyIDMap,
	}
	t.setPartyInfo(partyInfo)
	t.blameMgr.SetPartyInfo(partyInfo)

	// 7 Start Keygen
	var keyGenWg sync.WaitGroup
	keyGenWg.Add(2)
	// 7.1 First thread: Call the library method to start keyGen
	go func() {
		defer keyGenWg.Done()
		defer t.logger.Infof(">>>>>>>>>>>>>. keyGenParty started: n-%d t-%d, localParty: %s", len(partiesID), t.threshold, t.localPartyID)
		if err := keyGenParty.Start(); nil != err {
			t.logger.Errorf("keygen: fail to start keygen party: %v", err)
			close(errChan)
		}
	}()

	// 7.2 Second thread: process received p2p messages - in this case, received messages are those sent by 7.3
	go t.ProcessInboundMessages(&keyGenWg)

	// 7.3 Current main process: advance the execution of keyGen process - send out the pending p2p messages given in the library method as required
	newPubKey, newPubAddr, err := t.processKeyGen(errChan, outCh, endCh, keyGenLocalStateItem)
	if err != nil {
		close(t.inMsgHandleStopChan)
		return nil, fmt.Errorf("fail to process keygen: %w", err)
	}

	select {
	case <-time.After(time.Second * t.conf.KeyGenTimeout):
		close(t.inMsgHandleStopChan)
	case <-t.taskDoneChan:
		close(t.inMsgHandleStopChan)
	}
	keyGenWg.Wait()

	pid, _ := conversion.GetPIDFromPartyID(keyGenParty.PartyID())
	t.logger.WithFields(logrus.Fields{
		"partyID":    t.localPartyID,
		"partyPID":   pid,
		"newPubkey":  newPubKey,
		"newPubAddr": newPubAddr,
	}).Info(">>>>>>>>>>>>>. keygen success!!!")

	return keygen.NewResponse(
		newPubKey,
		newPubAddr,
		status,
		t.blameMgr.Blame,
	), nil
}

// Handles messages returned by KeyGen library methods that need to be sent
// - outCh: the messages that need to be sent for each turn
// - endCh：the message that ultimately needs to be stored
func (t *TssManager) processKeyGen(errChan chan struct{},
	outCh <-chan btss.Message,
	endCh <-chan bkg.LocalPartySaveData,
	keyGenLocalStateItem *storage.KeygenLocalState) (*ecdsa.PublicKey, string, error) {
	defer t.logger.Debug("finished keygen process")
	t.logger.Debug("start to read messages from local party")

	for {
		select {
		case <-t.stopChan: // when TSS processor receive signal to quit
			return nil, "", errors.New("received exit signal")
		case <-errChan: // when keyGenParty return
			t.logger.Error("key gen failed")
			return nil, "", fmt.Errorf("error channel closed fail to start local party")
		case <-time.After(t.conf.KeyGenTimeout): // key gen timeout
			// we bail out after KeyGenTimeoutSeconds
			t.logger.Errorf("fail to generate key in time %s", t.conf.KeyGenTimeout.String())

			// 1. get fail reason
			failReason := t.blameMgr.Blame.FailReason
			if failReason == "" {
				failReason = blame.TssTimeout
			}

			// 2. get last msg
			lastMsg := t.blameMgr.GetLastMsg()
			if lastMsg == nil {
				t.logger.Errorf("fail to start the keygen, the last produced message of this node is none")
				return nil, "", fmt.Errorf("timeout before shared message is generated")
			}

			// 3. blame problem node in unicast
			//   only KEYGEN2aUnicast is unicast during keygen
			blameNodesUnicast, err := t.blameMgr.GetUnicastBlame(message.KEYGEN2aUnicast)
			if err != nil {
				t.logger.Errorf("error in get unicast blame")
			}
			if len(blameNodesUnicast) > 0 && len(blameNodesUnicast) <= t.threshold {
				t.blameMgr.Blame.SetBlame(failReason, blameNodesUnicast, true)
			}

			// 4. blame problem node in broadcast
			blameNodesBroadcast, err := t.blameMgr.GetBroadcastBlame(lastMsg.Type())
			if err != nil {
				t.logger.Errorf("error in get broadcast blame")
			}
			t.blameMgr.Blame.AddBlameNodes(blameNodesBroadcast...)

			// 5. blame the node fail to send the shares to the node with batch signing
			if len(t.blameMgr.Blame.BlameNodes) == 0 {
				blameNodesMisingShare, isUnicast, err := t.blameMgr.TssMissingShareBlame(message.TSSKEYGENROUNDS)
				if err != nil {
					t.logger.Errorf("fail to get the node of missing share ")
				}
				if len(blameNodesMisingShare) > 0 && len(blameNodesMisingShare) <= t.threshold {
					t.blameMgr.Blame.AddBlameNodes(blameNodesMisingShare...)
					t.blameMgr.Blame.IsUnicast = isUnicast
				}
			}

			return nil, "", blame.ErrTssTimeOut
		case tssMsg := <-outCh: // get msg to send
			t.logger.Debugf(">>>>>>>>>> key gen msg: %s", tssMsg.String())
			t.blameMgr.SetLastMsg(tssMsg)
			err := t.ProcessOutCh(tssMsg, pb.Message_TSS_KEY_GEN)
			if err != nil {
				t.logger.Errorf("fail to process the message")
				return nil, "", err
			}
		case tssSaveData := <-endCh: // end
			t.logger.Debugf("keygen finished successfully: %s", tssSaveData.ECDSAPub.Y().String())

			// 1. notify task done
			if err := t.NotifyTaskDone(); err != nil {
				t.logger.Errorf("fail to broadcast the keygen done")
			}

			// 2. save local state to file
			ecdsaPk, err := conversion.GetTssPubKey(tssSaveData.ECDSAPub)
			if err != nil {
				t.logger.Errorf("fail to get threshold pubkey: %v", err)
				return nil, "", fmt.Errorf("fail to get threshold pubkey: %w", err)
			}
			pubAddr, pubData, err := conversion.GetPubKeyInfoFromECDSAPubkey(ecdsaPk)
			if err != nil {
				t.logger.Errorf("fail to convert ecdsa pubkey to pubkey addr and byte: %v", err)
				return nil, "", fmt.Errorf("fail to convert ecdsa pubkey to pubkey addr and byte: %w", err)
			}
			keyGenLocalStateItem.LocalData = tssSaveData
			keyGenLocalStateItem.PubKeyData = pubData
			keyGenLocalStateItem.PubKeyAddr = pubAddr
			if err = t.stateMgr.SaveLocalState(keyGenLocalStateItem); err != nil {
				t.logger.Errorf("fail to save keygen result to storage: %w", err)
			}

			// 3. save tss info to memory
			t.keygenLocalState = keyGenLocalStateItem
			t.logger.Infof("processKeyGen end")
			return ecdsaPk, pubAddr, nil
		}
	}
}
