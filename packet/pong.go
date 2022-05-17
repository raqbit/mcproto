package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=PongPacket -output=pong_gen.go

const PongPacketID int32 = 0x01

// PongPacket is sent by the server as a response to a PingPacket.
// https://wiki.vg/Protocol?oldid=16067#Pong
type PongPacket struct {
	Payload enc.Long
}

func (p *PongPacket) ID() int32 {
	return PongPacketID
}

func (p *PongPacket) Direction() Direction {
	return ClientBound
}

func (p *PongPacket) State() game.ConnectionState {
	return game.StatusState
}

func (*PongPacket) String() string {
	return "Pong"
}
