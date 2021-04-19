package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
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
	MainHand           types.Hand
}

func (*ClientSettingsPacket) String() string {
	return "ClientSettings"
}

func (*ClientSettingsPacket) Info() Info {
	return Info{
		ID:              ClientSettingsPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
