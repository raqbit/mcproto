package encoding

import (
	"encoding/binary"
	"io"
)

type Long int64

func (l *Long) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, l)
}

func (l *Long) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, l)
}
