package mcproto

const PlayerPositionLookPacketID int32 = 0x36

type PlayerPositionLookPacket struct {
	X          float64
	Y          float64
	Z          float64
	Yaw        float32
	Pitch      float32
	Flags      uint8
	TeleportID int32
}

func (*PlayerPositionLookPacket) String() string {
	return "PlayerPositionAndLook"
}

func (*PlayerPositionLookPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              PlayerPositionLookPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}

func (p *PlayerPositionLookPacket) Marshal(w PacketWriter) error {
	var err error

	if err = w.WriteDouble(p.X); err != nil {
		return err
	}

	if err = w.WriteDouble(p.Y); err != nil {
		return err
	}

	if err = w.WriteDouble(p.Z); err != nil {
		return err
	}

	if err = w.WriteFloat(p.Yaw); err != nil {
		return err
	}

	if err = w.WriteFloat(p.Pitch); err != nil {
		return err
	}

	if err = w.WriteUnsignedByte(p.Flags); err != nil {
		return err
	}

	if err = w.WriteVarInt(p.TeleportID); err != nil {
		return err
	}

	return nil
}

func (p *PlayerPositionLookPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.X, err = r.ReadDouble(); err != nil {
		return err
	}

	if p.Y, err = r.ReadDouble(); err != nil {
		return err
	}

	if p.Z, err = r.ReadDouble(); err != nil {
		return err
	}

	if p.Yaw, err = r.ReadFloat(); err != nil {
		return err
	}

	if p.Pitch, err = r.ReadFloat(); err != nil {
		return err
	}

	if p.Flags, err = r.ReadUnsignedByte(); err != nil {
		return err
	}

	if p.TeleportID, err = r.ReadVarInt(); err != nil {
		return err
	}

	return nil
}
