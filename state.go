package mcproto

//go:generate stringer -type=ConnectionState -output state_string.go -linecomment

// The state of a connection
type ConnectionState uint8

const (
	ConnectionStateHandshake ConnectionState = 0x00 // Handshake
	ConnectionStateStatus    ConnectionState = 0x01 // Status
	ConnectionStateLogin     ConnectionState = 0x02 // Login
	ConnectionStatePlay      ConnectionState = 0x03 // Play
)
