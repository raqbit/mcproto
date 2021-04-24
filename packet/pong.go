package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PongPacket -output=pong_gen.go

const PongPacketID int32 = 0x01

// PongPacket is sent by the server as a response to a PingPacket.
// https://wiki.vg/Protocol?oldid=16067#Pong
type PongPacket struct {
	Payload enc.Long
}

func (*PongPacket) Info() Info {
	return Info{
		ID:              PongPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateStatus,
	}
}

func (*PongPacket) String() string {
	return "Pong"
}
