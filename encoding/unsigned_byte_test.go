package encoding

import (
	"bytes"
	"math"
	"testing"
)

func TestWriteUnsignedByte(t *testing.T) {
	tests := []struct {
		Value    UnsignedByte
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00}},
		{Value: 5, Expected: []byte{0x05}},
		{Value: math.MaxUint8 / 2, Expected: []byte{0x7f}},
		{Value: math.MaxUint8, Expected: []byte{0xff}},
	}

	for _, test := range tests {
		actual := WriteUnsignedByte(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, actual, test.Expected)
		}
	}
}

func TestReadUnsignedByte(t *testing.T) {
	tests := []struct {
		Expected UnsignedByte
		Value    []byte
	}{
		{Expected: 0, Value: []byte{0x00}},
		{Expected: 5, Value: []byte{0x05}},
		{Expected: math.MaxUint8 / 2, Value: []byte{0x7f}},
		{Expected: math.MaxUint8, Value: []byte{0xff}},
	}

	var buff bytes.Buffer

	for _, test := range tests {

		buff.Write(test.Value)

		actual, err := ReadUnsignedByte(&buff)

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
