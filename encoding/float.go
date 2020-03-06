package encoding

import (
	"encoding/binary"
	"io"
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
	return WriteFloat(w, *l)
}

// WriteFloat writes the passed Float to the writer
func WriteFloat(buff io.Writer, value Float) error {
	return binary.Write(buff, binary.BigEndian, float32(value))
}

// ReadFloat reads a Float from the reader
func ReadFloat(buff io.Reader) (Float, error) {
	var float float32
	err := binary.Read(buff, binary.BigEndian, &float)
	return Float(float), err
}
