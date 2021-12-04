package mcproto

import (
	"github.com/Raqbit/mcproto/game"
	"github.com/Raqbit/mcproto/packet"
)

// Option is a configuration option for mcproto.
type Option func(conn *connection)

// WithLogger configures the logger for
// mcproto.
func WithLogger(logger Logger) Option {
	return func(conn *connection) {
		conn.logger = logger
	}
}

// WithPackets allows you to add custom
// packets to mcproto's packet registry.
// This will make mcproto automatically parse
// the packet when it is received.
func WithPackets(packets []packet.Packet) Option {
	return func(conn *connection) {
		for _, pkt := range packets {
			conn.registerPacket(pkt)
		}
	}
}

// WithSide configures the connection side
// of the mcproto Connection.
func WithSide(side Side) Option {
	return func(conn *connection) {
		conn.side = side
	}
}

func WithState(state game.ConnectionState) Option {
	return func(conn *connection) {
		conn.state = state
	}
}
