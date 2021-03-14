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

type VarInt int32

func (vi *VarInt) Write(w io.Writer) error {
	value := *vi
	for cont := true; cont; cont = value != 0 {
		temp := UnsignedByte(value & 0x7F)

		// Casting value to a uint to get a logical shift
		value = VarInt(uint32(value) >> 7)

		if value != 0 {
			temp |= 0x80
		}

		if err := temp.Write(w); err != nil {
			return err
		}
	}

	return nil
}

func (vi *VarInt) Read(r io.Reader) error {
	var numRead int
	var result int32
	var read UnsignedByte

	for cont := true; cont; cont = (read & 0x80) != 0 {
		var err error
		err = read.Read(r)

		if err != nil {
			return err
		}

		value := read & 0x7F

		result |= int32(value) << (7 * numRead)

		numRead++

		if numRead > VarIntMaxByteSize {
			return ErrVarIntTooLarge
		}
	}

	*vi = VarInt(result)

	return nil
}
