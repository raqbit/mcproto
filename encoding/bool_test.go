package encoding

import (
	"bytes"
	"io"
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
	_ = io.Writer(&buff)

	for _, test := range tests {
		err := WriteBool(&buff, test.Value)

		if err != nil {
			t.Error(err)
		}

		if bytes.Compare(test.Expected, buff.Bytes()) != 0 {
			// Not equal
			t.Errorf("Unable to convert %v: %v != %v", test.Value, buff.Bytes(), test.Expected)
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
	_ = io.Writer(&buff)

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
