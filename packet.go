package mcproto

import (
	"errors"
	"fmt"
	"io"
)

const MaxPacketLength = 1048576

type PacketInfo struct {
	ID              int
	Direction       PacketDirection
	ConnectionState ConnectionState
}

// Packet which has been decoded into a struct.
type Packet interface {
	fmt.Stringer
	Info() PacketInfo
	Marshal() ([]byte, error)
	Unmarshal(io.Reader) (Packet, error)
}

// The direction a packet
type PacketDirection byte

func (p PacketDirection) String() string {
	names := []string{
		"client-bound",
		"server-bound",
	}

	return names[p]
}

const (
	ClientBound PacketDirection = iota
	ServerBound
)

var (
	ErrInvalidPacketLength = errors.New("packet has a malformed length")
)
