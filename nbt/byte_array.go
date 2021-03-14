package nbt

import "io"

func WriteTagByteArray(value []byte) []byte {
	b := WriteTagInt(int32(len(value)))
	b = append(b, value...)
	return b
}

func ReadTagByteArray(r io.Reader) ([]byte, error) {
	l, err := ReadTagInt(r)

	if err != nil {
		return nil, err
	}

	buf := make([]byte, int(l))
	_, err = io.ReadFull(r, buf)
	return buf, err
}
