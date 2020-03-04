package mcproto

import (
	"bytes"
	"errors"
	"fmt"
)

const MaxPacketLength = 1048576

// Packet which has been decoded into a struct.
type Packet interface {
	fmt.Stringer
	ID() int
	Marshal() (*bytes.Buffer, error)
	Unmarshal(*bytes.Buffer) (Packet, error)
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

// All handled types of packet.
// Used for decoding packets by direction, connection state & id.
var PacketTypes = map[PacketDirection]map[ConnectionState]map[int]Packet{
	ServerBound: {
		HandshakeState: {
			HandshakePacket{}.ID(): HandshakePacket{},
		},
		StatusState: {
			RequestPacket{}.ID(): RequestPacket{},
			PingPacket{}.ID():    PingPacket{},
		},
		LoginState: {
			LoginPacket{}.ID(): LoginPacket{},
		},
	},
	ClientBound: {
		StatusState: {
			ResponsePacket{}.ID(): ResponsePacket{},
			PongPacket{}.ID():     PongPacket{},
		},
		LoginState: {
			DisconnectPacket{}.ID():   DisconnectPacket{},
			LoginSuccessPacket{}.ID(): LoginSuccessPacket{},
		},
	},
}
