package storage

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/btcsuite/btcd/btcec"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	ecdsa2 "github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
	. "gopkg.in/check.v1"
)

var (
	pids = []string{
		"QmQW3bFn8XX1t4W14Pmn37bPJUpUVBrBjnPuBZwPog3Qdy",
		"QmQUcDYCtqbpn5Nhaw4FAGxQaSSNvdWfAFcpQT9SPiezbS",
		"QmbmD1kzdsxRiawxu7bRrteDgW1ituXupR8GH6E2EUAHY4",
	}
)

type FileStateMgrTestSuite struct{}

var _ = Suite(&FileStateMgrTestSuite{})

func TestPackage(t *testing.T) { TestingT(t) }

func (s *FileStateMgrTestSuite) SetUpTest(c *C) {
}

func (s *FileStateMgrTestSuite) TestNewFileStateMgr(c *C) {
	folder := os.TempDir()
	f := filepath.Join(folder, "test", "test1", "test2")
	defer func() {
		err := os.RemoveAll(f)
		c.Assert(err, IsNil)
	}()
	fsm, err := NewFileStateMgr(f)
	c.Assert(err, IsNil)
	c.Assert(fsm, NotNil)
	_, err = os.Stat(f)
	c.Assert(err, IsNil)
}

func (s *FileStateMgrTestSuite) TestSaveLocalState(c *C) {
	priv1, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	c.Assert(err, IsNil)
	pubkey1, err := ecdsa2.NewPublicKey(priv1.PublicKey)
	c.Assert(err, IsNil)
	pubkeyData1, err := pubkey1.Bytes()
	c.Assert(err, IsNil)

	priv2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	c.Assert(err, IsNil)
	pubkey2, err := ecdsa2.NewPublicKey(priv2.PublicKey)
	c.Assert(err, IsNil)
	pubkeyData2, err := pubkey2.Bytes()
	c.Assert(err, IsNil)

	pubAddr, err := pubkey1.Address()
	c.Assert(err, IsNil)

	testPubKeys := []crypto.PubKey{}
	testPubKeysDataMap := map[string][]byte{}
	testPeers := []peer.ID{}
	for i := 0; i < 4; i++ {
		_, public, err := crypto.GenerateKeyPair(crypto.ECDSA, 1024)
		c.Assert(err, IsNil)
		testPubKeys = append(testPubKeys, public)
		pkData, err := public.Raw()
		c.Assert(err, IsNil)
		testPubKeysDataMap[strconv.Itoa(i+1)] = pkData

		pid, err := conversion.GetPIDFromPubKey(public)
		c.Assert(err, IsNil)
		testPeers = append(testPeers, pid)
	}

	stateItem := &KeygenLocalState{
		PubKeyData:        pubkeyData2,
		PubKeyAddr:        pubAddr.String(),
		LocalData:         keygen.NewLocalPartySaveData(5),
		ParticipantPksMap: testPubKeysDataMap,
		LocalPartyPk:      testPubKeysDataMap["1"],
	}
	folder := os.TempDir()
	f := filepath.Join(folder, "test", "test1", "test2")
	defer func() {
		err := os.RemoveAll(f)
		c.Assert(err, IsNil)
	}()
	fsm, err := NewFileStateMgr(f)
	c.Assert(err, IsNil)
	c.Assert(fsm, NotNil)
	// not in Curve
	c.Assert(fsm.SaveLocalState(stateItem), NotNil)

	stateItem.PubKeyData = pubkeyData1
	c.Assert(fsm.SaveLocalState(stateItem), IsNil)
	filePathName := filepath.Join(f, "localstate-"+stateItem.PubKeyAddr+".json")
	_, err = os.Stat(filePathName)
	c.Assert(err, IsNil)
	item, err := fsm.GetLocalState(pubAddr.String())
	c.Assert(err, IsNil)
	c.Assert(reflect.DeepEqual(stateItem, item), Equals, true)
}
