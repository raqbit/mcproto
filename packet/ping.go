package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=PingPacket -output=ping_gen.go

const PingPacketID = 0x01

// PingPacket is sent by the client to get a PongPacket response.
// https://wiki.vg/Protocol?oldid=16067#Ping
type PingPacket struct {
	Payload enc.Long
}

func (p *PingPacket) ID() int32 {
	return PingPacketID
}

func (p *PingPacket) Direction() Direction {
	return ServerBound
}

func (p *PingPacket) State() game.ConnectionState {
	return game.StatusState
}

func (*PingPacket) String() string {
	return "Ping"
}
