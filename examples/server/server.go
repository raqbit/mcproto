package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
	"log"
	"net"
)

const (
	ProtocolVersion = 578
)

func main() {
	listener, err := net.Listen("tcp", ":25565")

	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("tcp server accept error", err)
		}

		// spawn off goroutine to able to accept new connections
		go handleConnection(conn)
	}
}

// TODO: timeouts
// TODO: clean shutdown

func handleConnection(tcpConn net.Conn) {
	conn := mcproto.NewConnection(tcpConn, mcproto.ServerSide)
	defer conn.Close()

	for {
		p, err := conn.ReadPacket()

		if err != nil {
			if errors.Unwrap(err) == io.EOF {
				log.Println("Client closed the connection")
				return
			}

			log.Printf("Invalid packet: %s", err)
			return
		}

		err = handlePacket(conn, p)

		if err != nil {
			log.Printf("Closing connection: %s\n", err)
			return
		}
	}
}

func handlePacket(conn *mcproto.Connection, p mcproto.Packet) error {
	switch v := p.(type) {
	case *mcproto.HandshakePacket:
		return handleHandshakePacket(conn, v)
	case *mcproto.RequestPacket:
		return handleRequestPacket(conn, v)
	case *mcproto.PingPacket:
		return handlePingPacket(conn, v)
	case *mcproto.LoginStartPacket:
		return handleLoginPacket(conn, v)
	default:
		return fmt.Errorf("unhandled packet: %s", p)
	}
}

func handleLoginPacket(conn *mcproto.Connection, v *mcproto.LoginStartPacket) error {
	conn.State = mcproto.PlayState
	return conn.WritePacket(mcproto.LoginSuccessPacket{UUID: "f2bf38cd-0073-4703-94fa-d49d406a4885", Username: v.Name})
}

func handlePingPacket(conn *mcproto.Connection, v *mcproto.PingPacket) error {
	return conn.WritePacket(mcproto.PongPacket{Payload: v.Payload})
}

func handleRequestPacket(conn *mcproto.Connection, _ *mcproto.RequestPacket) error {
	status := &mcproto.ServerInfo{
		Version: mcproto.Version{
			Name:     "raqbit-custom",
			Protocol: ProtocolVersion,
		},
		Players: mcproto.Players{
			Max:    9001,
			Online: 69,
			Sample: []mcproto.Player{},
		},
		Description: mcproto.ChatComponent{
			Text:  ":partyparrot:",
			Color: "green",
		},
	}

	jsonStatus, err := json.Marshal(status)

	if err != nil {
		return fmt.Errorf("could not marshal server info: %w", err)
	}

	err = conn.WritePacket(mcproto.ResponsePacket{
		Json: enc.String(jsonStatus),
	})

	if err != nil {
		return fmt.Errorf("could not write response packet: %w", err)
	}

	return nil
}

func handleHandshakePacket(conn *mcproto.Connection, p *mcproto.HandshakePacket) error {
	if p.ProtoVer != ProtocolVersion {
		return fmt.Errorf("unsupported protocol version: %d", p.ProtoVer)
	}

	switch p.NextState {
	case 1:
		conn.State = mcproto.StatusState
	case 2:
		conn.State = mcproto.LoginState
	default:
		return fmt.Errorf("handshake packet with invalid next state")
	}

	return nil
}
