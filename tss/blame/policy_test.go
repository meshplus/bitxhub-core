package blame

import (
	"sort"
	"sync"
	"testing"

	bkg "github.com/binance-chain/tss-lib/ecdsa/keygen"
	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/meshplus/bitxhub-core/tss/message"
	. "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { TestingT(t) }

type policyTestSuite struct {
	blameMgr    *Manager
	testPubKeys []crypto.PubKey
}

var _ = Suite(&policyTestSuite{})

func (p *policyTestSuite) SetUpTest(c *C) {
	testPubKeys := []crypto.PubKey{}
	testPeers := []string{}
	for i := 0; i < 4; i++ {
		_, public, err := crypto.GenerateKeyPair(crypto.ECDSA, 1024)
		c.Assert(err, IsNil)
		testPubKeys = append(testPubKeys, public)

		pid, err := conversion.GetPIDFromPubKey(public)
		c.Assert(err, IsNil)
		testPeers = append(testPeers, pid.String())
	}
	p.blameMgr = NewBlameManager(nil)
	p1, err := peer.Decode(testPeers[0])
	c.Assert(err, IsNil)
	p2, err := peer.Decode(testPeers[1])
	c.Assert(err, IsNil)
	p3, err := peer.Decode(testPeers[2])
	c.Assert(err, IsNil)
	p4, err := peer.Decode(testPeers[3])
	c.Assert(err, IsNil)
	p.blameMgr.SetLastUnicastParty("1", "testType")
	p.blameMgr.SetLastUnicastParty("2", "testType")
	p.blameMgr.SetLastUnicastParty("3", "testType")

	peerMap := map[string]*peer.AddrInfo{}
	peerMap["1"] = &peer.AddrInfo{ID: p1}
	peerMap["2"] = &peer.AddrInfo{ID: p2}
	peerMap["3"] = &peer.AddrInfo{ID: p3}
	peerMap["4"] = &peer.AddrInfo{ID: p4}
	partiesID, localPartyID, err := conversion.GetParties(testPubKeys, testPubKeys[0], peerMap)
	c.Assert(err, IsNil)
	outCh := make(chan btss.Message, len(partiesID))
	endCh := make(chan bkg.LocalPartySaveData, len(partiesID))
	ctx := btss.NewPeerContext(partiesID)
	params := btss.NewParameters(ctx, localPartyID, len(partiesID), 3)
	keyGenParty := bkg.NewLocalParty(nil, params, outCh, endCh)
	testPartyMap := new(sync.Map)
	testPartyMap.Store("", keyGenParty)
	partyIDMap, err := conversion.GetPatyIDInfoMap(partiesID)
	c.Assert(err, IsNil)
	p.blameMgr.SetPartyInfo(&conversion.PartyInfo{
		PartyMap:   testPartyMap,
		PartyIDMap: partyIDMap,
	})
	p.testPubKeys = testPubKeys
}

func (p *policyTestSuite) TestGetUnicastBlame(c *C) {
	_, err := p.blameMgr.GetUnicastBlame("testTypeWrong")
	c.Assert(err, NotNil)
	_, err = p.blameMgr.GetUnicastBlame("testType")
	c.Assert(err, IsNil)
}

// 1 - localParty
// 2 - has broadcasted the msg
// 3 & 4 - have not broadcasted the msg - blame parties ID
func (p *policyTestSuite) TestGetBroadcastBlame(c *C) {
	pi := p.blameMgr.partyInfo

	r1 := btss.MessageRouting{
		From:                    pi.PartyIDMap["2"],
		To:                      nil,
		IsBroadcast:             false,
		IsToOldCommittee:        false,
		IsToOldAndNewCommittees: false,
	}
	msg := message.WireMessage{
		Routing:   &r1,
		RoundInfo: "key1",
		Message:   nil,
	}

	p.blameMgr.RoundMgr.Set("key1", &msg)
	blames, err := p.blameMgr.GetBroadcastBlame("key1")
	c.Assert(err, IsNil)
	var blamePartyIDs []string
	for _, el := range blames {
		blamePartyIDs = append(blamePartyIDs, el.PartyID)
	}
	sort.Strings(blamePartyIDs)
	var expected []string
	expected = append(expected, "3")
	expected = append(expected, "4")
	sort.Strings(expected)
	c.Assert(blamePartyIDs, DeepEquals, expected)
}

// 2 - provide wrong share
func (p *policyTestSuite) TestTssWrongShareBlame(c *C) {
	pi := p.blameMgr.partyInfo

	r1 := btss.MessageRouting{
		From:                    pi.PartyIDMap["2"],
		To:                      nil,
		IsBroadcast:             false,
		IsToOldCommittee:        false,
		IsToOldAndNewCommittees: false,
	}
	msg := message.WireMessage{
		Routing:   &r1,
		RoundInfo: "key2",
		Message:   nil,
	}
	target, err := p.blameMgr.TssWrongShareBlame(&msg)
	c.Assert(err, IsNil)
	c.Assert(target, Equals, "2")
}

func (p *policyTestSuite) TestTssMissingShareBlame(c *C) {
	blameMgr := p.blameMgr
	acceptedShares := blameMgr.acceptedShares
	// we only allow a message be updated only once.
	blameMgr.acceptShareLocker.Lock()
	acceptedShares[conversion.RoundInfo{0, "testRound", "123:0"}] = []string{"2", "3"}
	acceptedShares[conversion.RoundInfo{1, "testRound", "123:0"}] = []string{"2"}
	blameMgr.acceptShareLocker.Unlock()
	peers, _, err := blameMgr.TssMissingShareBlame(2)
	c.Assert(err, IsNil)
	// Check the number of rounds whose index is 0 first so no problem with index1 was found
	// Found that 4 did not accept direct abort in index 0 round
	c.Assert(peers[0].PartyID, Equals, "4")

	// we test if the missing share happens in round2
	blameMgr.acceptShareLocker.Lock()
	// Set all participants of a round with index 0 to accept so no problem with index0 was founc
	acceptedShares[conversion.RoundInfo{0, "testRound", "123:0"}] = []string{"2", "3", "4"}
	blameMgr.acceptShareLocker.Unlock()
	nodes, _, err := blameMgr.TssMissingShareBlame(2)
	c.Assert(err, IsNil)
	results := []string{nodes[0].PartyID, nodes[1].PartyID}
	// It was found that only 2 were accepted of the rounds with index1 , so 3 and 4 were blame parties
	excepted := []string{"3", "4"}
	sort.Strings(excepted)
	sort.Strings(results)
	c.Assert(results, DeepEquals, excepted)
}
