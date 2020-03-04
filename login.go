package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
)

type LoginPacket struct {
	Name enc.String
}

func (LoginPacket) String() string {
	return "Login"
}

func (LoginPacket) ID() int {
	return 0x00
}

func (l LoginPacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := l.Name.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (LoginPacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	lp := &LoginPacket{}

	if err := lp.Name.Decode(data); err != nil {
		return nil, err
	}

	return lp, nil
}

type DisconnectPacket struct {
	Reason enc.String
}

func (DisconnectPacket) String() string {
	return "Disconnect"
}

func (DisconnectPacket) ID() int {
	return 0x00
}

func (d DisconnectPacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := d.Reason.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (DisconnectPacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	dp := &DisconnectPacket{}

	if err := dp.Reason.Decode(data); err != nil {
		return nil, err
	}

	return dp, nil
}

type LoginSuccessPacket struct {
	UUID     enc.String
	Username enc.String
}

func (LoginSuccessPacket) String() string {
	return "LoginSucess"
}

func (LoginSuccessPacket) ID() int {
	return 0x02
}

func (l LoginSuccessPacket) Marshal() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	if err := l.UUID.Encode(buffer); err != nil {
		return nil, err
	}

	if err := l.Username.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (LoginSuccessPacket) Unmarshal(data *bytes.Buffer) (Packet, error) {
	lsp := &LoginSuccessPacket{}

	if err := lsp.UUID.Decode(data); err != nil {
		return nil, err
	}

	if err := lsp.Username.Decode(data); err != nil {
		return nil, err
	}

	return lsp, nil
}
