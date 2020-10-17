package mcproto

import (
	"encoding/json"
)

// https://wiki.vg/Protocol#Chat_Message_.28clientbound.29
type SChatMessagePacket struct {
	Message  TextComponent
	Position int8
}

func (mp *SChatMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x0f,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (*SChatMessagePacket) String() string {
	return "ChatMessage"
}

func (mp *SChatMessagePacket) Marshal(w PacketWriter) error {
	var err error
	var response []byte

	if response, err = json.Marshal(mp.Message); err != nil {
		return err
	}

	if err := w.WriteString(string(response)); err != nil {
		return err
	}

	if err := w.WriteByte(mp.Position); err != nil {
		return err
	}

	return nil
}

func (mp *SChatMessagePacket) Unmarshal(r PacketReader) error {
	var err error
	var message string

	if message, err = r.ReadMaxString(); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(message), &mp.Message); err != nil {
		return err
	}

	if mp.Position, err = r.ReadByte(); err != nil {
		return err
	}

	return nil
}
