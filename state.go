package mcproto

import enc "github.com/Raqbit/mcproto/encoding"

//go:generate stringer -type=ConnectionState -output state_string.go -linecomment

// ConnectionState is the state of a Minecraft connection
type ConnectionState int32

const (
	HandshakeState ConnectionState = 0x00 // Handshake
	StatusState    ConnectionState = 0x01 // Status
	LoginState     ConnectionState = 0x02 // Login
	PlayState      ConnectionState = 0x03 // Play
)

func (s ConnectionState) ID() enc.VarInt {
	return enc.VarInt(s)
}

func StateFromID(varInt enc.VarInt) ConnectionState {
	return ConnectionState(varInt)
}
