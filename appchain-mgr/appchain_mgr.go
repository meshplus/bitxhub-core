package appchain_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
	"github.com/sirupsen/logrus"
)

type AppchainStatus string

const (
	PREFIX = "appchain-"

	REGISTERED = 0
	APPROVED   = 1

	RelaychainType = "relaychain"
	AppchainType   = "appchain"
	FabricType     = "fabric"

	AppchainRegisting   AppchainStatus = "registing"
	AppchainAvailable   AppchainStatus = "available"
	AppchainUnavailable AppchainStatus = "unavailable"
	AppchainUpdating    AppchainStatus = "updating"
	AppchainFreezing    AppchainStatus = "freezing"
	AppchainActivating  AppchainStatus = "activating"
	AppchainFrozen      AppchainStatus = "frozen"
	AppchainLogouting   AppchainStatus = "logouting"

	EventRegister = "register"
	EventUpdate   = "update"
	EventFreeze   = "freeze"
	EventActivate = "activate"
	EventLogout   = "logout"
	EventApprove  = "approve"
	EventReject   = "reject"
	EventClose    = "close"
)

type AppchainManager struct {
	Persister
}

type Appchain struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Validators    string         `json:"validators"`
	ConsensusType int32          `json:"consensus_type"`
	Status        AppchainStatus `json:"status"`
	ChainType     string         `json:"chain_type"`
	Desc          string         `json:"desc"`
	Version       string         `json:"version"`
	PublicKey     string         `json:"public_key"`
	FSM           *fsm.FSM       `json:"fsm"`
}

type auditRecord struct {
	Appchain   *Appchain `json:"appchain"`
	IsApproved bool      `json:"is_approved"`
	Desc       string    `json:"desc"`
}

func New(persister Persister) AppchainMgr {
	return &AppchainManager{persister}
}

