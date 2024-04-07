package protocol

import (
	"encoding/binary"
	"net"
)

type Packet struct {
	id         int64
	length     int32
	packetType int16
	body       string
}

func (p Packet) Id() int64 {
	return p.id
}

func (p Packet) Length() int32 {
	return p.length
}

func (p Packet) PacketType() int16 {
	return p.packetType
}

func (p Packet) Body() string {
	return p.body
}

func ReadMessage(c net.Conn) Packet {
	packetId := make([]byte, 8)
	packetIdN, _ := c.Read(packetId)
	packetIdM := binary.BigEndian.Uint64(packetId[:packetIdN])

	packetLength := make([]byte, 4)
	packetLengthN, _ := c.Read(packetLength)
	packetLengthM := binary.BigEndian.Uint32(packetLength[:packetLengthN])

	packetType := make([]byte, 2)
	packetTypeN, _ := c.Read(packetType)
	packetTypeM := binary.BigEndian.Uint16(packetType[:packetTypeN])

	packetBody := make([]byte, packetLengthM)
	packetBodyN, _ := c.Read(packetBody)

	packet := Packet{
		id:         int64(packetIdM),
		length:     int32(packetLengthM),
		packetType: int16(packetTypeM),
		body:       string(packetBody[:packetBodyN]),
	}
	return packet
}
