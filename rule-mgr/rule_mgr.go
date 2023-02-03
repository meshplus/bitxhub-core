package rule_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
	appchain_mgr "github.com/meshplus/bitxhub-core/appchain-mgr"
	"github.com/meshplus/bitxhub-core/boltvm"
	"github.com/meshplus/bitxhub-core/governance"
	"github.com/meshplus/bitxhub-core/validator"
	"github.com/sirupsen/logrus"
)

const (
	RulePrefix = "rule"
)

type RuleManager struct {
	governance.Persister
}

type Rule struct {
	Address    string                      `json:"address"`
	RuleUrl    string                      `json:"rule_url"`
	ChainID    string                      `json:"chain_id"`
	Master     bool                        `json:"master"`
	Default    bool                        `json:"builtIn"`
	CreateTime int64                       `json:"create_time"`
	Status     governance.GovernanceStatus `json:"status"`
	FSM        *fsm.FSM                    `json:"fsm"`
}

func (r *Rule) GetChainRuleID() string {
	return fmt.Sprintf("%s:%s", r.ChainID, r.Address)
}

var ruleStateMap = map[governance.EventType][]governance.GovernanceStatus{
	governance.EventBind:   {governance.GovernanceBindable},
	governance.EventLogout: {governance.GovernanceBindable},
	governance.EventUpdate: {governance.GovernanceBindable},
	governance.EventUnbind: {governance.GovernanceAvailable},
	governance.EventCLear:  {governance.GovernanceBindable, governance.GovernanceAvailable, governance.GovernanceBinding, governance.GovernanceUnbinding, governance.GovernanceForbidden},
}

var ruleAvailableMap = map[governance.GovernanceStatus]struct{}{
	governance.GovernanceAvailable: {},
	governance.GovernanceUnbinding: {},
}

var defaultRuleMap = map[string]map[string]struct{}{
	appchain_mgr.ChainTypeFabric1_4_3: {
		validator.FabricRuleAddr:    {},
		validator.SimFabricRuleAddr: {},
	},
	appchain_mgr.ChainTypeFabric1_4_4: {
		validator.FabricRuleAddr:    {},
		validator.SimFabricRuleAddr: {},
	},
	appchain_mgr.ChainTypeHyperchain1_8_3: {},
	appchain_mgr.ChainTypeHyperchain1_8_6: {},
	appchain_mgr.ChainTypeFlato1_0_0:      {},
	appchain_mgr.ChainTypeFlato1_0_3:      {},
	appchain_mgr.ChainTypeFlato1_0_6:      {},
	appchain_mgr.ChainTypeBCOS2_6_0:       {},
	appchain_mgr.ChainTypeCITA20_2_2:      {},
	appchain_mgr.ChainTypeETH:             {},
}

func New(persister governance.Persister) RuleMgr {
	return &RuleManager{persister}
}

func (r *Rule) IsAvailable() bool {
	if _, ok := ruleAvailableMap[r.Status]; ok {
		return true
	} else {
		return false
	}
}

func IsDefault(addr, chainType string) bool {
	if addr == validator.HappyRuleAddr {
		return true
	}
	if _, ok := defaultRuleMap[chainType][addr]; ok {
		return true
	} else {
		return false
	}
}

func (rule *Rule) setFSM(lastStatus governance.GovernanceStatus) {
	rule.FSM = fsm.NewFSM(
		string(rule.Status),
		fsm.Events{
			// bind(bind first)
			{Name: string(governance.EventBind), Src: []string{string(governance.GovernanceBindable)}, Dst: string(governance.GovernanceAvailable)},

			// update(bind)
			{Name: string(governance.EventUpdate), Src: []string{string(governance.GovernanceBindable)}, Dst: string(governance.GovernanceBinding)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceBinding)}, Dst: string(governance.GovernanceAvailable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceBinding)}, Dst: string(lastStatus)},

			// unbind
			{Name: string(governance.EventUnbind), Src: []string{string(governance.GovernanceAvailable)}, Dst: string(governance.GovernanceUnbinding)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceUnbinding)}, Dst: string(governance.GovernanceBindable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceUnbinding)}, Dst: string(lastStatus)},

			// logout
			{Name: string(governance.EventLogout), Src: []string{string(governance.GovernanceBindable)}, Dst: string(governance.GovernanceForbidden)},

			// clear
			{Name: string(governance.EventCLear), Src: []string{string(governance.GovernanceBindable), string(governance.GovernanceAvailable), string(governance.GovernanceBinding), string(governance.GovernanceUnbinding), string(governance.GovernanceForbidden)}, Dst: string(governance.GovernanceUnavailable)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) {
				rule.Status = governance.GovernanceStatus(rule.FSM.Current())

				// After the status change, if the rule is bound or the master authentication rule is updated successfully,
				// we need to enable the master identifier of the master rule
				if e.Event == string(governance.EventApprove) {
					if rule.Status == governance.GovernanceAvailable {
						rule.Master = true
					} else {
						rule.Master = false
					}
				}

				if e.Event == string(governance.EventBind) {
					rule.Master = true
				}

				if e.Event == string(governance.EventCLear) && rule.Default {
					rule.Status = governance.GovernanceBindable
				}
			},
		},
	)
}

