package mcproto

import (
	"encoding/json"
)

// https://wiki.vg/Protocol#Request
type CServerQueryPacket struct{}

func (r CServerQueryPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ServerBound,
		ConnectionState: StatusState,
	}
}

func (*CServerQueryPacket) String() string {
	return "ServerQuery"
}

func (*CServerQueryPacket) Marshal(_ PacketWriter) error {
	return nil
}

func (*CServerQueryPacket) Unmarshal(_ PacketReader) error {
	return nil
}

// https://wiki.vg/Protocol#Response
type SServerInfoPacket struct {
	Response ServerInfo
}

func (*SServerInfoPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ClientBound,
		ConnectionState: StatusState,
	}
}

func (*SServerInfoPacket) String() string {
	return "ServerInfo"
}

func (si *SServerInfoPacket) Marshal(w PacketWriter) error {
	var err error
	var response []byte

	if response, err = json.Marshal(si.Response); err != nil {
		return err
	}

	if err := w.WriteString(string(response)); err != nil {
		return err
	}

	return nil
}

func (si *SServerInfoPacket) Unmarshal(r PacketReader) error {
	var err error
	var response string

	if response, err = r.ReadMaxString(); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(response), &si.Response); err != nil {
		return err
	}

	return nil
}

// https://wiki.vg/Protocol#Ping
type CPingPacket struct {
	Payload int64
}

func (*CPingPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x01,
		Direction:       ServerBound,
		ConnectionState: StatusState,
	}
}

func (*CPingPacket) String() string {
	return "Ping"
}

func (p *CPingPacket) Marshal(w PacketWriter) error {
	if err := w.WriteLong(p.Payload); err != nil {
		return err
	}

	return nil
}

func (p *CPingPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Payload, err = r.ReadLong(); err != nil {
		return err
	}

	return nil
}

// https://wiki.vg/Protocol#Pong
type SPongPacket struct {
	Payload int64
}

func (*SPongPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x01,
		Direction:       ClientBound,
		ConnectionState: StatusState,
	}
}

func (*SPongPacket) String() string {
	return "Pong"
}

func (p *SPongPacket) Marshal(w PacketWriter) error {
	if err := w.WriteLong(p.Payload); err != nil {
		return err
	}

	return nil
}

func (p *SPongPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Payload, err = r.ReadLong(); err != nil {
		return err
	}

	return nil
}
