package encoding

import (
	"encoding/binary"
	"io"
	"math"
)

func WriteDouble(value float64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(value))
	return buf
}

func ReadDouble(buff io.Reader) (float64, error) {
	var float float64
	err := binary.Read(buff, binary.BigEndian, &float)
	return float, err
}
