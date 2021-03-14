package mcproto

const LoginStartPacketID int32 = 0x00

// https://wiki.vg/Protocol#Login_Start
type LoginStartPacket struct {
	Name string
}

func (*LoginStartPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginStartPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateLogin,
	}
}

func (*LoginStartPacket) String() string {
	return "LoginStart"
}

func (l *LoginStartPacket) Marshal(w PacketWriter) error {
	return w.WriteString(l.Name)
}

func (l *LoginStartPacket) Unmarshal(r PacketReader) error {
	var err error

	if l.Name, err = r.ReadString(16); err != nil {
		return err
	}

	return nil
}
