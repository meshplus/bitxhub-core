package validator

var _ Validator = (*Fab600Validator)(nil)

type Fab600Validator struct {
}

func (h *Fab600Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 600; i++ {
		num++
	}
	return true, 0, nil
}
