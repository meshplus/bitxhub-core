package boltvm

import "github.com/meshplus/bitxhub-model/pb"

type TransactionRecord struct {
	Status        pb.TransactionStatus
	TimeoutHeight uint64
}
