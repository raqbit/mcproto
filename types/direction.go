package types

//go:generate stringer -type=PacketDirection -output direction_string.go -linecomment

// PacketDirection is the direction of a packet
type PacketDirection byte

const (
	ClientBound PacketDirection = iota // client-bound
	ServerBound                        // server-bound
)
