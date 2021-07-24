PHONY:
	gen


# https://grpc.io/docs/languages/go/quickstart/
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto