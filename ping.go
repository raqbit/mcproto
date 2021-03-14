package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PingPacket -output=ping_gen.go

const PingPacketID = 0x01

// https://wiki.vg/Protocol#Ping
type PingPacket struct {
	Payload enc.Long
}

func (*PingPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PingPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateStatus,
	}
}

func (*PingPacket) String() string {
	return "Ping"
}
