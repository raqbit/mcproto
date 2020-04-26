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

func WriteVarInt(value int32) []byte {
	buf := make([]byte, 0)
	for cont := true; cont; cont = value != 0 {
		temp := byte(value & 0x7F)

		// Casting value to a uint to get a logical shift
		value = int32(uint32(value) >> 7)

		if value != 0 {
			temp |= 0x80
		}

		buf = append(buf, WriteUnsignedByte(temp)[0])
	}

	return buf
}

func ReadVarInt(r io.Reader) (int32, error) {
	var numRead uint
	var result int32
	var read uint8

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

	return result, nil
}
