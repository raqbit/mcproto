package mcproto

import (
	"bytes"
	"errors"
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
	"github.com/Raqbit/mcproto/packet"
	"io"
	"net"
	"time"
)

var (
	ErrConnectionState  = errors.New("connection has invalid state for packet type")
	ErrDirection        = errors.New("packet being sent in wrong direction")
	ErrLegacyServerPing = errors.New("not implemented: legacy server ping")
)

// Connection represents a connection
// to a Minecraft server or client.
type Connection interface {
	ReadPacket() (packet.Packet, error)
	WritePacket(packet.Packet) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	Close() error
	SetState(state game.ConnectionState)
}

type connection struct {
	conn     net.Conn
	state    game.ConnectionState
	side     Side
	packets  map[packet.Info]packet.Packet
	writeBuf *bytes.Buffer
	readBuf  *bytes.Buffer
	logger   Logger
}

// ReadPacket reads a Minecraft protocol packet from the connection.
// ReadPacket will also try to parse the contents of this packet and return
// an error if the given packet ID is unknown or if the known packet format
// did not decode correctly.
func (c *connection) ReadPacket() (packet.Packet, error) {
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
	if length < 1 || length > packet.MaxPacketLength {
		return nil, packet.ErrInvalidPacketLength
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
	pkt, err := c.getPacketTypeByID(int32(pID))

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
func (c *connection) WritePacket(packetToWrite packet.Packet) error {
	if packetToWrite.State() != c.state {
		return ErrConnectionState
	}

	if packetToWrite.Direction() != c.side.WriteDirection() {
		return ErrDirection
	}

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
func (c *connection) SetState(state game.ConnectionState) {
	c.state = state
}

// Close closes the connection. After this has been called,
// the Connection should not be used anymore.
func (c *connection) Close() error {
	return c.conn.Close()
}

func (c *connection) registerPacket(pkt packet.Packet) {
	c.packets[packet.Info{
		ID:              pkt.ID(),
		ConnectionState: pkt.State(),
		Direction:       pkt.Direction(),
	}] = pkt
}

type UnknownPacketError struct {
	id        int32
	direction packet.Direction
	state     game.ConnectionState
}

func (u UnknownPacketError) Error() string {
	return fmt.Sprintf("unknown packet ID, direction: %s, state: %s, ID: %#x", u.direction, u.state, u.id)
}

func (c *connection) getPacketTypeByID(id int32) (packet.Packet, error) {
	p, ok := c.packets[packet.Info{
		ID:              id,
		Direction:       c.side.ReadDirection(),
		ConnectionState: c.state,
	}]

	if !ok {
		return nil, UnknownPacketError{
			id:        id,
			direction: c.side.WriteDirection(),
			state:     c.state,
		}
	}

	return p, nil
}
