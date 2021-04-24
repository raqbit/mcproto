package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginStartPacket -output=login_start_gen.go

const LoginStartPacketID int32 = 0x00

// LoginStartPacket is sent by the client to start logging in
// https://wiki.vg/Protocol?oldid=16067#Login_Start
type LoginStartPacket struct {
	Name enc.String
}

func (*LoginStartPacket) Info() Info {
	return Info{
		ID:              LoginStartPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateLogin,
	}
}

func (*LoginStartPacket) String() string {
	return "LoginStart"
}
