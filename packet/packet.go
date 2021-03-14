package packet

import (
	"errors"
	"fmt"
	"github.com/Raqbit/mcproto/types"
	"io"
)

const MaxPacketLength = 1048576

var (
	ErrInvalidPacketLength = errors.New("packet has a malformed length")
)

type (
	// A type that can be encoded
	Encoding interface {
		Write(w io.Writer) error
		Read(r io.Reader) error
	}

	PacketData Encoding

	// A packet type
	Packet interface {
		fmt.Stringer
		PacketData
		Info() PacketInfo
	}

	// The info which identifies a packet
	PacketInfo struct {
		ID              int32
		Direction       types.PacketDirection
		ConnectionState types.ConnectionState
	}
)
