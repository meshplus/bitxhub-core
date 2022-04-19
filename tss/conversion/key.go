package conversion

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sort"

	bcrypto "github.com/binance-chain/tss-lib/crypto"
	"github.com/btcsuite/btcd/btcec"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ecdsa2 "github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
)

// TSS PoolPubkey: "crypto/ecdsa".PublicKey (Spe256k1) =================================================================

// ECPoint -> ecdsa.PublicKey
func GetTssPubKey(pubKeyPoint *bcrypto.ECPoint) (*ecdsa.PublicKey, error) {
	// we check whether the point is on curve according to Kudelski report
	if pubKeyPoint == nil || !isOnCurve(pubKeyPoint.X(), pubKeyPoint.Y()) {
		return nil, fmt.Errorf("invalid points")
	}
	tssPubKey := btcec.PublicKey{
		Curve: btcec.S256(),
		X:     pubKeyPoint.X(),
		Y:     pubKeyPoint.Y(),
	}
	return tssPubKey.ToECDSA(), nil
}

// ecdsa.PublicKey -> addr, byte
func GetPubKeyInfoFromECDSAPubkey(ecdsaPk *ecdsa.PublicKey) (string, []byte, error) {
	pubKey, err := ecdsa2.NewPublicKey(*ecdsaPk)
	if err != nil {
		return "", nil, err
	}

	pubAddr, err := pubKey.Address()
	if err != nil {
		return "", nil, fmt.Errorf("pubkey to address error: %w", err)
	}
	pubData, err := pubKey.Bytes()
	if err != nil {
		return "", nil, fmt.Errorf("pubkey to byte error: %w", err)
	}

	return pubAddr.String(), pubData, nil
}

// pkData -> ecdsa.PublicKey
func GetECDSAPubKeyFromPubKeyData(pkData []byte) (*ecdsa.PublicKey, error) {
	if len(pkData) == 0 {
		return nil, fmt.Errorf("empty public key raw bytes")
	}

	ppk, err := crypto.UnmarshalSecp256k1PublicKey(pkData)
	if err != nil {
		return nil, fmt.Errorf("fail to convert pubkey to the ec pubkey used in Spec256k1: %w", err)
	}

	k1Pk := ppk.(*crypto.Secp256k1PublicKey)
	ecdsaPk := (*ecdsa.PublicKey)(k1Pk)

	return ecdsaPk, nil
}

// CheckKeyOnCurve check if a SECP256K1 public key is on the elliptic curve
func CheckKeyOnCurve(pkData []byte) (bool, error) {
	bPk, err := btcec.ParsePubKey(pkData, btcec.S256())
	if err != nil {
		return false, err
	}
	return isOnCurve(bPk.X, bPk.Y), nil
}

func isOnCurve(x, y *big.Int) bool {
	curve := btcec.S256()
	return curve.IsOnCurve(x, y)
}

// party p2p pubkey: "github.com/libp2p/go-libp2p-core/crypto".pubkey ==================================================

// pubkey -> peer.ID
func GetPIDFromPubKey(pk crypto.PubKey) (peer.ID, error) {
	return peer.IDFromPublicKey(pk)
}

// pkData -> pubkey
func GetPubKeyFromPubKeyData(pk []byte) (crypto.PubKey, error) {
	if len(pk) == 0 {
		return nil, fmt.Errorf("empty public key raw bytes")
	}
	ppk, err := crypto.UnmarshalECDSAPublicKey(pk)
	if err != nil {
		return nil, fmt.Errorf("fail to convert pubkey to the crypto pubkey used in libp2p: %w", err)
	}
	return ppk, nil
}

// []pkData -> []pubkey
func GetPubKeysFromPubKeyDatas(pksData [][]byte) ([]crypto.PubKey, error) {
	pks := []crypto.PubKey{}
	for _, data := range pksData {
		pk, err := GetPubKeyFromPubKeyData(data)
		if err != nil {
			return nil, fmt.Errorf("fail to convert pubkeyData to pubkey: %w", err)
		}
		pks = append(pks, pk)
	}

	return pks, nil
}

// pkData -> peer.ID
func GetPIDFromPubKeyData(pk []byte) (peer.ID, error) {
	if len(pk) == 0 {
		return "", fmt.Errorf("empty public key raw bytes")
	}
	ppk, err := crypto.UnmarshalECDSAPublicKey(pk)
	if err != nil {
		return "", fmt.Errorf("fail to convert pubkey to the crypto pubkey used in libp2p: %w", err)
	}
	return peer.IDFromPublicKey(ppk)
}

// []pkData -> []peer.ID
func GetPIDsFromPubKeys(pks [][]byte) ([]peer.ID, error) {
	var peerIDs []peer.ID
	for _, item := range pks {
		peerID, err := GetPIDFromPubKeyData(item)
		if err != nil {
			return nil, err
		}
		peerIDs = append(peerIDs, peerID)
	}
	return peerIDs, nil
}

func SortPubKey(keys []crypto.PubKey) {
	sort.SliceStable(keys, func(i, j int) bool {
		keyDataI, _ := keys[i].Raw()
		keyDataJ, _ := keys[j].Raw()
		ma, _ := MsgToHashInt(keyDataI)
		mb, _ := MsgToHashInt(keyDataJ)
		if ma.Cmp(mb) == -1 {
			return false
		}
		return true
	})
}
