package encoding

import (
	"encoding/binary"
	"io"
)

type Byte int8

func (b *Byte) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, b)
}

func (b *Byte) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, b)
}
