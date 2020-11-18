package agency

import (
	"fmt"

	"github.com/meshplus/bitxhub-kit/log"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

var logger = log.NewWithModule("agency")

type ContractInfo struct {
	Name        string
	Constructor ContractConstructor
}

type ApplyTxFunc func(int, *pb.Transaction, *TxOpt) *pb.Receipt

type RegisterContractFunc func() map[string]Contract

type TxsExecutorConstructor func(ApplyTxFunc, RegisterContractFunc, logrus.FieldLogger) TxsExecutor

type ContractConstructor func() Contract

var (
	TxsExecutorConstructorM = make(map[string]TxsExecutorConstructor)
	ContractConstructorM    = make(map[string]*ContractInfo)
)

func GetExecutorConstructor(typ string) (TxsExecutorConstructor, error) {
	con, ok := TxsExecutorConstructorM[typ]
	if !ok {
		return nil, fmt.Errorf("type %s is unsupported", typ)
	}
	return con, nil
}

func RegisterExecutorConstructor(typ string, f TxsExecutorConstructor) {
	TxsExecutorConstructorM[typ] = f
}

func RegisterContractConstructor(name string, addr *types.Address, f ContractConstructor) {
	ContractConstructorM[addr.String()] = &ContractInfo{
		Name:        name,
		Constructor: f,
	}
	logger.WithFields(logrus.Fields{
		"name": name,
		"addr": addr.String(),
	}).Info("contract registered")
}

func GetContractInfo(addr *types.Address) (*ContractInfo, error) {
	info, ok := ContractConstructorM[addr.String()]
	if !ok {
		return nil, fmt.Errorf("contract address %s is not registered", addr.String())
	}
	return info, nil
}

func GetRegisteredContractInfo() map[string]*ContractInfo {
	return ContractConstructorM
}
