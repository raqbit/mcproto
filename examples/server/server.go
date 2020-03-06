package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
	"log"
	"math"
	"math/rand"
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

			log.Printf("Error reading packet: %s", err)
			continue
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
	case *mcproto.ClientSettingsPacket:
		return handleClientSettingsPacket(conn, v)
	default:
		return fmt.Errorf("unhandled packet: %s", p)
	}
}

func handleClientSettingsPacket(conn *mcproto.Connection, v *mcproto.ClientSettingsPacket) error {
	conn.WritePacket(mcproto.PlayerPositionAndLookPacket{
		X:          0,
		Y:          0,
		Z:          0,
		Yaw:        0,
		Pitch:      0,
		Flags:      0,
		TeleportID: 0,
	})
	return nil
}

func handleLoginPacket(conn *mcproto.Connection, v *mcproto.LoginStartPacket) error {
	conn.State = mcproto.PlayState
	err := conn.WritePacket(mcproto.LoginSuccessPacket{UUID: "f2bf38cd-0073-4703-94fa-d49d406a4885", Username: v.Name})

	if err != nil {
		return fmt.Errorf("could not write login success packet: %w", err)
	}

	err = conn.WritePacket(mcproto.JoinGamePacket{
		EntityID:            enc.Int(genRandomEid()),
		GameMode:            1,
		Dimension:           enc.Int(0),
		HashedSeed:          0,
		MaxPlayers:          10,
		LevelType:           "flat",
		ViewDistance:        2,
		ReducedDebug:        false,
		EnableRespawnScreen: true,
	})

	if err != nil {
		return fmt.Errorf("could not write join game packet: %w", err)
	}

	buffer := new(bytes.Buffer)

	enc.WriteString(buffer, "Raqbit custom")

	err = conn.WritePacket(mcproto.PluginMessagePacket{
		Channel: "minecraft:brand",
		Data:    buffer.Bytes(),
	})

	if err != nil {
		return fmt.Errorf("could not write brand packet: %w", err)
	}

	return nil
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

	if p.NextState != mcproto.StatusState && p.NextState != mcproto.LoginState {
		return fmt.Errorf("handshake packet with invalid next state")
	} else {
		conn.State = mcproto.ConnectionState(p.NextState)
		return nil
	}
}

func genRandomEid() int32 {
	return int32(rand.Intn(math.MaxInt32))
}
