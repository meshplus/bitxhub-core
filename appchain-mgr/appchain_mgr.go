package appchain_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/looplab/fsm"
	"github.com/meshplus/bitxhub-core/boltvm"
	g "github.com/meshplus/bitxhub-core/governance"
	"github.com/sirupsen/logrus"
)

const (
	Prefix                = "appchain"
	ChainNumPrefix        = "chain-num"
	ChainOccupyIdPrefix   = "chain-occupy-id"
	ChainOccupyNamePrefix = "chain-occupy-name"
	NameChainPrefix       = "name-chain"
	AdminChainPrefix      = "admin-chain"
	ChainAdminsPrefix     = "chain-admins"

	RelaychainType = "relaychain"
	AppchainType   = "appchain"
	FabricType     = "fabric"

	ChainTypeFabric1_4_3     = "Fabric V1.4.3"
	ChainTypeFabric1_4_4     = "Fabric V1.4.4"
	ChainTypeHyperchain1_8_3 = "Hyperchain V1.8.3"
	ChainTypeHyperchain1_8_6 = "Hyperchain V1.8.6"
	ChainTypeFlato1_0_0      = "Flato V1.0.0"
	ChainTypeFlato1_0_3      = "Flato V1.0.3"
	ChainTypeFlato1_0_6      = "Flato V1.0.6"
	ChainTypeBCOS2_6_0       = "BCOS V2.6.0"
	ChainTypeCITA20_2_2      = "CITA V20.2.2"
	ChainTypeETH             = "ETH"
)

type AppchainManager struct {
	g.Persister
}

type Appchain struct {
	ID        string `json:"id"`
	ChainName string `json:"chain_name"`
	ChainType string `json:"chain_type"`
	TrustRoot []byte `json:"trust_root"`
	Broker    []byte `json:"broker"`
	Desc      string `json:"desc"`
	Version   uint64 `json:"version"`

	Status g.GovernanceStatus `json:"status"`
	FSM    *fsm.FSM           `json:"fsm"`
}

type FabricBroker struct {
	ChannelID     string `json:"channel_id"`
	ChaincodeID   string `json:"chaincode_id"`
	BrokerVersion string `json:"broker_version"`
}

type auditRecord struct {
	Appchain   *Appchain `json:"appchain"`
	IsApproved bool      `json:"is_approved"`
	Desc       string    `json:"desc"`
}

var appchainStateMap = map[g.EventType][]g.GovernanceStatus{
	g.EventUpdate:   {g.GovernanceAvailable, g.GovernanceFrozen},
	g.EventFreeze:   {g.GovernanceAvailable},
	g.EventActivate: {g.GovernanceFrozen},
	g.EventLogout:   {g.GovernanceAvailable, g.GovernanceUpdating, g.GovernanceFreezing, g.GovernanceActivating, g.GovernanceFrozen},

	g.EventPause:   {g.GovernanceAvailable, g.GovernanceFrozen},
	g.EventUnpause: {g.GovernanceFrozen},
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

func (a *Appchain) IsBitXHub() bool {
	return strings.EqualFold(a.ChainType, RelaychainType)
}

func (chain *Appchain) setFSM(lastStatus g.GovernanceStatus) {
	chain.FSM = fsm.NewFSM(
		string(chain.Status),
		fsm.Events{
			// update 1
			{Name: string(g.EventUpdate), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceFrozen), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceUpdating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceUpdating)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceUpdating)}, Dst: string(g.GovernanceFrozen)},

			// freeze 2
			{Name: string(g.EventFreeze), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceFreezing)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceFreezing)}, Dst: string(g.GovernanceFrozen)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceFreezing)}, Dst: string(lastStatus)},

			// activate 1
			{Name: string(g.EventActivate), Src: []string{string(g.GovernanceFrozen), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceActivating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceActivating)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceActivating)}, Dst: string(lastStatus)},

			// logout 3
			{Name: string(g.EventLogout), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceUpdating), string(g.GovernanceFreezing), string(g.GovernanceFrozen), string(g.GovernanceActivating)}, Dst: string(g.GovernanceLogouting)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceLogouting)}, Dst: string(g.GovernanceForbidden)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceLogouting)}, Dst: string(lastStatus)},

			// pause
			{Name: string(g.EventPause), Src: []string{string(g.GovernanceAvailable)}, Dst: string(g.GovernanceFrozen)},

			// unpause
			{Name: string(g.EventUnpause), Src: []string{string(g.GovernanceFrozen)}, Dst: string(g.GovernanceAvailable)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { chain.Status = g.GovernanceStatus(chain.FSM.Current()) },
		},
	)
}

