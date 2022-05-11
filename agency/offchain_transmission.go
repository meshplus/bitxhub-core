package agency

type OffChainTransmission interface {
	Start() error

	Stop() error
}
