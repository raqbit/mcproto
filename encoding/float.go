package encoding

import (
	"encoding/binary"
	"io"
	"math"
)

func WriteFloat(value float32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, math.Float32bits(value))
	return buf
}

func ReadFloat(buff io.Reader) (float32, error) {
	var float float32
	err := binary.Read(buff, binary.BigEndian, &float)
	return float, err
}
