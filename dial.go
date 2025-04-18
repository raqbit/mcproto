package mcproto

import (
	"bytes"
	"context"
	"github.com/Raqbit/mcproto/packet"
	"net"
)

// Dial connects to the specified address and creates
// a new Connection for it. Returns the connection &
// the address that was connected to.
//
// See DialContext for more information.
func Dial(address string, opts ...Option) (Connection, string, error) {
	return DialContext(context.Background(), address, opts...)
}

// DialContext creates a TCP connection with specified address
// and creates a new Connection for it. If no port is specified,
// the Minecraft default of 25565 will be used.
//
// Just like the vanilla Minecraft client, DialContext will do an SRV
// DNS lookup before connecting to a Minecraft server with the default port.
// If an SRV record is found, the contained target and port will be used
// to connect instead.
//
// The provided Context must be non-nil. If the context expires before
// the connection is established, an error is returned. Once successfully
// connected, any expiration of this context will not affect the connection.
//
// The given options will be applied to the created connection.
//
// Will return the created Connection and the resolved address that
// was connected to.
func DialContext(ctx context.Context, address string, opts ...Option) (Connection, string, error) {
	var dialer net.Dialer

	if ctx == nil {
		panic("provided nil context")
	}

	// TODO: separate add timeout to SRV lookup
	// Resolve server address from host & port, by checking SRV record & joining host & port
	resolvedAddress := resolveServerAddress(ctx, address)

	// Make TCP connection
	tcpConn, err := dialer.DialContext(ctx, "tcp", resolvedAddress)

	if err != nil {
		return nil, resolvedAddress, err
	}

	// Combine default options with given options
	opts = append([]Option{
		WithSide(ClientSide),
		WithState(HandshakeState),
	}, opts...)

	// Wrap TCP connection in a wrapper for sending/receiving Minecraft packets
	return Wrap(tcpConn, opts...), resolvedAddress, nil
}

// Wrap wraps the given connection with a Connection,
// so it can be used for sending/receiving Minecraft packets.
//
// Wrap defaults its side to ServerSide, which means it will
// interpret incoming packets as being server-bound.
// If you are using Wrap for the client side, use
// the WithSide option to configure this. Wrap will also
// default the connection state to packet.HandshakeState.
// If you are wrapping a connection which is in a different
// state, use the WithState option to configure this.
//
// The given options will be applied to the created connection.
func Wrap(conn net.Conn, opts ...Option) Connection {
	mcConn := &connection{
		conn:     conn,
		side:     ServerSide,
		state:    HandshakeState,
		packets:  make(map[packetKey]Packet),
		readBuf:  new(bytes.Buffer),
		writeBuf: new(bytes.Buffer),
		logger:   noopLogger{},
	}

	// Server bound packets
	mcConn.registerPacket(new(packet.HandshakePacket), serverBound, HandshakeState)
	mcConn.registerPacket(new(packet.ServerQueryPacket), serverBound, StatusState)
	mcConn.registerPacket(new(packet.PingPacket), serverBound, StatusState)
	mcConn.registerPacket(new(packet.LoginStartPacket), serverBound, LoginState)
	mcConn.registerPacket(new(packet.ClientSettingsPacket), serverBound, PlayState)

	// Client bound packets
	mcConn.registerPacket(new(packet.PongPacket), clientBound, StatusState)
	mcConn.registerPacket(new(packet.LoginSuccessPacket), clientBound, LoginState)
	mcConn.registerPacket(new(packet.ChatMessagePacket), clientBound, PlayState)
	mcConn.registerPacket(new(packet.ServerInfoPacket), clientBound, StatusState)
	mcConn.registerPacket(new(packet.LoginDisconnectPacket), clientBound, LoginState)
	mcConn.registerPacket(new(packet.JoinGamePacket), clientBound, PlayState)
	mcConn.registerPacket(new(packet.PlayerPositionLookPacket), clientBound, PlayState)
	mcConn.registerPacket(new(packet.PluginMessagePacket), clientBound, PlayState)

	// Apply options
	for _, opt := range opts {
		opt(mcConn)
	}

	return mcConn
}
