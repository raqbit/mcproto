package encoding

import "io"

// Minecraft Protocol UnsignedByte type
type UnsignedByte byte

func (ub *UnsignedByte) Decode(r io.Reader) error {
	bt, err := ReadUnsignedByte(r)

	if err != nil {
		return err
	}

	*ub = bt
	return nil
}

func (ub *UnsignedByte) Encode(w io.Writer) error {
	return WriteUnsignedByte(w, *ub)
}

func ReadUnsignedByte(r io.Reader) (UnsignedByte, error) {
	var bytes [1]byte
	_, err := r.Read(bytes[:1])
	return UnsignedByte(bytes[0]), err
}

func WriteUnsignedByte(w io.Writer, value UnsignedByte) error {
	var bytes [1]byte
	bytes[0] = byte(value)
	_, err := w.Write(bytes[:1])
	return err
}

