package validator

var _ Validator = (*Fab400Validator)(nil)

type Fab400Validator struct {
}

func (h *Fab400Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 400; i++ {
		num++
	}
	return true, 0, nil
}
