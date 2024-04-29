// client.go
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/rpc"
	"github.com/davecgh/go-xdr/xdr2"
)

type Args struct {
	A, B int // Define two fields to hold the numbers to be added
}

func main() {
	// Declare a variable to store the sum returned by the server
	serverAddress := "localhost"

	// Connect to the RPC server using DialHTTP
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Encode the Args struct into XDR format
	args := Args{A: 7, B: 8}
	var buf bytes.Buffer
	_, err = xdr.Marshal(&buf, args)
	if err != nil {
		log.Fatal("Error encoding:", err)
	}

	// Call the RPC function "Add" on the server, passing the encoded args struct
	var sum int

	fmt.Printf("Encoded data: %v\n", buf.Bytes())

	err = client.Call("Arith.Add", buf.Bytes(), &sum)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	// Print the sum returned by the server
	log.Printf("Sum: %d", sum)
}

