package conversion

import (
	"strings"

	"github.com/meshplus/bitxhub-core/tss/message"
)

// GetBroadcastMessageType gets the broadcast message type of the message
func GetBroadcastMessageType(msgType message.TssMsgType) message.TssMsgType {
	switch msgType {
	case message.TSSKeyGenMsg:
		return message.TSSKeyGenVerMsg
	case message.TSSKeySignMsg:
		return message.TSSKeySignVerMsg
	default:
		return message.TssUnknown // this should not happen
	}
}

// GetPreviousKeySignUicast eets the last unicast message for the current message
func GetPreviousKeySignUicast(current string) string {
	if strings.HasSuffix(current, message.KEYSIGN1b) {
		return message.KEYSIGN1aUnicast
	}
	return message.KEYSIGN2Unicast
}
