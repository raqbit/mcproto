package mcproto

import "github.com/google/uuid"

type GameProfile struct {
	UUID uuid.UUID
	Name string
}

func NewGameProfile(uuid uuid.UUID, name string) *GameProfile {
	return &GameProfile{UUID: uuid, Name: name}
}
