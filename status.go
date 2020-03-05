package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

// https://wiki.vg/Protocol#Request
type RequestPacket struct{}

func (r RequestPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ServerBound,
		ConnectionState: StatusState,
	}
}

func (RequestPacket) String() string {
	return "Request"
}

func (RequestPacket) Marshal() ([]byte, error) {
	return nil, nil
}

func (RequestPacket) Unmarshal(_ io.Reader) (Packet, error) {
	return &RequestPacket{}, nil
}

// https://wiki.vg/Protocol#Response
type ResponsePacket struct {
	Json enc.String
}

func (r ResponsePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ClientBound,
		ConnectionState: StatusState,
	}
}

func (ResponsePacket) String() string {
	return "Response"
}

func (r ResponsePacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := r.Json.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (ResponsePacket) Unmarshal(data io.Reader) (Packet, error) {
	rp := &ResponsePacket{}

	if err := rp.Json.Decode(data); err != nil {
		return nil, err
	}

	return rp, nil
}

// https://wiki.vg/Protocol#Ping
type PingPacket struct {
	Payload enc.Long
}

func (p PingPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x01,
		Direction:       ServerBound,
		ConnectionState: StatusState,
	}
}

func (PingPacket) String() string {
	return "Ping"
}

func (p PingPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := p.Payload.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (PingPacket) Unmarshal(data io.Reader) (Packet, error) {
	pp := &PingPacket{}

	if err := pp.Payload.Decode(data); err != nil {
		return nil, err
	}

	return pp, nil
}

// https://wiki.vg/Protocol#Pong
type PongPacket struct {
	Payload enc.Long
}

func (p PongPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x01,
		Direction:       ClientBound,
		ConnectionState: StatusState,
	}
}

func (PongPacket) String() string {
	return "Pong"
}

func (p PongPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := p.Payload.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (PongPacket) Unmarshal(data io.Reader) (Packet, error) {
	pp := &PongPacket{}

	if err := pp.Payload.Decode(data); err != nil {
		return nil, err
	}

	return pp, nil
}
