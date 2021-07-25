# mTLS for gRPC



- [grpc-go examples](https://github.com/grpc/grpc-go/blob/master/examples/README.md)
- [Logging](https://github.com/grpc/grpc-go#how-to-turn-on-logging)

```sh
export GRPC_GO_LOG_VERBOSITY_LEVEL=99
export GRPC_GO_LOG_SEVERITY_LEVEL=info
```

## Start the server

```sh
make server
```

## Start the client 
```sh
make name=foobar client
```