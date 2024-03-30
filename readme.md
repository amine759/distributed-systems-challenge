# Distributed Systems Challenge

This repository contains an example of a simple client-server application implementing remote procedure calls (RPC) in C and Go. The purpose of this project is to demonstrate the basic principles of distributed systems using RPC.

## Overview

The project consists of two main components:

- **Server**: Implements a server program that exposes an RPC function for performing addition operations.
- **Client**: Implements a client program that connects to the server and makes RPC calls to perform addition operations remotely.

## Requirements

- Unix-like operating system (e.g., Linux, macOS)
- C compiler (e.g., GCC)
- RPCGEN utility

## To run Go examples 
for go code you can directly run the files
```bash
go run server.go
go run client.go
```
## To run C examples
check out the readme file in the `C/` directory