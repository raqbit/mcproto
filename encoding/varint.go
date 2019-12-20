package encoding

import (
	"errors"
	"io"
)

const (
	// Maximum size for a varint
	VarIntMaxByteSize = 5
)

var (
	// ErrVarIntTooLarge is returned when a read varint was too large
	ErrVarIntTooLarge = errors.New("VarInt too large")
)

// Minecraft Protocol VarInt type
type VarInt int32

func (vi *VarInt) Decode(r io.Reader) error {
	i, err := ReadVarInt(r)

	if err != nil {
		return err
	}

	*vi = i
	return nil
}

func (vi *VarInt) Encode(w io.Writer) error {
	return WriteVarInt(w, *vi)
}

// WriteVarInt writes the passed VarInt encoded integer to the writer.
func WriteVarInt(w io.Writer, value VarInt) error {
	for cont := true; cont; cont = value != 0 {
		temp := byte(value & 0x7F)

		// Casting value to a uint to get a logical shift
		value = VarInt(uint32(value) >> 7)

		if value != 0 {
			temp |= 0x80
		}

		if err := WriteUnsignedByte(w, UnsignedByte(temp)); err != nil {
			return err
		}
	}

	return nil
}

// ReadVarInt reads a VarInt encoded integer from the reader.
func ReadVarInt(r io.Reader) (VarInt, error) {
	var numRead uint
	var result int32
	var read UnsignedByte

	for cont := true; cont; cont = (read & 0x80) != 0 {
		var err error
		read, err = ReadUnsignedByte(r)

		if err != nil {
			return 0, err
		}

		value := read & 0x7F

		result |= int32(value) << (7 * numRead)

		numRead++

		if numRead > VarIntMaxByteSize {
			return 0, ErrVarIntTooLarge
		}
	}

	return VarInt(result), nil
}
