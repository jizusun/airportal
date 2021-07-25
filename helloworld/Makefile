.PHONY: protobuf-gen server client

server:
	go run server/main.go

client:
	go run client/main.go $(name)

# https://grpc.io/docs/languages/go/quickstart/
protobuf-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

cfssl-cert-gen:
	cd certs ;\
	rm *.pem ;\
	sh ../scripts/cfssl.sh ;\
	# cp client-key.pem api/client.key ;\
	# cp client.pem  api/client.crt ;\
	# cp server-key.pem api/server.key ;\
	# cp server.pem  api/server.crt

rsync: 
	rsync -azP certs/api server.json virmach:~/trojan-go-linux-amd64
