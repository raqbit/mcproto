package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ServerInfoPacket -output=server_info_gen.go

const ServerInfoPacketID int32 = 0x00

// ServerInfoPacket is sent by the server as a response to a ServerQueryPacket.
// https://wiki.vg/Protocol?oldid=16067#Response
type ServerInfoPacket struct {
	Response types.ServerInfo
}

func (*ServerInfoPacket) Info() Info {
	return Info{
		ID:              ServerInfoPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateStatus,
	}
}

func (*ServerInfoPacket) String() string {
	return "ServerInfo"
}
