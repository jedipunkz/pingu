package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on UDP port
	addr := net.UDPAddr{
		Port: 12345,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		// Read from UDP connection
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr)

		// Echo back to client
		if _, err := conn.WriteToUDP(buffer[:n], addr); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
