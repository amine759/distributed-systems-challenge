package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Data struct represents the data received from the client
type Data struct {
	ID   int
	Name string
	// Add more fields as needed
}
type StorageNodes struct {
	Node1Addr string
	Node2Addr string
	Node3Addr string
}

// SendData sends data to a storage node via gRPC
func SendDatagrpcHander(ctx context.Context, data *Data) error {
	// TODO: Implement gRPC logic to send data to storage node

	return nil
}

func getClientDataHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from the request body
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Failed to decode data", http.StatusBadRequest)
		return
	}

	// Print the received data
	fmt.Println("Received data from client:")
	fmt.Printf("ID: %d\n", data.ID)
	fmt.Printf("Name: %s\n", data.Name)
	// Add more fields as needed

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Data received successfully")
}

func main() {
	// Initialize gRPC server for storage node communication
	// TODO: Implement gRPC server logic for storage node communication
	// Start HTTP server to receive data from clients

	storageNodes := StorageNodes{
		Node1Addr: "0.0.0.0:8081",
		Node2Addr: "0.0.0.0:8082",
		Node3Addr: "0.0.0.0:8083",
	}
	// Establish connections with all storage nodes
	for i, nodeAddr := range []string{storageNodes.Node1Addr, storageNodes.Node2Addr, storageNodes.Node3Addr} {
		conn, err := grpc.Dial(nodeAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		// TODO here must be smtng like grpc.send(master,all_nodes) but sending grpc must be after receiving the http from client
		if err != nil {
			log.Fatalf("Failed to connect to node %s: %v", nodeAddr, err)
		}
		defer conn.Close()

		fmt.Printf("gRPC connection established with node %d: %s\n", i+1, nodeAddr)

		// Now you can use 'conn' for gRPC communication with the current node
	}

	http.HandleFunc("/send", getClientDataHTTPHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
