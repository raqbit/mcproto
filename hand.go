package mcproto

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

//go:generate stringer -type=Hand -output hand_string.go -linecomment
type Hand int32

const (
	HandLeft  Hand = 0 // Left
	HandRight Hand = 1 // Right
)

func (h *Hand) Write(w io.Writer) error {
	num := enc.VarInt(*h)
	return num.Write(w)
}

func (h *Hand) Read(r io.Reader) error {
	num := enc.VarInt(*h)
	return num.Read(r)
}
