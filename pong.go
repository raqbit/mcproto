package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PongPacket -output=pong_gen.go

const PongPacketID int32 = 0x01

// https://wiki.vg/Protocol#Pong
type PongPacket struct {
	Payload enc.Long
}

func (*PongPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PongPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStateStatus,
	}
}

func (*PongPacket) String() string {
	return "Pong"
}