func SetFSM(chain *Appchain) {
	chain.FSM = fsm.NewFSM(
		string(chain.Status),
		fsm.Events{
			{Name: EventUpdate, Src: []string{string(AppchainAvailable)}, Dst: string(AppchainUpdating)},
			{Name: EventFreeze, Src: []string{string(AppchainAvailable)}, Dst: string(AppchainFreezing)},
			{Name: EventActivate, Src: []string{string(AppchainFrozen)}, Dst: string(AppchainActivating)},
			{Name: EventLogout, Src: []string{string(AppchainAvailable)}, Dst: string(AppchainLogouting)},
			{Name: EventApprove, Src: []string{string(AppchainRegisting), string(AppchainUpdating), string(AppchainActivating)}, Dst: string(AppchainAvailable)},
			{Name: EventApprove, Src: []string{string(AppchainFreezing)}, Dst: string(AppchainFrozen)},
			{Name: EventApprove, Src: []string{string(AppchainLogouting)}, Dst: string(AppchainUnavailable)},
			{Name: EventReject, Src: []string{string(AppchainRegisting)}, Dst: string(AppchainUnavailable)},
			{Name: EventReject, Src: []string{string(AppchainUpdating)}, Dst: string(AppchainAvailable)},
			{Name: EventReject, Src: []string{string(AppchainFreezing)}, Dst: string(AppchainAvailable)},
			{Name: EventReject, Src: []string{string(AppchainActivating)}, Dst: string(AppchainFrozen)},
			{Name: EventReject, Src: []string{string(AppchainLogouting)}, Dst: string(AppchainAvailable)},
			{Name: EventClose, Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { chain.Status = AppchainStatus(chain.FSM.Current()) },
		},
	)
}

// Register appchain manager registers appchain info caller is the appchain
// manager address return appchain id and error
func (am *AppchainManager) Register(id, validators string, consensusType int32, chainType, name, desc, version, pubkey string) (bool, []byte) {
	chain := &Appchain{
		ID:            id,
		Name:          name,
		Validators:    validators,
		ConsensusType: consensusType,
		Status:        AppchainRegisting,
		ChainType:     chainType,
		Desc:          desc,
		Version:       version,
		PublicKey:     pubkey,
	}
	isRegister := false

	ok := am.Has(am.appchainKey(id))
	if ok {
		am.Persister.Logger().WithFields(logrus.Fields{
			"id": id,
		}).Debug("Appchain has registered")
		am.GetObject(am.appchainKey(id), chain)
		isRegister = true
	} else {
		am.SetObject(am.appchainKey(id), chain)
		am.Logger().WithFields(logrus.Fields{
			"id": id,
		}).Info("Appchain is registering")
	}

	return isRegister, []byte(chain.ID)
}

func (am *AppchainManager) UpdateAppchain(id, validators string, consensusType int32, chainType, name, desc, version, pubkey string) (bool, []byte) {
	ok := am.Has(am.appchainKey(id))
	if !ok {
		return false, []byte("register appchain firstly")
	}

	chain := &Appchain{}
	am.GetObject(am.appchainKey(id), chain)

	if chain.Status != AppchainAvailable {
		return false, []byte("this appchain is " + chain.Status + ", can not be updated")
	}

	chain = &Appchain{
		ID:            id,
		Name:          name,
		Validators:    validators,
		ConsensusType: consensusType,
		Status:        AppchainAvailable,
		ChainType:     chainType,
		Desc:          desc,
		Version:       version,
		PublicKey:     pubkey,
	}

	am.SetObject(am.appchainKey(id), chain)

	return true, nil
}

// Audit bitxhub manager audit appchain register info
func (am *AppchainManager) Audit(proposer string, isApproved int32, desc string) (bool, []byte) {
	chain := &Appchain{}
	ok := am.GetObject(am.appchainKey(proposer), chain)
	if !ok {
		return false, []byte("this appchain does not exist")
	}

	chain.Status = AppchainAvailable

	record := &auditRecord{
		Appchain:   chain,
		IsApproved: isApproved == APPROVED,
		Desc:       desc,
	}

	var records []*auditRecord
	am.GetObject(am.auditRecordKey(proposer), &records)
	records = append(records, record)

	am.SetObject(am.auditRecordKey(proposer), records)
	am.SetObject(am.appchainKey(proposer), chain)

	return true, []byte(fmt.Sprintf("audit %s successfully", proposer))
}

func (am *AppchainManager) FetchAuditRecords(id string) (bool, []byte) {
	var records []*auditRecord
	am.GetObject(am.auditRecordKey(id), &records)

	body, err := json.Marshal(records)
	if err != nil {
		return false, []byte(err.Error())
	}

	return true, body
}

func (am *AppchainManager) ChangeStatus(id, trigger string) (bool, []byte) {
	ok, data := am.Get(am.appchainKey(id))
	if !ok {
		return false, []byte(fmt.Errorf("this appchain does not exist").Error())
	}

	chain := &Appchain{}
	if err := json.Unmarshal(data, chain); err != nil {
		return false, []byte(fmt.Sprintf("unmarshal json error: %v", err))
	}

	SetFSM(chain)
	err := chain.FSM.Event(trigger)
	if err != nil {
		return false, []byte(fmt.Sprintf("change status error: %v, %s, %s", err, chain.FSM.Current(), trigger))
	}

	am.SetObject(am.appchainKey(id), *chain)
	return true, nil
}

// CountAvailableAppchains counts all available appchains
func (am *AppchainManager) CountAvailableAppchains() (bool, []byte) {
	ok, value := am.Query(PREFIX)
	if !ok {
		return true, []byte("0")
	}

	count := 0
	for _, v := range value {
		a := &Appchain{}
		if err := json.Unmarshal(v, a); err != nil {
			return false, []byte(fmt.Sprintf("unmarshal json error: %v", err))
		}
		if a.Status == AppchainAvailable {
			count++
		}
	}
	return true, []byte(strconv.Itoa(count))
}

// CountAppchains counts all appchains including approved, rejected or registered
func (am *AppchainManager) CountAppchains() (bool, []byte) {
	ok, value := am.Query(PREFIX)
	if !ok {
		return true, []byte("0")
	}
	return true, []byte(strconv.Itoa(len(value)))
}

// Appchains returns all appchains
func (am *AppchainManager) Appchains() (bool, []byte) {
	ok, value := am.Query(PREFIX)
	if !ok {
		return true, nil
	}

	ret := make([]*Appchain, 0)
	for _, data := range value {
		chain := &Appchain{}
		if err := json.Unmarshal(data, chain); err != nil {
			return false, []byte(err.Error())
		}
		ret = append(ret, chain)
	}

	data, err := json.Marshal(ret)
	if err != nil {
		return false, []byte(err.Error())
	}
	return true, data
}

func (am *AppchainManager) DeleteAppchain(id string) (bool, []byte) {
	am.Delete(am.appchainKey(id))
	am.Logger().Infof("delete appchain:%s", id)
	return true, nil
}

func (am *AppchainManager) Appchain() (bool, []byte) {
	ok, data := am.Get(am.appchainKey(am.Caller()))
	if !ok {
		return false, []byte(fmt.Errorf("this appchain does not exist").Error())
	}
	return true, data
}

func (am *AppchainManager) GetAppchain(id string) (bool, []byte) {
	ok, data := am.Get(am.appchainKey(id))
	if !ok {
		return false, []byte(fmt.Errorf("this appchain does not exist").Error())
	}

	return true, data
}

// GetPubKeyByChainID can get aim chain's public key using aim chain ID
func (am *AppchainManager) GetPubKeyByChainID(id string) (bool, []byte) {
	ok := am.Has(am.appchainKey(id))
	if !ok {
		return false, []byte("chain is not existed")
	} else {
		chain := &Appchain{}
		am.GetObject(am.appchainKey(id), chain)
		return true, []byte(chain.PublicKey)
	}
}

func (am *AppchainManager) appchainKey(id string) string {
	return PREFIX + id
}

func (am *AppchainManager) auditRecordKey(id string) string {
	return "audit-record-" + id
}

func (am *AppchainManager) indexMapKey(id string) string {
	return fmt.Sprintf("index-tx-%s", id)
}
