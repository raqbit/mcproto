package encoding

import (
	"bytes"
	"github.com/google/uuid"
	"testing"
)

func TestWriteUUID(t *testing.T) {
	tests := []struct {
		Value    UUID
		Expected []byte
	}{
		{
			Value:    UUID(uuid.Nil),
			Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			Value:    UUID(uuid.MustParse("70d651be-e8c1-4214-951e-fe92dea07cb6")),
			Expected: []byte{0x70, 0xd6, 0x51, 0xbe, 0xe8, 0xc1, 0x42, 0x14, 0x95, 0x1e, 0xfe, 0x92, 0xde, 0xa0, 0x7c, 0xb6},
		},
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

func TestReadUUID(t *testing.T) {
	tests := []struct {
		Expected UUID
		Value    []byte
	}{
		{
			Expected: UUID(uuid.Nil),
			Value:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			Expected: UUID(uuid.MustParse("70d651be-e8c1-4214-951e-fe92dea07cb6")),
			Value:    []byte{0x70, 0xd6, 0x51, 0xbe, 0xe8, 0xc1, 0x42, 0x14, 0x95, 0x1e, 0xfe, 0x92, 0xde, 0xa0, 0x7c, 0xb6},
		},
	}
	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		var actual UUID
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
