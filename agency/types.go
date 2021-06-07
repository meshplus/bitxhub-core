package agency

type Contract interface{}
type Registry interface{}
type License interface {
	Verify(licensePath string) error
}
