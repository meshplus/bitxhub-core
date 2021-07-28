package validatorlib

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	mb "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/meshplus/bitxhub-model/pb"
)

const (
	FABRIC_EVALUATOR = "fabric_evaluator"
)

var (
	evaluatorMap map[string]*PolicyEvaluator
)

type valiadationArtifacts struct {
	rwset        []byte
	prp          []byte
	endorsements []*peer.Endorsement
	cap          *peer.ChaincodeActionPayload
	payload      payloadInfo
}

type ValidatorInfo struct {
	ChainId  string   `json:"chain_id"`
	ConfByte []string `json:"conf_byte"`
	Policy   string   `json:"policy"`
	Cid      string   `json:"cid"`
}

type payloadInfo struct {
	Index          uint64 `json:"index"`
	DstContractDID string `json:"dst_contract_did"`
	SrcContractID  string `json:"src_contract_id"`
	Func           string `json:"func"`
	Args           string `json:"args"`
	Callback       string `json:"callback"`
	Argscb         string `json:"argscb"`
	Rollback       string `json:"rollback"`
	Argsrb         string `json:"argsrb"`
}

func GetPolicyEnvelope(policy string) ([]byte, error) {
	policyEnv, err := cauthdsl.FromString(policy)
	if err != nil {
		return nil, err
	}
	policyBytes, err := proto.Marshal(policyEnv)
	if err != nil {
		return nil, err
	}
	return policyBytes, nil
}

func UnmarshalValidatorInfo(validatorBytes []byte) (*ValidatorInfo, error) {
	vInfo := &ValidatorInfo{}
	if err := json.Unmarshal(validatorBytes, vInfo); err != nil {
		return nil, err
	}
	return vInfo, nil
}

func ExtractValidationArtifacts(proof []byte) (*valiadationArtifacts, error) {
	return extractValidationArtifacts(proof)
}

func extractValidationArtifacts(proof []byte) (*valiadationArtifacts, error) {
	cap, err := protoutil.UnmarshalChaincodeActionPayload(proof)
	if err != nil {
		return nil, err
	}

	pRespPayload, err := protoutil.UnmarshalProposalResponsePayload(cap.Action.ProposalResponsePayload)
	if err != nil {
		err = fmt.Errorf("GetProposalResponsePayload error %s", err)
		return nil, err
	}
	if pRespPayload.Extension == nil {
		err = fmt.Errorf("nil pRespPayload.Extension")
		return nil, err
	}
	respPayload, err := protoutil.UnmarshalChaincodeAction(pRespPayload.Extension)
	if err != nil {
		err = fmt.Errorf("GetChaincodeAction error %s", err)
		return nil, err
	}

	var (
		payload      payloadInfo
		payloadArray []payloadInfo
	)
	err = json.Unmarshal(respPayload.Response.Payload, &payloadArray)
	if err != nil {
		// try if it is from getOutMessage
		if err = json.Unmarshal(respPayload.Response.Payload, &payload); err != nil {
			return nil, err
		}
	} else {
		payload = payloadArray[len(payloadArray)-1]
	}

	return &valiadationArtifacts{
		rwset:        respPayload.Results,
		prp:          cap.Action.ProposalResponsePayload,
		endorsements: cap.Action.Endorsements,
		cap:          cap,
		payload:      payload,
	}, nil
}

func PreCheck(proof, payload []byte, cid string) (*valiadationArtifacts, error) {
	// Get the validation artifacts that help validate the chaincodeID and policy
	artifact, err := extractValidationArtifacts(proof)
	if err != nil {
		return nil, err
	}

	err = ValidateChainCodeID(artifact.prp, cid)
	if err != nil {
		return nil, err
	}

	err = ValidatePayload(artifact.payload, payload)
	if err != nil {
		return nil, err
	}

	return artifact, nil
}

func ValidateV14(proof, payload, policyBytes []byte, confByte []string, cid, from string) error {
	// Get the validation artifacts that help validate the chaincodeID and policy
	artifact, err := extractValidationArtifacts(proof)
	if err != nil {
		return err
	}

	err = ValidateChainCodeID(artifact.prp, cid)
	if err != nil {
		return err
	}

	err = ValidatePayload(artifact.payload, payload)
	if err != nil {
		return err
	}

	signatureSet := GetSignatureSet(artifact)

	pe, err := NewPolicyEvaluator(confByte)
	if err != nil {
		return err
	}

	return pe.Evaluate(policyBytes, signatureSet)
}

