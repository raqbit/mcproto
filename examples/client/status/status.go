package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
	"github.com/Raqbit/mcproto/packet"
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

	conn, addr, err := mcproto.DialContext(context.Background(), os.Args[1])

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

	err = conn.WritePacket(&packet.HandshakePacket{
		ProtoVer:   ProtocolVersion,
		ServerAddr: enc.String(host),
		ServerPort: enc.UnsignedShort(port),
		NextState:  game.StatusState,
	})

	// Actually update our connection as well
	conn.SetState(game.StatusState)

	if err != nil {
		log.Fatalf("could not write handshake packet: %s", err)
	}

	err = conn.WritePacket(&packet.ServerQueryPacket{})

	if err != nil {
		log.Fatalf("could not write request packet: %s", err)
	}

	resp, err := conn.ReadPacket()

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
