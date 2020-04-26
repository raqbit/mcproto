package encoding

import (
	"bytes"
	"testing"
)

func TestWriteVarInt(t *testing.T) {
	tests := []struct {
		Value    int32
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00}},
		{Value: 5, Expected: []byte{0x05}},
		{Value: 2147483647, Expected: []byte{0xff, 0xff, 0xff, 0xff, 0x07}},
		{Value: -1, Expected: []byte{0xff, 0xff, 0xff, 0xff, 0x0f}},
		{Value: -2147483648, Expected: []byte{0x80, 0x80, 0x80, 0x80, 0x08}},
	}

	for _, test := range tests {
		actual := WriteVarInt(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, actual, test.Expected)
		}

	}
}

func TestReadVarInt(t *testing.T) {
	tests := []struct {
		Expected int32
		Value    []byte
	}{
		{Expected: 0, Value: []byte{0x00}},
		{Expected: 2147483647, Value: []byte{0xff, 0xff, 0xff, 0xff, 0x07}},
		{Expected: -1, Value: []byte{0xff, 0xff, 0xff, 0xff, 0x0f}},
		{Expected: -2147483648, Value: []byte{0x80, 0x80, 0x80, 0x80, 0x08}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		actual, err := ReadVarInt(&buff)

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
