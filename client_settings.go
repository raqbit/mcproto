package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ClientSettingsPacket -output=client_settings_gen.go

const ClientSettingsPacketID int32 = 0x05

type ClientSettingsPacket struct {
	Lang               enc.String
	ViewDistance       enc.Byte
	ChatVisibility     enc.VarInt
	EnableChatColors   enc.Bool
	DisplayedSkinParts enc.UnsignedByte
	MainHand           Hand
}

func (*ClientSettingsPacket) String() string {
	return "ClientSettings"
}

func (*ClientSettingsPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ClientSettingsPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStatePlay,
	}
}
