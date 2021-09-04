package mcproto

import (
	"context"
	"net"
	"strconv"
	"strings"
)

const (
	defaultMinecraftPort = "25565"
	minecraftSRVService  = "minecraft"
	minecraftSRVProtocol = "tcp"
)

func ResolveServerAddress(ctx context.Context, address string) string {
	var resolver net.Resolver

	// Split host & port, port being optional
	host, port := splitHostPort(address)

	// If no port is given, use the default Minecraft port.
	if port == "" {
		port = defaultMinecraftPort
	}

	// If no port is given or the given port is the default,
	// do a DNS SRV record lookup.
	if port == defaultMinecraftPort {
		// Do DNS SRV record lookup on given hostname
		_, srvRecords, err := resolver.LookupSRV(ctx, minecraftSRVService, minecraftSRVProtocol, host)

		if err == nil && len(srvRecords) > 0 {
			// Override host & port with details from the first SRV record returned
			record := srvRecords[0]
			host = record.Target
			port = strconv.Itoa(int(record.Port))
		}
	}

	// Join host & port for connecting to the server but also for returning
	// the resolved server address.
	//
	// Note: If the host was resolved via an SRV record, it will have a
	// trailing period. This is kept so the returned address can be used for
	// a handshake packet, which the vanilla client also sends with a trailing period.
	return net.JoinHostPort(host, port)
}

// Like net.SplitHostPort but with the port being optional
// Handles IPv6, IPv4 and hostnames
func splitHostPort(address string) (host string, port string) {
	// Ipv6 with port
	if address[0] == '[' {
		endIdx := strings.LastIndexByte(address, ']')
		host = address[1:endIdx]
		port = address[endIdx+2:]
		return
	}

	lastColon := -1

	i := len(address)
	for i--; i >= 0; i-- {
		if address[i] == ':' {
			if lastColon != -1 {
				// We found multiple colons, so ipv6 (without port)
				host = address
				return
			}

			lastColon = i
		}

		if address[i] == '.' {
			// We found a dot, so we've found an ipv4 address or hostname
			if lastColon != -1 {
				// We found a port
				host = address[0:lastColon]
				port = address[lastColon+1:]
			} else {
				// We did not find a port
				host = address
			}
			return
		}
	}

	// Ipv4 or hostname without port
	if lastColon == -1 {
		host = address
		return
	}

	// ipv4 or hostname with port
	host = address[0:lastColon]
	port = address[lastColon+1:]

	return
}
