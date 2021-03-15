package encoding

import (
	"encoding/binary"
	"io"
)

type UnsignedByte uint8

func (ub *UnsignedByte) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, ub)
}

func (ub *UnsignedByte) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, ub)
}