func ValidateChainCodeID(prp []byte, name string) error {
	payload, err := protoutil.UnmarshalProposalResponsePayload(prp)
	if err != nil {
		err = fmt.Errorf("GetProposalResponsePayload error %s", err)
		return err
	}
	chaincodeAct, err := protoutil.UnmarshalChaincodeAction(payload.Extension)
	if err != nil {
		err = fmt.Errorf("GetChaincodeAction error %s", err)
		return err
	}
	if name != chaincodeAct.ChaincodeId.Name {
		return fmt.Errorf("chaincode id does not match")
	}

	return nil
}

func ValidatePayload(info payloadInfo, payloadByte []byte) error {
	payload := &pb.Payload{}
	if err := payload.Unmarshal(payloadByte); err != nil {
		return err
	}

	if payload.Encrypted {
		return nil
	}

	content := &pb.Content{}
	if err := content.Unmarshal(payload.Content); err != nil {
		return fmt.Errorf("unmarshal ibtp payload content: %w", err)
	}

	//if bitxid.DID(info.DstContractDID).GetAddress() != content.DstContractId {
	//	return fmt.Errorf("dst contrct id not correct")
	//}
	//if info.SrcContractID != content.SrcContractId {
	//	return fmt.Errorf("src contrct id not correct")
	//}
	if info.Func != content.Func {
		return fmt.Errorf("interchain function name not correct")
	}
	if info.Callback != content.Callback {
		return fmt.Errorf("callback not correct")
	}
	if info.Rollback != content.Rollback {
		return fmt.Errorf("rollback not correct")
	}
	if !checkArgs(info.Argsrb, content.ArgsRb) {
		return fmt.Errorf("args for rollback not correct")
	}
	if !checkArgs(info.Argscb, content.ArgsCb) {
		return fmt.Errorf("args for callback not correct")
	}
	if !checkArgs(info.Args, content.Args) {
		return fmt.Errorf("args for interchain not correct")
	}
	return nil
}

func checkArgs(args string, argArr [][]byte) bool {
	if args == "" && len(argArr) == 0 {
		return true
	}
	argsSplit := strings.Split(args, ",")
	if len(argsSplit) != len(argArr) {
		return false
	}
	for index, arg := range argsSplit {
		if arg != string(argArr[index]) {
			return false
		}
	}
	return true
}

type PolicyEvaluator struct {
	msp.IdentityDeserializer
}

func NewPolicyEvaluator(confBytes []string) (*PolicyEvaluator, error) {
	mspList := make([]msp.MSP, len(confBytes))
	for i, confByte := range confBytes {
		tempBccsp, err := msp.New(
			&msp.BCCSPNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_3}},
			factory.GetDefault(),
		)
		if err != nil {
			return nil, err
		}
		conf := &mb.MSPConfig{}
		if err := proto.UnmarshalText(confByte, conf); err != nil {
			return nil, err
		}
		err = tempBccsp.Setup(conf)
		if err != nil {
			return nil, err
		}
		mspList[i] = tempBccsp
	}

	manager := msp.NewMSPManager()
	err := manager.Setup(mspList)
	if err != nil {
		return nil, err
	}
	deserializer := &dynamicDeserializer{mspm: manager}
	pe := &PolicyEvaluator{IdentityDeserializer: deserializer}

	return pe, nil
}

func (id *PolicyEvaluator) Evaluate(policyBytes []byte, signatureSet []*protoutil.SignedData) error {
	pp := cauthdsl.NewPolicyProvider(id.IdentityDeserializer)
	policy, _, err := pp.NewPolicy(policyBytes)
	if err != nil {
		return err
	}
	return policy.EvaluateSignedData(signatureSet)
}

func GetSignatureSet(artifact *valiadationArtifacts) []*protoutil.SignedData {
	signatureSet := []*protoutil.SignedData{}
	for _, endorsement := range artifact.endorsements {
		data := make([]byte, len(artifact.prp)+len(endorsement.Endorser))
		copy(data, artifact.prp)
		copy(data[len(artifact.prp):], endorsement.Endorser)

		signatureSet = append(signatureSet, &protoutil.SignedData{
			// set the data that is signed; concatenation of proposal response bytes and endorser ID
			Data: data,
			// set the identity that signs the message: it's the endorser
			Identity: endorsement.Endorser,
			// set the signature
			Signature: endorsement.Signature})
	}
	return signatureSet
}

type dynamicDeserializer struct {
	mspm msp.MSPManager
}

func (ds *dynamicDeserializer) DeserializeIdentity(serializedIdentity []byte) (msp.Identity, error) {
	return ds.mspm.DeserializeIdentity(serializedIdentity)
}

func (ds *dynamicDeserializer) IsWellFormed(identity *mb.SerializedIdentity) error {
	return ds.mspm.IsWellFormed(identity)
}
