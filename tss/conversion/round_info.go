package conversion

import (
	"fmt"
	"strings"

	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/meshplus/bitxhub-core/tss/message"
)

type RoundInfo struct {
	Index         int
	RoundMsg      string
	MsgIdentifier string
}

func GetMsgRound(msg []byte, partyID *btss.PartyID, isBroadcast bool) (RoundInfo, error) {
	parsedMsg, err := btss.ParseWireMessage(msg, partyID, isBroadcast)
	if err != nil {
		return RoundInfo{}, err
	}
	switch parsedMsg.Content().(type) {
	case *keygen.KGRound1Message:
		return RoundInfo{
			Index:    0,
			RoundMsg: message.KEYGEN1,
		}, nil

	case *keygen.KGRound2Message1:
		return RoundInfo{
			Index:    1,
			RoundMsg: message.KEYGEN2aUnicast,
		}, nil

	case *keygen.KGRound2Message2:
		return RoundInfo{
			Index:    2,
			RoundMsg: message.KEYGEN2b,
		}, nil

	case *keygen.KGRound3Message:
		return RoundInfo{
			Index:    3,
			RoundMsg: message.KEYGEN3,
		}, nil

	case *signing.SignRound1Message1:
		return RoundInfo{
			Index:    0,
			RoundMsg: message.KEYSIGN1aUnicast,
		}, nil

	case *signing.SignRound1Message2:
		return RoundInfo{
			Index:    1,
			RoundMsg: message.KEYSIGN1b,
		}, nil

	case *signing.SignRound2Message:
		return RoundInfo{
			Index:    2,
			RoundMsg: message.KEYSIGN2Unicast,
		}, nil

	case *signing.SignRound3Message:
		return RoundInfo{
			Index:    3,
			RoundMsg: message.KEYSIGN3,
		}, nil

	case *signing.SignRound4Message:
		return RoundInfo{
			Index:    4,
			RoundMsg: message.KEYSIGN4,
		}, nil

	case *signing.SignRound5Message:
		return RoundInfo{
			Index:    5,
			RoundMsg: message.KEYSIGN5,
		}, nil

	case *signing.SignRound6Message:
		return RoundInfo{
			Index:    6,
			RoundMsg: message.KEYSIGN6,
		}, nil

	case *signing.SignRound7Message:
		return RoundInfo{
			Index:    7,
			RoundMsg: message.KEYSIGN7,
		}, nil

	default:
		return RoundInfo{}, fmt.Errorf("unknown round")
	}
}

// due to the nature of tss, we may find the invalid share of the previous round only
// when we get the shares from the peers in the current round. So, when we identify
// an error in this round, we check whether the previous round is the unicast
func CheckUnicast(round RoundInfo) bool {
	index := round.Index
	isKeyGen := strings.Contains(round.RoundMsg, "KGR")
	// keygen unicast blame
	if isKeyGen {
		if index == 1 || index == 2 {
			return true
		}
		return false
	}
	// keysign unicast blame
	if index < 5 {
		return true
	}
	return false
}
