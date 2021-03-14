package encoding

import (
	"encoding/binary"
	"io"
)

type Bool bool

func (b *Bool) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, b)
}

func (b *Bool) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, b)
}
