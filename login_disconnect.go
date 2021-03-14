package mcproto

import "encoding/json"

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

func (d *LoginDisconnectPacket) Marshal(w PacketWriter) error {
	var err error
	var reason []byte

	if reason, err = json.Marshal(d.Reason); err != nil {
		return err
	}

	if err = w.WriteString(string(reason)); err != nil {
		return err
	}

	return nil
}

func (d *LoginDisconnectPacket) Unmarshal(r PacketReader) error {
	var err error
	var reason string

	if reason, err = r.ReadMaxString(); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(reason), &d.Reason); err != nil {
		return err
	}

	return nil
}
