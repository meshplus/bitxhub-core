package message

import (
	"fmt"

	"github.com/meshplus/bitxhub-model/pb"

	btss "github.com/binance-chain/tss-lib/tss"
)

type SendMsgChan struct {
	P2PMsg    *pb.Message
	PartiesID []uint64
}

// pb.Message_TSS_KEY_GEN, pb.Message_TSS_KEY_SIGN
// WireMessage the message that produced by tss-lib package
type WireMessage struct {
	Routing   *btss.MessageRouting `json:"routing"`
	RoundInfo string               `json:"round_info"`
	Message   []byte               `json:"message"`
	Sig       []byte               `json:"signature"`
}

// pb.Message_TSS_KEY_GEN_VER, pb.Message_TSS_KEY_SIGN_VER
// BroadcastConfirmMessage is used to broadcast to all parties what message they receive
type BroadcastConfirmMessage struct {
	FromID string `json:"from_id"`
	Key    string `json:"key"`
	Hash   string `json:"hash"`
}

// pb.Message_TSS_CONTROL
type TssControl struct {
	FromID      string          `json:"from_id"`
	ReqHash     string          `json:"reqest_hash"`
	ReqKey      string          `json:"request_key"`
	RequestType pb.Message_Type `json:"request_type"`
	Msg         *WireMessage    `json:"message_body"`
}

// pb.Message_TSS_TASK_DONE
type TssTaskNotifier struct {
	FromID   string `json:"from_id"`
	TaskDone bool   `json:"task_done"`
}

// GetCacheKey return the key we used to cache it locally
func (m *WireMessage) GetCacheKey() string {
	return fmt.Sprintf("%s-%s", m.Routing.From.Id, m.RoundInfo)
}
