package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Raqbit/mcproto"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	ProtocolVersion = 578
)

var (
	ServerHost = flag.String("host", "", "server host")
	ServerPort = flag.Uint("port", 25565, "server port")
)

func main() {
	flag.Parse()

	if *ServerHost == "" {
		flag.Usage()
		os.Exit(2)
	}

	ctx := context.Background()

	address := net.JoinHostPort(*ServerHost, strconv.Itoa(int(*ServerPort)))

	conn, err := mcproto.Dial(address, mcproto.ClientSide)

	if err != nil {
		log.Fatalf("tcp dial error: %s", err)
	}

	err = conn.WritePacket(ctx,
		&mcproto.HandshakePacket{
			ProtoVer:   ProtocolVersion,
			ServerAddr: *ServerHost,
			ServerPort: uint16(*ServerPort),
			NextState:  mcproto.ConnectionStateStatus,
		})

	// Actually update our connection as well
	conn.SwitchState(mcproto.ConnectionStateStatus)

	if err != nil {
		log.Fatalf("could not write handshake packet: %s", err)
	}

	err = conn.WritePacket(ctx, &mcproto.ServerQueryPacket{})

	if err != nil {
		log.Fatalf("could not write request packet: %s", err)
	}

	packet, err := conn.ReadPacket(ctx)

	if err != nil {
		log.Fatalf("could not read response packet: %s", err)
	}

	response, ok := packet.(*mcproto.ServerInfoPacket)

	if !ok {
		log.Fatalf("Server sent unexpected packet: %s", packet.String())
	}

	jsonRes, _ := json.Marshal(response.Response)

	fmt.Println(string(jsonRes))
}
