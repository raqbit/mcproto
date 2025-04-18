package mcproto

// Option is a configuration option for mcproto.
type Option func(conn *connection)

// WithLogger configures the logger for
// mcproto.
func WithLogger(logger Logger) Option {
	return func(conn *connection) {
		conn.logger = logger
	}
}

// WithPackets allows you to add custom packets to mcproto's packet registry.
// This will make mcproto automatically parse the packet when it is received.
// Because Packet only specifies the ID, the direction and state should also be given
func WithPackets(packets []Packet, direction PacketDirection, state ConnectionState) Option {
	return func(conn *connection) {
		for _, pkt := range packets {
			conn.registerPacket(pkt, direction, state)
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

func WithState(state ConnectionState) Option {
	return func(conn *connection) {
		conn.state = state
	}
}
