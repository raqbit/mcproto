package mcproto

//go:generate go run ../tools/genpacket/genpacket.go -packet=LoginDisconnectPacket -output=login_disconnect_gen.go

const LoginDisconnectPacketID int32 = 0x00

// https://wiki.vg/Protocol#Disconnect_.28login.29
type LoginDisconnectPacket struct {
	Reason TextComponent
}

func (*LoginDisconnectPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              LoginDisconnectPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStateLogin,
	}
}

func (*LoginDisconnectPacket) String() string {
	return "LoginDisconnect"
}
