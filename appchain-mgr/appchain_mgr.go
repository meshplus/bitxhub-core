package appchain_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	PREFIX = "appchain-"

	REGISTERED = 0
	APPROVED   = 1
)

type AppchainManager struct {
	Persister
}

type Appchain struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Validators    string `json:"validators"`
	ConsensusType int32  `json:"consensus_type"`
	// 0 => registered, 1 => approved, -1 => rejected
	Status    int32  `json:"status"`
	ChainType string `json:"chain_type"`
	Desc      string `json:"desc"`
	Version   string `json:"version"`
	PublicKey string `json:"public_key"`
}

type auditRecord struct {
	Appchain   *Appchain `json:"appchain"`
	IsApproved bool      `json:"is_approved"`
	Desc       string    `json:"desc"`
}

func New(persister Persister) AppchainMgr {
	return &AppchainManager{persister}
}

// Register appchain manager registers appchain info caller is the appchain
// manager address return appchain id and error
func (am *AppchainManager) Register(id, validators string, consensusType int32, chainType, name, desc, version, pubkey string) (bool, []byte) {
	chain := &Appchain{
		ID:            id,
		Name:          name,
		Validators:    validators,
		ConsensusType: consensusType,
		ChainType:     chainType,
		Desc:          desc,
		Version:       version,
		PublicKey:     pubkey,
	}

	ok := am.Has(am.appchainKey(id))
	if ok {
		am.Persister.Logger().WithFields(logrus.Fields{
			"id": am.Caller(),
		}).Debug("Appchain has registered")
		am.GetObject(am.appchainKey(am.Caller()), chain)
	} else {
		// logger.Info(am.Caller())
		am.SetObject(am.appchainKey(am.Caller()), chain)
		am.Logger().WithFields(logrus.Fields{
			"id": am.Caller(),
		}).Info("Appchain register successfully")
	}
	body, err := json.Marshal(chain)
	if err != nil {
		return false, []byte(err.Error())
	}

	return true, body
}

func (am *AppchainManager) UpdateAppchain(id, validators string, consensusType int32, chainType, name, desc, version, pubkey string) (bool, []byte) {
	ok := am.Has(am.appchainKey(id))
	if !ok {
		return false, []byte("register appchain firstly")
	}

	chain := &Appchain{}
	am.GetObject(am.appchainKey(id), chain)

	if chain.Status == REGISTERED {
		return false, []byte("this appchain is being audited")
	}

	chain = &Appchain{
		ID:            id,
		Name:          name,
		Validators:    validators,
		ConsensusType: consensusType,
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

	chain.Status = isApproved

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

// CountApprovedAppchains counts all approved appchains
func (am *AppchainManager) CountApprovedAppchains() (bool, []byte) {
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
		if a.Status == APPROVED {
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
