package mcproto

const PingPacketID = 0x01

// https://wiki.vg/Protocol#Ping
type PingPacket struct {
	Payload int64
}

func (*PingPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PingPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateStatus,
	}
}

func (*PingPacket) String() string {
	return "Ping"
}

func (p *PingPacket) Marshal(w PacketWriter) error {
	return w.WriteLong(p.Payload)
}

func (p *PingPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Payload, err = r.ReadLong(); err != nil {
		return err
	}

	return nil
}
