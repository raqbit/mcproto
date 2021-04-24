package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ServerQueryPacket -output=server_query_gen.go

const ServerQueryPacketID int32 = 0x00

// ServerQueryPacket is sent by the client to query the Minecraft server
// for protocol version, message of the day and online player information.
// https://wiki.vg/Protocol?oldid=16067#Request
type ServerQueryPacket struct{}

func (r ServerQueryPacket) Info() Info {
	return Info{
		ID:              ServerQueryPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStateStatus,
	}
}

func (*ServerQueryPacket) String() string {
	return "ServerQuery"
}
