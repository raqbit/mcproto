package mcproto

import "encoding/json"

const ServerInfoPacketID int32 = 0x00

// https://wiki.vg/Protocol#Response
type ServerInfoPacket struct {
	Response ServerInfo
}

func (*ServerInfoPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ServerInfoPacketID,
		Direction:       ClientBound,
		ConnectionState: ConnectionStateStatus,
	}
}

func (*ServerInfoPacket) String() string {
	return "ServerInfo"
}

func (si *ServerInfoPacket) Marshal(w PacketWriter) error {
	var err error
	var response []byte

	if response, err = json.Marshal(si.Response); err != nil {
		return err
	}

	if err = w.WriteString(string(response)); err != nil {
		return err
	}

	return nil
}

func (si *ServerInfoPacket) Unmarshal(r PacketReader) error {
	var err error
	var response string

	if response, err = r.ReadMaxString(); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(response), &si.Response); err != nil {
		return err
	}

	return nil
}

type Version struct {
	Name     string `json:"name"`     // Version name
	Protocol int32  `json:"protocol"` // Version protocol number
}

// Server info player
type Player struct {
	Name string `json:"name"` // Player name
	ID   string `json:"id"`   // Player UUID
}

// Server info players
type Players struct {
	Max    int32    `json:"max"`    // Max amount of players allowed
	Online int32    `json:"online"` // Amount of players online
	Sample []Player `json:"sample"` // Sample of online players
}

// Server ping response
// https://wiki.vg/Server_List_Ping#Response
type ServerInfo struct {
	Version     Version           `json:"version"`           // Server version info
	Players     Players           `json:"players"`           // Server player info
	Description ServerDescription `json:"description"`       // Server description
	Favicon     string            `json:"favicon,omitempty"` // Server favicon
}

// ServerDescription can be both a string (legacy) or TextComponent JSON structure
type ServerDescription TextComponent

// Lenient server description parser,
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
