package main

import (
	"bytes"
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
	case *mcproto.CHandshakePacket:
		return handleHandshakePacket(conn, v)
	case *mcproto.CServerQueryPacket:
		return handleRequestPacket(conn, v)
	case *mcproto.CPingPacket:
		return handlePingPacket(conn, v)
	case *mcproto.CLoginStartPacket:
		return handleLoginPacket(conn, v)
	case *mcproto.CClientSettingsPacket:
		return handleClientSettingsPacket(conn, v)
	default:
		return fmt.Errorf("unhandled packet: %s", p)
	}
}

func handleClientSettingsPacket(_ *mcproto.Connection, _ *mcproto.CClientSettingsPacket) error {
	return nil
}

func handleLoginPacket(conn *mcproto.Connection, v *mcproto.CLoginStartPacket) error {
	conn.State = mcproto.PlayState
	playerUuid, _ := uuid.NewRandom()
	err := conn.WritePacket(&mcproto.SLoginSuccessPacket{Profile: mcproto.NewGameProfile(playerUuid, "Raqbit")})

	if err != nil {
		return fmt.Errorf("could not write login success packet: %w", err)
	}

	err = conn.WritePacket(&mcproto.SJoinGamePacket{
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
	err = data.WriteString("Raqbit custom")

	if err != nil {
		return fmt.Errorf("could not write channel data")
	}

	err = conn.WritePacket(&mcproto.PluginMessagePacket{
		Channel: mcproto.NewResourceLocation("minecraft", "brand"),
		Data:    data,
	})

	if err != nil {
		return fmt.Errorf("could not write brand packet: %w", err)
	}

	err = conn.WritePacket(&mcproto.SPlayerPositionLookPacket{
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

	return nil
}

func handlePingPacket(conn *mcproto.Connection, v *mcproto.CPingPacket) error {
	return conn.WritePacket(&mcproto.SPongPacket{Payload: v.Payload})
}

func handleRequestPacket(conn *mcproto.Connection, _ *mcproto.CServerQueryPacket) error {
	status := mcproto.ServerInfo{
		Version: mcproto.Version{
			Name:     "raqbit-custom",
			Protocol: ProtocolVersion,
		},
		Players: mcproto.Players{
			Max:    9001,
			Online: 69,
			Sample: []mcproto.Player{},
		},
		Description: mcproto.TextComponent{
			Text:  ":partyparrot:",
			Color: "green",
		},
	}

	err := conn.WritePacket(&mcproto.SServerInfoPacket{
		Response: status,
	})

	if err != nil {
		return fmt.Errorf("could not write response packet: %w", err)
	}

	return nil
}

func handleHandshakePacket(conn *mcproto.Connection, p *mcproto.CHandshakePacket) error {
	if p.ProtoVer != ProtocolVersion {
		return fmt.Errorf("unsupported protocol version: %d", p.ProtoVer)
	}

	if p.NextState != mcproto.StatusState && p.NextState != mcproto.LoginState {
		return fmt.Errorf("handshake packet with invalid next state")
	} else {
		conn.State = p.NextState
		return nil
	}
}

func genRandomEid() int32 {
	return int32(rand.Intn(math.MaxInt32))
}
