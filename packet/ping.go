package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PingPacket -output=ping_gen.go

const PingPacketID = 0x01

// PingPacket is sent by the client to get a PongPacket response.
// https://wiki.vg/Protocol?oldid=16067#Ping
type PingPacket struct {
	Payload enc.Long
}

func (*PingPacket) Info() Info {
	return Info{
		ID:              PingPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateStatus,
	}
}

func (*PingPacket) String() string {
	return "Ping"
}
