package appchain_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
	g "github.com/meshplus/bitxhub-core/governance"
	"github.com/sirupsen/logrus"
)

const (
	PREFIX = "appchain-"

	RelaychainType = "relaychain"
	AppchainType   = "appchain"
	FabricType     = "fabric"
)

type AppchainManager struct {
	g.Persister
}

type Appchain struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Validators    string             `json:"validators"`
	ConsensusType string             `json:"consensus_type"`
	Status        g.GovernanceStatus `json:"status"`
	ChainType     string             `json:"chain_type"`
	Desc          string             `json:"desc"`
	Version       string             `json:"version"`
	PublicKey     string             `json:"public_key"`
	OwnerDID      string             `json:"owner_did"`
	DidDocAddr    string             `json:"did_doc_addr"`
	DidDocHash    string             `json:"did_doc_hash"`
	FSM           *fsm.FSM           `json:"fsm"`
}

type auditRecord struct {
	Appchain   *Appchain `json:"appchain"`
	IsApproved bool      `json:"is_approved"`
	Desc       string    `json:"desc"`
}

func New(persister g.Persister) AppchainMgr {
	return &AppchainManager{persister}
}

func setFSM(chain *Appchain) {
	chain.FSM = fsm.NewFSM(
		string(chain.Status),
		fsm.Events{
			// register 3
			{Name: string(g.EventRegister), Src: []string{string(g.GovernanceUnavailable)}, Dst: string(g.GovernanceRegisting)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceRegisting)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceRegisting)}, Dst: string(g.GovernanceUnavailable)},

			// update 1
			{Name: string(g.EventUpdate), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceFrozen), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceUpdating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceUpdating)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceUpdating)}, Dst: string(g.GovernanceAvailable)},

			// freeze 2
			{Name: string(g.EventFreeze), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceUpdating), string(g.GovernanceActivating), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceFreezing)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceFreezing)}, Dst: string(g.GovernanceFrozen)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceFreezing)}, Dst: string(g.GovernanceAvailable)},

			// active 1
			{Name: string(g.EventActivate), Src: []string{string(g.GovernanceFrozen), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceActivating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceActivating)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceActivating)}, Dst: string(g.GovernanceFrozen)},

			// logout 3
			{Name: string(g.EventLogout), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceUpdating), string(g.GovernanceFreezing), string(g.GovernanceFrozen), string(g.GovernanceActivating)}, Dst: string(g.GovernanceLogouting)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceLogouting)}, Dst: string(g.GovernanceForbidden)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceLogouting)}, Dst: string(g.GovernanceAvailable)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { chain.Status = g.GovernanceStatus(chain.FSM.Current()) },
		},
	)
}

// Register registers appchain info return appchain id and error
func (am *AppchainManager) Register(info []byte) (bool, []byte) {
	chain := &Appchain{}
	if err := json.Unmarshal(info, chain); err != nil {
		return false, []byte(err.Error())
	}

	res := &g.RegisterResult{}
	res.ID = chain.ID

	tmpChain := &Appchain{}
	ok := am.GetObject(am.appchainKey(chain.ID), tmpChain)

	if ok && tmpChain.Status != g.GovernanceUnavailable {
		am.Persister.Logger().WithFields(logrus.Fields{
			"id": chain.ID,
		}).Info("Appchain has registered")
		res.IsRegistered = true
	} else {
		am.SetObject(am.appchainKey(chain.ID), chain)
		am.Logger().WithFields(logrus.Fields{
			"id": chain.ID,
		}).Info("Appchain is registering")
		res.IsRegistered = false
	}

	resData, err := json.Marshal(res)
	if err != nil {
		return false, []byte(err.Error())
	}

	return true, resData
}

func (am *AppchainManager) Update(info []byte) (bool, []byte) {
	chain := &Appchain{}
	if err := json.Unmarshal(info, chain); err != nil {
		return false, []byte(err.Error())
	}

	ok := am.Has(am.appchainKey(chain.ID))
	if !ok {
		return false, []byte("this appchain does not exist")
	}

	am.SetObject(am.appchainKey(chain.ID), chain)

	return true, nil
}

func (am *AppchainManager) ChangeStatus(id, trigger string, _ []byte) (bool, []byte) {
	ok, data := am.Get(am.appchainKey(id))
	if !ok {
		return false, []byte(fmt.Errorf("this appchain does not exist").Error())
	}

	chain := &Appchain{}
	if err := json.Unmarshal(data, chain); err != nil {
		return false, []byte(fmt.Sprintf("unmarshal json error: %v", err))
	}

	setFSM(chain)
	err := chain.FSM.Event(trigger)
	if err != nil {
		return false, []byte(fmt.Sprintf("change status error: %v", err))
	}

	am.SetObject(am.appchainKey(id), *chain)
	return true, nil
}

func (am *AppchainManager) DeleteAppchain(id string) (bool, []byte) {
	am.Delete(am.appchainKey(id))
	am.Logger().Infof("delete appchain:%s", id)
	return true, nil
}

// Audit bitxhub manager audit appchain register info
func (am *AppchainManager) Audit(proposer string, isApproved int32, desc string) (bool, []byte) {
	chain := &Appchain{}
	ok := am.GetObject(am.appchainKey(proposer), chain)
	if !ok {
		return false, []byte("this appchain does not exist")
	}

	chain.Status = g.GovernanceAvailable

	record := &auditRecord{
		Appchain:   chain,
		IsApproved: isApproved == g.APPROVED,
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

// CountAvailableAppchains counts all available appchains
func (am *AppchainManager) CountAvailable(_ []byte) (bool, []byte) {
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
		if a.Status == g.GovernanceAvailable {
			count++
		}
	}
	return true, []byte(strconv.Itoa(count))
}

// CountAppchains counts all appchains including approved, rejected or registered
func (am *AppchainManager) CountAll(_ []byte) (bool, []byte) {
	ok, value := am.Query(PREFIX)
	if !ok {
		return true, []byte("0")
	}
	return true, []byte(strconv.Itoa(len(value)))
}

// Appchains returns all appchains
func (am *AppchainManager) All(_ []byte) (bool, []byte) {
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

func (am *AppchainManager) QueryById(id string, _ []byte) (bool, []byte) {
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
