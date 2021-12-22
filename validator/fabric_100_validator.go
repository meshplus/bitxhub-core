package validator

var _ Validator = (*Fab100Validator)(nil)

type Fab100Validator struct {
}

func (h *Fab100Validator) Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	num := 0
	for i := 0; i < 100; i++ {
		num++
	}
	return true, 0, nil
}
