package agency

import "github.com/meshplus/bitxhub-model/pb"

type PierHA interface {
	Start() error

	Stop() error

	IsMain() <-chan bool
}

type HAClient interface {
	//Check whethe there is a master pier connect to the BitXHub.
	CheckMasterPier(address string) (*pb.Response, error)

	//Set the master pier connect to the BitXHub.
	SetMasterPier(address string, index string, timeout int64) (*pb.Response, error)

	//Update the master pier status
	HeartBeat(address string, index string) (*pb.Response, error)
}
