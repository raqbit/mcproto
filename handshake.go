package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=HandshakePacket -output=handshake_gen.go

const HandshakePacketID int32 = 0x00

// https://wiki.vg/Protocol#Handshake
type HandshakePacket struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  ConnectionState
}

func (h *HandshakePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              HandshakePacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateHandshake,
	}
}

func (*HandshakePacket) String() string {
	return "Handshake"
}
