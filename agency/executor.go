package agency

import (
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
)

type TxsExecutor interface {
	ApplyTransactions(txs []*pb.Transaction) []*pb.Receipt

	GetBoltContracts() map[string]Contract

	AddNormalTx(hash types.Hash)

	GetNormalTxs() []types.Hash

	AddInterchainCounter(to string, index uint64)

	GetInterchainCounter() map[string][]uint64
}

type TxOpt struct {
	Contracts map[string]Contract
}
