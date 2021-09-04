package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ClientSettingsPacket -output=client_settings_gen.go

const ClientSettingsPacketID int32 = 0x05

// ClientSettingsPacket is sent by the client to inform the server
// of the (updated) client settings
// https://wiki.vg/Protocol?oldid=16067#Client_Settings
type ClientSettingsPacket struct {
	Lang               enc.String
	ViewDistance       enc.Byte
	ChatVisibility     enc.VarInt
	EnableChatColors   enc.Bool
	DisplayedSkinParts enc.UnsignedByte
	MainHand           game.Hand
}

func (c *ClientSettingsPacket) ID() int32 {
	return ClientSettingsPacketID
}

func (c *ClientSettingsPacket) Direction() Direction {
	return ServerBound
}

func (c *ClientSettingsPacket) State() game.ConnectionState {
	return game.PlayState
}

func (*ClientSettingsPacket) String() string {
	return "ClientSettings"
}
