package encoding

import (
	"encoding/binary"
	"io"
)

func WriteByte(value byte) []byte {
	return []byte{value}
}

func ReadByte(r io.Reader) (byte, error) {
	var b byte
	err := binary.Read(r, binary.BigEndian, &b)
	return b, err
}
