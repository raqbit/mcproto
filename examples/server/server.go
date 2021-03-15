package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/packet"
	"github.com/Raqbit/mcproto/packet/channel"
	"github.com/Raqbit/mcproto/types"
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
	conn := mcproto.WrapConnection(tcpConn, types.ServerSide)
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

func handlePacket(conn mcproto.Connection, p packet.Packet) error {
	switch v := p.(type) {
	case *packet.HandshakePacket:
		return handleHandshakePacket(conn, v)
	case *packet.ServerQueryPacket:
		return handleRequestPacket(conn, v)
	case *packet.PingPacket:
		return handlePingPacket(conn, v)
	case *packet.LoginStartPacket:
		return handleLoginPacket(conn, v)
	case *packet.ClientSettingsPacket:
		return handleClientSettingsPacket(conn, v)
	default:
		return fmt.Errorf("unhandled packet: %s", p)
	}
}

func handleClientSettingsPacket(_ mcproto.Connection, _ *packet.ClientSettingsPacket) error {
	return nil
}

func handleLoginPacket(conn mcproto.Connection, v *packet.LoginStartPacket) error {
	playerUuid, _ := uuid.NewRandom()
	err := conn.WritePacket(context.Background(), &packet.LoginSuccessPacket{UUID: enc.UUID(playerUuid), Username: v.Name})

	if err != nil {
		return fmt.Errorf("could not write login success packet: %w", err)
	}

	conn.SetState(types.ConnectionStatePlay)

	err = conn.WritePacket(context.Background(), &packet.JoinGamePacket{
		PlayerID:            enc.Int(genRandomEid()),
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

	err = conn.WritePacket(context.Background(), &packet.PluginMessagePacket{
		Channel: channel.BrandChannelID,
		Data: &channel.BrandChannel{
			Brand: "mcproto custom",
		},
	})

	if err != nil {
		return fmt.Errorf("could not write brand packet: %w", err)
	}

	err = conn.WritePacket(context.Background(), &packet.PlayerPositionLookPacket{
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
		err = conn.WritePacket(context.Background(), &packet.ChatMessagePacket{
			Message: types.TextComponent{
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

func handlePingPacket(conn mcproto.Connection, v *packet.PingPacket) error {
	return conn.WritePacket(context.Background(), &packet.PongPacket{Payload: v.Payload})
}

func handleRequestPacket(conn mcproto.Connection, _ *packet.ServerQueryPacket) error {
	status := packet.ServerInfo{
		Version: packet.Version{
			Name:     "mcproto-custom",
			Protocol: ProtocolVersion,
		},
		Players: packet.Players{
			Max:    9001,
			Online: 69,
			Sample: []packet.Player{},
		},
		Description: packet.ServerDescription{
			Text:  "mcproto example server",
			Color: "red",
		},
	}

	err := conn.WritePacket(context.Background(), &packet.ServerInfoPacket{
		Response: status,
	})

	if err != nil {
		return fmt.Errorf("could not write response packet: %w", err)
	}

	return nil
}

func handleHandshakePacket(conn mcproto.Connection, p *packet.HandshakePacket) error {
	if p.ProtoVer != ProtocolVersion {
		return fmt.Errorf("unsupported protocol version: %d", p.ProtoVer)
	}

	if p.NextState != types.ConnectionStateStatus && p.NextState != types.ConnectionStateLogin {
		return fmt.Errorf("handshake packet with invalid next state")
	} else {
		conn.SetState(p.NextState)
		return nil
	}
}

func genRandomEid() int32 {
	return int32(rand.Intn(math.MaxInt32))
}
