package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=PlayerPositionLookPacket -output=player_position_look_gen.go

const PlayerPositionLookPacketID int32 = 0x36

// PlayerPositionLookPacket is sent by the server to inform the client
// of its location
// https://wiki.vg/Protocol?oldid=16067#Player_Position_And_Look_.28clientbound.29
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

func (*PlayerPositionLookPacket) Info() Info {
	return Info{
		ID:              PlayerPositionLookPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStatePlay,
	}
}
