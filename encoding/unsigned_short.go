package encoding

import (
	"encoding/binary"
	"io"
)

type UnsignedShort uint16

func (us *UnsignedShort) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, us)
}

func (us *UnsignedShort) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, us)
}
