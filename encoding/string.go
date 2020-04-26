package encoding

import (
	"errors"
	"io"
)

var (
	// ErrStringLengthTooLarge is returned when the length of a string
	// was too large
	ErrStringLengthTooLarge = errors.New("length of String is too large")
)

func WriteString(str string) []byte {
	b := WriteVarInt(int32(len(str)))
	b = append(b, []byte(str)...)
	return b
}

func ReadString(r io.Reader, maxLength int32) (string, error) {
	l, err := ReadVarInt(r)
	if err != nil {
		return "", nil
	}

	// Checking if string size is valid
	if l < 0 || l > maxLength {
		return "", ErrStringLengthTooLarge
	}

	stringBuff := make([]byte, int(l))
	_, err = io.ReadFull(r, stringBuff)
	return string(stringBuff), err
}
