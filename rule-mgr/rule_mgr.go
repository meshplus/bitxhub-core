package rule_mgr

import (
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
	"github.com/meshplus/bitxhub-core/governance"
	"github.com/meshplus/bitxhub-core/validator"
)

const (
	RULEPREFIX = "rule-"
)

type RuleManager struct {
	governance.Persister
}

type Rule struct {
	Address string                      `json:"address"`
	RuleUrl string                      `json:"rule_url"`
	ChainID string                      `json:"chain_id"`
	Master  bool                        `json:"master"`
	Status  governance.GovernanceStatus `json:"status"`
	FSM     *fsm.FSM                    `json:"fsm"`
}

var ruleStateMap = map[governance.EventType][]governance.GovernanceStatus{
	governance.EventBind:   {governance.GovernanceBindable},
	governance.EventLogout: {governance.GovernanceBindable},
	governance.EventUpdate: {governance.GovernanceBindable},
	governance.EventUnbind: {governance.GovernanceAvailable},
}

var ruleAvailableMap = map[governance.GovernanceStatus]struct{}{
	governance.GovernanceAvailable: {},
	governance.GovernanceUnbinding: {},
}

var defaultRuleMap = map[string]struct{}{
	validator.FabricRuleAddr:    {},
	validator.SimFabricRuleAddr: {},
	validator.HappyRuleAddr:     {},
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

func IsDefault(addr string) bool {
	if _, ok := defaultRuleMap[addr]; ok {
		return true
	} else {
		return false
	}
}

func (rule *Rule) setFSM(lastStatus governance.GovernanceStatus) {
	rule.FSM = fsm.NewFSM(
		string(rule.Status),
		fsm.Events{
			// bind
			{Name: string(governance.EventBind), Src: []string{string(governance.GovernanceBindable), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceAvailable)},

			// update(bind) 1
			{Name: string(governance.EventUpdate), Src: []string{string(governance.GovernanceBindable), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceBinding)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceBinding)}, Dst: string(governance.GovernanceAvailable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceBinding)}, Dst: string(lastStatus)},

			// unbind 1
			{Name: string(governance.EventUnbind), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceUnbinding)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceUnbinding)}, Dst: string(governance.GovernanceBindable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceUnbinding)}, Dst: string(lastStatus)},

			// logout 3
			{Name: string(governance.EventLogout), Src: []string{string(governance.GovernanceBindable)}, Dst: string(governance.GovernanceForbidden)},
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
			},
		},
	)
}

// Register record rule
func (rm *RuleManager) Register(chainID, ruleAddress, ruleUrl string) (bool, []byte) {
	rules := make([]*Rule, 0)
	_ = rm.GetObject(RuleKey(chainID), &rules)

	rules = append(rules, &Rule{ruleAddress, ruleUrl, chainID, false, governance.GovernanceBindable, nil})
	rm.SetObject(RuleKey(chainID), rules)

	return true, nil
}

// GovernancePre checks if the rule address can do event with appchain id and record rule. (only check, not modify infomation)
func (rm *RuleManager) GovernancePre(ruleAddress string, event governance.EventType, chainID []byte) (interface{}, error) {
	rules := make([]*Rule, 0)
	if ok := rm.GetObject(RuleKey(string(chainID)), &rules); !ok {
		if event == governance.EventRegister {
			return nil, nil
		} else {
			return nil, fmt.Errorf("this appchain's rules do not exist")
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
			return nil, fmt.Errorf("the rule (%s) does not exist ", ruleAddress)
		}
	}

	if !statusOk {
		return nil, fmt.Errorf("the rule (%s) can not be %s", string(rule.Status), string(event))
	}

	switch event {
	case governance.EventUpdate:
		for _, r := range rules {
			if true == r.Master && governance.GovernanceAvailable != r.Status {
				return nil, fmt.Errorf("The master rule is changing(%s) now. Please wait until the proposal close before updating new rule", r.Status)
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
		return false, []byte("the rule does not exist ")
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
		if true == r.Master {
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
		if true == r.Master {
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

func RuleKey(chainID string) string {
	return fmt.Sprintf("%s-%s", RULEPREFIX, chainID)
}
