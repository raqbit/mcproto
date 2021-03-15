package encoding

import (
	"github.com/google/uuid"
	"io"
)

type UUID uuid.UUID

func (u *UUID) Write(w io.Writer) error {
	str := String(uuid.UUID(*u).String())
	return str.Write(w)
}

func (u *UUID) Read(r io.Reader) error {
	var err error
	var str String

	if err = str.Read(r); err != nil {
		return err
	}

	uid, err := uuid.Parse(string(str))

	if err != nil {
		return err
	}

	*u = UUID(uid)

	return nil
}
