package validatorlib

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFabricSM2Validation(t *testing.T) {
	smProof, err := ioutil.ReadFile("./test_data/proof_peer1")
	require.Nil(t, err)
	validator, err := ioutil.ReadFile("./test_data/peer1.pem")
	require.Nil(t, err)
	_, err = fabric_sm2_validate(smProof, validator)

	require.Nil(t, err)
}
