package main

import (
	"fmt"
	"net"
)

func main() {

	// Listen for incoming UDP packets
	addr, err := net.ResolveUDPAddr("udp", ":8000")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening for UDP packets:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening for UDP packets on", addr)

	// Read incoming UDP packets and print them
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading UDP packet:", err)
			continue
		}
		fmt.Printf("Received %d bytes from %s: %s\n", n, addr.String(), string(buf[:n]))
	}
}
