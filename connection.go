package mcproto

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
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
	ReadPacket(context.Context) (Packet, error)
	WritePacket(context.Context, Packet) error
	Close() error
	SwitchState(state ConnectionState)
}

type connection struct {
	transport net.Conn
	state     ConnectionState
	side      Side
	packets   map[PacketInfo]Packet
}

// Dial connects to the specified address and creates
// a new Connection for it.
func Dial(address string, side Side) (Connection, error) {
	return DialContext(context.Background(), address, side)
}

// Dial creates a TCP connection with specified address
// and creates a new Connection for it.
//
// The provided Context must be non-nil. If the context expires before
// the connection is complete, an error is returned. Once successfully
// connected, any expiration of the context will not affect the connection.
func DialContext(ctx context.Context, address string, side Side) (Connection, error) {
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
func WrapConnection(transport net.Conn, side Side) Connection {
	conn := &connection{
		transport: transport,
		state:     ConnectionStateHandshake,
		side:      side,
		packets:   make(map[PacketInfo]Packet),
	}

	// Server bound packets
	conn.registerPacket(&HandshakePacket{})
	conn.registerPacket(&ServerQueryPacket{})
	conn.registerPacket(&PingPacket{})
	conn.registerPacket(&LoginStartPacket{})
	conn.registerPacket(&ClientSettingsPacket{})

	// Client bound packets
	conn.registerPacket(&PongPacket{})
	conn.registerPacket(&LoginSuccessPacket{})
	conn.registerPacket(&ChatMessagePacket{})
	conn.registerPacket(&ServerInfoPacket{})
	conn.registerPacket(&LoginDisconnectPacket{})
	conn.registerPacket(&JoinGamePacket{})
	conn.registerPacket(&PlayerPositionLookPacket{})

	return conn
}

// TODO: Make cancel-based contexts work for ReadPacket & WritePacket

// Reads a packet from the connection
func (c *connection) ReadPacket(ctx context.Context) (Packet, error) {
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
	length, err := enc.ReadVarInt(c.transport)

	if err != nil {
		return nil, fmt.Errorf("could not read packet length: %w", err)
	}

	// TODO: Maybe handle legacy server ping? (https://wiki.vg/Server_List_Ping#1.6)
	if length == 0xFE {
		return nil, fmt.Errorf("not implemented: Legacy server ping")
	}

	// Catch invalid packet lengths
	if length < 1 || length > MaxPacketLength {
		return nil, ErrInvalidPacketLength
	}

	// Read complete packet into memory
	data := make([]byte, length)

	if _, err = io.ReadAtLeast(c.transport, data, int(length)); err != nil {
		return nil, fmt.Errorf("could not read packet from connection: %w", err)
	}

	// Create packetbuffer from data
	packetBuffer := NewPacketReader(bytes.NewReader(data))

	pID, err := packetBuffer.ReadVarInt()

	if err != nil {
		return nil, fmt.Errorf("could not decode packet ID: %w", err)
	}

	packet, err := c.getPacketTypeByID(pID)

	if err != nil {
		return nil, err
	}

	log.Printf("[Recv] %s, %d: %s\n", c.state, pID, packet.String())

	// Decode packet
	err = packet.Unmarshal(packetBuffer)

	if err != nil {
		return nil, fmt.Errorf("could not decode packet data: %w", err)
	}

	return packet, nil
}

// Writes a packet to the connection
func (c *connection) WritePacket(ctx context.Context, packetToWrite Packet) error {
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

	buffer := new(bytes.Buffer)
	packetBuffer := NewPacketBuffer(buffer)

	// Write packet ID to buffer
	if err := packetBuffer.WriteVarInt(packetInfo.ID); err != nil {
		return fmt.Errorf("unable to write packet id to buffer: %w", err)
	}

	// Write packet data to buffer
	if err := packetToWrite.Marshal(packetBuffer); err != nil {
		return fmt.Errorf("could not encode packet data: %w", err)
	}

	// Write packet length to connection
	if _, err := c.transport.Write(enc.WriteVarInt(int32(buffer.Len()))); err != nil {
		return fmt.Errorf("could not write packet length: %w", err)
	}

	// Write buffer to connection
	if _, err := buffer.WriteTo(c.transport); err != nil {
		return fmt.Errorf("could not write packet data: %w", err)
	}

	return nil
}

func (c *connection) SwitchState(state ConnectionState) {
	c.state = state
}

func (c *connection) getReadDirection() PacketDirection {
	if c.side == ServerSide {
		return ServerBound
	} else {
		return ClientBound
	}
}

func (c *connection) getWriteDirection() PacketDirection {
	if c.side == ServerSide {
		return ClientBound
	} else {
		return ServerBound
	}
}

func (c *connection) Close() error {
	return c.transport.Close()
}

func (c *connection) registerPacket(packet Packet) {
	c.packets[packet.Info()] = packet
}

func (c *connection) getPacketTypeByID(id int32) (Packet, error) {
	p, ok := c.packets[PacketInfo{
		ID:              id,
		Direction:       c.getReadDirection(),
		ConnectionState: c.state,
	}]

	if !ok {
		return nil, fmt.Errorf("unknown packet ID, direction: %s, state: %s, ID: %#x", c.getReadDirection(), c.state, id)
	}

	return p, nil
}
