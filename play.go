package mcproto

type SJoinGamePacket struct {
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

func (*SJoinGamePacket) String() string {
	return "JoinGame"
}

func (*SJoinGamePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x26,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (j *SJoinGamePacket) Marshal(w PacketWriter) error {
	if err := w.WriteInt(j.PlayerID); err != nil {
		return err
	}

	if err := w.WriteUnsignedByte(j.GameMode); err != nil {
		return err
	}

	if err := w.WriteInt(j.Dimension); err != nil {
		return err
	}

	if err := w.WriteLong(j.HashedSeed); err != nil {
		return err
	}

	if err := w.WriteUnsignedByte(j.MaxPlayers); err != nil {
		return err
	}

	if err := w.WriteString(j.LevelType); err != nil {
		return err
	}

	if err := w.WriteVarInt(j.ViewDistance); err != nil {
		return err
	}

	if err := w.WriteBool(j.ReducedDebug); err != nil {
		return err
	}

	if err := w.WriteBool(j.EnableRespawnScreen); err != nil {
		return err
	}

	return nil
}

func (j *SJoinGamePacket) Unmarshal(r PacketReader) error {
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

type PluginMessagePacket struct {
	Channel *ResourceLocation
	Data    *PacketBuffer
}

func (*PluginMessagePacket) String() string {
	return "PluginMessage"
}

func (*PluginMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x19,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (p *PluginMessagePacket) Marshal(w PacketWriter) error {
	if err := w.WriteResourceLocation(p.Channel); err != nil {
		return err
	}

	if err := w.WriteBytes(p.Data); err != nil {
		return err
	}

	return nil
}

func (p *PluginMessagePacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Channel, err = r.ReadResourceLocation(); err != nil {
		return err
	}

	return nil
}

type CClientSettingsPacket struct {
	Lang               string
	ViewDistance       int8
	ChatVisibility     int32
	EnableChatColors   bool
	DisplayedSkinParts uint8
	MainHand           int32
}

func (*CClientSettingsPacket) String() string {
	return "ClientSettings"
}

func (*CClientSettingsPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x05,
		Direction:       ServerBound,
		ConnectionState: PlayState,
	}
}

func (cs *CClientSettingsPacket) Marshal(w PacketWriter) error {
	if err := w.WriteString(cs.Lang); err != nil {
		return err
	}

	if err := w.WriteByte(cs.ViewDistance); err != nil {
		return err
	}

	if err := w.WriteVarInt(cs.ChatVisibility); err != nil {
		return err
	}

	if err := w.WriteBool(cs.EnableChatColors); err != nil {
		return err
	}

	if err := w.WriteUnsignedByte(cs.DisplayedSkinParts); err != nil {
		return err
	}

	if err := w.WriteVarInt(cs.MainHand); err != nil {
		return err
	}

	return nil
}

func (cs *CClientSettingsPacket) Unmarshal(r PacketReader) error {
	var err error

	if cs.Lang, err = r.ReadString(16); err != nil {
		return err
	}

	if cs.ViewDistance, err = r.ReadByte(); err != nil {
		return err
	}

	if cs.ChatVisibility, err = r.ReadVarInt(); err != nil {
		return err
	}

	if cs.EnableChatColors, err = r.ReadBool(); err != nil {
		return err
	}

	if cs.DisplayedSkinParts, err = r.ReadUnsignedByte(); err != nil {
		return err
	}

	if cs.MainHand, err = r.ReadVarInt(); err != nil {
		return err
	}

	return nil
}

type SPlayerPositionLookPacket struct {
	X          float64
	Y          float64
	Z          float64
	Yaw        float32
	Pitch      float32
	Flags      uint8
	TeleportID int32
}

func (*SPlayerPositionLookPacket) String() string {
	return "PlayerPositionAndLook"
}

func (*SPlayerPositionLookPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x36,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (p *SPlayerPositionLookPacket) Marshal(w PacketWriter) error {
	if err := w.WriteDouble(p.X); err != nil {
		return err
	}

	if err := w.WriteDouble(p.Y); err != nil {
		return err
	}

	if err := w.WriteDouble(p.Z); err != nil {
		return err
	}

	if err := w.WriteFloat(p.Yaw); err != nil {
		return err
	}

	if err := w.WriteFloat(p.Pitch); err != nil {
		return err
	}

	if err := w.WriteUnsignedByte(p.Flags); err != nil {
		return err
	}

	if err := w.WriteVarInt(p.TeleportID); err != nil {
		return err
	}

	return nil
}

func (p *SPlayerPositionLookPacket) Unmarshal(r PacketReader) error {
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
