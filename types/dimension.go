package types

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

//go:generate stringer -type=Dimension -output dimension_string.go -linecomment

// Dimension is a Minecraft dimension
type Dimension int32

const (
	DimensionNether    Dimension = -1 // Nether
	DimensionOverworld Dimension = 0  // Overworld
	DimensionEnd       Dimension = 1  // End
)

func (d *Dimension) Write(w io.Writer) error {
	num := enc.Int(*d)
	return num.Write(w)
}

func (d *Dimension) Read(r io.Reader) error {
	var num enc.Int

	if err := num.Read(r); err != nil {
		return err
	}

	*d = Dimension(num)
	return nil
}
