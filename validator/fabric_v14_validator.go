package validator

import (
	"sort"
	"sync"

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
	evMap  *sync.Map
	peMap  *sync.Map
	ppMap  *sync.Map
	idMap  *sync.Map
	keyMap *sync.Map
}

// New a validator instance
func NewFabV14Validator(logger logrus.FieldLogger) *FabV14Validator {
	return &FabV14Validator{
		logger: logger,
		evMap:  &sync.Map{},
		peMap:  &sync.Map{},
		ppMap:  &sync.Map{},
		idMap:  &sync.Map{},
		keyMap: &sync.Map{},
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabV14Validator) Verify(from string, proof, payload []byte, validators string) (bool, uint64, error) {
	var (
		vInfo  *validatorlib.ValidatorInfo
		policy policies.Policy
		err    error
	)
	raw, ok := vlt.evMap.Load(from)
	if !ok {
		vInfo, err = validatorlib.UnmarshalValidatorInfo([]byte(validators))
		if err != nil {
			return false, 0, err
		}
		vlt.evMap.Store(from, vInfo)
	} else {
		vInfo = raw.(*validatorlib.ValidatorInfo)
	}

	// Get the validation artifacts that help validate the chaincodeID and policy
	artifact, err := validatorlib.PreCheck(proof, payload, vInfo.Cid)
	if err != nil {
		return false, 0, err
	}

	signatureSet := validatorlib.GetSignatureSet(artifact)

	raw, ok = vlt.ppMap.Load(vInfo.ChainId)
	if !ok {
		pe, err := validatorlib.NewPolicyEvaluator(vInfo.ConfByte)
		if err != nil {
			return false, 0, err
		}
		pp := cauthdsl.NewPolicyProvider(pe.IdentityDeserializer)
		policy, _, err = pp.NewPolicy([]byte(vInfo.Policy))
		if err != nil {
			return false, 0, err
		}
		vlt.ppMap.Store(vInfo.ChainId, policy)
		vlt.peMap.Store(vInfo.ChainId, pe)
	} else {
		policy = raw.(policies.Policy)
	}

	pe, _ := vlt.peMap.Load(vInfo.ChainId)
	return vlt.EvaluateSignedData(signatureSet, pe.(*validatorlib.PolicyEvaluator).IdentityDeserializer, policy)
}

func (vlt *FabV14Validator) EvaluateSignedData(signedData []*protoutil.SignedData, identityDeserializer mspi.IdentityDeserializer, policy policies.Policy) (bool, uint64, error) {
	idMap := map[string]struct{}{}
	identities := make([]mspi.Identity, 0, len(signedData))
	ids := make([]string, 0, len(signedData))

	for _, sd := range signedData {
		var (
			identity mspi.Identity
			err      error
		)
		raw, ok := vlt.idMap.Load(string(sd.Identity))
		if !ok {
			identity, err = identityDeserializer.DeserializeIdentity(sd.Identity)
			if err != nil {
				continue
			}
			vlt.idMap.Store(string(sd.Identity), identity)
		} else {
			identity = raw.(mspi.Identity)
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

	_, ok := vlt.keyMap.Load(idStr)
	if !ok {
		if err := policy.EvaluateIdentities(identities); err != nil {
			return false, 0, err
		}
		vlt.keyMap.Store(idStr, struct{}{})
		return true, 0, nil
	}

	return true, 0, nil
}
