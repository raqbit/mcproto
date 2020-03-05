package main

import (
	"flag"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"log"
	"net"
	"os"
)

const (
	ProtocolVersion = 578
)

var (
	ServerHost = flag.String("host", "", "server host")
	ServerPort = flag.Int("port", 25565, "server port")
)

func main() {
	flag.Parse()

	if *ServerHost == "" {
		flag.Usage()
		os.Exit(2)
	}

	tcpConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *ServerHost, *ServerPort))

	if err != nil {
		log.Fatalf("tcp dial error: %s", err)
	}

	conn := mcproto.NewConnection(tcpConn, mcproto.ClientSide)

	err = conn.WritePacket(mcproto.HandshakePacket{
		ProtoVer:   ProtocolVersion,
		ServerAddr: enc.String(*ServerHost),
		ServerPort: enc.UnsignedShort(*ServerPort),
		NextState:  mcproto.StatusState,
	})

	// Actually update our connection as well
	conn.State = mcproto.StatusState

	if err != nil {
		log.Fatalf("could not write handshake packet: %s", err)
	}

	err = conn.WritePacket(mcproto.RequestPacket{})

	if err != nil {
		log.Fatalf("could not write request packet: %s", err)
	}

	packet, err := conn.ReadPacket()

	if err != nil {
		log.Fatalf("could not read response packet: %s", err)
	}

	response, ok := packet.(*mcproto.ResponsePacket)

	if !ok {
		log.Fatalf("Server sent unexpected packet: %s", packet.String())
	}

	fmt.Println(response.Json)
}
