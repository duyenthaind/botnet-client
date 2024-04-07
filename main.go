package main

import (
	client "example/botnet-tcp-client"
	"fmt"
	"strconv"
)

func main() {
	// connect to server
	config := client.Connection{}
	config.SetHost("127.0.0.1")
	config.SetPort(int32(8888))

	client.Connect(config)
	fmt.Println("Connected to socket on ", config.Host()+":"+strconv.Itoa(int(config.Port())))
}
