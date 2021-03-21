package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginStartPacket -output=login_start_gen.go

const LoginStartPacketID int32 = 0x00

// https://wiki.vg/Protocol#Login_Start
type LoginStartPacket struct {
	Name enc.String `pkt:"strLen(16)"`
}

func (*LoginStartPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginStartPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateLogin,
	}
}

func (*LoginStartPacket) String() string {
	return "LoginStart"
}
