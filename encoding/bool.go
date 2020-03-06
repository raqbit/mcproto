package encoding

import (
	"encoding/binary"
	"io"
)

// Minecraft Protocol Bool type
type Bool bool

func (b *Bool) Decode(r io.Reader) error {
	readBool, err := ReadBool(r)

	if err != nil {
		return err
	}

	*b = readBool
	return nil
}

func (b *Bool) Encode(w io.Writer) error {
	return WriteBool(w, *b)
}

// WriteBool writes the passed Bool to the writer
func WriteBool(buff io.Writer, value Bool) error {
	return binary.Write(buff, binary.BigEndian, bool(value))
}

// ReadBool reads a Bool from the reader
func ReadBool(buff io.Reader) (Bool, error) {
	var bl bool
	err := binary.Read(buff, binary.BigEndian, &bl)
	return Bool(bl), err
}
