package agency

import (
	"fmt"

	"github.com/meshplus/bitxhub-kit/log"
	"github.com/meshplus/bitxhub-kit/storage"
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

type RegistryConstructor func(storage.Storage, logrus.FieldLogger) Registry

type PierHAConstructor func(client HAClient, pierID string) PierHA

var (
	TxsExecutorConstructorM = make(map[string]TxsExecutorConstructor)
	ContractConstructorM    = make(map[string]*ContractInfo)
	RegisterConstructorM    = make(map[string]RegistryConstructor)

	PierHAConstructorM = make(map[string]PierHAConstructor)
)

func RegisterRegistryConstructor(typ string, f RegistryConstructor) {
	RegisterConstructorM[typ] = f
}

func GetRegistryConstructor(typ string) (RegistryConstructor, error) {
	registry, ok := RegisterConstructorM[typ]
	if !ok {
		return nil, fmt.Errorf("type %s registry is unsupported", typ)
	}
	return registry, nil
}

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

func RegisterPierHAConstructor(typ string, f PierHAConstructor) {
	PierHAConstructorM[typ] = f
}

func GetPierHAConstructor(typ string) (PierHAConstructor, error) {
	con, ok := PierHAConstructorM[typ]
	if !ok {
		return nil, fmt.Errorf("type %s is unsupported", typ)
	}
	return con, nil
}
