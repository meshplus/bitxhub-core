package validator

var _ Validator = (*Fab50Validator)(nil)

type Fab50Validator struct {
}

func (h *Fab50Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 50; i++ {
		num++
	}
	return true, 0, nil
}
