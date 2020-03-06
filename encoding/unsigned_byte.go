package encoding

import "io"

// Minecraft Protocol UnsignedByte type
type UnsignedByte uint8

func (ub *UnsignedByte) Decode(r io.Reader) error {
	bt, err := ReadUnsignedByte(r)

	if err != nil {
		return err
	}

	*ub = bt
	return nil
}

func (ub *UnsignedByte) Encode(w io.Writer) error {
	_, err := w.Write(WriteUnsignedByte(*ub))
	return err
}

func ReadUnsignedByte(r io.Reader) (UnsignedByte, error) {
	var bytes [1]byte
	_, err := r.Read(bytes[:1])
	return UnsignedByte(bytes[0]), err
}

func WriteUnsignedByte(value UnsignedByte) []byte {
	return []byte{byte(value)}
}
