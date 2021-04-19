package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ChatMessagePacket -output=chat_gen.go

const ChatMessagePacketID = 0x0f

// ChatMessagePacket is sent by the server for incoming chat messages
// https://wiki.vg/Protocol?oldid=16067#Chat_Message_.28clientbound.29
type ChatMessagePacket struct {
	Message  types.TextComponent
	Position enc.Byte
}

func (c *ChatMessagePacket) Info() Info {
	return Info{
		ID:              ChatMessagePacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}

func (c *ChatMessagePacket) String() string {
	return "ChatMessage"
}
