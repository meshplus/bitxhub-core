package validator

var _ Validator = (*Fab1500Validator)(nil)

type Fab1500Validator struct {
}

func (h *Fab1500Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 1500; i++ {
		num++
	}
	return true, 0, nil
}
