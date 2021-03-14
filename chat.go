package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ChatMessagePacket -output=chat_gen.go

const ChatMessagePacketID = 0x0f

// https://wiki.vg/Protocol#Chat_Message_.28clientbound.29
type ChatMessagePacket struct {
	Message  TextComponent
	Position enc.Byte
}

func (c *ChatMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ChatMessagePacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}

func (c *ChatMessagePacket) String() string {
	return "ChatMessage"
}
