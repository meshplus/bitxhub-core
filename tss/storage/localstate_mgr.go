package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/meshplus/bitxhub-core/tss/conversion"
)

const (
	LocalstateFileName = "localstate-%s.json"
	PoolPkAddrFileName = "poolPk.addr"
)

// KeygenLocalState is a structure used to represent the data we saved locally for different keygen
// A TSS public key address stores a corresponding KeygenLocalState structure
type KeygenLocalState struct {
	// TSS pool pubkey raw byte
	PubKeyData []byte `json:"pub_key_data"`
	// TSS pool pubkey addr
	PubKeyAddr string `json:"pub_key_addr"`
	// the library method returns information that needs to be stored at the end of the keyGen
	LocalData keygen.LocalPartySaveData `json:"local_data"`

	// partyID.id -> pubkey
	ParticipantPksMap map[string][]byte `json:"participant_pks_map"` // the paticipant of last key gen
	LocalPartyPk      []byte            `json:"local_party_pk"`
}

// LocalStateManager provide necessary methods to manage the local state, save it , and read it back
// LocalStateManager doesn't have any opinion in regards to where it should be persistent to
type LocalStateManager interface {
	SaveLocalState(state *KeygenLocalState) error
	GetLocalState(pubAddr string) (*KeygenLocalState, error)
}

var _ LocalStateManager = (*FileStateMgr)(nil)

// FileStateMgr save the local state to file
type FileStateMgr struct {
	folder    string
	writeLock *sync.RWMutex
}

// NewFileStateMgr create a new instance of the FileStateMgr which implements LocalStateManager
func NewFileStateMgr(folder string) (*FileStateMgr, error) {
	if len(folder) > 0 {
		_, err := os.Stat(folder)
		if err != nil && os.IsNotExist(err) {
			if err := os.MkdirAll(folder, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}
	return &FileStateMgr{
		folder:    folder,
		writeLock: &sync.RWMutex{},
	}, nil
}

// SaveLocalState save the local state to file
func (fsm *FileStateMgr) SaveLocalState(state *KeygenLocalState) error {
	// 1.Check if the public key is on the elliptic curve
	isOnCurve, err := conversion.CheckKeyOnCurve(state.PubKeyData)
	if err != nil {
		return err
	}
	if !isOnCurve {
		return fmt.Errorf("invalid pubkey")
	}

	// 2.get LocalState file name and content
	filePathName := fmt.Sprintf(LocalstateFileName, state.PubKeyAddr)
	if len(fsm.folder) > 0 {
		filePathName = filepath.Join(fsm.folder, filePathName)
	}
	buf, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("fail to marshal KeygenLocalState to json: %w", err)
	}

	// 3.write LocalState to file
	err = ioutil.WriteFile(filePathName, buf, 0o655)
	if err != nil {
		return fmt.Errorf("fail to write KeygenLocalState to file: %w", err)
	}

	// 4.get PoolPk file name
	filePathName = PoolPkAddrFileName
	if len(fsm.folder) > 0 {
		filePathName = filepath.Join(fsm.folder, filePathName)
	}

	// 5.write PoolPk to file
	err = ioutil.WriteFile(filePathName, []byte(state.PubKeyAddr), 0o655)
	if err != nil {
		return fmt.Errorf("fail to write pool pubkey to file: %w", err)
	}
	return nil
}

// GetLocalState read the local state from file system
func (fsm *FileStateMgr) GetLocalState(pubAddr string) (*KeygenLocalState, error) {
	// 1. check param
	if len(pubAddr) == 0 {
		return nil, fmt.Errorf("pub key is empty")
	}

	// 2. check file
	filePathName := fmt.Sprintf("localstate-%s.json", pubAddr)
	if len(fsm.folder) > 0 {
		filePathName = filepath.Join(fsm.folder, filePathName)
	}
	if _, err := os.Stat(filePathName); os.IsNotExist(err) {
		return nil, err
	}

	// 3. read file
	buf, err := ioutil.ReadFile(filePathName)
	if err != nil {
		return nil, fmt.Errorf("file to read from file(%s): %w", filePathName, err)
	}
	localState := &KeygenLocalState{}
	if err := json.Unmarshal(buf, localState); nil != err {
		return nil, fmt.Errorf("fail to unmarshal KeygenLocalState: %w", err)
	}
	return localState, nil
}
