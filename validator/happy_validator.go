package validator

var _ Validator = (*HappyValidator)(nil)

type HappyValidator struct {
}

func (h *HappyValidator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	return true, 0, nil
}
