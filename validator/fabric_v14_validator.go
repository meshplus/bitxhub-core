package validator

import (
	"sort"

	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/common/policies"
	mspi "github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabV14Validator struct {
	logger logrus.FieldLogger
	evMap  map[string]*validatorlib.ValidatorInfo
	peMap  map[string]*validatorlib.PolicyEvaluator
	ppMap  map[string]policies.Policy
	idMap  map[string]mspi.Identity
	keyMap map[string]struct{}
}

// New a validator instance
func NewFabV14Validator(logger logrus.FieldLogger) *FabV14Validator {
	return &FabV14Validator{
		logger: logger,
		evMap:  make(map[string]*validatorlib.ValidatorInfo),
		peMap:  make(map[string]*validatorlib.PolicyEvaluator),
		ppMap:  make(map[string]policies.Policy),
		idMap:  make(map[string]mspi.Identity),
		keyMap: make(map[string]struct{}),
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabV14Validator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	var err error
	vInfo, ok := vlt.evMap[from]
	if !ok {
		vInfo, err = validatorlib.UnmarshalValidatorInfo([]byte(validators))
		if err != nil {
			return false, err
		}
		vlt.evMap[from] = vInfo
	}
	// Get the validation artifacts that help validate the chaincodeID and policy
	artifact, err := validatorlib.PreCheck(proof, payload, vInfo.Cid)
	if err != nil {
		return false, err
	}

	signatureSet := validatorlib.GetSignatureSet(artifact)

	policy, ok := vlt.ppMap[vInfo.ChainId]
	if !ok {
		pe, err := validatorlib.NewPolicyEvaluator(vInfo.ConfByte)
		if err != nil {
			return false, err
		}
		pp := cauthdsl.NewPolicyProvider(pe.IdentityDeserializer)
		policy, _, err = pp.NewPolicy([]byte(vInfo.Policy))
		if err != nil {
			return false, err
		}
		vlt.ppMap[vInfo.ChainId] = policy
		vlt.peMap[vInfo.ChainId] = pe
	}

	pe := vlt.peMap[vInfo.ChainId]
	return vlt.EvaluateSignedData(signatureSet, pe.IdentityDeserializer, policy)
}

func (vlt *FabV14Validator) EvaluateSignedData(signedData []*protoutil.SignedData, identityDeserializer mspi.IdentityDeserializer, policy policies.Policy) (bool, error) {
	idMap := map[string]struct{}{}
	identities := make([]mspi.Identity, 0, len(signedData))
	ids := make([]string, 0, len(signedData))

	for _, sd := range signedData {
		var err error
		identity, ok := vlt.idMap[string(sd.Identity)]
		if !ok {
			identity, err = identityDeserializer.DeserializeIdentity(sd.Identity)
			if err != nil {
				continue
			}
			vlt.idMap[string(sd.Identity)] = identity
		}

		key := identity.GetIdentifier().Mspid + identity.GetIdentifier().Id

		if _, ok := idMap[key]; ok {
			continue
		}

		err = identity.Verify(sd.Data, sd.Signature)
		if err != nil {
			continue
		}

		idMap[key] = struct{}{}
		ids = append(ids, key)
		identities = append(identities, identity)
	}
	nids := sort.StringSlice(ids)
	var idStr string
	for _, id := range nids {
		idStr = idStr + id
	}
	if _, ok := vlt.keyMap[idStr]; !ok {
		if err := policy.EvaluateIdentities(identities); err != nil {
			return false, err
		}
		vlt.keyMap[idStr] = struct{}{}
		return true, nil
	}

	return true, nil
}
