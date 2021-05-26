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
	v := NewValidationEngine(nil, nil, logger)

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
	v := NewValidationEngine(nil, nil, logger)

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
func TestFabSimValidator_Verify(t *testing.T) {
	logger := log.NewWithModule("validator")
	v := NewValidationEngine(nil, nil, logger)

	proof, err := ioutil.ReadFile("./testdata/proof_1.0.0_rc")
	require.Nil(t, err)

	validators, err := ioutil.ReadFile("./testdata/single_validator")
	require.Nil(t, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "0x30c5D3aeb4681af4D13384DBc2a717C51cb1cc11",
		Func:          "interchainCharge",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "",
		Rollback:      "interchainRollback",
		ArgsRb:        [][]byte{[]byte("Alice"), []byte("1")},
	}

	bytes, err := content.Marshal()
	require.Nil(t, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(t, err)

	ok, err := v.Validate(SimFabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
	require.Nil(t, err)
	require.True(t, ok)
}

func BenchmarkFabV14Validator_Verify(b *testing.B) {
	logger := log.NewWithModule("validator")

	proof, err := ioutil.ReadFile("./testdata/proof_1.0.0_rc")
	require.Nil(b, err)

	validators, err := ioutil.ReadFile("./testdata/validator_1.0.0_rc")
	require.Nil(b, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "mychannel&transfer",
		Func:          "interchainCharge",
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

	v := NewValidationEngine(nil, nil, logger)
	ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
	require.Nil(b, err)
	require.True(b, ok)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
		require.Nil(b, err)
		require.True(b, ok)
	}
}
func BenchmarkFabSimValidator_Verify(b *testing.B) {
	logger := log.NewWithModule("validator")

	proof, err := ioutil.ReadFile("./testdata/proof_1.0.0_rc")
	require.Nil(b, err)

	validators, err := ioutil.ReadFile("./testdata/single_validator")
	require.Nil(b, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "0x30c5D3aeb4681af4D13384DBc2a717C51cb1cc11",
		Func:          "interchainCharge",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "",
		Rollback:      "interchainRollback",
		ArgsRb:        [][]byte{[]byte("Alice"), []byte("1")},
	}

	bytes, err := content.Marshal()
	require.Nil(b, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(b, err)

	v := NewValidationEngine(nil, nil, logger)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok, err := v.Validate(SimFabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
		require.Nil(b, err)
		require.True(b, ok)
	}
}

func BenchmarkFabComplexValidator_Verify(b *testing.B) {
	logger := log.NewWithModule("validator")

	proof, err := ioutil.ReadFile("./testdata/proof_1.0.0_rc_complex")
	require.Nil(b, err)

	validators, err := ioutil.ReadFile("./testdata/validator_1.0.0_rc_complex")
	require.Nil(b, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "mychannel&transfer",
		Func:          "interchainCharge",
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

	v := NewValidationEngine(nil, nil, logger)
	ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
	require.Nil(b, err)
	require.True(b, ok)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok, err := v.Validate(FabricRuleAddr, "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, body, string(validators))
		require.Nil(b, err)
		require.True(b, ok)
	}
}
