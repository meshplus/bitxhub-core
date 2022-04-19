package blame

import "bytes"

type Node struct {
	PartyID        string `json:"party_id"`
	BlameData      []byte `json:"blame_data"`
	BlameSignature []byte `json:"blame_signature"`
}

func NewNode(partyID string, blameData, blameSig []byte) Node {
	return Node{
		PartyID:        partyID,
		BlameData:      blameData,
		BlameSignature: blameSig,
	}
}

func (bn *Node) Equal(node Node) bool {
	if bn.PartyID == node.PartyID && bytes.Equal(bn.BlameSignature, node.BlameSignature) {
		return true
	}
	return false
}
