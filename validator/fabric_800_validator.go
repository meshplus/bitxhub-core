package validator

var _ Validator = (*Fab800Validator)(nil)

type Fab800Validator struct {
}

func (h *Fab800Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 800; i++ {
		num++
	}
	return true, 0, nil
}
