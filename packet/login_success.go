package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginSuccessPacket -output=login_success_gen.go

const LoginSuccessPacketID = 0x02

// LoginSuccessPacket is sent by the server
// when the login was successful
// https://wiki.vg/Protocol?oldid=16067#Login_Success
type LoginSuccessPacket struct {
	UUID     enc.UUID
	Username enc.String
}

func (*LoginSuccessPacket) Info() Info {
	return Info{
		ID:              LoginSuccessPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateLogin,
	}
}

func (*LoginSuccessPacket) String() string {
	return "LoginSuccess"
}
