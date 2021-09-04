package game

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

//go:generate stringer -type=Mode -output mode_string.go -linecomment

// Mode is a Minecraft gamemode
type Mode enc.UnsignedByte

const (
	Survival  Mode = 0 // Survival
	Creative  Mode = 1 // Creative
	Adventure Mode = 2 // Adventure
	Spectator Mode = 3 // Spectator
)

// Is check if Mode is equal to other Mode, ignoring the hardcore bit
func (m Mode) Is(other Mode) bool {
	return m&other == other
}

// IsHardcore checks if the Mode is hardcore
func (m Mode) IsHardcore() bool {
	return m&0x8 == 0x8
}

func (m *Mode) Write(w io.Writer) error {
	num := enc.UnsignedByte(*m)
	return num.Write(w)
}

func (m *Mode) Read(r io.Reader) error {
	var num enc.UnsignedByte

	if err := num.Read(r); err != nil {
		return err
	}

	*m = Mode(num)
	return nil
}
