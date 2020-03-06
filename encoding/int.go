package encoding

import (
	"encoding/binary"
	"io"
)

// Minecraft Protocol Int type
type Int int32

func (i *Int) Decode(r io.Reader) error {
	readInt, err := ReadInt(r)

	if err != nil {
		return err
	}

	*i = readInt
	return nil
}

func (i *Int) Encode(w io.Writer) error {
	return WriteInt(w, *i)
}

// WriteInt writes the passed integer to the writer.
func WriteInt(w io.Writer, value Int) error {
	return binary.Write(w, binary.BigEndian, int32(value))
}

// ReadInt reads an integer from the reader.
func ReadInt(r io.Reader) (Int, error) {
	var readInt int32
	err := binary.Read(r, binary.BigEndian, &readInt)
	return Int(readInt), err
}
