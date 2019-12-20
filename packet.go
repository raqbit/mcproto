package mcproto

import (
	"bytes"
	"errors"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/packet"
	"golang.org/x/xerrors"
	"reflect"
)

const MaxPacketLength = 1048576

// Packet which has been decoded into a struct.
// With ID() you can figure out what type of packet it is
// and using a type conversion you can get the underlying packet.
type Packet interface {
	ID() int
}

// The direction a packet
type packetDirection byte

const (
	clientBound packetDirection = iota
	serverBound
)

var (
	ErrUnknownPacket       = errors.New("unknown packet")
	ErrInvalidPacketLength = errors.New("packet has a malformed length")
	ErrInvalidPacketFieldType = errors.New("invalid packet field type")
)

// All handled types of packet.
// Used for decoding packets by direction, connection state & id.
var packetTypes = map[packetDirection]map[ConnectionState]map[int]reflect.Type{
	clientBound: {
		Handshake: {
			packet.Handshake{}.ID(): reflect.TypeOf(packet.Handshake{}),
		},
		Status: {
			packet.Request{}.ID(): reflect.TypeOf(packet.Request{}),
		},
	},
	serverBound: {
		Status: {
			packet.Response{}.ID(): reflect.TypeOf(packet.Response{}),
		},
	},
}

// Encodes the given packet into a buffer
func encodePacket(p Packet) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	// Write packet id
	if err := enc.WriteVarInt(buffer, enc.VarInt(p.ID())); err != nil {
		return nil, err
	}

	value := reflect.ValueOf(p)

	// Loop over every field
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		// Assert that field has Encodable type
		encoding, ok := field.Interface().(enc.Encodable)

		if !ok {
			return nil, ErrInvalidPacketFieldType
		}

		// Encode field into buffer
		if err := encoding.Encode(buffer); err != nil {
			return nil, err
		}
	}

	return buffer, nil
}

// Decodes a packet from the given buffer with given connection parameters
func decodePacket(direction packetDirection, state ConnectionState, data *bytes.Buffer) (Packet, error) {
	// Read packet ID
	id, err := enc.ReadVarInt(data)

	if err != nil {
		return nil, xerrors.Errorf("could not read packet id from packet data: %w", err)
	}

	// Detect packet type
	typ, ok := packetTypes[direction][state][int(id)]

	if !ok {
		// Packet type is unknown
		return nil, ErrUnknownPacket
	}

	// Create a pointer to the zero value of detected type
	inst := reflect.New(typ)

	for i := 0; i < inst.NumField(); i++ {
		field := inst.Field(i)

		// Assert that field has Decodable type
		encoding, ok := field.Interface().(enc.Decodable)

		if !ok {
			return nil, ErrInvalidPacketFieldType
		}

		// Decode from buffer into field
		if err := encoding.Decode(data); err != nil {
			return nil, err
		}
	}

	return inst.Interface().(Packet), nil
}
