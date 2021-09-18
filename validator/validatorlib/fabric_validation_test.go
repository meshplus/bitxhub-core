package validatorlib

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReceiptVerify(t *testing.T) {

	proof, err := ioutil.ReadFile("./test_data/proof_receipt")
	require.Nil(t, err)

	payload, err := ioutil.ReadFile("./test_data/payload_receipt")
	require.Nil(t, err)

	_, err = PreCheck(proof, payload, "broker")
	require.Nil(t, err)

}
