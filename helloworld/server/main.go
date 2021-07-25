package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/jizusun/trojan-manager/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50055"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) ServerStreamingSayHello(in *pb.HelloRequest, stream pb.Greeter_ServerStreamingSayHelloServer) error {
	greetings := []string{
		"hello",
		"hi",
	}
	for _, greeting := range greetings {
		g := &pb.HelloReply{Message: greeting + ", " + in.GetName()}
		if err := stream.Send(g); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var (
		tlsCert = "certs/server.pem"
		tlsKey  = "certs/server-key.pem"
		caCert  = "certs/ca.pem"
	)
	cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
	if err != nil {
		log.Fatal(err)
	}
	rawCaCert, err := ioutil.ReadFile(caCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCaCert)
	// https://github.com/kelseyhightower/grpc-hello-service/blob/4be93eb92cbc1f21b3ed2519cdc9f1d3ef5e9143/hello-server/main.go#L55
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
