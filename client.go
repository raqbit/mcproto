package mcproto

import (
	"net"
)

type Client struct {
	conn *Connection
}

// Creates a new client
func NewClient(tcpConn net.Conn) *Client {
	return &Client{conn: NewConnection(tcpConn, ClientSide)}
}
