package packet

import (
	enc "github.com/Raqbit/mcproto/encoding"
)

// https://wiki.vg/Protocol#Handshake
type Handshake struct {
	ProtoVer   enc.VarInt
	ServerAddr enc.String
	ServerPort enc.UnsignedShort
	NextState  enc.VarInt
}

func (Handshake) ID() int {
	return 0x00
}
