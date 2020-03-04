package mcproto

import (
	"bytes"
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
	"log"
)

// The state of a connection
type ConnectionState uint8

const (
	HandshakeState ConnectionState = iota // Client is making a handshake
	StatusState                           // Client is receiving the server status
	LoginState                            // Client is logging in
	PlayState                             // Client is playing
)

func (c ConnectionState) String() string {
	names := []string{
		"handshake",
		"status",
		"login",
		"play",
	}

	return names[c]
}

type ConnectionSide uint8

const (
	ClientSide ConnectionSide = iota
	ServerSide
)

type Connection struct {
	rw    io.ReadWriteCloser
	State ConnectionState
	side  ConnectionSide
}

// Creates a new connection using the provided tcp connection
func NewConnection(conn io.ReadWriteCloser, side ConnectionSide) *Connection {
	return &Connection{rw: conn, side: side, State: HandshakeState}
}

// Reads a packet from the connection
func (c *Connection) ReadPacket() (Packet, error) {
	// Read packet length
	length, err := enc.ReadVarInt(c.rw)

	if err != nil {
		return nil, fmt.Errorf("could not read packet length: %w", err)
	}

	// TODO: Maybe handle legacy server ping? (https://wiki.vg/Server_List_Ping#1.6)
	if length == 0xFE {
		return nil, fmt.Errorf("not implemented: Legacy server ping")
	}

	// Catch invalid packet lengths
	if length < 0 || length > MaxPacketLength {
		return nil, ErrInvalidPacketLength
	}

	data := make([]byte, length)

	// Read complete packet into data slice
	_, err = io.ReadFull(c.rw, data)

	if err != nil {
		return nil, fmt.Errorf("could not read packet data: %w", err)
	}
	// Create buffer
	buff := bytes.NewBuffer(data)

	pID, err := enc.ReadVarInt(buff)

	if err != nil {
		return nil, fmt.Errorf("could not decode packet ID: %w", err)
	}

	// Decode packet
	packetType, ok := PacketTypes[c.getReadDirection()][c.State][int(pID)]

	if !ok {
		return nil, fmt.Errorf("unknown packet ID, direction: %s, state: %s, ID: %d", c.getReadDirection(), c.State, pID)
	}

	log.Printf("[Recv] %s, %d: %s\n", c.State, pID, packetType.String())

	decodedPacket, err := packetType.Unmarshal(buff)

	if err != nil {
		return nil, fmt.Errorf("could not decode packet data: %w", err)
	}

	return decodedPacket, nil
}

// Writes a packet to the connection
func (c *Connection) WritePacket(packetToWrite Packet) error {
	log.Printf("[Send] %s, %d: %s\n", c.State, packetToWrite.ID(), packetToWrite.String())

	idBuff := new(bytes.Buffer)

	// Write packet ID to packet buffer
	if err := enc.WriteVarInt(idBuff, enc.VarInt(packetToWrite.ID())); err != nil {
		return fmt.Errorf("could not write packet id: %w", err)
	}

	dataBuff, err := packetToWrite.Marshal()

	if err != nil {
		return fmt.Errorf("could not encode packet data: %w", err)
	}

	// Write packet length to connection
	if err = enc.WriteVarInt(c.rw, enc.VarInt(idBuff.Len()+dataBuff.Len())); err != nil {
		return fmt.Errorf("could not write packet length: %w", err)
	}

	// Write packet ID to connection
	if _, err = c.rw.Write(idBuff.Bytes()); err != nil {
		return fmt.Errorf("could not write packet id: %w", err)
	}

	// Write packet data to connection
	if _, err = c.rw.Write(dataBuff.Bytes()); err != nil {
		return fmt.Errorf("could not write packet data: %w", err)
	}

	return nil
}

func (c *Connection) getReadDirection() PacketDirection {
	if c.side == ServerSide {
		return ServerBound
	} else {
		return ClientBound
	}
}

func (c *Connection) getWriteDirection() PacketDirection {
	if c.side == ServerSide {
		return ClientBound
	} else {
		return ServerBound
	}
}

func (c *Connection) Close() error {
	return c.rw.Close()
}
