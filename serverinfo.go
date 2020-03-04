package mcproto

type Version struct {
	Name     string `json:"name"`     // Version name
	Protocol uint   `json:"protocol"` // Version protocol number
}

// Server info player
type Player struct {
	Name string `json:"name"` // Player name
	ID   string `json:"id"`   // Player UUID
}

// Server info players
type Players struct {
	Max    uint     `json:"max"`    // Max amount of players allowed
	Online uint     `json:"online"` // Amount of players online
	Sample []Player `json:"sample"` // Sample of online players
}

// Server ping response
// https://wiki.vg/Server_List_Ping#Response
type ServerInfo struct {
	Version     Version       `json:"version"`           // Server version info
	Players     Players       `json:"players"`           // Server player info
	Description ChatComponent `json:"description"`       // Server description
	Favicon     string        `json:"favicon,omitempty"` // Server favicon
}
