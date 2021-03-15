package packet

import (
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PluginMessagePacket -output=plugin_message_gen.go

const PluginMessagePacketID int32 = 0x19

type PluginMessagePacket struct {
	Channel types.Identifier
	Data    PacketData
}

func (*PluginMessagePacket) String() string {
	return "PluginMessage"
}

func (*PluginMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PluginMessagePacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
