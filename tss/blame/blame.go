package blame

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Blame is used to store the blame Peers and the fail reason
type Blame struct {
	FailReason string `json:"fail_reason"`
	IsUnicast  bool   `json:"is_unicast"`
	BlameNodes []Node `json:"blame_peers,omitempty"`
	blameLock  *sync.RWMutex
}

// NewBlame create a new instance of Blame
func NewBlame(reason string, blameNodes []Node) Blame {
	return Blame{
		FailReason: reason,
		BlameNodes: blameNodes,
		blameLock:  &sync.RWMutex{},
	}
}

// String implement fmt.Stringer
func (b Blame) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("reason: %s is_unicast: %s \n", b.FailReason, strconv.FormatBool(b.IsUnicast)))
	sb.WriteString(fmt.Sprintf("nodes:%+v\n", b.BlameNodes))
	return sb.String()
}

// SetBlame update the field values of Blame
func (b *Blame) SetBlame(reason string, nodes []Node, isUnicast bool) {
	b.blameLock.Lock()
	defer b.blameLock.Unlock()
	b.FailReason = reason
	b.IsUnicast = isUnicast
	b.BlameNodes = append(b.BlameNodes, nodes...)
}

func (b *Blame) AlreadyBlame() bool {
	b.blameLock.RLock()
	defer b.blameLock.RUnlock()
	return len(b.BlameNodes) > 0
}

// AddBlameNodes add Nodes to the blame list
func (b *Blame) AddBlameNodes(newBlamePeers ...Node) {
	b.blameLock.Lock()
	defer b.blameLock.Unlock()
	for _, nodeNew := range newBlamePeers {
		found := false
		for _, nodeBlame := range b.BlameNodes {
			if nodeNew.Equal(nodeBlame) {
				found = true
				break
			}
		}
		if !found {
			b.BlameNodes = append(b.BlameNodes, nodeNew)
		}
	}
}
