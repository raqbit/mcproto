package encoding

import (
	"bytes"
	"testing"
)

func TestWriteBool(t *testing.T) {
	tests := []struct {
		Value    bool
		Expected []byte
	}{
		{Value: true, Expected: []byte{0x01}},
		{Value: false, Expected: []byte{0x00}},
	}

	for _, test := range tests {
		actual := WriteBool(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %v: %v != %v", test.Value, actual, test.Expected)
		}
	}
}

func TestReadBool(t *testing.T) {
	tests := []struct {
		Expected bool
		Value    []byte
	}{
		{Expected: true, Value: []byte{0x01}},
		{Expected: false, Value: []byte{0x00}},
	}

	var buff bytes.Buffer

	for _, test := range tests {

		buff.Write(test.Value)

		actual, err := ReadBool(&buff)

		if err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf("Unable to convert %v: %v != %v", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
