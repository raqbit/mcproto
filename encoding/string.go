package encoding

import (
	"errors"
	"io"
	"math"
)

var (
	// ErrStringLengthTooLarge is returned when the length of a string
	// was too large (More than 32767 bytes)
	ErrStringLengthTooLarge = errors.New("length of String is too large")
)

// Minecraft Protocol String type
type String string

func (st *String) Decode(r io.Reader) error {
	s, err := ReadString(r)

	if err != nil {
		return err
	}

	*st = s
	return nil
}

func (st *String) Encode(w io.Writer) error {
	return WriteString(w, *st)
}

// WriteString writes a VarInt prefixed utf-8 string to the
// writer.
func WriteString(w io.Writer, str String) error {

	// Creating buffer from string
	b := []byte(str)

	// Writing string length as varint to output buffer
	err := WriteVarInt(w, VarInt(len(b)))

	if err != nil {
		return err
	}

	// Writing string to buffer
	_, err = w.Write(b)

	return err
}

// ReadString reads a VarInt prefixed utf-8 string to the
// reader. It uses io.ReadFull to ensure all bytes are read.
func ReadString(r io.Reader) (String, error) {

	// Reading string size encoded as VarInt
	l, err := ReadVarInt(r)

	if err != nil {
		return "", nil
	}

	// Checking if string size is valid
	if l < 0 || l > math.MaxInt16 {
		return "", ErrStringLengthTooLarge
	}

	// Creating string buffer with the specified size
	stringBuff := make([]byte, int(l))

	// Reading l amount of bytes from the buffer
	_, err = io.ReadFull(r, stringBuff)

	return String(stringBuff), err
}
