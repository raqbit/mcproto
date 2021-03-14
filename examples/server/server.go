package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/Raqbit/mcproto"
	"github.com/google/uuid"
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
	conn := mcproto.WrapConnection(tcpConn, mcproto.ServerSide)
	defer conn.Close()

	for {
		p, err := conn.ReadPacket(context.Background())

		if err != nil {
			if errors.Is(err, io.EOF) {
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

func handlePacket(conn mcproto.Connection, p mcproto.Packet) error {
	switch v := p.(type) {
	case *mcproto.HandshakePacket:
		return handleHandshakePacket(conn, v)
	case *mcproto.ServerQueryPacket:
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

func handleClientSettingsPacket(_ mcproto.Connection, _ *mcproto.ClientSettingsPacket) error {
	return nil
}

func handleLoginPacket(conn mcproto.Connection, v *mcproto.LoginStartPacket) error {
	playerUuid, _ := uuid.NewRandom()
	err := conn.WritePacket(context.Background(), &mcproto.LoginSuccessPacket{UUID: playerUuid, Username: v.Name})

	if err != nil {
		return fmt.Errorf("could not write login success packet: %w", err)
	}

	conn.SwitchState(mcproto.ConnectionStatePlay)

	err = conn.WritePacket(context.Background(), &mcproto.JoinGamePacket{
		PlayerID:            genRandomEid(),
		GameMode:            1,
		Dimension:           0,
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

	data := mcproto.NewPacketBuffer(new(bytes.Buffer))
	err = data.WriteString("mcproto custom")

	if err != nil {
		return fmt.Errorf("could not write channel data")
	}

	err = conn.WritePacket(context.Background(), &mcproto.PluginMessagePacket{
		Channel: mcproto.NewIdentifier("minecraft", "brand"),
		Data:    data,
	})

	if err != nil {
		return fmt.Errorf("could not write brand packet: %w", err)
	}

	err = conn.WritePacket(context.Background(), &mcproto.PlayerPositionLookPacket{
		X:          0,
		Y:          0,
		Z:          0,
		Yaw:        0,
		Pitch:      0,
		Flags:      0,
		TeleportID: 0,
	})

	if err != nil {
		return fmt.Errorf("could not write player posision and look packet: %w", err)
	}

	for i := 0; i < 40; i += 1 {
		err = conn.WritePacket(context.Background(), &mcproto.ChatMessagePacket{
			Message: mcproto.TextComponent{
				Text:  "Sent from mcproto",
				Color: "red",
			},
		})

		if err != nil {
			return fmt.Errorf("could not chat message packet: %w", err)
		}
	}

	return nil
}

func handlePingPacket(conn mcproto.Connection, v *mcproto.PingPacket) error {
	return conn.WritePacket(context.Background(), &mcproto.PongPacket{Payload: v.Payload})
}

func handleRequestPacket(conn mcproto.Connection, _ *mcproto.ServerQueryPacket) error {
	status := mcproto.ServerInfo{
		Version: mcproto.Version{
			Name:     "mcproto-custom",
			Protocol: ProtocolVersion,
		},
		Players: mcproto.Players{
			Max:    9001,
			Online: 69,
			Sample: []mcproto.Player{},
		},
		Description: mcproto.ServerDescription{
			Text:  "mcproto example server",
			Color: "red",
		},
	}

	err := conn.WritePacket(context.Background(), &mcproto.ServerInfoPacket{
		Response: status,
	})

	if err != nil {
		return fmt.Errorf("could not write response packet: %w", err)
	}

	return nil
}

func handleHandshakePacket(conn mcproto.Connection, p *mcproto.HandshakePacket) error {
	if p.ProtoVer != ProtocolVersion {
		return fmt.Errorf("unsupported protocol version: %d", p.ProtoVer)
	}

	if p.NextState != mcproto.ConnectionStateStatus && p.NextState != mcproto.ConnectionStateLogin {
		return fmt.Errorf("handshake packet with invalid next state")
	} else {
		conn.SwitchState(p.NextState)
		return nil
	}
}

func genRandomEid() int32 {
	return int32(rand.Intn(math.MaxInt32))
}
