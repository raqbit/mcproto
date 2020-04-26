package mcproto

import "encoding/json"

// Minecraft chat component
// See: https://wiki.vg/Chat#Current_system_.28JSON_Chat.29
type TextComponent struct {
	Text          string          `json:"text"`                    // Text content
	Bold          *bool           `json:"bold,omitempty"`          // Component is emboldened
	Italic        *bool           `json:"italic,omitempty"`        // Component is italicized
	Underlined    *bool           `json:"underlined,omitempty"`    // Component is underlined
	Strikethrough *bool           `json:"strikethrough,omitempty"` // Component is struck out
	Obfuscated    *bool           `json:"obfuscated,omitempty"`    // Component randomly switches between characters of the same width
	Color         string          `json:"color,omitempty"`         // Contains the color for the component
	Extra         []TextComponent `json:"extra,omitempty"`         // TextComponent siblings
}

// Lenient text component parser,
func (c *TextComponent) UnmarshalJSON(data []byte) error {
	// The data starts with quotes which means it's a string, not an object
	if data[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}

		c.Text = text
	} else {
		if err := json.Unmarshal(data, &c); err != nil {
			return err
		}
	}

	return nil
}
