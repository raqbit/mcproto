package encoding

import (
	"errors"
	"io"
	"math"
)

var (
	// ErrStringLengthTooLarge is returned when the length of a string
	// was too large
	ErrStringLengthTooLarge = errors.New("length of String is too large")
)

type String string

func (s *String) Write(w io.Writer) error {
	var err error
	l := VarInt(len(*s))

	if err = l.Write(w); err != nil {
		return err
	}

	_, err = w.Write([]byte(*s))
	return err
}

func (s *String) Read(r io.Reader) error {
	var err error
	var l VarInt

	if err = l.Read(r); err != nil {
		return err
	}

	// Checking if string size is valid
	if l < 0 || int(l) > math.MaxInt16 {
		return ErrStringLengthTooLarge
	}

	stringBuff := make([]byte, int(l))
	_, err = io.ReadAtLeast(r, stringBuff, int(l))

	if err != nil {
		return err
	}

	*s = String(stringBuff)

	return nil
}
