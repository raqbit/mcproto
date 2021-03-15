package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginSuccessPacket -output=login_success_gen.go

const LoginSuccessPacketID = 0x02

// https://wiki.vg/Protocol#Login_Success
type LoginSuccessPacket struct {
	UUID     enc.UUID
	Username enc.String
}

func (*LoginSuccessPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginSuccessPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateLogin,
	}
}

func (*LoginSuccessPacket) String() string {
	return "LoginSuccess"
}
