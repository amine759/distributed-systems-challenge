package main

import (
    "fmt"
    "net"
    "os"
)

const (
    bufferSize = 1024
)

func handleRequest(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, bufferSize)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Handle the request and send the response
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }

    defer listener.Close()
    fmt.Println("Listening on :8080...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err)
            os.Exit(1)
        }

        go handleRequest(conn)
    }
}