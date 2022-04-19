package conversion

import (
	"strings"

	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/meshplus/bitxhub-model/pb"
)

// GetBroadcastMessageType gets the broadcast message type of the message
func GetBroadcastMessageType(msgType pb.Message_Type) pb.Message_Type {
	switch msgType {
	case pb.Message_TSS_KEY_GEN:
		return pb.Message_TSS_KEY_GEN_VER
	case pb.Message_TSS_KEY_SIGN:
		return pb.Message_TSS_KEY_SIGN_VER
	default:
		return pb.Message_TSS_UNKNOW // this should not happen
	}
}

// GetPreviousKeySignUicast eets the last unicast message for the current message
func GetPreviousKeySignUicast(current string) string {
	if strings.HasSuffix(current, message.KEYSIGN1b) {
		return message.KEYSIGN1aUnicast
	}
	return message.KEYSIGN2Unicast
}
