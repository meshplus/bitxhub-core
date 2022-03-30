package blame

import (
	"fmt"

	btss "github.com/binance-chain/tss-lib/tss"
	mapset "github.com/deckarep/golang-set"
	"github.com/meshplus/bitxhub-core/tss/message"
)

// GetUnicastBlame blames the node who cause the timeout in unicast message
func (m *Manager) GetUnicastBlame(lastMsgType string) ([]Node, error) {
	m.lastMsgLocker.RLock()

	if len(m.lastUnicastParty) == 0 {
		m.lastMsgLocker.RUnlock()
		m.logger.Debugf("we do not have any unicast message received yet")
		return nil, nil
	}
	partysMap := make(map[string]bool)
	partysID, ok := m.lastUnicastParty[lastMsgType]
	m.lastMsgLocker.RUnlock()
	if !ok {
		return nil, fmt.Errorf("fail to find peers of the given msg type %w", ErrTssTimeOut)
	}
	for _, el := range partysID {
		partysMap[el] = true
	}

	// online nodes are the ones that sent the last unicast message
	var onlinePartys []string
	for key := range partysMap {
		onlinePartys = append(onlinePartys, key)
	}

	// blamePartys are the participants not in the onlinePartye, that is, the problem node in the unicast message
	_, blamePartys, err := m.GetBlamePartyIDsLists(onlinePartys)
	if err != nil {
		m.logger.Errorf("fail to get the blamed peers: %v", err)
		return nil, fmt.Errorf("fail to get the blamed peers: %w", ErrTssTimeOut)
	}

	var blameNodes []Node
	for _, el := range blamePartys {
		blameNodes = append(blameNodes, NewNode(el, nil, nil))
	}
	return blameNodes, nil
}

// GetBroadcastBlame blames the node who cause the timeout in broadcast message
func (m *Manager) GetBroadcastBlame(lastMessageType string) ([]Node, error) {
	blameParties, err := m.tssTimeoutBlame(lastMessageType, m.partyInfo.PartyIDMap)
	if err != nil {
		m.logger.Errorf("fail to get the blamed peers: %v", err)
		return nil, fmt.Errorf("fail to get the blamed peers: %w", ErrTssTimeOut)
	}

	var blameNodes []Node
	for _, el := range blameParties {
		blameNodes = append(blameNodes, NewNode(el, nil, nil))
	}
	return blameNodes, nil
}

// tssTimeoutBlame queries parties that do not participate in messages of the specified type
func (m *Manager) tssTimeoutBlame(lastMessageType string, partyIDMap map[string]*btss.PartyID) ([]string, error) {
	peersSet := mapset.NewSet()
	for _, el := range partyIDMap {
		if el.Id != m.localPartyID {
			peersSet.Add(el.Id)
		}
	}

	// Query the sender of a message of a specified type
	standbyNodes := m.RoundMgr.GetByRound(lastMessageType)
	if len(standbyNodes) == 0 {
		return nil, nil
	}
	s := make([]interface{}, len(standbyNodes))
	for i, v := range standbyNodes {
		s[i] = v
	}
	standbySet := mapset.NewSetFromSlice(s)

	var blames []string
	diff := peersSet.Difference(standbySet).ToSlice()
	for _, el := range diff {
		blames = append(blames, el.(string))
	}
	return blames, nil
}

// TssWrongShareBlame blames the node who provide the wrong share
func (m *Manager) TssWrongShareBlame(wiredMsg *message.WireMessage) (string, error) {
	shareOwner := wiredMsg.Routing.From
	return shareOwner.Id, nil
}

// TssMissingShareBlame blames the node fail to send the shares to the node
// with batch signing, we need to put the accepted shares into different message group
// then search the missing share for each keysign message
func (m *Manager) TssMissingShareBlame(rounds int) ([]Node, bool, error) {
	var blameNodes []Node
	acceptedShareForMsg := make(map[string][][]string)
	var partyIDs []string
	isUnicast := false
	m.acceptShareLocker.Lock()

	for roundInfo, value := range m.acceptedShares {
		cachedShares, ok := acceptedShareForMsg[roundInfo.MsgIdentifier]
		if !ok {
			cachedShares := make([][]string, rounds)
			cachedShares[roundInfo.Index] = value
			acceptedShareForMsg[roundInfo.MsgIdentifier] = cachedShares
			continue
		}
		cachedShares[roundInfo.Index] = value
	}
	m.acceptShareLocker.Unlock()

	for _, cachedShares := range acceptedShareForMsg {
		// we search from the first round to find the missing
		for index, el := range cachedShares {
			// the len of el(who accepted the share) plus 1 is equal to the number of participants,
			// that is, everyone has accepted the shared information
			if len(el)+1 == len(m.partyInfo.PartyIDMap) {
				continue
			}
			// we find whether the missing share is in unicast
			if rounds == message.TSSKEYGENROUNDS {
				// we are processing the keygen and if the missing shares is in second round(index=1)
				// we mark it as the unicast.
				if index == 1 {
					isUnicast = true
				}
			}

			if rounds == message.TSSKEYSIGNROUNDS {
				// we are processing the keysign and if the missing shares is in the 5 round(index<1)
				// we all mark it as the unicast, because in some cases, the error will be detected
				// in the following round, so we cannot "trust" the node stops at the current round.
				if index < 5 {
					isUnicast = true
				}
			}
			// we add our own id to avoid blame ourselves
			// since all the local parties have the same id, so we just need to take one of them to get the peer

			el = append(el, m.localPartyID)
			for _, partyIDid := range el {
				partyIDs = append(partyIDs, partyIDid)
			}
			break
		}

		blameParties, err := m.getBlamePartyIDsNotInList(partyIDs)
		if err != nil {
			return nil, isUnicast, err
		}
		for _, el := range blameParties {
			node := Node{
				el,
				nil,
				nil,
			}
			blameNodes = append(blameNodes, node)
		}
	}
	return blameNodes, isUnicast, nil
}
