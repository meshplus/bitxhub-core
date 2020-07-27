package validator

import (
	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabV14Validator struct {
	logger logrus.FieldLogger
	evMap  map[string]*validatorlib.PolicyEvaluator
}

// New a validator instance
func NewFabV14Validator(logger logrus.FieldLogger) *FabV14Validator {
	return &FabV14Validator{
		logger: logger,
		evMap:  make(map[string]*validatorlib.PolicyEvaluator),
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabV14Validator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	vInfo, err := validatorlib.UnmarshalValidatorInfo([]byte(validators))
	if err != nil {
		return false, err
	}
	// Get the validation artifacts that help validate the chaincodeID and policy
	artifact, err := validatorlib.PreCheck(proof, payload, vInfo.Cid)
	if err != nil {
		return false, err
	}

	signatureSet := validatorlib.GetSignatureSet(artifact)

	pe, ok := vlt.evMap[vInfo.ChainId]
	if !ok {
		pe, err = validatorlib.NewPolicyEvaluator(vInfo.ConfByte)
		if err != nil {
			return false, err
		}
		vlt.evMap[vInfo.ChainId] = pe
	}

	err = pe.Evaluate([]byte(vInfo.Policy), signatureSet)
	if err != nil {
		return false, err
	}

	return true, nil
}
