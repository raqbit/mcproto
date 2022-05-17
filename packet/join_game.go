package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=JoinGamePacket -output=join_game_gen.go

const JoinGamePacketID int32 = 0x26

// JoinGamePacket is sent by the server to inform the client
// of various game & world parameters
// https://wiki.vg/Protocol?oldid=16067#Join_Game
type JoinGamePacket struct {
	PlayerID            enc.Int
	GameMode            game.Mode
	Dimension           game.Dimension
	HashedSeed          enc.Long
	MaxPlayers          enc.UnsignedByte
	LevelType           enc.String `len:"16"`
	ViewDistance        enc.VarInt
	ReducedDebug        enc.Bool
	EnableRespawnScreen enc.Bool
}

func (j *JoinGamePacket) ID() int32 {
	return JoinGamePacketID
}

func (j *JoinGamePacket) Direction() Direction {
	return ClientBound
}

func (j *JoinGamePacket) State() game.ConnectionState {
	return game.PlayState
}

func (*JoinGamePacket) String() string {
	return "JoinGame"
}
