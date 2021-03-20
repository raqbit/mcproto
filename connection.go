package mcproto

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/packet"
	"github.com/Raqbit/mcproto/types"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	DefaultMinecraftPort = "25565"
	MinecraftSRVService  = "minecraft"
	MinecraftSRVProtocol = "tcp"
)

var (
	ErrConnectionState  = errors.New("connection has invalid state for packet type")
	ErrDirection        = errors.New("packet being sent in wrong direction")
	ErrLegacyServerPing = errors.New("not implemented: legacy server ping")
)

// Connection represents a connection
// to a Minecraft server or client.
type Connection interface {
	ReadPacket(context.Context) (packet.Packet, error)
	WritePacket(context.Context, packet.Packet) error
	Close() error
	SetState(state types.ConnectionState)
}

type connection struct {
	transport net.Conn
	state     types.ConnectionState
	side      types.Side
	packets   map[packet.PacketInfo]packet.Packet
}

// Dial connects to the specified address without timeout
// and creates a new Connection for it. Returns the connection &
// the address that was connected to.
//
// See DialContext for more information.
func Dial(host string, port string) (Connection, string, error) {
	return DialContext(context.Background(), host, port)
}

// DialContext creates a TCP connection with specified host & port
// and creates a new Connection for it. If the specified port is
// an empty string, the Minecraft default of 25565 will be used.
//
// Just like the vanilla Minecraft client, DialContext will do an SRV
// DNS lookup before connecting to a Minecraft server with the default port.
// If an SRV record is found, the contained target and port will be used
// to connect instead.
//
// The provided Context must be non-nil. If the context expires before
// the connection is complete, an error is returned. Once successfully
// connected, any expiration of this context will not affect the connection.
//
// Will return the created Connection and the resolved address that
// was connected to.
func DialContext(ctx context.Context, host string, port string) (Connection, string, error) {
	var resolver net.Resolver
	var dialer net.Dialer

	// If no port is given, use the default Minecraft port.
	if port == "" {
		port = DefaultMinecraftPort
	}

	// If no port is given or the given port is the default,
	// do a DNS SRV record lookup.
	if port == DefaultMinecraftPort {
		// Do DNS SRV record lookup on given hostname
		_, srvRecords, err := resolver.LookupSRV(ctx, MinecraftSRVService, MinecraftSRVProtocol, host)

		if err == nil && len(srvRecords) > 0 {
			// Override host & port with details from first SRV record returned
			record := srvRecords[0]
			host = record.Target
			port = strconv.Itoa(int(record.Port))
		}
	}

	// Join host & port for connecting to the server but also for returning
	// the resolved server address.
	// Note: If the host was resolved via an SRV record, it will have a
	// trailing period. This is kept so the returned address can be used for
	// a handshake packet, which the vanilla client also sends with a trailing period.
	resolvedAddress := net.JoinHostPort(host, port)

	// Make TCP connection
	tcpConn, err := dialer.DialContext(ctx, "tcp", resolvedAddress)

	if err != nil {
		return nil, resolvedAddress, err
	}

	// Wrap TCP connection in a wrapper for sending/receiving Minecraft packets
	return WrapConnection(tcpConn, types.ClientSide), resolvedAddress, nil
}

// WrapConnection wraps the given connection with a Connection,
// so it can be used for sending/receiving Minecraft packets.
func WrapConnection(transport net.Conn, side types.Side) Connection {
	conn := &connection{
		transport: transport,
		state:     types.ConnectionStateHandshake,
		side:      side,
		packets:   make(map[packet.PacketInfo]packet.Packet),
	}

	// Server bound packets
	conn.registerPacket(&packet.HandshakePacket{})
	conn.registerPacket(&packet.ServerQueryPacket{})
	conn.registerPacket(&packet.PingPacket{})
	conn.registerPacket(&packet.LoginStartPacket{})
	conn.registerPacket(&packet.ClientSettingsPacket{})

	// Client bound packets
	conn.registerPacket(&packet.PongPacket{})
	conn.registerPacket(&packet.LoginSuccessPacket{})
	conn.registerPacket(&packet.ChatMessagePacket{})
	conn.registerPacket(&packet.ServerInfoPacket{})
	conn.registerPacket(&packet.LoginDisconnectPacket{})
	conn.registerPacket(&packet.JoinGamePacket{})
	conn.registerPacket(&packet.PlayerPositionLookPacket{})
	conn.registerPacket(&packet.PluginMessagePacket{})

	return conn
}

// TODO: Make cancel-based contexts work for ReadPacket & WritePacket

