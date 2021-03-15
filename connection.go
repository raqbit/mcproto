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
	"time"
)

var (
	ErrConnectionState = errors.New("connection has invalid state for packet type")
	ErrDirection       = errors.New("packet being sent in wrong direction")
)

// Connection represents a connection
// to a Minecraft server or client.
type Connection interface {
	ReadPacket(context.Context) (packet.Packet, error)
	WritePacket(context.Context, packet.Packet) error
	Close() error
	SwitchState(state types.ConnectionState)
}

type connection struct {
	transport net.Conn
	state     types.ConnectionState
	side      types.Side
	packets   map[packet.PacketInfo]packet.Packet
}

// Dial connects to the specified address and creates
// a new Connection for it.
func Dial(address string, side types.Side) (Connection, error) {
	return DialContext(context.Background(), address, side)
}

// Dial creates a TCP connection with specified address
// and creates a new Connection for it.
//
// The provided Context must be non-nil. If the context expires before
// the connection is complete, an error is returned. Once successfully
// connected, any expiration of the context will not affect the connection.
func DialContext(ctx context.Context, address string, side types.Side) (Connection, error) {
	// Make TCP connection
	var d net.Dialer
	tcpConn, err := d.DialContext(ctx, "tcp", address)

	if err != nil {
		return nil, err
	}

	return WrapConnection(tcpConn, side), nil
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

// Reads a packet from the connection
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
		return nil, fmt.Errorf("not implemented: Legacy server ping")
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
	packet, err := c.getPacketTypeByID(int32(pID))

	if err != nil {
		return nil, err
	}

	log.Printf("[Recv] %s, %d: %s\n", c.state, pID, packet.String())

	// Decode packet
	if err = packet.Read(reader); err != nil {
		return nil, fmt.Errorf("could not decode packet data: %w", err)
	}

	return packet, nil
}

// Writes a packet to the connection
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

func (c *connection) SwitchState(state types.ConnectionState) {
	c.state = state
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

func (c *connection) Close() error {
	return c.transport.Close()
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