// GovernancePre checks if the appchain can do the event. (only check, not modify infomation)
// return *appchain, extra info, error
func (am *AppchainManager) GovernancePre(chainId string, event g.EventType, _ []byte) (interface{}, *boltvm.BxhError) {
	chain := &Appchain{}
	if ok := am.GetObject(AppchainKey(chainId), chain); !ok {
		return nil, boltvm.BError(boltvm.AppchainNonexistentChainCode, fmt.Sprintf(string(boltvm.AppchainNonexistentChainMsg), chainId, ""))
	}

	for _, s := range appchainStateMap[event] {
		if chain.Status == s {
			return chain, nil
		}
	}

	return nil, boltvm.BError(boltvm.AppchainStatusErrorCode, fmt.Sprintf(string(boltvm.AppchainStatusErrorMsg), chainId, chain.Status, string(event)))
}

// Register registers appchain info, return chain id
func (am *AppchainManager) Register(chainInfo *Appchain) {
	am.SetObject(AppchainKey(chainInfo.ID), *chainInfo)
	am.SetObject(AppchainNameKey(chainInfo.ChainName), chainInfo.ID)

	am.Logger().WithFields(logrus.Fields{
		"id":   chainInfo.ID,
		"name": chainInfo.ChainName,
	}).Info("Appchain is registering")
}

func (am *AppchainManager) Update(updateInfo *Appchain) (bool, []byte) {
	chain := &Appchain{}
	ok := am.GetObject(AppchainKey(updateInfo.ID), chain)
	if !ok {
		return false, []byte(fmt.Sprintf("this appchain(%s) does not exist", updateInfo.ID))
	}

	oldName := chain.ChainName
	chain.ChainName = updateInfo.ChainName
	chain.Desc = updateInfo.Desc
	chain.TrustRoot = updateInfo.TrustRoot
	chain.Version++
	am.SetObject(AppchainKey(updateInfo.ID), chain)
	if oldName != chain.ChainName {
		am.Delete(AppchainNameKey(oldName))
		am.SetObject(AppchainNameKey(chain.ChainName), chain.ID)
	}

	am.Logger().WithFields(logrus.Fields{
		"id":   chain.ID,
		"name": chain.ChainName,
	}).Debug("Appchain is updating")

	return true, nil
}

func (am *AppchainManager) ChangeStatus(id, trigger, lastStatus string, _ []byte) (bool, []byte) {
	ok, data := am.Get(AppchainKey(id))
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

	am.SetObject(AppchainKey(id), *chain)
	return true, nil
}

func (am *AppchainManager) DeleteAppchain(id string) (bool, []byte) {
	am.Delete(AppchainKey(id))
	am.Logger().Infof("delete appchain:%s", id)
	return true, nil
}

// Audit bitxhub manager audit appchain register info
func (am *AppchainManager) Audit(proposer string, isApproved int32, desc string) (bool, []byte) {
	chain := &Appchain{}
	ok := am.GetObject(AppchainKey(proposer), chain)
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
	am.SetObject(AppchainKey(proposer), chain)

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

// CountAvailable counts all available appchains
func (am *AppchainManager) CountAvailable(_ []byte) (bool, []byte) {
	ok, value := am.Query(Prefix)
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

// CountAll counts all appchains including approved, rejected or registered
func (am *AppchainManager) CountAll(_ []byte) (bool, []byte) {
	ok, value := am.Query(Prefix)
	if !ok {
		return true, []byte("0")
	}
	return true, []byte(strconv.Itoa(len(value)))
}

// All returns all appchains
func (am *AppchainManager) All(_ []byte) (interface{}, error) {
	ret := make([]*Appchain, 0)
	ok, value := am.Query(Prefix)
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
	ok := am.GetObject(AppchainKey(id), &appchain)
	if !ok {
		return nil, fmt.Errorf("this appchain(%s) does not exist", id)
	}

	return &appchain, nil
}

func (am *AppchainManager) GetChainIdByName(name string) (string, error) {
	id := ""
	ok := am.GetObject(AppchainNameKey(name), &id)
	if !ok {
		return "", fmt.Errorf("not found chain %s", name)
	}
	return id, nil
}

func AppchainKey(id string) string {
	return fmt.Sprintf("%s-%s", Prefix, id)
}

func AppchainOccupyNameKey(name string) string {
	return fmt.Sprintf("%s-%s", ChainOccupyNamePrefix, name)
}

func AppchainNameKey(name string) string {
	return fmt.Sprintf("%s-%s", NameChainPrefix, name)
}

func AppchainAdminKey(addr string) string {
	return fmt.Sprintf("%s-%s", AdminChainPrefix, addr)
}

func AppAdminsChainKey(id string) string {
	return fmt.Sprintf("%s-%s", ChainAdminsPrefix, id)
}

func (am *AppchainManager) auditRecordKey(id string) string {
	return "audit-record-" + id
}

func (am *AppchainManager) indexMapKey(id string) string {
	return fmt.Sprintf("index-tx-%s", id)
}
