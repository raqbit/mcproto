package encoding

import (
	"bytes"
	"math"
	"testing"
)

func TestWriteShort(t *testing.T) {
	tests := []struct {
		Value    Short
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00, 0x00}},
		{Value: 5, Expected: []byte{0x00, 0x05}},
		{Value: math.MaxInt8 + 1, Expected: []byte{0x00, 0x80}},
		{Value: math.MaxInt16 / 2, Expected: []byte{0x3f, 0xff}},
		{Value: math.MaxInt16, Expected: []byte{0x7f, 0xff}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		if err := test.Value.Write(&buff); err != nil {
			t.Error(err)
		}

		actual := buff.Bytes()

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, actual, test.Expected)
		}
	}
}

func TestReadShort(t *testing.T) {
	tests := []struct {
		Expected Short
		Value    []byte
	}{
		{Expected: 0, Value: []byte{0x00, 0x00}},
		{Expected: 5, Value: []byte{0x00, 0x05}},
		{Expected: math.MaxInt8 + 1, Value: []byte{0x00, 0x80}},
		{Expected: math.MaxInt16 / 2, Value: []byte{0x3f, 0xff}},
		{Expected: math.MaxInt16, Value: []byte{0x7f, 0xff}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		var actual Short
		if err := actual.Read(&buff); err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf("Unable to convert %v: %d != %d", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
