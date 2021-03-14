package encoding

import (
	"encoding/binary"
	"github.com/google/uuid"
	"io"
)

type UUID uuid.UUID

func (u *UUID) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, u)
}

func (u *UUID) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, u)

}
