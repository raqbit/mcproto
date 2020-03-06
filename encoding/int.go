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
	_, err := w.Write(WriteInt(*i))
	return err
}

// WriteInt writes the passed integer to the writer.
func WriteInt(value Int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(value))
	return buf
}

// ReadInt reads an integer from the reader.
func ReadInt(r io.Reader) (Int, error) {
	var readInt int32
	err := binary.Read(r, binary.BigEndian, &readInt)
	return Int(readInt), err
}
