package mcproto

const PluginMessagePacketID int32 = 0x19

type PluginMessagePacket struct {
	Channel Identifier
	Data    *PacketBuffer
}

func (*PluginMessagePacket) String() string {
	return "PluginMessage"
}

func (*PluginMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PluginMessagePacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}

func (p *PluginMessagePacket) Marshal(w PacketWriter) error {
	var err error

	if err = w.WriteIdentifier(p.Channel); err != nil {
		return err
	}

	if err = w.WriteBytes(p.Data); err != nil {
		return err
	}

	return nil
}

func (p *PluginMessagePacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Channel, err = r.ReadIdentifier(); err != nil {
		return err
	}

	return nil
}
