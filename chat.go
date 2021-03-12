package mcproto

import (
	"encoding/json"
)

const ChatMessagePacketID = 0x0f

// https://wiki.vg/Protocol#Chat_Message_.28clientbound.29
type ChatMessagePacket struct {
	Message  TextComponent
	Position int8
}

func (mp *ChatMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ChatMessagePacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStatePlay,
	}
}

func (*ChatMessagePacket) String() string {
	return "ChatMessage"
}

func (mp *ChatMessagePacket) Marshal(w PacketWriter) error {
	var err error
	var response []byte

	if response, err = json.Marshal(mp.Message); err != nil {
		return err
	}

	if err = w.WriteString(string(response)); err != nil {
		return err
	}

	if err = w.WriteByte(mp.Position); err != nil {
		return err
	}

	return nil
}

func (mp *ChatMessagePacket) Unmarshal(r PacketReader) error {
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
