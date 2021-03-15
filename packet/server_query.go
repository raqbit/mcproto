package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ServerQueryPacket -output=server_query_gen.go

const ServerQueryPacketID int32 = 0x00

// https://wiki.vg/Protocol#Request
type ServerQueryPacket struct{}

func (r ServerQueryPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ServerQueryPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateStatus,
	}
}

func (*ServerQueryPacket) String() string {
	return "ServerQuery"
}
