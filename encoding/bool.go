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
	_, err := w.Write(WriteBool(*b))
	return err
}

// WriteBool writes the passed Bool to the writer
func WriteBool(value Bool) []byte {
	var b byte
	if value {
		b = 0x01
	}
	return []byte{b}
}

// ReadBool reads a Bool from the reader
func ReadBool(buff io.Reader) (Bool, error) {
	var bl bool
	err := binary.Read(buff, binary.BigEndian, &bl)
	return Bool(bl), err
}
