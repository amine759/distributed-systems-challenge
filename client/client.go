package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// Data struct represents the data to be sent
type Data struct {
    ID   int
    Name string
    // Add more fields as needed
}

func main() {
    // Data to be sent
    data := Data{
        ID:   1,
        Name: "Example",
    }

    // Convert data to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error marshalling data:", err)
        return
    }

    // Send data via HTTP POST request to the master node
    resp, err := http.Post("http://localhost:8080/send", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error sending data to master node:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Data sent successfully to master node")
}
