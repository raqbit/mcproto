package mcproto

import "testing"

func Test_splitHostPort(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		wantHost string
		wantPort string
	}{
		{
			name:     "ipv4",
			address:  "127.0.0.1",
			wantHost: "127.0.0.1",
			wantPort: "",
		},
		{
			name:     "ipv4_port",
			address:  "127.0.0.1:25565",
			wantHost: "127.0.0.1",
			wantPort: "25565",
		},
		{
			name:     "host_short",
			address:  "localhost",
			wantHost: "localhost",
			wantPort: "",
		},
		{
			name:     "host_short_port",
			address:  "localhost:25565",
			wantHost: "localhost",
			wantPort: "25565",
		},
		{
			name:     "host",
			address:  "mc.raqb.it",
			wantHost: "mc.raqb.it",
			wantPort: "",
		},
		{
			name:     "host_port",
			address:  "mc.raqb.it:25565",
			wantHost: "mc.raqb.it",
			wantPort: "25565",
		},
		{
			name:     "ipv6_short",
			address:  "::1",
			wantHost: "::1",
			wantPort: "",
		},
		{
			name:     "ipv6_short_port",
			address:  "[::1]:25565",
			wantHost: "::1",
			wantPort: "25565",
		},
		{
			name:     "ipv6",
			address:  "c0ec:5d71:7830:511d:ae76:dbf0:d659:9754",
			wantHost: "c0ec:5d71:7830:511d:ae76:dbf0:d659:9754",
			wantPort: "",
		},
		{
			name:     "ipv6_port",
			address:  "[c0ec:5d71:7830:511d:ae76:dbf0:d659:9754]:25565",
			wantHost: "c0ec:5d71:7830:511d:ae76:dbf0:d659:9754",
			wantPort: "25565",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHost, gotPort := splitHostPort(tt.address)
			if gotHost != tt.wantHost {
				t.Errorf("splitHostPort() gotHost = %v, want %v", gotHost, tt.wantHost)
			}
			if gotPort != tt.wantPort {
				t.Errorf("splitHostPort() gotPort = %v, want %v", gotPort, tt.wantPort)
			}
		})
	}
}
