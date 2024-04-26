package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Data struct represents the data to be sent
type Data struct {
	ID   int
	Name string
	// Add more fields as needed
}

type FileMetadata struct {
	Name      string // Name of the file
	Size      int64  // Size of the file in bytes
	Timestamp string // Last modified timestamp
}

// Connect to the master node
func main() {
	// Start listening for incoming connections from master
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("node1 is listening on the master :8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle incoming connection in a separate goroutine
		go handleConnection(conn)
	}
}

// handleConnection handles an incoming connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Initialize a decoder to decode incoming data
	decoder := gob.NewDecoder(conn)

	// Decode the data received from the client
	var data Data
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}

	// Process the received data
	fmt.Printf("Received data: %+v\n", data)

	// Handle the data further as needed (e.g., store in the distributed file system)
}
