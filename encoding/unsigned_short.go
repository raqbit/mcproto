package encoding

import (
	"encoding/binary"
	"io"
)

func WriteUnsignedShort(value uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, value)
	return buf
}

func ReadUnsignedShort(buff io.Reader) (uint16, error) {
	var short uint16
	err := binary.Read(buff, binary.BigEndian, &short)
	return short, err
}
