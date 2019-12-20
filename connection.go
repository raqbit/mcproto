package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
	"golang.org/x/xerrors"
	"io"
	"net"
)

const (
// Current supported Minecraft protocol version
)

// The state of a connection
type ConnectionState uint8

const (
	Handshake ConnectionState = iota // Client is making a handshake
	Status                           // Client is receiving the server status
	Login                            // Client is logging in
	Play                             // Client is playing
)

type ConnectionSide uint8

const (
	ClientSide ConnectionSide = iota
	ServerSide
)

type Connection struct {
	rw    io.ReadWriteCloser
	State ConnectionState
	Side  ConnectionSide
}

// Creates a new connection using the provided tcp connection
func NewConnection(conn net.Conn, side ConnectionSide) *Connection {
	return &Connection{rw: conn, Side: side}
}

// Reads a packet from the connection
func (c *Connection) readPacket() (Packet, error) {
	// Read packet length
	length, err := enc.ReadVarInt(c.rw)

	if err != nil {
		return nil, xerrors.Errorf("could not read packet length: %w", err)
	}

	// Catch invalid packet lengths
	if length < 0 || length > MaxPacketLength {
		return nil, ErrInvalidPacketLength
	}

	data := make([]byte, length)

	// Read complete packet into data slice
	_, err = io.ReadFull(c.rw, data)

	if err != nil {
		return nil, xerrors.Errorf("could not read packet data: %w", err)
	}

	// Create buffer
	buff := bytes.NewBuffer(data)

	// Decode packet
	packet, err := decodePacket(c.getReadDirection(), c.State, buff)

	if err != nil {
		return nil, xerrors.Errorf("could not decode packet data: %w", err)
	}

	return packet, nil
}

// Writes a packet to the connection
func (c *Connection) writePacket(packet Packet) error {
	// Encode packet
	dataBuff, err := encodePacket(packet)

	if err != nil {
		return xerrors.Errorf("could not encode packet data: %w", err)
	}

	// Write packet length to connection
	if err = enc.WriteVarInt(c.rw, enc.VarInt(dataBuff.Len())); err != nil {
		return xerrors.Errorf("could not write packet length: %w", err)
	}

	// Write packet data to connection
	if _, err = c.rw.Write(dataBuff.Bytes()); err != nil {
		return xerrors.Errorf("could not write packet data: %w", err)
	}

	return nil
}

func (c *Connection) getReadDirection() packetDirection {
	if c.Side == ServerSide {
		return serverBound
	} else {
		return clientBound
	}
}

func (c *Connection) getWriteDirection() packetDirection {
	if c.Side == ServerSide {
		return clientBound
	} else {
		return serverBound
	}
}
