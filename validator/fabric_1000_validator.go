package validator

var _ Validator = (*Fab1000Validator)(nil)

type Fab1000Validator struct {
}

func (h *Fab1000Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 1000; i++ {
		num++
	}
	return true, 0, nil
}
