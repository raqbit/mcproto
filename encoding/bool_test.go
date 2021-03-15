package encoding

import (
	"bytes"
	"testing"
)

func TestWriteBool(t *testing.T) {
	tests := []struct {
		Value    Bool
		Expected []byte
	}{
		{Value: true, Expected: []byte{0x01}},
		{Value: false, Expected: []byte{0x00}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		if err := test.Value.Write(&buff); err != nil {
			t.Error(err)
		}

		actual := buff.Bytes()

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %v: %v != %v", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}

func TestReadBool(t *testing.T) {
	tests := []struct {
		Expected Bool
		Value    []byte
	}{
		{Expected: true, Value: []byte{0x01}},
		{Expected: false, Value: []byte{0x00}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		var actual Bool
		if err := actual.Read(&buff); err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf("Unable to convert %v: %v != %v", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
