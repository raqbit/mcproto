package types

//go:generate stringer -type=Side -output side_string.go -linecomment

// Side is the side of a Minecraft connection
type Side uint8

const (
	ClientSide Side = iota // Client
	ServerSide             // Server
)
