package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket -packet=LoginStartPacket -output=login_start_gen.go

const LoginStartPacketID int32 = 0x00

// LoginStartPacket is sent by the client to start logging in
// https://wiki.vg/Protocol?oldid=16067#Login_Start
type LoginStartPacket struct {
	Name enc.String
}

func (l *LoginStartPacket) ID() int32 {
	return LoginStartPacketID
}

func (*LoginStartPacket) String() string {
	return "LoginStart"
}
