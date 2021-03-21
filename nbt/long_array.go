package nbt

import "io"

func WriteTagLongArray(value []int64) []byte {
	b := WriteTagInt(int32(len(value)))
	for _, i := range value {
		b = append(b, WriteTagLong(i)...)
	}
	return b
}

func ReadTagLongArray(r io.Reader) ([]int64, error) {
	length, err := ReadTagInt(r)

	if err != nil {
		return nil, err
	}

	buf := make([]int64, int(length))

	for i := 0; i < int(length); i++ {
		if buf[i], err = ReadTagLong(r); err != nil {
			return nil, err
		}
	}

	return buf, err
}
