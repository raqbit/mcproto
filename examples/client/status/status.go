package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/packet"
	"github.com/Raqbit/mcproto/types"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	ProtocolVersion = 578
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify the address to get the status of")
		os.Exit(2)
	}

	ctx := context.Background()

	conn, addr, err := mcproto.DialContext(ctx, os.Args[1])

	if err != nil {
		log.Fatalf("mcproto dial error: %s", err)
	}

	host, portStr, err := net.SplitHostPort(addr)

	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.WritePacket(ctx,
		&packet.HandshakePacket{
			ProtoVer:   ProtocolVersion,
			ServerAddr: enc.String(host),
			ServerPort: enc.UnsignedShort(port),
			NextState:  types.ConnectionStateStatus,
		})

	// Actually update our connection as well
	conn.SetState(types.ConnectionStateStatus)

	if err != nil {
		log.Fatalf("could not write handshake packet: %s", err)
	}

	err = conn.WritePacket(ctx, &packet.ServerQueryPacket{})

	if err != nil {
		log.Fatalf("could not write request packet: %s", err)
	}

	resp, err := conn.ReadPacket(ctx)

	if err != nil {
		log.Fatalf("could not read response packet: %s", err)
	}

	response, ok := resp.(*packet.ServerInfoPacket)

	if !ok {
		log.Fatalf("Server sent unexpected packet: %s", resp.String())
	}

	jsonRes, _ := json.Marshal(response.Response)

	fmt.Println(string(jsonRes))
}
