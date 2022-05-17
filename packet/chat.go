package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=ChatMessagePacket -output=chat_gen.go

const ChatMessagePacketID = 0x0f

// ChatMessagePacket is sent by the server for incoming chat messages
// https://wiki.vg/Protocol?oldid=16067#Chat_Message_.28clientbound.29
type ChatMessagePacket struct {
	Message  game.TextComponent
	Position enc.Byte
}

func (c *ChatMessagePacket) ID() int32 {
	return ChatMessagePacketID
}

func (c *ChatMessagePacket) Direction() Direction {
	return ClientBound
}

func (c *ChatMessagePacket) State() game.ConnectionState {
	return game.PlayState
}

func (c *ChatMessagePacket) String() string {
	return "ChatMessage"
}
