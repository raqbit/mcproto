package encoding

import (
	"encoding/binary"
	"io"
)

// Minecraft Protocol Byte type
type Byte int8

func (b *Byte) Decode(r io.Reader) error {
	bt, err := ReadByte(r)

	if err != nil {
		return err
	}

	*b = bt
	return nil
}

func (b *Byte) Encode(w io.Writer) error {
	_, err := w.Write(WriteByte(*b))
	return err
}

// WriteByte writes the passed Byte to the writer
func WriteByte(value Byte) []byte {
	return []byte{byte(value)}
}

// ReadByte reads a Byte from the reader
func ReadByte(r io.Reader) (Byte, error) {
	var b int8
	err := binary.Read(r, binary.BigEndian, &b)
	return Byte(b), err
}
