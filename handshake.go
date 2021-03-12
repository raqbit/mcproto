package mcproto

const HandshakePacketID int32 = 0x00

// https://wiki.vg/Protocol#Handshake
type HandshakePacket struct {
	ProtoVer   int32
	ServerAddr string
	ServerPort uint16
	NextState  ConnectionState
}

func (h *HandshakePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              HandshakePacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateHandshake,
	}
}

func (*HandshakePacket) String() string {
	return "Handshake"
}

func (h *HandshakePacket) Marshal(w PacketWriter) error {
	var err error

	if err = w.WriteVarInt(h.ProtoVer); err != nil {
		return err
	}

	if err = w.WriteString(h.ServerAddr); err != nil {
		return err
	}

	if err = w.WriteUnsignedShort(h.ServerPort); err != nil {
		return err
	}

	if err = w.WriteVarInt(int32(h.NextState)); err != nil {
		return err
	}

	return nil
}

func (h *HandshakePacket) Unmarshal(r PacketReader) error {
	var err error

	if h.ProtoVer, err = r.ReadVarInt(); err != nil {
		return err
	}

	if h.ServerAddr, err = r.ReadString(255); err != nil {
		return err
	}

	if h.ServerPort, err = r.ReadUnsignedShort(); err != nil {
		return err
	}

	var nextState int32
	if nextState, err = r.ReadVarInt(); err != nil {
		return err
	}
	h.NextState = ConnectionState(nextState)

	return nil
}
