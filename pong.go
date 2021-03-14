package mcproto

const PongPacketID int32 = 0x01

// https://wiki.vg/Protocol#Pong
type PongPacket struct {
	Payload int64
}

func (*PongPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PongPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStateStatus,
	}
}

func (*PongPacket) String() string {
	return "Pong"
}

func (p *PongPacket) Marshal(w PacketWriter) error {
	return w.WriteLong(p.Payload)
}

func (p *PongPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Payload, err = r.ReadLong(); err != nil {
		return err
	}

	return nil
}
