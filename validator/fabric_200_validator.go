package validator

var _ Validator = (*Fab200Validator)(nil)

type Fab200Validator struct {
}

func (h *Fab200Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 200; i++ {
		num++
	}
	return true, 0, nil
}
