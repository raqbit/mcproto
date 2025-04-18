package mcproto

//go:generate stringer -type=Side -output side_string.go -linecomment

// Side is the side of a Minecraft connection
type Side uint8

const (
	ClientSide Side = iota // Client
	ServerSide             // Server
)

func (s Side) readDirection() PacketDirection {
	if s == ServerSide {
		return serverBound
	} else {
		return clientBound
	}
}

func (s Side) writeDirection() PacketDirection {
	if s == ServerSide {
		return clientBound
	} else {
		return serverBound
	}
}
