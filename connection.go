package mcproto

import (
	"bytes"
	"errors"
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
	"net"
	"time"
)

const MaxPacketLength = 1048576
const (
	clientBound PacketDirection = iota
	serverBound PacketDirection = iota
)

var (
	ErrInvalidPacketLength = errors.New("packet has a malformed length")
	ErrLegacyServerPing    = errors.New("legacy server ping not supported")
)

type (
	// Packet is a packet with side-bound ID
	Packet interface {
		fmt.Stringer
		enc.MinecraftEncodable
		enc.MinecraftDecodable

		ID() int32
	}

	// Connection represents a connection
	// to a Minecraft server or client.
	Connection interface {
		Read() (Packet, error)
		Write(Packet) error
		SetReadDeadline(time.Time) error
		SetWriteDeadline(time.Time) error
		Close() error
		SetState(state ConnectionState)
	}

	// packetKey uniquely identifies a packet
	packetKey struct {
		id        int32
		direction PacketDirection
		state     ConnectionState
	}

	connection struct {
		conn     net.Conn
		state    ConnectionState
		side     Side
		packets  map[packetKey]Packet
		writeBuf *bytes.Buffer
		readBuf  *bytes.Buffer
		logger   Logger
	}
)

func (c *connection) Read() (Packet, error) {
	var err error

	// Read packet length
	var length enc.VarInt
	if err = length.Read(c.conn); err != nil {
		return nil, fmt.Errorf("could not read packet length: %w", err)
	}

	// TODO: Handle legacy server ping (https://wiki.vg/Server_List_Ping#1.6)
	if length == 0xFE {
		return nil, ErrLegacyServerPing
	}

	// Catch invalid packet lengths
	if length < 1 || length > MaxPacketLength {
		return nil, ErrInvalidPacketLength
	}

	// Read complete packet into read buffer
	if _, err = c.readBuf.ReadFrom(io.LimitReader(c.conn, int64(length))); err != nil {
		return nil, fmt.Errorf("could not read packet from connection: %w", err)
	}

	// Read packet ID
	var pID enc.VarInt
	if err = pID.Read(c.readBuf); err != nil {
		return nil, fmt.Errorf("could not decode packet ID: %w", err)
	}

	// Lookup packet type
	pkt, err := c.packetTypeFromID(int32(pID))

	if err != nil {
		return nil, err
	}

	c.logger.Debugf("[Recv] %s, %d: %s\n", c.state, pID, pkt.String())

	// Decode packet
	if err = pkt.Read(c.readBuf); err != nil {
		return nil, fmt.Errorf("could not decode packet data: %w", err)
	}

	// Reset read buffer
	c.readBuf.Reset()

	return pkt, nil
}

// WritePacket writes a Minecraft protocol packet to the connection.
//
// If the provided deadline expires before a packet could be written,
// an error is returned. Providing a zero value for the deadline will
// make the read not time out.
func (c *connection) Write(packetToWrite Packet) error {
	c.logger.Debugf("[Send] %s, %d: %s\n", c.state, packetToWrite.ID(), packetToWrite.String())

	packetId := enc.VarInt(packetToWrite.ID())

	// Write packet ID to buffer
	if err := packetId.Write(c.writeBuf); err != nil {
		return fmt.Errorf("unable to write packet id to buffer: %w", err)
	}

	// Write packet data to buffer
	if err := packetToWrite.Write(c.writeBuf); err != nil {
		return fmt.Errorf("could not encode packet data: %w", err)
	}

	// Get packet length
	length := enc.VarInt(c.writeBuf.Len())

	// Write packet length to connection
	if err := length.Write(c.conn); err != nil {
		return fmt.Errorf("could not write packet length: %w", err)
	}

	// Flush write buffer to connection
	if _, err := c.writeBuf.WriteTo(c.conn); err != nil {
		return fmt.Errorf("could not flush write buffer: %w", err)
	}

	// Reset write buffer
	c.writeBuf.Reset()

	return nil
}

// SetReadDeadline sets the deadline for future ReadPacket calls
// and any currently-blocked ReadPacket call.
// A zero value for t means ReadPacket will not time out.
func (c *connection) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future WritePacket calls
// and any currently-blocked WritePacket call.
// A zero value for t means WritePacket will not time out.
func (c *connection) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

// SetState switches the protocol to a different state,
// which changes the meaning of packet IDs.
func (c *connection) SetState(state ConnectionState) {
	c.state = state
}

// Close closes the connection. After this has been called,
// the Connection should not be used anymore.
func (c *connection) Close() error {
	return c.conn.Close()
}

func (c *connection) registerPacket(pkt Packet, dir PacketDirection, state ConnectionState) {
	c.packets[packetKey{id: pkt.ID(), direction: dir, state: state}] = pkt
}

type UnknownPacketError struct {
	id        int32
	direction PacketDirection
	state     ConnectionState
}

func (u UnknownPacketError) Error() string {
	return fmt.Sprintf("unknown packet id, direction: %s, state: %s, id: %#x", u.direction, u.state, u.id)
}

func (c *connection) packetTypeFromID(id int32) (Packet, error) {
	p, ok := c.packets[packetKey{
		id:        id,
		direction: c.side.readDirection(),
		state:     c.state,
	}]

	if !ok {
		return nil, UnknownPacketError{
			id:        id,
			direction: c.side.writeDirection(),
			state:     c.state,
		}
	}

	return p, nil
}
