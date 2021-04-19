package types

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

//go:generate stringer -type=ConnectionState -output state_string.go -linecomment

// ConnectionState is the state of a Minecraft connection
type ConnectionState int32

const (
	ConnectionStateHandshake ConnectionState = 0x00 // Handshake
	ConnectionStateStatus    ConnectionState = 0x01 // Status
	ConnectionStateLogin     ConnectionState = 0x02 // Login
	ConnectionStatePlay      ConnectionState = 0x03 // Play
)

func (c *ConnectionState) Write(w io.Writer) error {
	num := enc.VarInt(*c)
	return num.Write(w)
}

func (c *ConnectionState) Read(r io.Reader) error {
	var num enc.VarInt

	if err := num.Read(r); err != nil {
		return err
	}

	*c = ConnectionState(num)
	return nil
}
