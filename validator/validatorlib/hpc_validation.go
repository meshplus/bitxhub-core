package validatorlib

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"

	"github.com/cbergoon/merkletree"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"github.com/meshplus/bitxhub-kit/types"
)

type HpcProof struct {
	Signatures [][]byte `json:"signatures"`
	SpvHash    []byte   `json:"spv_hash"`
	PreHashes  [][]byte `json:"pre_hashes"`
	PostHashes [][]byte `json:"post_hashes"`
	RootHash   []byte   `json:"root_hash"`
}

type HpcValidator struct {
	Total      int      `json:"total"`
	Validators [][]byte `json:"validators"`
}

func ValidateHpc(proof []byte, validators []byte, payload []byte) (bool, error) {
	hpcProof := &HpcProof{}
	hpcValidator := &HpcValidator{}
	if err := json.Unmarshal(proof, hpcProof); err != nil {
		return false, err
	}
	if err := json.Unmarshal(validators, hpcValidator); err != nil {
		return false, err
	}
	count := 0
	for _, sig := range hpcProof.Signatures {
		validSig := false
		for index, validator := range hpcValidator.Validators {
			address := types.NewAddress(validator)
			ok, err := asym.Verify(crypto.Secp256k1, sig, hpcProof.SpvHash, *address)
			if err != nil {
				return false, err
			}
			if ok {
				validSig = true
				hpcValidator.Validators = append(hpcValidator.Validators[:index], hpcValidator.Validators[index+1:]...)
				break
			}
		}
		if validSig {
			count += 1
			if count == (hpcValidator.Total-1)/3+1 {
				return VerifyMerkle(hpcProof, payload)
			}
		}
	}
	return false, nil
}

func VerifyMerkle(hpcProof *HpcProof, payload []byte) (bool, error) {
	h := sha256.Sum256(payload)
	if !bytes.Equal(h[:], hpcProof.SpvHash) {
		return false, nil
	}
	hashes := make([]merkletree.Content, 0, len(hpcProof.PreHashes)+len(hpcProof.PostHashes)+1)
	for _, hash := range hpcProof.PreHashes {
		hashes = append(hashes, types.NewHash(hash))
	}
	hashes = append(hashes, types.NewHash(hpcProof.SpvHash))
	for _, hash := range hpcProof.PostHashes {
		hashes = append(hashes, types.NewHash(hash))
	}
	root, err := calcMerkleRoot(hashes)
	if err != nil {
		return false, err
	}
	return root.Equals(types.NewHash(hpcProof.RootHash))
}

func calcMerkleRoot(contents []merkletree.Content) (*types.Hash, error) {
	if len(contents) == 0 {
		return &types.Hash{}, nil
	}

	tree, err := merkletree.NewTree(contents)
	if err != nil {
		return nil, err
	}

	return types.NewHash(tree.MerkleRoot()), nil
}
