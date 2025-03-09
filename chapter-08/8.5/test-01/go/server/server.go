package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const MAX_CONNECTION_NUM = 1100000

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("\n Usage: %s <server ip> <server port>\n", os.Args[0])
		return
	}

	ip := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("\n Error: Invalid server port")
		return
	}

	// Create server
	addr := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("\n Error: Could not create listener")
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", addr)

	// Accept connections
	sockets := make([]net.Conn, MAX_CONNECTION_NUM)
	i := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("\n Error: Accept Failed")
			continue
		}
		sockets[i] = conn
		i++
		fmt.Printf("%s %d accept success: %d\n", ip, port, i)
	}
}
