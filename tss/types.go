package tss

import (
	"crypto/ecdsa"
	"time"

	"github.com/meshplus/bitxhub-core/tss/keygen"
	"github.com/meshplus/bitxhub-core/tss/keysign"
	"github.com/meshplus/bitxhub-model/pb"
)

//go:generate mockgen -destination mock_tss/mock_tss.go -package mock_tss -source types.go
type Tss interface {
	Start(t uint64)

	Stop()

	Keygen(req keygen.Request) (*keygen.Response, error)

	Keysign(req keysign.Request) (*keysign.Response, error)

	PutTssMsg(msg *pb.Message)

	// GetTssPubkey returns tss pool pubkey addr and pubkey
	GetTssPubkey() (string, *ecdsa.PublicKey, error)

	// GetTssInfo returns tss pubkey and participants pubkey info
	GetTssInfo() (*pb.TssInfo, error)

	// LoadTssPubkey returns tss pool pubkey addr and pubkey from file
	LoadTssLoaclState() error

	//// LoadTssInfo returns tss pubkey and participants pubkey info from file
	//LoadTssInfo() (*pb.TssInfo, error)

	DeleteCulpritsFromLocalState(culprits []string) error
}

type TssConfig struct {
	EnableTSS       bool          `mapstructure:"enable_tss" json:"enable_tss"`
	PartyTimeout    time.Duration `mapstructure:"party_timeout" json:"party_timeout"`
	KeyGenTimeout   time.Duration `mapstructure:"key_gen_timeout" json:"key_gen_timeout"`
	KeySignTimeout  time.Duration `mapstructure:"key_sign_timeout" json:"key_sign_timeout"`
	PreParamTimeout time.Duration `mapstructure:"pre_param_timeout" json:"pre_param_timeout"`
	TssConfPath     string        `mapstructure:"tss_conf_path" json:"tss_conf_path"`
}
