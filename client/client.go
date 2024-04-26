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

func main() {
	// Connect to the master node
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Create a new Data instance to send
	data := Data{
		ID:   1,
		Name: "Example Data",
	}

	// Initialize an encoder to send data over the network
	encoder := gob.NewEncoder(conn)

	// Send the data
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}

	fmt.Println("Data sent successfully!")
}
