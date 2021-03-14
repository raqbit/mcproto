package types

//go:generate stringer -type=Gamemode -output gamemode_string.go -linecomment
type Gamemode uint8

const (
	Survival  Gamemode = 0 // Survival
	Creative  Gamemode = 1 // Creative
	Adventure Gamemode = 2 // Adventure
	Spectator Gamemode = 3 // Spectator
)

// Check if Gamemode is equal to other Gamemode, ignoring
// hardcore bit
func (m Gamemode) Is(other Gamemode) bool {
	return m&other == other
}

// Check if Gamemode is also hardcore
func (m Gamemode) IsHardcore() bool {
	return m&0x8 == 0x8
}
