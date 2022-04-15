# go-grpc

$ go mod init github.com/mg52/go-grpc

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ export PATH="$PATH:$(go env GOPATH)/bin"

// create grpc-stream/greetpb/greet.proto

$ protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
grpc-stream/greetpb/greet.proto

$ go get google.golang.org/grpc

// add sever and client folder and files.

---
// run server in a terminal

$ go run grpc-stream/server/server.go

---
// run client in another terminal for server streaming

$ go run grpc-stream/client/client.go -o 0 -f name -l lastname

---
// run client in another terminal for client streaming

$ go run grpc-stream/client/client.go -o 1