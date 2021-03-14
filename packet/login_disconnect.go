package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginDisconnectPacket -output=login_disconnect_gen.go

const LoginDisconnectPacketID int32 = 0x00

// https://wiki.vg/Protocol#Disconnect_.28login.29
type LoginDisconnectPacket struct {
	Reason types.TextComponent
}

func (*LoginDisconnectPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginDisconnectPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateLogin,
	}
}

func (*LoginDisconnectPacket) String() string {
	return "LoginDisconnect"
}
