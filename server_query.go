package mcproto

const ServerQueryPacketID int32 = 0x00

// https://wiki.vg/Protocol#Request
type ServerQueryPacket struct{}

func (r ServerQueryPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ServerQueryPacketID,
		Direction:       ServerBound,
		ConnectionState: ConnectionStateStatus,
	}
}

func (*ServerQueryPacket) String() string {
	return "ServerQuery"
}

func (*ServerQueryPacket) Marshal(_ PacketWriter) error {
	return nil
}

func (*ServerQueryPacket) Unmarshal(_ PacketReader) error {
	return nil
}