// ReadPacket reads a Minecraft protocol packet from the connection.
// ReadPacket will also try to parse the contents of this packet and return
// an error if the given packet ID is unknown or if the known packet format
// did not decode correctly.
//
// The provided Context must be non-nil. If the context expires before
// a packet was read, an error is returned. Cancelling this context
// currently does not stop the packet read, only a deadline (or timeout) works.
func (c *connection) ReadPacket(ctx context.Context) (packet.Packet, error) {
	var err error

	if ctx == nil {
		panic("nil context")
	}

	deadline, hasDeadline := ctx.Deadline()

	if hasDeadline && !deadline.IsZero() {
		_ = c.transport.SetReadDeadline(deadline)
	} else {
		_ = c.transport.SetReadDeadline(time.Time{})
	}

	// Read packet length
	var length enc.VarInt
	if err = length.Read(c.transport); err != nil {
		return nil, fmt.Errorf("could not read packet length: %w", err)
	}

	// TODO: Maybe handle legacy server ping? (https://wiki.vg/Server_List_Ping#1.6)
	if length == 0xFE {
		return nil, ErrLegacyServerPing
	}

	// Catch invalid packet lengths
	if length < 1 || length > packet.MaxPacketLength {
		return nil, packet.ErrInvalidPacketLength
	}

	// Read complete packet into memory
	data := make([]byte, length)
	if _, err = io.ReadAtLeast(c.transport, data, int(length)); err != nil {
		return nil, fmt.Errorf("could not read packet from connection: %w", err)
	}

	// Reader for data buffer
	reader := bytes.NewReader(data)

	// Read packet ID
	var pID enc.VarInt
	if err = pID.Read(reader); err != nil {
		return nil, fmt.Errorf("could not decode packet ID: %w", err)
	}

	// Lookup packet type
	pkt, err := c.getPacketTypeByID(int32(pID))

	if err != nil {
		return nil, err
	}

	log.Printf("[Recv] %s, %d: %s\n", c.state, pID, pkt.String())

	// Decode packet
	if err = pkt.Read(reader); err != nil {
		return nil, fmt.Errorf("could not decode packet data: %w", err)
	}

	return pkt, nil
}

// WritePacket writes a Minecraft protocol packet to the connection.
//
// The provided Context must be non-nil. If the context expires before
// a packet was written, an error is returned. Cancelling this context
// currently does not stop the packet write, only a deadline (or timeout) works.
func (c *connection) WritePacket(ctx context.Context, packetToWrite packet.Packet) error {
	if ctx == nil {
		panic("nil context")
	}

	deadline, hasDeadline := ctx.Deadline()

	if hasDeadline && !deadline.IsZero() {
		_ = c.transport.SetWriteDeadline(deadline)
	} else {
		_ = c.transport.SetWriteDeadline(time.Time{})
	}

	packetInfo := packetToWrite.Info()

	if packetInfo.ConnectionState != c.state {
		return ErrConnectionState
	}

	if packetInfo.Direction != c.getWriteDirection() {
		return ErrDirection
	}

	log.Printf("[Send] %s, %d: %s\n", c.state, packetInfo.ID, packetToWrite.String())

	var buffer bytes.Buffer

	packetId := enc.VarInt(packetInfo.ID)

	// Write packet ID to buffer
	if err := packetId.Write(&buffer); err != nil {
		return fmt.Errorf("unable to write packet id to buffer: %w", err)
	}

	// Write packet data to buffer
	if err := packetToWrite.Write(&buffer); err != nil {
		return fmt.Errorf("could not encode packet data: %w", err)
	}

	// Get packet length
	length := enc.VarInt(buffer.Len())

	// Write packet length to connection
	if err := length.Write(c.transport); err != nil {
		return fmt.Errorf("could not write packet length: %w", err)
	}

	// Write buffer to connection
	if _, err := buffer.WriteTo(c.transport); err != nil {
		return fmt.Errorf("could not write packet data: %w", err)
	}

	return nil
}

// SetState switches the protocol to a different state,
// which changes the meaning of packet IDs.
func (c *connection) SetState(state types.ConnectionState) {
	c.state = state
}

// Close closes the connection. After this has been called,
// the Connection should not be used anymore.
func (c *connection) Close() error {
	return c.transport.Close()
}

func (c *connection) getReadDirection() types.PacketDirection {
	if c.side == types.ServerSide {
		return types.ServerBound
	} else {
		return types.ClientBound
	}
}

func (c *connection) getWriteDirection() types.PacketDirection {
	if c.side == types.ServerSide {
		return types.ClientBound
	} else {
		return types.ServerBound
	}
}

func (c *connection) registerPacket(packet packet.Packet) {
	c.packets[packet.Info()] = packet
}

func (c *connection) getPacketTypeByID(id int32) (packet.Packet, error) {
	p, ok := c.packets[packet.PacketInfo{
		ID:              id,
		Direction:       c.getReadDirection(),
		ConnectionState: c.state,
	}]

	if !ok {
		return nil, fmt.Errorf("unknown packet ID, direction: %s, state: %s, ID: %#x", c.getReadDirection(), c.state, id)
	}

	return p, nil
}
