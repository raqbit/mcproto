package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=HandshakePacket -output=handshake_gen.go

const HandshakePacketID int32 = 0x00

// https://wiki.vg/Protocol#Handshake
type HandshakePacket struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  types.ConnectionState
}

func (h *HandshakePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              HandshakePacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateHandshake,
	}
}

func (*HandshakePacket) String() string {
	return "Handshake"
}
