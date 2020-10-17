package mcproto

import (
	"encoding/json"
)

// https://wiki.vg/Protocol#Request
type SServerQueryPacket struct{}

func (r SServerQueryPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ServerBound,
		ConnectionState: StatusState,
	}
}

func (*SServerQueryPacket) String() string {
	return "ServerQuery"
}

func (*SServerQueryPacket) Marshal(_ PacketWriter) error {
	return nil
}

func (*SServerQueryPacket) Unmarshal(_ PacketReader) error {
	return nil
}

// https://wiki.vg/Protocol#Response
type CServerInfoPacket struct {
	Response ServerInfo
}

func (*CServerInfoPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x00,
		Direction:       ClientBound,
		ConnectionState: StatusState,
	}
}

func (*CServerInfoPacket) String() string {
	return "ServerInfo"
}

func (si *CServerInfoPacket) Marshal(w PacketWriter) error {
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

func (si *CServerInfoPacket) Unmarshal(r PacketReader) error {
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
type SPingPacket struct {
	Payload int64
}

func (*SPingPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x01,
		Direction:       ServerBound,
		ConnectionState: StatusState,
	}
}

func (*SPingPacket) String() string {
	return "Ping"
}

func (p *SPingPacket) Marshal(w PacketWriter) error {
	if err := w.WriteLong(p.Payload); err != nil {
		return err
	}

	return nil
}

func (p *SPingPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Payload, err = r.ReadLong(); err != nil {
		return err
	}

	return nil
}

// https://wiki.vg/Protocol#Pong
type CPongPacket struct {
	Payload int64
}

func (*CPongPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              0x01,
		Direction:       ClientBound,
		ConnectionState: StatusState,
	}
}

func (*CPongPacket) String() string {
	return "Pong"
}

func (p *CPongPacket) Marshal(w PacketWriter) error {
	if err := w.WriteLong(p.Payload); err != nil {
		return err
	}

	return nil
}

func (p *CPongPacket) Unmarshal(r PacketReader) error {
	var err error

	if p.Payload, err = r.ReadLong(); err != nil {
		return err
	}

	return nil
}
