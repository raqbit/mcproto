package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

// https://wiki.vg/Protocol#Handshake
type HandshakePacket struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  enc.VarInt
}

func (h HandshakePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ServerBound,
		ConnectionState: HandshakeState,
	}
}

func (HandshakePacket) String() string {
	return "Handshake"
}

func (h HandshakePacket) Marshal() ([]byte, error) {
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

	return buffer.Bytes(), nil
}

func (HandshakePacket) Unmarshal(r io.Reader) (Packet, error) {
	h := &HandshakePacket{}

	if err := h.ProtoVer.Decode(r); err != nil {
		return nil, err
	}

	if err := h.ServerAddr.Decode(r); err != nil {
		return nil, err
	}

	if err := h.ServerPort.Decode(r); err != nil {
		return nil, err
	}

	if err := h.NextState.Decode(r); err != nil {
		return nil, err
	}

	return h, nil
}
