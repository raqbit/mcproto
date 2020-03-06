package encoding

import (
	"encoding/binary"
	"io"
)

// Minecraft Protocol Double type
type Double float64

func (l *Double) Decode(r io.Reader) error {
	Double, err := ReadDouble(r)

	if err != nil {
		return err
	}

	*l = Double
	return nil
}

func (l *Double) Encode(w io.Writer) error {
	return WriteDouble(w, *l)
}

// WriteDouble writes the passed Double to the writer
func WriteDouble(buff io.Writer, value Double) error {
	return binary.Write(buff, binary.BigEndian, float64(value))
}

// ReadDouble reads a Double from the reader
func ReadDouble(buff io.Reader) (Double, error) {
	var float float64
	err := binary.Read(buff, binary.BigEndian, &float)
	return Double(float), err
}
