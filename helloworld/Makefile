.PHONY: protobuf server client

server:
	go run server/main.go

client:
	go run client/main.go $(name)

# https://grpc.io/docs/languages/go/quickstart/
protobuf:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
