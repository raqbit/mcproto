package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=JoinGamePacket -output=join_game_gen.go

const JoinGamePacketID int32 = 0x26

// JoinGamePacket is sent by the server to inform the client
// of various game & world parameters
// https://wiki.vg/Protocol?oldid=16067#Join_Game
type JoinGamePacket struct {
	PlayerID            enc.Int
	GameMode            types.Gamemode
	Dimension           types.Dimension
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

func (*JoinGamePacket) Info() Info {
	return Info{
		ID:              JoinGamePacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
