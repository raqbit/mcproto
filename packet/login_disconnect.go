package packet

import (
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../tools/genpacket -packet=LoginDisconnectPacket -output=login_disconnect_gen.go

const LoginDisconnectPacketID int32 = 0x00

// LoginDisconnectPacket is sent by the server when
// the user is disconnected while they are logging in (e.g. banned)
// https://wiki.vg/Protocol?oldid=16067#Disconnect_.28login.29
type LoginDisconnectPacket struct {
	Reason game.TextComponent
}

func (l *LoginDisconnectPacket) ID() int32 {
	return LoginDisconnectPacketID
}

func (l *LoginDisconnectPacket) Direction() Direction {
	return ClientBound
}

func (l *LoginDisconnectPacket) State() game.ConnectionState {
	return game.LoginState
}

func (*LoginDisconnectPacket) String() string {
	return "LoginDisconnect"
}
