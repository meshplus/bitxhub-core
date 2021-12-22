package validator

var _ Validator = (*Fab2000Validator)(nil)

type Fab2000Validator struct {
}

func (h *Fab2000Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 2000; i++ {
		num++
	}
	return true, 0, nil
}
