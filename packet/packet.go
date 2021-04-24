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
	// Encoding is a type that can be encoded
	Encoding interface {
		Write(w io.Writer) error
		Read(r io.Reader) error
	}

	// Packet is a packet with side-bound ID & direction that can be encoded and sent
	Packet interface {
		fmt.Stringer
		Encoding
		Info() Info
	}

	// Info is info that identifies a packet
	Info struct {
		ID              int32
		Direction       types.PacketDirection
		ConnectionState types.ConnectionState
	}
)
