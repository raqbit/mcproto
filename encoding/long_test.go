package encoding

import (
	"bytes"
	"math"
	"testing"
)

func TestWriteLong(t *testing.T) {
	tests := []struct {
		Value    int64
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Value: 5, Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05}},
		{Value: -1, Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{Value: math.MinInt64, Expected: []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Value: math.MaxInt64, Expected: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, test := range tests {
		actual := WriteLong(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, actual, test.Expected)
		}

	}
}

func TestReadLong(t *testing.T) {
	tests := []struct {
		Expected int64
		Value    []byte
	}{
		{Expected: 0, Value: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Expected: 5, Value: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05}},
		{Expected: -1, Value: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{Expected: math.MinInt64, Value: []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{Expected: math.MaxInt64, Value: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		actual, err := ReadLong(&buff)

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
