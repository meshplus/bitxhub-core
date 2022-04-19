package agency

import "github.com/meshplus/bitxhub-model/pb"

type OffChainTransmission interface {
	Start() error

	Stop() error

	ParseIBTPOffChain(ibtp *pb.IBTP) (*pb.IBTP, error)

	CheckIBTPOffChain(ibtp *pb.IBTP) (bool, error)
}
