package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ClientSettingsPacket -output=client_settings_gen.go

const ClientSettingsPacketID int32 = 0x05

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

func (*ClientSettingsPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ClientSettingsPacketID,
		Direction:       types.ServerBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
