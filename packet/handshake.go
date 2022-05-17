package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=HandshakePacket -output=handshake_gen.go

const HandshakePacketID int32 = 0x00

// HandshakePacket is sent by the client to initiate
// a handshake by changing the connection state
// https://wiki.vg/Protocol?oldid=16067#Handshake
type HandshakePacket struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  game.ConnectionState
}

func (h *HandshakePacket) ID() int32 {
	return HandshakePacketID
}

func (h *HandshakePacket) Direction() Direction {
	return ServerBound
}

func (h *HandshakePacket) State() game.ConnectionState {
	return game.HandshakeState
}

func (*HandshakePacket) String() string {
	return "Handshake"
}
