package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run client.go <server IP>")
		os.Exit(1)
	}

	serverIP := os.Args[1]
	serverAddr := fmt.Sprintf("%s:12345", serverIP)

	// Resolve UDP address
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		panic(err)
	}

	// Connect to server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// Handle Ctrl+C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nCaught interrupt. Exiting...")
		conn.Close()
		os.Exit(1)
	}()

	for {
		// Write a message to server
		message := []byte("hello, server!")
		sendTime := time.Now()
		_, err = conn.Write(message)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sent message to server at ", sendTime)

		// Read response from server
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		receiveTime := time.Now()
		fmt.Println("Received response from server at ", receiveTime)

		// Calculate RTT
		rtt := receiveTime.Sub(sendTime)
		fmt.Println("RTT is ", rtt)

		fmt.Println("Received: ", string(buffer[:n]))

		// Sleep for a second before next round
		time.Sleep(1 * time.Second)
	}
}
