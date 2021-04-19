package types

import (
	"encoding/json"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

type ServerInfoVersion struct {
	Name     string `json:"name"`     // Version name
	Protocol int32  `json:"protocol"` // Version protocol number
}

type ServerInfoPlayer struct {
	Name string `json:"name"` // Player name
	ID   string `json:"id"`   // Player UUID
}

type ServerInfoPlayers struct {
	Max    int32              `json:"max"`    // Max amount of players allowed
	Online int32              `json:"online"` // Amount of players online
	Sample []ServerInfoPlayer `json:"sample"` // Sample of online players
}

// ServerInfo is the JSON datastructure returned by the Server List Ping response
// https://wiki.vg/Server_List_Ping#Response
type ServerInfo struct {
	Version     ServerInfoVersion `json:"version"`           // Server version info
	Players     ServerInfoPlayers `json:"players"`           // Server player info
	Description ServerDescription `json:"description"`       // Server description
	Favicon     string            `json:"favicon,omitempty"` // Server favicon
}

func (s *ServerInfo) Write(w io.Writer) error {
	bytes, err := json.Marshal(s)

	if err != nil {
		return err
	}

	str := enc.String(bytes)
	return str.Write(w)
}

func (s *ServerInfo) Read(r io.Reader) error {
	var str enc.String

	if err := str.Read(r); err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), s)
}

// ServerDescription can be both a string (legacy) or TextComponent JSON structure
type ServerDescription TextComponent

// UnmarshalJSON implementation for ServerDescription which accepts both a TextComponent or a string (legacy format)
func (c *ServerDescription) UnmarshalJSON(data []byte) error {
	// The data starts with quotes which means it's a string, not an object
	if data[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}

		c.Text = text
	} else {
		// Data to unmarshal is not a string, we can parse it as a regular text component
		var out TextComponent
		if err := json.Unmarshal(data, &out); err != nil {
			return err
		}
		*c = ServerDescription(out)
	}

	return nil
}
