package mcproto

const JoinGamePacketID int32 = 0x26

type JoinGamePacket struct {
	PlayerID            int32
	GameMode            uint8
	Dimension           int32
	HashedSeed          int64
	MaxPlayers          uint8
	LevelType           string
	ViewDistance        int32
	ReducedDebug        bool
	EnableRespawnScreen bool
}

func (*JoinGamePacket) String() string {
	return "JoinGame"
}

func (*JoinGamePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              JoinGamePacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}

func (j *JoinGamePacket) Marshal(w PacketWriter) error {
	var err error

	if err = w.WriteInt(j.PlayerID); err != nil {
		return err
	}

	if err = w.WriteUnsignedByte(j.GameMode); err != nil {
		return err
	}

	if err = w.WriteInt(j.Dimension); err != nil {
		return err
	}

	if err = w.WriteLong(j.HashedSeed); err != nil {
		return err
	}

	if err = w.WriteUnsignedByte(j.MaxPlayers); err != nil {
		return err
	}

	if err = w.WriteString(j.LevelType); err != nil {
		return err
	}

	if err = w.WriteVarInt(j.ViewDistance); err != nil {
		return err
	}

	if err = w.WriteBool(j.ReducedDebug); err != nil {
		return err
	}

	if err = w.WriteBool(j.EnableRespawnScreen); err != nil {
		return err
	}

	return nil
}

func (j *JoinGamePacket) Unmarshal(r PacketReader) error {
	var err error

	if j.PlayerID, err = r.ReadInt(); err != nil {
		return err
	}

	if j.GameMode, err = r.ReadUnsignedByte(); err != nil {
		return err
	}

	if j.Dimension, err = r.ReadInt(); err != nil {
		return err
	}

	if j.HashedSeed, err = r.ReadLong(); err != nil {
		return err
	}

	if j.MaxPlayers, err = r.ReadUnsignedByte(); err != nil {
		return err
	}

	if j.LevelType, err = r.ReadString(16); err != nil {
		return err
	}

	if j.ViewDistance, err = r.ReadVarInt(); err != nil {
		return err
	}

	if j.ReducedDebug, err = r.ReadBool(); err != nil {
		return err
	}

	if j.EnableRespawnScreen, err = r.ReadBool(); err != nil {
		return err
	}

	return nil
}
