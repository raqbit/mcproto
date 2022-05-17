package packet

import (
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=PluginMessagePacket -output=plugin_message_gen.go

const PluginMessagePacketID int32 = 0x19

// PluginMessagePacket is sent by the client
// as part of a plugin's message channel
// https://wiki.vg/Protocol?oldid=16067#Plugin_Message_.28clientbound.29
type PluginMessagePacket struct {
	Channel game.Identifier
	Data    Encoding
}

func (p *PluginMessagePacket) ID() int32 {
	return PluginMessagePacketID
}

func (p *PluginMessagePacket) Direction() Direction {
	return ClientBound
}

func (p *PluginMessagePacket) State() game.ConnectionState {
	return game.PlayState
}

func (*PluginMessagePacket) String() string {
	return "PluginMessage"
}
