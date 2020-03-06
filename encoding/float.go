package encoding

import (
	"encoding/binary"
	"io"
	"math"
)

// Minecraft Protocol Float type
type Float float32

func (l *Float) Decode(r io.Reader) error {
	Float, err := ReadFloat(r)

	if err != nil {
		return err
	}

	*l = Float
	return nil
}

func (l *Float) Encode(w io.Writer) error {
	_, err := w.Write(WriteFloat(*l))
	return err
}

// WriteFloat writes the passed Float to the writer
func WriteFloat(value Float) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, math.Float32bits(float32(value)))
	return buf
}

// ReadFloat reads a Float from the reader
func ReadFloat(buff io.Reader) (Float, error) {
	var float float32
	err := binary.Read(buff, binary.BigEndian, &float)
	return Float(float), err
}
