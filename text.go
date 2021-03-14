package mcproto

import (
	"encoding/json"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
)

// Minecraft text component
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

func (t *TextComponent) Write(w io.Writer) error {
	bytes, err := json.Marshal(t)

	if err != nil {
		return err
	}

	str := enc.String(bytes)
	return str.Write(w)
}

func (t *TextComponent) Read(r io.Reader) error {
	var str enc.String

	if err := str.Read(r); err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), t)
}
