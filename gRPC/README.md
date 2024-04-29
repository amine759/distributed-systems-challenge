# a basic gRPC client - server call with protobuff
to run :
install grpc and protobuffer package dependencies
```bash
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
```
```bash
sudo apt install -y protobuf-compiler
```
to generate grpc code from .proto
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative contract/contract.proto
```
run client and server 
```
go run server/server.go
go run client/client.go
```