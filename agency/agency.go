package agency

import (
	"fmt"

	"github.com/meshplus/bitxhub-core/order"
	"github.com/meshplus/bitxhub-kit/log"
	"github.com/meshplus/bitxhub-kit/storage"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

var logger = log.NewWithModule("agency")

type InvalidReason string

type ContractInfo struct {
	Name        string
	Constructor ContractConstructor
}

type ApplyTxFunc func(int, pb.Transaction, InvalidReason, *TxOpt) *pb.Receipt

type RegisterContractFunc func() map[string]Contract

type TxsExecutorConstructor func(ApplyTxFunc, RegisterContractFunc, logrus.FieldLogger) TxsExecutor

type LicenseConstructor func(pubKey, verifier string) License

type ContractConstructor func() Contract

type RegistryConstructor func(storage.Storage, logrus.FieldLogger) Registry

type PierHAConstructor func(client HAClient, pierID string) PierHA

type OrderConstructor func(opt ...order.Option) (order.Order, error)

type OffchainTransmissionConstructor func(appchainID string, peerMgr PeerManager, client Client) OffChainTransmission

var (
	TxsExecutorConstructorM         = make(map[string]TxsExecutorConstructor)
	ContractConstructorM            = make(map[string]*ContractInfo)
	RegisterConstructorM            = make(map[string]RegistryConstructor)
	LicenseConstructorM             = make(map[string]LicenseConstructor)
	PierHAConstructorM              = make(map[string]PierHAConstructor)
	OrderConstructorM               = make(map[string]OrderConstructor)
	OffchainTransmissionContructorM = make(map[string]OffchainTransmissionConstructor)
)

func RegisterOrderConstructor(typ string, f OrderConstructor) {
	OrderConstructorM[typ] = f
}

func GetOrderConstructor(typ string) (OrderConstructor, error) {
	con, ok := OrderConstructorM[typ]
	if !ok {
		return nil, fmt.Errorf("the order type %s is unsupported", typ)
	}
	return con, nil
}

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

func GetLicenseConstructor(typ string) (LicenseConstructor, error) {
	con, ok := LicenseConstructorM[typ]
	if !ok {
		return nil, fmt.Errorf("type %s is unsupported", typ)
	}
	return con, nil
}

func RegisterLicenseConstructor(typ string, f LicenseConstructor) {
	LicenseConstructorM[typ] = f
}

func RegisterContractConstructor(name string, addr *types.Address, f ContractConstructor) {
	ContractConstructorM[addr.String()] = &ContractInfo{
		Name:        name,
		Constructor: f,
	}
	logger.WithFields(logrus.Fields{
		"name": name,
		"addr": addr.String(),
	}).Debug("contract registered")
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

func RegisterOffchainTransmissionConstructor(typ string, f OffchainTransmissionConstructor) {
	OffchainTransmissionContructorM[typ] = f
}

func GetOffchainTransmissionConstructor(typ string) (OffchainTransmissionConstructor, error) {
	con, ok := OffchainTransmissionContructorM[typ]
	if !ok {
		return nil, fmt.Errorf("type %s is unsupported", typ)
	}
	return con, nil
}
