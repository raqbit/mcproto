package nbt

import "io"

func WriteTagIntArray(value []int32) []byte {
	b := WriteTagInt(int32(len(value)))
	for _, i := range value {
		b = append(b, WriteTagInt(i)...)
	}
	return b
}

func ReadTagIntArray(r io.Reader) ([]int32, error) {
	length, err := ReadTagInt(r)

	if err != nil {
		return nil, err
	}

	buf := make([]int32, int(length))

	for i := 0; i < int(length); i++ {
		if buf[i], err = ReadTagInt(r); err != nil {
			return nil, err
		}
	}

	return buf, err
}
