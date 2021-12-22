package validator

var _ Validator = (*Fab10Validator)(nil)

type Fab10Validator struct {
}

func (h *Fab10Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 10; i++ {
		num++
	}
	return true, 0, nil
}
