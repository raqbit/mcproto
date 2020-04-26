package encoding

import (
	"encoding/binary"
	"io"
)

func WriteBool(value bool) []byte {
	var b byte
	if value {
		b = 0x01
	}
	return []byte{b}
}

func ReadBool(buff io.Reader) (bool, error) {
	var bl bool
	err := binary.Read(buff, binary.BigEndian, &bl)
	return bl, err
}
