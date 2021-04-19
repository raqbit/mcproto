package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PluginMessagePacket -output=plugin_message_gen.go

const PluginMessagePacketID int32 = 0x19

// PluginMessagePacket is sent by the client
// as part of a plugin's message channel
// https://wiki.vg/Protocol?oldid=16067#Plugin_Message_.28clientbound.29
type PluginMessagePacket struct {
	Channel types.Identifier
	Data    Encoding
}

func (*PluginMessagePacket) String() string {
	return "PluginMessage"
}

func (*PluginMessagePacket) Info() Info {
	return Info{
		ID:              PluginMessagePacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
