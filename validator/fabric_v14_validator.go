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
	logger     logrus.FieldLogger
	evMap      map[string]*validatorlib.ValidatorInfo
	peMap      map[string]*validatorlib.PolicyEvaluator
	ppMap      map[string]policies.Policy
	idMap      map[string]mspi.Identity
	keyMap     map[string]struct{}
	evMapLock  sync.Mutex
	peMapLock  sync.Mutex
	ppMapLock  sync.Mutex
	idMapLock  sync.Mutex
	keyMapLock sync.Mutex
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
	vlt.evMapLock.Lock()
	vInfo, ok := vlt.evMap[from]
	vlt.evMapLock.Unlock()
	if !ok {
		vInfo, err = validatorlib.UnmarshalValidatorInfo([]byte(validators))
		if err != nil {
			return false, err
		}
		vlt.evMapLock.Lock()
		vlt.evMap[from] = vInfo
		vlt.evMapLock.Unlock()
	}
	// Get the validation artifacts that help validate the chaincodeID and policy
	artifact, err := validatorlib.PreCheck(proof, payload, vInfo.Cid)
	if err != nil {
		return false, err
	}

	signatureSet := validatorlib.GetSignatureSet(artifact)

	vlt.ppMapLock.Lock()
	policy, ok := vlt.ppMap[vInfo.ChainId]
	vlt.ppMapLock.Unlock()
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
		vlt.ppMapLock.Lock()
		vlt.ppMap[vInfo.ChainId] = policy
		vlt.ppMapLock.Unlock()

		vlt.peMapLock.Lock()
		vlt.peMap[vInfo.ChainId] = pe
		vlt.peMapLock.Unlock()
	}

	vlt.peMapLock.Lock()
	pe := vlt.peMap[vInfo.ChainId]
	vlt.peMapLock.Unlock()
	return vlt.EvaluateSignedData(signatureSet, pe.IdentityDeserializer, policy)
}

func (vlt *FabV14Validator) EvaluateSignedData(signedData []*protoutil.SignedData, identityDeserializer mspi.IdentityDeserializer, policy policies.Policy) (bool, error) {
	idMap := map[string]struct{}{}
	identities := make([]mspi.Identity, 0, len(signedData))
	ids := make([]string, 0, len(signedData))

	for _, sd := range signedData {
		var err error
		vlt.idMapLock.Lock()
		identity, ok := vlt.idMap[string(sd.Identity)]
		vlt.idMapLock.Unlock()
		if !ok {
			identity, err = identityDeserializer.DeserializeIdentity(sd.Identity)
			if err != nil {
				continue
			}
			vlt.idMapLock.Lock()
			vlt.idMap[string(sd.Identity)] = identity
			vlt.idMapLock.Unlock()
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

	vlt.keyMapLock.Lock()
	_, ok := vlt.keyMap[idStr]
	vlt.keyMapLock.Unlock()
	if !ok {
		if err := policy.EvaluateIdentities(identities); err != nil {
			return false, err
		}
		vlt.keyMapLock.Lock()
		vlt.keyMap[idStr] = struct{}{}
		vlt.keyMapLock.Unlock()
		return true, nil
	}

	return true, nil
}
