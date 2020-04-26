package encoding

import (
	"encoding/binary"
	"io"
)

func WriteByte(value int8) []byte {
	return []byte{byte(value)}
}

func ReadByte(r io.Reader) (int8, error) {
	var b int8
	err := binary.Read(r, binary.BigEndian, &b)
	return b, err
}
