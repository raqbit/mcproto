package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginDisconnectPacket -output=login_disconnect_gen.go

const LoginDisconnectPacketID int32 = 0x00

// LoginDisconnectPacket is sent by the server when
// the user is disconnected while they are logging in (e.g. banned)
// https://wiki.vg/Protocol?oldid=16067#Disconnect_.28login.29
type LoginDisconnectPacket struct {
	Reason types.TextComponent
}

func (*LoginDisconnectPacket) Info() Info {
	return Info{
		ID:              LoginDisconnectPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateLogin,
	}
}

func (*LoginDisconnectPacket) String() string {
	return "LoginDisconnect"
}
