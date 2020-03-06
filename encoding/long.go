package encoding

import (
	"encoding/binary"
	"io"
)

// Minecraft Protocol Long type
type Long int64

func (l *Long) Decode(r io.Reader) error {
	long, err := ReadLong(r)

	if err != nil {
		return err
	}

	*l = long
	return nil
}

func (l *Long) Encode(w io.Writer) error {
	_, err := w.Write(WriteLong(*l))
	return err
}

// WriteLong writes the passed Long to the writer
func WriteLong(value Long) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(value))
	return buf
}

// ReadLong reads a Long from the reader
func ReadLong(buff io.Reader) (Long, error) {
	var long int64
	err := binary.Read(buff, binary.BigEndian, &long)
	return Long(long), err
}
