package encoding

import (
	"bytes"
	"io"
	"math"
	"testing"
)

func TestWriteDouble(t *testing.T) {
	tests := []struct {
		Value    Double
		Expected []byte
	}{
		{Value: 0.0000000002, Expected: []byte{0x3D, 0xEB, 0x7C, 0xDF, 0xD9, 0xD7, 0xBD, 0xBB}},
		{Value: 0, Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Value: math.MaxFloat64, Expected: []byte{0x7F, 0xEF, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	var buff bytes.Buffer
	_ = io.Writer(&buff)

	for _, test := range tests {
		err := WriteDouble(&buff, test.Value)

		if err != nil {
			t.Error(err)
		}

		if bytes.Compare(test.Expected, buff.Bytes()) != 0 {
			// Not equal
			t.Errorf("Unable to convert %f: %v != %v", test.Value, buff.Bytes(), test.Expected)
		}

		buff.Reset()
	}
}

func TestReadDouble(t *testing.T) {
	tests := []struct {
		Expected Double
		Value    []byte
	}{
		{Expected: 0.0000000002, Value: []byte{0x3D, 0xEB, 0x7C, 0xDF, 0xD9, 0xD7, 0xBD, 0xBB}},
		{Expected: 0, Value: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Expected: math.MaxFloat64, Value: []byte{0x7F, 0xEF, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	var buff bytes.Buffer
	_ = io.Writer(&buff)

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
