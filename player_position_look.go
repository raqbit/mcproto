package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PlayerPositionLookPacket -output=player_position_look_gen.go

const PlayerPositionLookPacketID int32 = 0x36

type PlayerPositionLookPacket struct {
	X          enc.Double
	Y          enc.Double
	Z          enc.Double
	Yaw        enc.Float
	Pitch      enc.Float
	Flags      enc.Byte
	TeleportID enc.VarInt
}

func (*PlayerPositionLookPacket) String() string {
	return "PlayerPositionAndLook"
}

func (*PlayerPositionLookPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PlayerPositionLookPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}
