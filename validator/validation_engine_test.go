package validator

import (
	"io/ioutil"
	"testing"

	"github.com/meshplus/bitxhub-kit/log"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/stretchr/testify/require"
)

func TestFabV14ValidatorWasm_Verify(t *testing.T) {
	logger := log.NewWithModule("validator")
	v := NewValidationEngine(nil, logger)

	proof, err := ioutil.ReadFile("./testdata/proof")
	require.Nil(t, err)

	validators, err := ioutil.ReadFile("./testdata/validators")
	require.Nil(t, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "0x668a209Dc6562707469374B8235e37b8eC25db08",
		Func:          "get",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "interchainConfirm",
	}

	bytes, err := content.Marshal()
	require.Nil(t, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(t, err)

	ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
	require.NotNil(t, err)
	require.False(t, ok)
}
func TestFabV14Validator_Verify(t *testing.T) {
	logger := log.NewWithModule("validator")
	v := NewValidationEngine(nil, logger)

	proof, err := ioutil.ReadFile("./testdata/proof")
	require.Nil(t, err)

	validators, err := ioutil.ReadFile("./testdata/validators")
	require.Nil(t, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "0x668a209Dc6562707469374B8235e37b8eC25db08",
		Func:          "get",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "interchainConfirm",
	}

	bytes, err := content.Marshal()
	require.Nil(t, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(t, err)

	ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
	require.NotNil(t, err)
	require.False(t, ok)
}

func BenchmarkFabV14Validator_Verify(b *testing.B) {
	logger := log.NewWithModule("validator")

	proof, err := ioutil.ReadFile("./testdata/proof")
	require.Nil(b, err)

	validators, err := ioutil.ReadFile("./testdata/validators")
	require.Nil(b, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "0x668a209Dc6562707469374B8235e37b8eC25db08",
		Func:          "get",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "interchainConfirm",
	}

	bytes, err := content.Marshal()
	require.Nil(b, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewValidationEngine(nil, logger)
		ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
		require.Nil(b, err)
		require.True(b, ok)
	}
}
