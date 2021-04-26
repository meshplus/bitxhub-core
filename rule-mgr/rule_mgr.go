package rule_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/meshplus/bitxhub-core/governance"

	"github.com/meshplus/bitxhub-core/validator"

	"github.com/looplab/fsm"
	g "github.com/meshplus/bitxhub-core/governance"
)

const (
	RULEPREFIX = "rule-"
)

type RuleManager struct {
	g.Persister
}

type Rule struct {
	Address string             `json:"address"`
	ChainId string             `json:"chain_id"`
	Status  g.GovernanceStatus `json:"status"`
	FSM     *fsm.FSM           `json:"fsm"`
}

func New(persister g.Persister) RuleMgr {
	return &RuleManager{persister}
}

func SetFSM(rule *Rule) {
	rule.FSM = fsm.NewFSM(
		string(rule.Status),
		fsm.Events{
			// bind 1
			{Name: string(g.EventBind), Src: []string{string(g.GovernanceBindable), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceBinding)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceBinding)}, Dst: string(g.GovernanceAvailable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceBinding)}, Dst: string(g.GovernanceBindable)},

			// unbind 1
			{Name: string(g.EventUnbind), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceUnbinding)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceUnbinding)}, Dst: string(g.GovernanceBindable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceUnbinding)}, Dst: string(g.GovernanceAvailable)},

			// freeze 2
			{Name: string(g.EventFreeze), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceBindable), string(g.GovernanceActivating), string(g.GovernanceBinding), string(g.GovernanceUnbinding), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceFreezing)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceFreezing)}, Dst: string(g.GovernanceFrozen)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceFreezing)}, Dst: string(g.GovernanceBindable)},

			// active 1
			{Name: string(g.EventActivate), Src: []string{string(g.GovernanceFrozen), string(g.GovernanceFreezing), string(g.GovernanceLogouting)}, Dst: string(g.GovernanceActivating)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceActivating)}, Dst: string(g.GovernanceBindable)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceActivating)}, Dst: string(g.GovernanceFrozen)},

			// logout 3
			{Name: string(g.EventLogout), Src: []string{string(g.GovernanceAvailable), string(g.GovernanceBindable), string(g.GovernanceFrozen), string(g.GovernanceFreezing), string(g.GovernanceActivating), string(g.GovernanceBinding), string(g.GovernanceUnbinding)}, Dst: string(g.GovernanceLogouting)},
			{Name: string(g.EventApprove), Src: []string{string(g.GovernanceLogouting)}, Dst: string(g.GovernanceForbidden)},
			{Name: string(g.EventReject), Src: []string{string(g.GovernanceLogouting)}, Dst: string(g.GovernanceBindable)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { rule.Status = g.GovernanceStatus(rule.FSM.Current()) },
		},
	)
}

// BindPre checks if the rule address can bind with appchain id and record rule
func (rm *RuleManager) BindPre(chainId, ruleAddress string) (bool, []byte) {
	flag := false
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(rm.ruleKey(chainId), rules); !ok {
		flag = true
	}

	for _, r := range rules {
		if ruleAddress == r.Address {
			if r.Status != g.GovernanceBindable {
				return false, []byte("The rule is in an unbindable state: " + r.Status)
			} else {
				flag = true
			}
		}
	}

	if flag {
		rules = append(rules, &Rule{ruleAddress, chainId, g.GovernanceBindable, nil})
		rm.SetObject(rm.ruleKey(chainId), rules)
	}

	return true, nil
}

func (rm *RuleManager) ChangeStatus(ruleAddress, trigger string, chainId []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(rm.ruleKey(string(chainId)), rules); !ok {
		return false, []byte("this appchain' rules does not exist")
	}

	flag := false
	for _, r := range rules {
		if ruleAddress == r.Address {
			flag = true
			SetFSM(r)
			err := r.FSM.Event(trigger)
			if err != nil {
				return false, []byte(fmt.Sprintf("change status error: %v", err))
			}
		}
	}

	if !flag {
		return false, []byte("this appchain' rules does not exist")
	}

	rm.SetObject(rm.ruleKey(string(chainId)), rules)

	return true, nil
}

// CountAvailable counts all rules of one appchain including available
func (rm *RuleManager) CountAvailable(chainId []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(rm.ruleKey(string(chainId)), rules); !ok {
		return false, []byte("this appchain' rules does not exist")
	}

	count := 0
	for _, r := range rules {
		if g.GovernanceAvailable == r.Status {
			count++
		}
	}

	return true, []byte(strconv.Itoa(count))
}

func (rm *RuleManager) CountAll(chainId []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(rm.ruleKey(string(chainId)), rules); !ok {
		return false, []byte("this appchain' rules does not exist")
	}

	return true, []byte(strconv.Itoa(len(rules)))
}

// Appchains returns all appchains
func (rm *RuleManager) All(chainId []byte) (bool, []byte) {
	ok, data := rm.Get(rm.ruleKey(string(chainId)))
	if !ok {
		return false, []byte("this appchain' rules does not exist")
	}

	return true, data
}

func (rm *RuleManager) QueryById(ruleAddress string, chainId []byte) (bool, []byte) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(rm.ruleKey(string(chainId)), rules); !ok {
		return false, []byte(fmt.Errorf("this appchain's rules do not exist").Error())
	}

	for _, r := range rules {
		if ruleAddress == r.Address {
			ruleData, err := json.Marshal(r)
			if err != nil {
				return false, []byte(fmt.Sprintf("marshal rule error: %v", err))
			}
			return true, ruleData
		}
	}

	return false, []byte(fmt.Errorf("this rule does not exist").Error())
}

func (rm *RuleManager) GetAvailableRuleAddress(chainId, chainType string) (bool, []byte) {
	rules := make([]*Rule, 0)
	_ = rm.GetObject(rm.ruleKey(chainId), rules)

	for _, r := range rules {
		if g.GovernanceAvailable == r.Status {
			return true, []byte(r.Address)
		}
	}

	if chainType == "fabric" {
		return true, []byte(validator.FabricRuleAddr)
	}

	return false, []byte("this appchain's available rule does not exist")
}

func (rm *RuleManager) IsAvailable(chainId, ruleAddress string) (bool, []byte) {
	is, data := rm.QueryById(ruleAddress, []byte(chainId))

	if !is {
		return false, []byte("get rule info error: " + string(data))
	}

	rule := &Rule{}
	if err := json.Unmarshal(data, rule); err != nil {
		return false, []byte("unmarshal rule error: " + err.Error())
	}

	if rule.Status != governance.GovernanceAvailable {
		return false, []byte("the rule status is " + string(rule.Status))
	}

	return true, nil
}

func (rm *RuleManager) ruleKey(id string) string {
	return RULEPREFIX + id
}
