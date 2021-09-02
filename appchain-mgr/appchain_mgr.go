package appchain_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
	"github.com/meshplus/bitxhub-core/governance"
	g "github.com/meshplus/bitxhub-core/governance"
	"github.com/sirupsen/logrus"
)

const (
	PREFIX = "appchain"

	RelaychainType = "relaychain"
	AppchainType   = "appchain"
	FabricType     = "fabric"
)

type AppchainManager struct {
	g.Persister
}

type Appchain struct {
	ID        string `json:"id"`
	TrustRoot []byte `json:"trust_root"`
	Broker    string `json:"broker"`
	Desc      string `json:"desc"`
	Version   uint64 `json:"version"`

	Status g.GovernanceStatus `json:"status"`
	FSM    *fsm.FSM           `json:"fsm"`
}

type auditRecord struct {
	Appchain   *Appchain `json:"appchain"`
	IsApproved bool      `json:"is_approved"`
	Desc       string    `json:"desc"`
}

var appchainStateMap = map[g.EventType][]g.GovernanceStatus{
	g.EventRegister: {g.GovernanceUnavailable},
	g.EventUpdate:   {g.GovernanceAvailable, g.GovernanceFrozen},
	g.EventFreeze:   {g.GovernanceAvailable, g.GovernanceUpdating, g.GovernanceActivating},
	g.EventActivate: {g.GovernanceFrozen},
	g.EventLogout:   {g.GovernanceAvailable, g.GovernanceUpdating, g.GovernanceFreezing, g.GovernanceActivating, g.GovernanceFrozen},
	g.EventPause:    {g.GovernanceAvailable, g.GovernanceFrozen},
	g.EventUnpause:  {g.GovernanceFrozen},
}

var appchainAvailableMap = map[g.GovernanceStatus]struct{}{
	g.GovernanceAvailable: {},
	g.GovernanceFreezing:  {},
}

func New(persister g.Persister) AppchainMgr {
	return &AppchainManager{persister}
}

func (a *Appchain) IsAvailable() bool {
	if _, ok := appchainAvailableMap[a.Status]; ok {
		return true
	} else {
		return false
	}
}

func (chain *Appchain) setFSM(lastStatus g.GovernanceStatus) {
	chain.FSM = fsm.NewFSM(
		string(chain.Status),
		fsm.Events{
			// register 3
			{Name: string(g.EventRegister), Src: []string{string(g.GovernanceUnavailable)}, Dst: string(g.GovernanceRegisting)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceRegisting)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceRegisting)}, Dst: string(lastStatus)},

			// update 1
			{Name: string(g.EventUpdate), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceFrozen), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceUpdating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceUpdating)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceUpdating)}, Dst: string(g.GovernanceFrozen)},

			// freeze 2
			{Name: string(g.EventFreeze), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceUpdating), string(g.GovernanceActivating), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceFreezing)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceFreezing)}, Dst: string(g.GovernanceFrozen)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceFreezing)}, Dst: string(lastStatus)},

			// activate 1
			{Name: string(g.EventActivate), Src: []string{string(g.GovernanceFrozen), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceActivating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceActivating)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceActivating)}, Dst: string(lastStatus)},

			// logout 3
			{Name: string(g.EventLogout), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceUpdating), string(g.GovernanceFreezing), string(g.GovernanceFrozen), string(g.GovernanceActivating)}, Dst: string(g.GovernanceLogouting)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceLogouting)}, Dst: string(g.GovernanceForbidden)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceLogouting)}, Dst: string(lastStatus)},

			// pause
			{Name: string(g.EventPause), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceFrozen)}, Dst: string(g.GovernanceFrozen)},

			// unpause
			{Name: string(g.EventUnpause), Src: []string{string(g.GovernanceFrozen)}, Dst: string(lastStatus)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { chain.Status = g.GovernanceStatus(chain.FSM.Current()) },
		},
	)
}

// GovernancePre checks if the appchain can do the event. (only check, not modify infomation)
// return *appchain, extra info, error
func (am *AppchainManager) GovernancePre(chainId string, event g.EventType, _ []byte) (interface{}, error) {
	chain := &Appchain{}
	if ok := am.GetObject(am.appchainKey(chainId), chain); !ok {
		if event == governance.EventRegister {
			return nil, nil
		} else {
			return nil, fmt.Errorf("the appchain does not exist")
		}
	}

	for _, s := range appchainStateMap[event] {
		if chain.Status == s {
			return chain, nil
		}
	}

	return nil, fmt.Errorf("the appchain (%s) can not be %s", string(chain.Status), string(event))
}

// Register registers appchain info
func (am *AppchainManager) Register(chainInfo *Appchain) (bool, []byte) {

	am.SetObject(am.appchainKey(chainInfo.ID), chainInfo)
	am.Logger().WithFields(logrus.Fields{
		"id": chainInfo.ID,
	}).Info("Appchain is registering")

	return true, nil
}

func (am *AppchainManager) Update(updateInfo *Appchain) (bool, []byte) {
	var chain Appchain
	ok := am.GetObject(am.appchainKey(updateInfo.ID), &chain)
	if !ok {
		return false, []byte("this appchain does not exist")
	}

	chain.Desc = updateInfo.Desc
	chain.Version++
	am.SetObject(am.appchainKey(updateInfo.ID), chain)

	return true, nil
}

func (am *AppchainManager) ChangeStatus(id, trigger, lastStatus string, _ []byte) (bool, []byte) {
	ok, data := am.Get(am.appchainKey(id))
	if !ok {
		return false, []byte(fmt.Sprintf("this appchain does not exist"))
	}

	chain := &Appchain{}
	if err := json.Unmarshal(data, chain); err != nil {
		return false, []byte(fmt.Sprintf("unmarshal json error: %v", err))
	}

	chain.setFSM(g.GovernanceStatus(lastStatus))
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
		if a.IsAvailable() {
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
func (am *AppchainManager) All(_ []byte) (interface{}, error) {
	ret := make([]*Appchain, 0)
	ok, value := am.Query(PREFIX)
	if ok {
		for _, data := range value {
			chain := &Appchain{}
			if err := json.Unmarshal(data, chain); err != nil {
				return nil, err
			}
			ret = append(ret, chain)
		}
	}

	return ret, nil
}

func (am *AppchainManager) QueryById(id string, _ []byte) (interface{}, error) {
	var appchain Appchain
	ok := am.GetObject(am.appchainKey(id), &appchain)
	if !ok {
		return nil, fmt.Errorf("this appchain does not exist")
	}

	return &appchain, nil
}

func (am *AppchainManager) appchainKey(id string) string {
	return fmt.Sprintf("%s-%s", PREFIX, id)
}

func (am *AppchainManager) auditRecordKey(id string) string {
	return "audit-record-" + id
}

func (am *AppchainManager) indexMapKey(id string) string {
	return fmt.Sprintf("index-tx-%s", id)
}
