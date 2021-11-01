package rule_mgr

import (
	"github.com/meshplus/bitxhub-core/governance"
)

//go:generate mockgen -destination mock_ruleMgr/mock_ruleMgr.go -package mock_ruleMgr -source types.go
type RuleMgr interface {
	governance.Governance

	GetMaster(chainId string) (*Rule, error)

	HasMaster(chainId string) bool

	IsAvailable(chainId, ruleId string) bool

	AllRules() ([]*Rule, error)
}
