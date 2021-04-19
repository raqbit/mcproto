package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=HandshakePacket -output=handshake_gen.go

const HandshakePacketID int32 = 0x00

// HandshakePacket is sent by the client to initiate
// a handshake by changing the connection state
// https://wiki.vg/Protocol?oldid=16067#Handshake
type HandshakePacket struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  types.ConnectionState
}

func (h *HandshakePacket) Info() Info {
	return Info{
		ID:              HandshakePacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateHandshake,
	}
}

func (*HandshakePacket) String() string {
	return "Handshake"
}
