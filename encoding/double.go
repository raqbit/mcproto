package encoding

import (
	"encoding/binary"
	"io"
	"math"
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
	_, err := w.Write(WriteDouble(*l))
	return err
}

// WriteDouble writes the passed Double to the writer
func WriteDouble(value Double) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(float64(value)))
	return buf
}

// ReadDouble reads a Double from the reader
func ReadDouble(buff io.Reader) (Double, error) {
	var float float64
	err := binary.Read(buff, binary.BigEndian, &float)
	return Double(float), err
}
