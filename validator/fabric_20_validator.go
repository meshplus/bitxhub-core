package validator

var _ Validator = (*Fab20Validator)(nil)

type Fab20Validator struct {
}

func (h *Fab20Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 20; i++ {
		num++
	}
	return true, 0, nil
}
