package mcproto

import "github.com/Raqbit/mcproto/packet"

//go:generate stringer -type=Side -output side_string.go -linecomment

// Side is the side of a Minecraft connection
type Side uint8

const (
	ClientSide Side = iota // Client
	ServerSide             // Server
)

func (s Side) ReadDirection() packet.Direction {
	if s == ServerSide {
		return packet.ServerBound
	} else {
		return packet.ClientBound
	}
}

func (s Side) WriteDirection() packet.Direction {
	if s == ServerSide {
		return packet.ClientBound
	} else {
		return packet.ServerBound
	}
}
