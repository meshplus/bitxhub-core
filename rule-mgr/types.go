package rule_mgr

import (
	"github.com/meshplus/bitxhub-core/governance"
)

//go:generate mockgen -destination mock_ruleMgr/mock_ruleMgr.go -package mock_ruleMgr -source types.go
type RuleMgr interface {
	governance.Governance

	BindPre(chainId, ruleAddress string, force bool) (bool, []byte)

	GetAvailableRuleAddress(chainId string) (bool, []byte)

	GetMaster(chainId string) (bool, []byte)

	HasMaster(chainId string) bool

	IsAvailable(chainId, ruleId string) (bool, []byte)
}
