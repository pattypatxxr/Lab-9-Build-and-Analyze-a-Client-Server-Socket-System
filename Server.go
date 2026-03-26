package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

const (
	PORT            = ":8080"
	READ_TIMEOUT    = 5 * time.Second
	WRITE_TIMEOUT   = 5 * time.Second
	MAX_BUFFER_SIZE = 1024
)

func handleConnection(conn net.Conn) {
	
	defer conn.Close()

	// Set read deadline (prevents hanging / slowloris)
	conn.SetReadDeadline(time.Now().Add(READ_TIMEOUT))

	buffer := make([]byte, MAX_BUFFER_SIZE)

	n, err := conn.Read(buffer)

	if len(buffer[:n]) > 100 {
    	return
	}

	if err != nil {
		if err == io.EOF {
			fmt.Println("Client disconnected")
		} else {
			fmt.Println("Read error:", err)
		}
		return
	}

	fmt.Println("Received:", string(buffer[:n]))

	// Set write deadline
	conn.SetWriteDeadline(time.Now().Add(WRITE_TIMEOUT))

	_, err = conn.Write([]byte("Hello from Go server"))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Println("Client:", conn.RemoteAddr())
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server running on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go handleConnection(conn) // still concurrent
	}
}
