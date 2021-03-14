package encoding

import (
	"encoding/binary"
	"io"
)

type Float float32

func (f *Float) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, f)
}

func (f *Float) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, f)
}