// Register record rule
func (rm *RuleManager) Register(chainID, ruleAddress, ruleUrl string, createTime int64, isDefault bool) {
	rules := make([]*Rule, 0)
	_ = rm.GetObject(RuleKey(chainID), &rules)

	rules = append(rules, &Rule{
		Address:    ruleAddress,
		RuleUrl:    ruleUrl,
		ChainID:    chainID,
		Master:     false,
		Default:    isDefault,
		CreateTime: createTime,
		Status:     governance.GovernanceBindable,
	})
	rm.SetObject(RuleKey(chainID), rules)

	rm.Logger().WithFields(logrus.Fields{
		"chainID":  chainID,
		"ruleAddr": ruleAddress,
	}).Info("Rule is registering")
}

// GovernancePre checks if the rule address can do event with appchain id and record rule. (only check, not modify information)
func (rm *RuleManager) GovernancePre(ruleAddress string, event governance.EventType, chainID []byte) (interface{}, *boltvm.BxhError) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(string(chainID)), &rules); !ok {
		if event == governance.EventRegister {
			return nil, nil
		} else {
			return nil, boltvm.BError(boltvm.RuleNonexistentRuleCode, fmt.Sprintf(string(boltvm.RuleNonexistentRuleMsg), ruleAddress))
		}
	}

	// check current rule status
	var rule *Rule
	statusOk := false
	exist := false
	for _, r := range rules {
		if ruleAddress == r.Address {
			exist = true
			rule = r
			for _, s := range ruleStateMap[event] {
				if r.Status == s {
					statusOk = true
					break
				}
			}
			break
		}
	}

	if !exist {
		if event == governance.EventRegister {
			return nil, nil
		} else {
			return nil, boltvm.BError(boltvm.RuleNonexistentRuleCode, fmt.Sprintf(string(boltvm.RuleNonexistentRuleMsg), ruleAddress))
		}
	}

	if !statusOk {
		return nil, boltvm.BError(boltvm.RuleStatusErrorCode, fmt.Sprintf(string(boltvm.RuleStatusErrorMsg), ruleAddress, string(rule.Status), string(event)))
	}

	switch event {
	case governance.EventUpdate:
		for _, r := range rules {
			if r.Master && governance.GovernanceAvailable != r.Status {
				return nil, boltvm.BError(boltvm.RuleMasterRuleUpdatingCode, fmt.Sprintf(string(boltvm.RuleMasterRuleUpdatingMsg), r.Address))
			}
		}
		fallthrough
	default:
		return rule, nil
	}
}

func (rm *RuleManager) ChangeStatus(ruleAddress, trigger, lastStatus string, chainID []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(string(chainID)), &rules); !ok {
		return false, []byte("this appchain's rules do not exist")
	}

	flag := false
	for _, r := range rules {
		if ruleAddress == r.Address {
			flag = true
			r.setFSM(governance.GovernanceStatus(lastStatus))
			err := r.FSM.Event(trigger)
			if err != nil {
				return false, []byte(fmt.Sprintf("change status error: %v", err))
			}
		}
	}

	if !flag {
		return false, []byte(fmt.Sprintf("the rule does not exist: %s", ruleAddress))
	}

	rm.SetObject(RuleKey(string(chainID)), rules)

	return true, nil
}

// CountAvailable counts all rules of one appchain including available
func (rm *RuleManager) CountAvailable(chainID []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(string(chainID)), &rules); !ok {
		return true, []byte(strconv.Itoa(0))
	}

	count := 0
	for _, r := range rules {
		if r.IsAvailable() {
			count++
		}
	}

	return true, []byte(strconv.Itoa(count))
}

func (rm *RuleManager) CountAll(chainID []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(string(chainID)), &rules); !ok {
		return true, []byte(strconv.Itoa(0))
	}

	return true, []byte(strconv.Itoa(len(rules)))
}

// Appchains returns all appchains
func (rm *RuleManager) All(chainID []byte) (interface{}, error) {
	ret := make([]*Rule, 0)
	_ = rm.GetObject(RuleKey(string(chainID)), &ret)

	return ret, nil
}

func (rm *RuleManager) QueryById(ruleAddress string, chainID []byte) (interface{}, error) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(string(chainID)), &rules); !ok {
		return nil, fmt.Errorf("this appchain's rules do not exist")
	}

	for _, r := range rules {
		if ruleAddress == r.Address {
			return r, nil
		}
	}

	return nil, fmt.Errorf("this rule does not exist")
}

func (rm *RuleManager) GetMaster(chainID string) (*Rule, error) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(chainID), &rules); !ok {
		return nil, fmt.Errorf("the master rule is not exist")
	}

	for _, r := range rules {
		if r.Master {
			return r, nil
		}
	}

	return nil, fmt.Errorf("the master rule is not exist")
}

func (rm *RuleManager) HasMaster(chainID string) bool {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(chainID), &rules); !ok {
		return false
	}

	for _, r := range rules {
		if r.Master {
			return true
		}
	}

	return false
}

func (rm *RuleManager) IsAvailable(chainID, ruleAddress string) bool {
	rule, err := rm.QueryById(ruleAddress, []byte(chainID))
	if err != nil {
		return false
	}

	return rule.(*Rule).IsAvailable()
}

func (rm *RuleManager) AllRules() ([]*Rule, error) {
	ret := make([]*Rule, 0)
	ok, value := rm.Query(RulePrefix)
	if ok {
		for _, data := range value {
			chainRules := make([]*Rule, 0)
			if err := json.Unmarshal(data, &chainRules); err != nil {
				return nil, err
			}
			ret = append(ret, chainRules...)
		}
	}
	return ret, nil
}

func RuleKey(chainID string) string {
	return fmt.Sprintf("%s-%s", RulePrefix, chainID)
}
