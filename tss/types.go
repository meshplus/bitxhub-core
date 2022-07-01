package tss

import (
	"time"

	"github.com/meshplus/bitxhub-core/tss/keygen"
	"github.com/meshplus/bitxhub-core/tss/keysign"
	"github.com/meshplus/bitxhub-model/pb"
)

//go:generate mockgen -destination mock_tss/mock_tss.go -package mock_tss -source types.go
type Tss interface {
	Keygen(req keygen.Request) (*keygen.Response, error)

	Keysign(req keysign.Request) (*keysign.Response, error)

	PutTssMsg(msg *pb.Message)
}

type TssConfig struct {
	EnableTSS bool `mapstructure:"enable_tss" json:"enable_tss"`
	//PartyTimeout    time.Duration `mapstructure:"party_timeout" json:"party_timeout"`
	KeyGenTimeout   time.Duration `mapstructure:"key_gen_timeout" json:"key_gen_timeout"`
	KeySignTimeout  time.Duration `mapstructure:"key_sign_timeout" json:"key_sign_timeout"`
	PreParamTimeout time.Duration `mapstructure:"pre_param_timeout" json:"pre_param_timeout"`
	TssConfPath     string        `mapstructure:"tss_conf_path" json:"tss_conf_path"`
}
