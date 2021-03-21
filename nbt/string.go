package nbt

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

func WriteTagString(str string) []byte {
	b := enc.WriteUnsignedShort(uint16(len(str)))
	b = append(b, []byte(str)...)
	return b
}

func ReadTagString(r io.Reader) (string, error) {
	s, err := enc.ReadUnsignedByte(r)
	if err != nil {
		return "", nil
	}

	stringBuff := make([]byte, int(s))
	_, err = io.ReadFull(r, stringBuff)
	return string(stringBuff), err
}
