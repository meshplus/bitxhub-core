package agency

import (
	"fmt"

	"github.com/meshplus/bitxhub-model/pb"
)

type ApplyTxFun func(int, *pb.Transaction, *TxOpt) *pb.Receipt

type RegisterContractFun func() map[string]Contract

type TxsExecutorConstructor func(ApplyTxFun, RegisterContractFun) TxsExecutor

var TxsExecutorConstructorM = make(map[string]TxsExecutorConstructor)

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
