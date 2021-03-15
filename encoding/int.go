package encoding

import (
	"encoding/binary"
	"io"
)

type Int int32

func (i *Int) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, i)
}

func (i *Int) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, i)
}
