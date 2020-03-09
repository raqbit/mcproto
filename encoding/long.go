package encoding

import (
	"encoding/binary"
	"io"
)

func WriteLong(value int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(value))
	return buf
}

func ReadLong(buff io.Reader) (int64, error) {
	var long int64
	err := binary.Read(buff, binary.BigEndian, &long)
	return long, err
}
