package mcproto

// Minecraft chat component
// See: https://wiki.vg/Chat#Current_system_.28JSON_Chat.29
type ChatComponent struct {
	Text          string          `json:"text"`                    // Text content
	Bold          *bool           `json:"bold,omitempty"`          // Component is emboldened
	Italic        *bool           `json:"italic,omitempty"`        // Component is italicized
	Underlined    *bool           `json:"underlined,omitempty"`    // Component is underlined
	Strikethrough *bool           `json:"strikethrough,omitempty"` // Component is struck out
	Obfuscated    *bool           `json:"obfuscated,omitempty"`    // Component randomly switches between characters of the same width
	Color         string          `json:"color,omitempty"`         // Contains the color for the component
	Extra         []ChatComponent `json:"extra,omitempty"`         // ChatComponent siblings
}
