package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
)

// https://wiki.vg/Protocol#Handshake
type HandshakePacket struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  enc.VarInt
}

func (HandshakePacket) String() string {
	return "Handshake"
}

func (HandshakePacket) ID() int {
	return 0x00
}

func (h HandshakePacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := h.ProtoVer.Encode(buffer); err != nil {
		return nil, err
	}

	if err := h.ServerAddr.Encode(buffer); err != nil {
		return nil, err
	}

	if err := h.ServerPort.Encode(buffer); err != nil {
		return nil, err
	}

	if err := h.NextState.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (HandshakePacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	h := &HandshakePacket{}

	if err := h.ProtoVer.Decode(data); err != nil {
		return nil, err
	}

	if err := h.ServerAddr.Decode(data); err != nil {
		return nil, err
	}

	if err := h.ServerPort.Decode(data); err != nil {
		return nil, err
	}

	if err := h.NextState.Decode(data); err != nil {
		return nil, err
	}

	return h, nil
}
