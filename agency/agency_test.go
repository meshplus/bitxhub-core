package agency

import (
	"testing"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/stretchr/testify/assert"
)

func TestGetContractInfo(t *testing.T) {
	addr0 := types.NewAddress([]byte{1})
	addr1 := types.NewAddress([]byte{1})

	RegisterContractConstructor("1", addr0, newContract)

	info, err := GetContractInfo(addr1)
	assert.Nil(t, err)
	assert.Equal(t, "1", info.Name)
	assert.Equal(t, newContract(), info.Constructor())
}

func newContract() Contract {
	return struct {
	}{}
}
