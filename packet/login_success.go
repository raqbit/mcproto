package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket -packet=LoginSuccessPacket -output=login_success_gen.go

const LoginSuccessPacketID = 0x02

// LoginSuccessPacket is sent by the server
// when the login was successful
// https://wiki.vg/Protocol?oldid=16067#Login_Success
type LoginSuccessPacket struct {
	UUID     enc.UUID
	Username enc.String
}

func (l *LoginSuccessPacket) ID() int32 {
	return LoginSuccessPacketID
}

func (*LoginSuccessPacket) String() string {
	return "LoginSuccess"
}
