package conversion

import (
	"fmt"
	"math/big"
	"sync"

	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

// PartyInfo the information used by tss key gen and key sign
type PartyInfo struct {
	// Message identifier -> localParty of the message
	// - only one in keygen
	// - maybe some in keysign. Each message to be signed corresponds to a corresponding current participant
	PartyMap *sync.Map
	// partyID.id -> partyID
	PartyIDMap map[string]*btss.PartyID
}

// GetParties get parties and localParty according the parties pubkey and peersInfo.
// PeersInfo is needed to match the partyid.id with the serial number ID of the P2P Peer.
// []pubkeys, pubkey, peersInfo -> []partyIDs, partyID
func GetParties(keys []crypto.PubKey, localPK crypto.PubKey, peerMap map[string]*peer.AddrInfo) ([]*btss.PartyID, *btss.PartyID, error) {
	var localPartyID *btss.PartyID
	var unSortedPartiesID []*btss.PartyID

	for _, pk := range keys {
		pkBytes, err := pk.Raw()
		if err != nil {
			return nil, nil, fmt.Errorf("fail to get pubkey bytes: %w", err)
		}
		key := new(big.Int).SetBytes(pkBytes)

		// set p2p peer id to party id
		partyPid, err := peer.IDFromPublicKey(pk)
		if err != nil {
			return nil, nil, fmt.Errorf("IDFromPublicKey error: %v", err)
		}
		partyIDID := ""
		for peerID, peerInfo := range peerMap {
			if string(peerInfo.ID) == partyPid.String() {
				partyIDID = peerID
				break
			}
		}
		if partyIDID == "" {
			return nil, nil, fmt.Errorf("the pubkey is not one of peers %v", pk)
		}

		// Note: The `id` and `moniker` fields are for convenience to allow you to easily track participants.
		// The `id` should be a unique string representing this party in the network and `moniker` can be anything (even left blank).
		// The `uniqueKey` is a unique identifying key for this peer (such as its p2p public key) as a big.Int.
		partyID := btss.NewPartyID(partyIDID, "", key)
		if pk.Equals(localPK) {
			localPartyID = partyID
		}

		unSortedPartiesID = append(unSortedPartiesID, partyID)
	}
	if localPartyID == nil {
		return nil, nil, fmt.Errorf("local party is not in the list")
	}

	// The index in the partyID is automatically populated after sort
	partiesID := btss.SortPartyIDs(unSortedPartiesID)
	return partiesID, localPartyID, nil
}

// partyID -> pid
func GetPIDFromPartyID(partyID *btss.PartyID) (peer.ID, error) {
	if partyID == nil || !partyID.ValidateBasic() {
		return "", fmt.Errorf("invalid partyID")
	}
	pkBytes := partyID.KeyInt().Bytes()
	return GetPIDFromPubKeyData(pkBytes)
}

// partyID -> pubkeyData
func GetPubKeyDataFromPartyID(partyID *btss.PartyID) ([]byte, error) {
	if partyID == nil || !partyID.ValidateBasic() {
		return nil, fmt.Errorf("invalid partyID")
	}
	pkBytes := partyID.KeyInt().Bytes()
	return pkBytes, nil
}

// []partyID -> []pubkeyData
func GetPubKeyDatasFromPartyIDs(partyIDs []*btss.PartyID) ([][]byte, error) {
	pks := [][]byte{}
	for _, el := range partyIDs {
		pk, err := GetPubKeyDataFromPartyID(el)
		if err != nil {
			return nil, fmt.Errorf("get pk from party id error: %v", err)
		}
		pks = append(pks, pk)
	}
	return pks, nil
}

// []partyID -> map[partyID.id][pubkeyData]
func GetPubKeyDatasMapFromPartyIDs(partyIDs []*btss.PartyID) (map[string][]byte, error) {
	pksMap := map[string][]byte{}
	for _, el := range partyIDs {
		pk, err := GetPubKeyDataFromPartyID(el)
		if err != nil {
			return nil, fmt.Errorf("get pk from party id error: %v", err)
		}
		pksMap[el.Id] = pk
	}
	return pksMap, nil
}

// partyID -> pubkey
func GetPubKeyFromPartyID(partyID *btss.PartyID) (crypto.PubKey, error) {
	if partyID == nil || !partyID.ValidateBasic() {
		return nil, fmt.Errorf("invalid partyID")
	}
	pkBytes := partyID.KeyInt().Bytes()
	return GetPubKeyFromPubKeyData(pkBytes)
}

// []partyID -> []pubkey
func GetPubKeysFromPartyIDs(partyIDs []*btss.PartyID) ([]crypto.PubKey, error) {
	pks := []crypto.PubKey{}
	for _, el := range partyIDs {
		pk, err := GetPubKeyFromPartyID(el)
		if err != nil {
			return nil, fmt.Errorf("get pk from party id error: %v", err)
		}
		pks = append(pks, pk)
	}
	return pks, nil
}

// []partyIDs -> map[partyID.id][partyID]
func GetPatyIDInfoMap(partiesID []*btss.PartyID) (map[string]*btss.PartyID, error) {
	partyIDMap := make(map[string]*btss.PartyID)
	for _, partyID := range partiesID {
		partyIDMap[partyID.Id] = partyID
	}
	return partyIDMap, nil
}

//// todo：(fbz) delete
//// []partyIDs -> []pids
//func GetPIDsFromPartyIDs(partyIDs []*btss.PartyID) ([]peer.ID, error) {
//	pids := []peer.ID{}
//	for _, el := range partyIDs {
//		pid, err := GetPIDFromPartyID(el)
//		if err != nil {
//			return nil, fmt.Errorf("get pid from party id error: %v", err)
//		}
//		pids = append(pids, pid)
//	}
//	return pids, nil
//}
//
//// todo: (fbz) delete
//// []partyIDs.id -> []pids
//func AccPIDsFromPartyIDs(partyIDs []string, partyIDMap map[string]*btss.PartyID) ([]string, error) {
//	pids := make([]string, 0)
//	for _, partyID := range partyIDs {
//		party, ok := partyIDMap[partyID]
//		if !ok {
//			return nil, fmt.Errorf("cannot find the blame party")
//		}
//		pid, err := GetPIDFromPartyID(party)
//		if err != nil {
//			return nil, err
//		}
//		pids = append(pids, pid.String())
//	}
//	return pids, nil
//}
//
//// todo: (fbz) delete
//// pid -> p2pPeer id
//func GetIDFromPID(pid peer.ID, peerInfo map[string]*peer.AddrInfo) (uint64, error) {
//	for id, info := range peerInfo {
//		if info.ID == peer.ID(pid.String()) {
//			ret, err := strconv.ParseUint(id, 10, 32)
//			if err != nil {
//				return 0, fmt.Errorf("GetIDFromPID parse int error: %v", err)
//			}
//			return ret, nil
//		}
//	}
//	return 0, fmt.Errorf("GetIDFromPID error")
//}
