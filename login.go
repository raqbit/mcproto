package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

// https://wiki.vg/Protocol#Login_Start
type LoginStartPacket struct {
	Name enc.String
}

func (l LoginStartPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ServerBound,
		ConnectionState: LoginState,
	}
}

func (LoginStartPacket) String() string {
	return "Login"
}

func (l LoginStartPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := l.Name.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (LoginStartPacket) Unmarshal(r io.Reader) (Packet, error) {
	lp := &LoginStartPacket{}

	if err := lp.Name.Decode(r); err != nil {
		return nil, err
	}

	return lp, nil
}

// https://wiki.vg/Protocol#Disconnect_.28login.29
type DisconnectPacket struct {
	Reason enc.String
}

func (d DisconnectPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ClientBound,
		ConnectionState: LoginState,
	}
}

func (DisconnectPacket) String() string {
	return "Disconnect"
}

func (d DisconnectPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := d.Reason.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (DisconnectPacket) Unmarshal(r io.Reader) (Packet, error) {
	dp := &DisconnectPacket{}

	if err := dp.Reason.Decode(r); err != nil {
		return nil, err
	}

	return dp, nil
}

// https://wiki.vg/Protocol#Login_Success
type LoginSuccessPacket struct {
	UUID     enc.String
	Username enc.String
}

func (l LoginSuccessPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x02,
		Direction:       ClientBound,
		ConnectionState: LoginState,
	}
}

func (LoginSuccessPacket) String() string {
	return "LoginSuccess"
}

func (l LoginSuccessPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := l.UUID.Encode(buffer); err != nil {
		return nil, err
	}

	if err := l.Username.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (LoginSuccessPacket) Unmarshal(r io.Reader) (Packet, error) {
	lsp := &LoginSuccessPacket{}

	if err := lsp.UUID.Decode(r); err != nil {
		return nil, err
	}

	if err := lsp.Username.Decode(r); err != nil {
		return nil, err
	}

	return lsp, nil
}
