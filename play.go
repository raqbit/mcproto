package mcproto

import (
	"bytes"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

type JoinGamePacket struct {
	EntityID            enc.Int
	GameMode            enc.UnsignedByte
	Dimension           enc.Int
	HashedSeed          enc.Long
	MaxPlayers          enc.UnsignedByte
	LevelType           enc.String
	ViewDistance        enc.VarInt
	ReducedDebug        enc.Bool
	EnableRespawnScreen enc.Bool
}

func (JoinGamePacket) String() string {
	return "JoinGame"
}

func (JoinGamePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x26,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (j JoinGamePacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := j.EntityID.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.GameMode.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.Dimension.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.HashedSeed.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.MaxPlayers.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.LevelType.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.ViewDistance.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.ReducedDebug.Encode(buffer); err != nil {
		return nil, err
	}

	if err := j.EnableRespawnScreen.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (JoinGamePacket) Unmarshal(r io.Reader) (Packet, error) {
	j := &JoinGamePacket{}

	if err := j.EntityID.Decode(r); err != nil {
		return nil, err
	}

	if err := j.GameMode.Decode(r); err != nil {
		return nil, err
	}

	if err := j.Dimension.Decode(r); err != nil {
		return nil, err
	}

	if err := j.HashedSeed.Decode(r); err != nil {
		return nil, err
	}

	if err := j.MaxPlayers.Decode(r); err != nil {
		return nil, err
	}

	if err := j.LevelType.Decode(r); err != nil {
		return nil, err
	}

	if err := j.ViewDistance.Decode(r); err != nil {
		return nil, err
	}

	if err := j.ReducedDebug.Decode(r); err != nil {
		return nil, err
	}

	if err := j.EnableRespawnScreen.Decode(r); err != nil {
		return nil, err
	}

	return j, nil
}

type PluginMessagePacket struct {
	Channel enc.String
	Data    []byte
}

func (PluginMessagePacket) String() string {
	return "PluginMessage"
}

func (PluginMessagePacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x19,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (p PluginMessagePacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := p.Channel.Encode(buffer); err != nil {
		return nil, err
	}

	if _, err := buffer.Write(p.Data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (PluginMessagePacket) Unmarshal(io.Reader) (Packet, error) {
	panic("implement me")
}

type ClientSettingsPacket struct {
	Locale             enc.String
	ViewDistance       enc.Byte
	ChatMode           enc.VarInt
	ChatColors         enc.Bool
	DisplayedSkinParts enc.UnsignedByte
	MainHand           enc.VarInt
}

func (ClientSettingsPacket) String() string {
	return "ClientSettings"
}

func (ClientSettingsPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x05,
		Direction:       ServerBound,
		ConnectionState: PlayState,
	}
}

func (cs ClientSettingsPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := cs.Locale.Encode(buffer); err != nil {
		return nil, err
	}

	if err := cs.ViewDistance.Encode(buffer); err != nil {
		return nil, err
	}

	if err := cs.ChatMode.Encode(buffer); err != nil {
		return nil, err
	}

	if err := cs.ChatColors.Encode(buffer); err != nil {
		return nil, err
	}

	if err := cs.DisplayedSkinParts.Encode(buffer); err != nil {
		return nil, err
	}

	if err := cs.MainHand.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (ClientSettingsPacket) Unmarshal(r io.Reader) (Packet, error) {
	cs := &ClientSettingsPacket{}

	if err := cs.Locale.Decode(r); err != nil {
		return nil, err
	}

	if err := cs.ViewDistance.Decode(r); err != nil {
		return nil, err
	}

	if err := cs.ChatMode.Decode(r); err != nil {
		return nil, err
	}

	if err := cs.ChatColors.Decode(r); err != nil {
		return nil, err
	}

	if err := cs.DisplayedSkinParts.Decode(r); err != nil {
		return nil, err
	}

	if err := cs.MainHand.Decode(r); err != nil {
		return nil, err
	}

	return cs, nil
}

type PlayerPositionAndLookPacket struct {
	X          enc.Double
	Y          enc.Double
	Z          enc.Double
	Yaw        enc.Float
	Pitch      enc.Float
	Flags      enc.Byte
	TeleportID enc.VarInt
}

func (PlayerPositionAndLookPacket) String() string {
	return "PlayerPositionAndLook"
}

func (PlayerPositionAndLookPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x36,
		Direction:       ClientBound,
		ConnectionState: PlayState,
	}
}

func (p PlayerPositionAndLookPacket) Marshal() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := p.X.Encode(buffer); err != nil {
		return nil, err
	}

	if err := p.Y.Encode(buffer); err != nil {
		return nil, err
	}

	if err := p.Z.Encode(buffer); err != nil {
		return nil, err
	}

	if err := p.Yaw.Encode(buffer); err != nil {
		return nil, err
	}

	if err := p.Pitch.Encode(buffer); err != nil {
		return nil, err
	}

	if err := p.Flags.Encode(buffer); err != nil {
		return nil, err
	}

	if err := p.TeleportID.Encode(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (PlayerPositionAndLookPacket) Unmarshal(r io.Reader) (Packet, error) {
	p := &PlayerPositionAndLookPacket{}

	if err := p.X.Decode(r); err != nil {
		return nil, err
	}

	if err := p.Y.Decode(r); err != nil {
		return nil, err
	}

	if err := p.Z.Decode(r); err != nil {
		return nil, err
	}

	if err := p.Yaw.Decode(r); err != nil {
		return nil, err
	}

	if err := p.Pitch.Decode(r); err != nil {
		return nil, err
	}

	if err := p.Flags.Decode(r); err != nil {
		return nil, err
	}

	if err := p.TeleportID.Decode(r); err != nil {
		return nil, err
	}

	return p, nil
}
