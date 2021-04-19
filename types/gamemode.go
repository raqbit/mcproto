package types

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

//go:generate stringer -type=Gamemode -output gamemode_string.go -linecomment

// Gamemode is a Minecraft gamemode
type Gamemode enc.UnsignedByte

const (
	Survival  Gamemode = 0 // Survival
	Creative  Gamemode = 1 // Creative
	Adventure Gamemode = 2 // Adventure
	Spectator Gamemode = 3 // Spectator
)

// Is check if Gamemode is equal to other Gamemode, ignoring the hardcore bit
func (m Gamemode) Is(other Gamemode) bool {
	return m&other == other
}

// IsHardcore checks if the Gamemode is hardcore
func (m Gamemode) IsHardcore() bool {
	return m&0x8 == 0x8
}

func (m *Gamemode) Write(w io.Writer) error {
	num := enc.UnsignedByte(*m)
	return num.Write(w)
}

func (m *Gamemode) Read(r io.Reader) error {
	var num enc.UnsignedByte

	if err := num.Read(r); err != nil {
		return err
	}

	*m = Gamemode(num)
	return nil
}
