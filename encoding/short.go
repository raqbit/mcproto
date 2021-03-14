package encoding

import (
	"encoding/binary"
	"io"
)

type Short int16

func (s *Short) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, s)
}

func (s *Short) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, s)
}
