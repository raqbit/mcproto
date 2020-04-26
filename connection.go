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
	HandshakeState = 0x00 // Client is making a handshake
	StatusState    = 0x01 // Client is receiving the server status
	LoginState     = 0x02 // Client is logging in
	PlayState      = 0x03 // Client is playing
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
	rw      io.ReadWriteCloser
	State   ConnectionState
	side    ConnectionSide
	packets map[PacketInfo]Packet
}

// Creates a new connection using the provided tcp connection
func NewConnection(rw io.ReadWriteCloser, side ConnectionSide) *Connection {
	conn := &Connection{
		rw:      rw,
		side:    side,
		State:   HandshakeState,
		packets: make(map[PacketInfo]Packet),
	}

	// FIXME: DO NOT use the same packet instance for each unmarshal
	conn.registerPacket(&CHandshakePacket{})
	conn.registerPacket(&CServerQueryPacket{})
	conn.registerPacket(&SServerInfoPacket{})
	conn.registerPacket(&CPingPacket{})
	conn.registerPacket(&SPongPacket{})
	conn.registerPacket(&SDisconnectLoginPacket{})
	conn.registerPacket(&CLoginStartPacket{})
	conn.registerPacket(&SLoginSuccessPacket{})
	conn.registerPacket(&SJoinGamePacket{})
	conn.registerPacket(&CClientSettingsPacket{})
	conn.registerPacket(&SPlayerPositionLookPacket{})

	return conn
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
	if length < 1 || length > MaxPacketLength {
		return nil, ErrInvalidPacketLength
	}

	// Read complete packet into memory
	data := make([]byte, length)
	_, err = io.ReadFull(c.rw, data)

	// Create packetbuffer from data
	packetBuffer := NewPacketReader(bytes.NewReader(data))

	if err != nil {
		return nil, fmt.Errorf("could not read packet data: %w", err)
	}

	pID, err := packetBuffer.ReadVarInt()

	if err != nil {
		return nil, fmt.Errorf("could not decode packet ID: %w", err)
	}

	packet, ok := c.packets[PacketInfo{
		ID:              pID,
		Direction:       c.getReadDirection(),
		ConnectionState: c.State,
	}]

	if !ok {
		return nil, fmt.Errorf("unknown packet ID, direction: %s, state: %s, ID: %#x", c.getReadDirection(), c.State, pID)
	}

	log.Printf("[Recv] %s, %d: %s\n", c.State, pID, packet.String())

	// Decode packet
	err = packet.Unmarshal(packetBuffer)

	if err != nil {
		return nil, fmt.Errorf("could not decode packet data: %w", err)
	}

	return packet, nil
}

// Writes a packet to the connection
func (c *Connection) WritePacket(packetToWrite Packet) error {
	packetInfo := packetToWrite.Info()

	log.Printf("[Send] %s, %d: %s\n", c.State, packetInfo.ID, packetToWrite.String())

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

	// TODO: handle case where less than x bytes were written

	// Write packet length to connection
	if _, err := c.rw.Write(enc.WriteVarInt(int32(buffer.Len()))); err != nil {
		return fmt.Errorf("could not write packet length: %w", err)
	}

	// Write buffer to connection
	if _, err := buffer.WriteTo(c.rw); err != nil {
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

func (c *Connection) registerPacket(packet Packet) {
	c.packets[packet.Info()] = packet
}
