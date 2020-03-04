package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
)

// https://wiki.vg/Server_List_Ping#Request
type RequestPacket struct{}

func (RequestPacket) String() string {
	return "Request"
}

func (RequestPacket) Marshal() (*bytes.Buffer, error) {
	return nil, nil
}

func (RequestPacket) Unmarshal(_ *bytes.Buffer) (Packet, error) {
	return &RequestPacket{}, nil
}

func (RequestPacket) ID() int {
	return 0x00
}

// https://wiki.vg/Server_List_Ping#Response
type ResponsePacket struct {
	Json enc.String
}

func (ResponsePacket) String() string {
	return "Response"
}

func (ResponsePacket) ID() int {
	return 0x00
}

func (r ResponsePacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := r.Json.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (ResponsePacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	rp := &ResponsePacket{}

	if err := rp.Json.Decode(data); err != nil {
		return nil, err
	}

	return rp, nil
}

// https://wiki.vg/Server_List_Ping#Ping
type PingPacket struct {
	Payload enc.Long
}

func (PingPacket) String() string {
	return "Ping"
}

func (PingPacket) ID() int {
	return 0x01
}

func (p PingPacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := p.Payload.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (PingPacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	pp := &PingPacket{}

	if err := pp.Payload.Decode(data); err != nil {
		return nil, err
	}

	return pp, nil
}

// https://wiki.vg/Server_List_Ping#Pong
type PongPacket struct {
	Payload enc.Long
}

func (PongPacket) String() string {
	return "Pong"
}

func (PongPacket) ID() int {
	return 0x01
}

func (p PongPacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := p.Payload.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (PongPacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	pp := &PongPacket{}

	if err := pp.Payload.Decode(data); err != nil {
		return nil, err
	}

	return pp, nil
}
