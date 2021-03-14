package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginStartPacket -output=login_start_gen.go

const LoginStartPacketID int32 = 0x00

// https://wiki.vg/Protocol#Login_Start
type LoginStartPacket struct {
	Name enc.String
}

func (*LoginStartPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginStartPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateLogin,
	}
}

func (*LoginStartPacket) String() string {
	return "LoginStart"
}
