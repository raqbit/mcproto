package mcproto

import (
	"encoding/json"
	"github.com/google/uuid"
)

// https://wiki.vg/Protocol#Login_Start
type SLoginStartPacket struct {
	Profile *GameProfile
}

func (*SLoginStartPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ServerBound,
		ConnectionState: LoginState,
	}
}

func (*SLoginStartPacket) String() string {
	return "LoginStart"
}

func (l *SLoginStartPacket) Marshal(w PacketWriter) error {
	if err := w.WriteString(l.Profile.Name); err != nil {
		return err
	}

	return nil
}

func (l *SLoginStartPacket) Unmarshal(r PacketReader) error {
	var err error
	var name string

	if name, err = r.ReadString(16); err != nil {
		return err
	}

	l.Profile = NewGameProfile(uuid.Nil, name)

	return nil
}

// https://wiki.vg/Protocol#Disconnect_.28login.29
type CDisconnectLoginPacket struct {
	Reason TextComponent
}

func (*CDisconnectLoginPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ClientBound,
		ConnectionState: LoginState,
	}
}

func (*CDisconnectLoginPacket) String() string {
	return "Disconnect"
}

func (d *CDisconnectLoginPacket) Marshal(w PacketWriter) error {
	var err error
	var reason []byte

	if reason, err = json.Marshal(d.Reason); err != nil {
		return err
	}

	if err := w.WriteString(string(reason)); err != nil {
		return err
	}

	return nil
}

func (d *CDisconnectLoginPacket) Unmarshal(r PacketReader) error {
	var err error
	var reason string

	if reason, err = r.ReadMaxString(); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(reason), &d.Reason); err != nil {
		return err
	}

	return nil
}

// https://wiki.vg/Protocol#Login_Success
type CLoginSuccessPacket struct {
	Profile *GameProfile
}

func (*CLoginSuccessPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x02,
		Direction:       ClientBound,
		ConnectionState: LoginState,
	}
}

func (*CLoginSuccessPacket) String() string {
	return "LoginSuccess"
}

func (l *CLoginSuccessPacket) Marshal(w PacketWriter) error {
	if err := w.WriteString(l.Profile.UUID.String()); err != nil {
		return err
	}

	if err := w.WriteString(l.Profile.Name); err != nil {
		return err
	}

	return nil
}

func (l *CLoginSuccessPacket) Unmarshal(r PacketReader) error {
	var err error
	var uuidStr string
	var name string

	if uuidStr, err = r.ReadString(36); err != nil {
		return err
	}

	if name, err = r.ReadString(16); err != nil {
		return err
	}

	var parsedUUID uuid.UUID
	if parsedUUID, err = uuid.Parse(uuidStr); err != nil {
		return err
	}
	l.Profile = NewGameProfile(parsedUUID, name)

	return nil
}
