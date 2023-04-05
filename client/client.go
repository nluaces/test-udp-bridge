package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	serverAddr, err := net.ResolveUDPAddr("udp", "udp-server:8000")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	// Send a message to the server
	message := []byte("Hello, server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending UDP packet:", err)
		return
	}

	fmt.Println("Sent message to server:", string(message))

}
