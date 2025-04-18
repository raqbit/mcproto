package encoding

import "io"

// MinecraftEncodable types can be encoded for the Minecraft protocol
type MinecraftEncodable interface {
	Write(w io.Writer) error
}

// MinecraftDecodable types can be decoded for the Minecraft protocol
type MinecraftDecodable interface {
	Read(r io.Reader) error
}
