package game

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const (
	ServerInfoTestDataDir = "testdata/serverinfo/"
)

var normalServerInfo = ServerInfo{
	Version: ServerInfoVersion{
		Name:     "1.13.2",
		Protocol: 404,
	},
	Players: ServerInfoPlayers{
		Max:    100,
		Online: 5,
		Sample: []ServerInfoPlayer{
			{
				Name: "Raqbit",
				ID:   "09bc745b-3679-4152-b96b-3f9c59c42059",
			},
		},
	},
	Description: ServerDescription{
		Text: "Hello world",
	},
	Favicon: "data:image/png;base64,<data>",
}

func TestParseServerInfo(t *testing.T) {
	tests := []struct {
		name     string
		expected ServerInfo
	}{
		{
			name:     "normal",
			expected: normalServerInfo,
		},
		{
			name:     "description_string",
			expected: normalServerInfo,
		},
		{
			name: "negative_protocol_version",
			expected: ServerInfo{
				Version: ServerInfoVersion{
					Name:     "Paper 1.16.5",
					Protocol: -1,
				},
				Players: ServerInfoPlayers{
					Max:    8,
					Online: 0,
				},
				Description: ServerDescription{
					Text: "A Vanilla Minecraft Server powered by Docker",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			infoJson, err := getServerInfoData(test.name)

			if err != nil {
				t.Fatalf("could not load test data: %s", err)
			}

			var parsedInfo ServerInfo

			if err = json.Unmarshal(infoJson, &parsedInfo); err != nil {
				t.Fatalf("could not parse server info: %s", err)
			}

			assert.Equal(t, test.expected, parsedInfo, "%s: did not parse correctly", test.name)
		})
	}
}

func getServerInfoData(name string) ([]byte, error) {
	path := filepath.Join(ServerInfoTestDataDir, fmt.Sprintf("%s.json", name))

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return data, nil
}
