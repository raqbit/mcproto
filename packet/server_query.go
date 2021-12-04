package packet

import "github.com/Raqbit/mcproto/game"

//go:generate go run ../tools/genpacket/genpacket.go -packet=ServerQueryPacket -output=server_query_gen.go

const ServerQueryPacketID int32 = 0x00

// ServerQueryPacket is sent by the client to query the Minecraft server
// for protocol version, message of the day and online player information.
// https://wiki.vg/Protocol?oldid=16067#Request
type ServerQueryPacket struct{}

func (s ServerQueryPacket) ID() int32 {
	return ServerQueryPacketID
}

func (s ServerQueryPacket) Direction() Direction {
	return ServerBound
}

func (s ServerQueryPacket) State() game.ConnectionState {
	return game.StatusState
}

func (*ServerQueryPacket) String() string {
	return "ServerQuery"
}
