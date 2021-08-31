package validator

var _ Validator = (*HappyValidator)(nil)

type HappyValidator struct {
}

func (h *HappyValidator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	return true, nil
}
