# Pingu

UDP Quality Measure Tool

This is a simple tool to measure the quality of UDP communication. The tool consists of a server and a client, both implemented in Go. The client sends a message to the server, and the server echoes back the message. The client then calculates the Round Trip Time (RTT) based on the sent and received time of the message.

## Requirements

- Go (version 1.16 or later)

## Usage

### Server

1. Navigate to the directory where the server.go file is located.
2. Run the server using the following command:

```bash
go run server.go
```

The server listens for incoming messages on UDP port 12345.

### Client

1. Open a new terminal and navigate to the directory where the client.go file is located.
2. Run the client using the following command:

```bash
go run client.go <IP_ADDR_OF_SERVER>
```

The client sends a message to the server, waits for the echo, and calculates the RTT. This process is repeated every second until you interrupt the process (Ctrl-C).

## Output

The client outputs the sent time, received time, and RTT for each message. Here is an example of the client's output:

```
Sent message to server at  2023-05-14 12:34:56.789123 +0000 UTC m=+0.000000001
Received response from server at  2023-05-14 12:34:56.789456 +0000 UTC m=+0.000000333
RTT is  333ns
Received:  hello, server!
```

The output will continue until you stop the client.

## Author

@jedipunkz

