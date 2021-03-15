package encoding

import (
	"encoding/binary"
	"io"
)

type Double float64

func (d *Double) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, d)
}

func (d *Double) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, d)
}
