package mcproto

import "github.com/google/uuid"

const LoginSuccessPacketID = 0x02

// https://wiki.vg/Protocol#Login_Success
type LoginSuccessPacket struct {
	UUID     uuid.UUID
	Username string
}

func (*LoginSuccessPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginSuccessPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStateLogin,
	}
}

func (*LoginSuccessPacket) String() string {
	return "LoginSuccess"
}

func (l *LoginSuccessPacket) Marshal(w PacketWriter) error {
	var err error

	if err = w.WriteString(l.UUID.String()); err != nil {
		return err
	}

	if err = w.WriteString(l.Username); err != nil {
		return err
	}

	return nil
}

func (l *LoginSuccessPacket) Unmarshal(r PacketReader) error {
	var err error
	var uuidStr string

	if uuidStr, err = r.ReadString(36); err != nil {
		return err
	}

	if l.UUID, err = uuid.Parse(uuidStr); err != nil {
		return err
	}

	if l.Username, err = r.ReadString(16); err != nil {
		return err
	}

	return nil
}
