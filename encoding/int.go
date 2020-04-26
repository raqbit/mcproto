package encoding

import (
	"encoding/binary"
	"io"
)

func WriteInt(value int32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(value))
	return buf
}

func ReadInt(r io.Reader) (int32, error) {
	var readInt int32
	err := binary.Read(r, binary.BigEndian, &readInt)
	return readInt, err
}
