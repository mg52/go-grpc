# go-grpc

$ go mod init github.com/mg52/go-grpc

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ export PATH="$PATH:$(go env GOPATH)/bin"

// create server-stream/greetpb/greet.proto

$ protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
server-stream/greetpb/greet.proto

$ go get google.golang.org/grpc

// add sever and client folder and files.

$ go run server-stream/server/server.go

$ go run server-stream/client/client.go -f name -l lastname