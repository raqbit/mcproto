package packet

//go:generate stringer -type=PacketDirection -output direction_string.go -linecomment

// Direction is the direction of a packet
type Direction byte

const (
	ClientBound Direction = iota // client-bound
	ServerBound                  // server-bound
)
