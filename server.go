package main

import (
    "bytes"
    "log"
    "net"
    "net/http"
    "net/rpc"
    "github.com/davecgh/go-xdr/xdr2"
)

type Args struct {
    A, B int
}

type Arith int
//this function receives encoded xdr, parses it into the pointed interface (args) 
func (t *Arith) Add(xdrData []byte, reply *int) error {
    var args Args
	
    log.Println("XDR data sent from client",xdrData)
	// Decode the received XDR data into Args struct
    _, err := xdr.Unmarshal(bytes.NewReader(xdrData), &args)
    if err != nil {
        return err
    }

    // Perform addition operation
    *reply = args.A + args.B
    return nil
}

func server() {

    arith := new(Arith)
    serverAddress := "0.0.0.0"

    rpc.Register(arith)
    rpc.HandleHTTP()

    listener, err := net.Listen("tcp", serverAddress+":1234")
    if err != nil {
        log.Fatal("Listener error: ", err)
    }
    log.Println("RPC server listening on", listener.Addr())

    err = http.Serve(listener, nil)
    if err != nil {
        log.Fatal("Serve error: ", err)
    }
}
func main() {
    server()
}

