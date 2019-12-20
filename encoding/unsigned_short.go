package encoding

import (
	"encoding/binary"
	"io"
)

// Minecraft Protocol UnsignedShort type
type UnsignedShort uint16

func (us *UnsignedShort) Decode(r io.Reader) error {
	sht, err := ReadUnsignedShort(r)

	if err != nil {
		return err
	}

	*us = sht
	return nil
}

func (us *UnsignedShort) Encode(w io.Writer) error {
	return WriteUnsignedShort(w, *us)
}

// WriteUnsignedShort writes the passed UnsignedShort to the writer
func WriteUnsignedShort(buff io.Writer, value UnsignedShort) error {
	return binary.Write(buff, binary.BigEndian, uint16(value))
}

//ReadUnsignedShort reads an UnsignedShort from the reader
func ReadUnsignedShort(buff io.Reader) (UnsignedShort, error) {
	var short uint16
	err := binary.Read(buff, binary.BigEndian, &short)
	return UnsignedShort(short), err
}

