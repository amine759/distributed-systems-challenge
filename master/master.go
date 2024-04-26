package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Data struct represents the data received from the client
type Data struct {
	ID   int
	Name string
}

type FileMetadata struct {
    Name        string   // Name of the file
    Size        int64    // Size of the file in bytes
    Timestamp   string   // Last modified timestamp
}

func main() {
	// Start listening for incoming connections from clients
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Master node is listening on :8080...")

	// Accept incoming connections indefinitely
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		data := &Data{}
		// Handle incoming connection in a separate goroutine
		go forwardData(conn, data)
	}
}

// handleConnection handles an incoming connection
func forwardData(conn net.Conn, data *Data) {
	defer conn.Close()
	storageNodeAddrs := [3]string{":8081", ":8082", ":8083"}
	// Forward the received data to all storage nodes
	for _, storageNodeAddr := range storageNodeAddrs {
		// Connect to the storage node
		storageConn, err := net.Dial("tcp", storageNodeAddr)
		if err != nil {
			fmt.Printf("Error connecting to storage node %s: %v\n", storageNodeAddr, err)
			continue
		}
		defer storageConn.Close()

		// Encode and send the data to the storage node
		encoder := gob.NewEncoder(storageConn)
		if err := encoder.Encode(data); err != nil {
			fmt.Printf("Error encoding and sending data to storage node %s: %v\n", storageNodeAddr, err)
			continue
		}

		fmt.Printf("Data forwarded to storage node %s\n", storageNodeAddr)
	}

	fmt.Println("Data forwarded to all storage nodes")
}

