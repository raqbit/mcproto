package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket -packet=JoinGamePacket -output=join_game_gen.go

const JoinGamePacketID int32 = 0x26

type JoinGamePacket struct {
	PlayerID         enc.Int
	IsHardcore       enc.Bool
	GameMode         enc.UnsignedByte
	PreviousGameMode enc.Byte
	WorldCount       enc.VarInt
	WorldNames       []types.Identifier `pkt:"lenFrom(WorldCount)"`
	//DimensionCodec nbt.Something // TODO
	//Dimension nbt.Something // TODO
	WorldName           types.Identifier
	HashedSeed          enc.Long
	MaxPlayers          enc.UnsignedByte
	ViewDistance        enc.VarInt
	ReducedDebugInfo    enc.Bool
	EnableRespawnScreen enc.Bool
	IsDebug             enc.Bool
	IsFlat              enc.Bool
}

func (*JoinGamePacket) String() string {
	return "JoinGame"
}

func (*JoinGamePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              JoinGamePacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
