package encoding

import "io"

func ReadUnsignedByte(r io.Reader) (uint8, error) {
	var bytes [1]byte
	_, err := r.Read(bytes[:1])
	return bytes[0], err
}

func WriteUnsignedByte(value uint8) []byte {
	return []byte{value}
}
