package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=JoinGamePacket -output=join_game_gen.go

const JoinGamePacketID int32 = 0x26

type JoinGamePacket struct {
	PlayerID            enc.Int
	GameMode            enc.UnsignedByte
	Dimension           enc.Int
	HashedSeed          enc.Long
	MaxPlayers          enc.UnsignedByte
	LevelType           enc.String `len:"16"`
	ViewDistance        enc.VarInt
	ReducedDebug        enc.Bool
	EnableRespawnScreen enc.Bool
}

func (*JoinGamePacket) String() string {
	return "JoinGame"
}

func (*JoinGamePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              JoinGamePacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}
