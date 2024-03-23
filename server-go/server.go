package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Decode XDR data from the request body
	var a, b int32
	buf := bytes.NewReader(body)
	if err := binary.Read(buf, binary.BigEndian, &a); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if err := binary.Read(buf, binary.BigEndian, &b); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Calculate the sum
	sum := int(a) + int(b)

	// Write the JSON-RPC response with the sum
	response := []byte(`{"jsonrpc": "2.0", "result": ` +  strconv.Itoa(sum) + `, "id": 1}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

