package mcproto

const ClientSettingsPacketID int32 = 0x05

type ClientSettingsPacket struct {
	Lang               string
	ViewDistance       int8
	ChatVisibility     int32
	EnableChatColors   bool
	DisplayedSkinParts uint8
	MainHand           int32
}

func (*ClientSettingsPacket) String() string {
	return "ClientSettings"
}

func (*ClientSettingsPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ClientSettingsPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStatePlay,
	}
}

func (cs *ClientSettingsPacket) Marshal(w PacketWriter) error {
	var err error

	if err = w.WriteString(cs.Lang); err != nil {
		return err
	}

	if err = w.WriteByte(cs.ViewDistance); err != nil {
		return err
	}

	if err = w.WriteVarInt(cs.ChatVisibility); err != nil {
		return err
	}

	if err = w.WriteBool(cs.EnableChatColors); err != nil {
		return err
	}

	if err = w.WriteUnsignedByte(cs.DisplayedSkinParts); err != nil {
		return err
	}

	if err = w.WriteVarInt(cs.MainHand); err != nil {
		return err
	}

	return nil
}

func (cs *ClientSettingsPacket) Unmarshal(r PacketReader) error {
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
