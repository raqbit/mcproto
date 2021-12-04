package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginStartPacket -output=login_start_gen.go

const LoginStartPacketID int32 = 0x00

// LoginStartPacket is sent by the client to start logging in
// https://wiki.vg/Protocol?oldid=16067#Login_Start
type LoginStartPacket struct {
	Name enc.String
}

func (l *LoginStartPacket) ID() int32 {
	return LoginStartPacketID
}

func (l *LoginStartPacket) Direction() Direction {
	return ServerBound
}

func (l *LoginStartPacket) State() game.ConnectionState {
	return game.LoginState
}

func (*LoginStartPacket) Info() Info {
	return Info{
		ID:              LoginStartPacketID,
		Direction:       ServerBound,
		ConnectionState: game.LoginState,
	}
}

func (*LoginStartPacket) String() string {
	return "LoginStart"
}
