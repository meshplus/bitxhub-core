package agency

import (
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
)

type TxsExecutor interface {
	ApplyTransactions(txs []*pb.Transaction, invalidTxs map[int]InvalidReason) []*pb.Receipt

	GetBoltContracts() map[string]Contract

	AddNormalTx(hash *types.Hash)

	GetNormalTxs() []*types.Hash

	AddInterchainCounter(to string, index uint64)

	GetInterchainCounter() map[string][]uint64

	GetDescription() string
}

type TxOpt struct {
	Contracts map[string]Contract
}
