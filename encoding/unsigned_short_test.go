package encoding

import (
	"bytes"
	"io"
	"math"
	"testing"
)

func TestWriteUnsignedShort(t *testing.T) {
	tests := []struct {
		Value    UnsignedShort
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00, 0x00}},
		{Value: 5, Expected: []byte{0x00, 0x05}},
		{Value: math.MaxUint8 + 1, Expected: []byte{0x01, 0x00}},
		{Value: math.MaxUint16 / 2, Expected: []byte{0x7f, 0xff}},
		{Value: math.MaxUint16, Expected: []byte{0xff, 0xff}},
	}

	var buff bytes.Buffer
	_ = io.Writer(&buff)

	for _, test := range tests {
		err := WriteUnsignedShort(&buff, test.Value)

		if err != nil {
			t.Error(err)
		}

		if bytes.Compare(test.Expected, buff.Bytes()) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, buff.Bytes(), test.Expected)
		}

		buff.Reset()
	}
}

func TestReadUnsignedShort(t *testing.T) {
	tests := []struct {
		Expected UnsignedShort
		Value    []byte
	}{
		{Expected: 0, Value: []byte{0x00, 0x00}},
		{Expected: 5, Value: []byte{0x00, 0x05}},
		{Expected: math.MaxUint8 + 1, Value: []byte{0x01, 0x00}},
		{Expected: math.MaxUint16 / 2, Value: []byte{0x7f, 0xff}},
		{Expected: math.MaxUint16, Value: []byte{0xff, 0xff}},
	}

	var buff bytes.Buffer
	_ = io.Writer(&buff)

	for _, test := range tests {

		buff.Write(test.Value)

		actual, err := ReadUnsignedShort(&buff)

		if err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf("Unable to convert %v: %d != %d", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
