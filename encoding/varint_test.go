package encoding

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestWriteVarInt(t *testing.T) {
	tests := []struct {
		Value    int32
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00}},
		{Value: 1, Expected: []byte{0x01}},
		{Value: 2, Expected: []byte{0x02}},
		{Value: 127, Expected: []byte{0x7f}},
		{Value: 128, Expected: []byte{0x80, 0x01}},
		{Value: 255, Expected: []byte{0xff, 0x01}},
		{Value: 2097151, Expected: []byte{0xff, 0xff, 0x7f}},
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
		{Expected: 1, Value: []byte{0x01}},
		{Expected: 2, Value: []byte{0x02}},
		{Expected: 127, Value: []byte{0x7f}},
		{Expected: 128, Value: []byte{0x80, 0x01}},
		{Expected: 255, Value: []byte{0xff, 0x01}},
		{Expected: 2097151, Value: []byte{0xff, 0xff, 0x7f}},
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

func TestReadVarIntReadError(t *testing.T) {
	var buff bytes.Buffer

	_, err := ReadVarInt(&buff)

	if !errors.Is(err, io.EOF) {
		t.Errorf("should get read error")
	}
}

func TestReadVarIntTooLarge(t *testing.T) {
	var buff bytes.Buffer

	buff.Write([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x01})

	_, err := ReadVarInt(&buff)

	if !errors.Is(err, ErrVarIntTooLarge) {
		t.Errorf("should get error for too large VarInt")
	}
}
