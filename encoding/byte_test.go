package encoding

import (
	"bytes"
	"math"
	"testing"
)

func TestWriteByte(t *testing.T) {
	tests := []struct {
		Value    Byte
		Expected []byte
	}{
		{Value: math.MinInt8, Expected: []byte{0x80}},
		{Value: 5, Expected: []byte{0x05}},
		{Value: math.MaxInt8 / 2, Expected: []byte{0x3f}},
		{Value: math.MaxInt8, Expected: []byte{0x7f}},
	}

	for _, test := range tests {
		actual := WriteByte(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, actual, test.Expected)
		}
	}
}

func TestReadByte(t *testing.T) {
	tests := []struct {
		Expected Byte
		Value    []byte
	}{
		{Expected: math.MinInt8, Value: []byte{0x80}},
		{Expected: 5, Value: []byte{0x05}},
		{Expected: math.MaxInt8 / 2, Value: []byte{0x3f}},
		{Expected: math.MaxInt8, Value: []byte{0x7f}},
	}

	var buff bytes.Buffer

	for _, test := range tests {

		buff.Write(test.Value)

		actual, err := ReadByte(&buff)

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
