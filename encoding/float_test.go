package encoding

import (
	"bytes"
	"math"
	"testing"
)

func TestWriteFloat(t *testing.T) {
	tests := []struct {
		Value    float32
		Expected []byte
	}{
		{Value: 0.0000000002, Expected: []byte{0x2F, 0x5B, 0xE6, 0xFF}},
		{Value: 0, Expected: []byte{0x00, 0x00, 0x00, 0x00}},
		{Value: math.MaxFloat32, Expected: []byte{0x7f, 0x7f, 0xff, 0xff}},
	}

	for _, test := range tests {
		actual := WriteFloat(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %f: %v != %v", test.Value, actual, test.Expected)
		}
	}
}

func TestReadFloat(t *testing.T) {
	tests := []struct {
		Expected float32
		Value    []byte
	}{
		{Expected: 0.0000000002, Value: []byte{0x2F, 0x5B, 0xE6, 0xFF}},
		{Expected: 0, Value: []byte{0x00, 0x00, 0x00, 0x00}},
		{Expected: math.MaxFloat32, Value: []byte{0x7f, 0x7f, 0xff, 0xff}},
	}

	var buff bytes.Buffer

	for _, test := range tests {

		buff.Write(test.Value)

		actual, err := ReadFloat(&buff)

		if err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf("Unable to convert %v: %f != %f", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
