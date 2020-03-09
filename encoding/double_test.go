package encoding

import (
	"bytes"
	"math"
	"testing"
)

func TestWriteDouble(t *testing.T) {
	tests := []struct {
		Value    float64
		Expected []byte
	}{
		{Value: 0.0000000002, Expected: []byte{0x3D, 0xEB, 0x7C, 0xDF, 0xD9, 0xD7, 0xBD, 0xBB}},
		{Value: 0, Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Value: math.MaxFloat64, Expected: []byte{0x7F, 0xEF, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, test := range tests {
		actual := WriteDouble(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %f: %v != %v", test.Value, actual, test.Expected)
		}
	}
}

func TestReadDouble(t *testing.T) {
	tests := []struct {
		Expected float64
		Value    []byte
	}{
		{Expected: 0.0000000002, Value: []byte{0x3D, 0xEB, 0x7C, 0xDF, 0xD9, 0xD7, 0xBD, 0xBB}},
		{Expected: 0, Value: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Expected: math.MaxFloat64, Value: []byte{0x7F, 0xEF, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	var buff bytes.Buffer

	for _, test := range tests {

		buff.Write(test.Value)

		actual, err := ReadDouble(&buff)

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
