package packet

import (
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ServerInfoPacket -output=server_info_gen.go

const ServerInfoPacketID int32 = 0x00

// ServerInfoPacket is sent by the server as a response to a ServerQueryPacket.
// https://wiki.vg/Protocol?oldid=16067#Response
type ServerInfoPacket struct {
	Response game.ServerInfo
}

func (s *ServerInfoPacket) ID() int32 {
	return ServerInfoPacketID
}

func (s *ServerInfoPacket) Direction() Direction {
	return ClientBound
}

func (s *ServerInfoPacket) State() game.ConnectionState {
	return game.StatusState
}

func (*ServerInfoPacket) String() string {
	return "ServerInfo"
}
