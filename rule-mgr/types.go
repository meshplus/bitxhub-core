package rule_mgr

import (
	"github.com/meshplus/bitxhub-core/governance"
)

//go:generate mockgen -destination mock_ruleMgr/mock_ruleMgr.go -package mock_ruleMgr -source types.go
type RuleMgr interface {
	governance.Governance

	BindPre(chainId, ruleAddress string) (bool, []byte)

	GetAvailableRuleAddress(chainId, chainType string) (bool, []byte)

	IsAvailable(chainId, ruleId string) (bool, []byte)
}
