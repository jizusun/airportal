package controllers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/p4gefau1t/trojan-go/api/service"
	"github.com/p4gefau1t/trojan-go/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ListUsers(host string, port int) error {
	// https://github.com/p4gefau1t/trojan-go/blob/07fec5eb263658766a5f3597c98bbeb7d7843f31/api/service/server_test.go#L160-L313
	pool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile("../certs/ca.pem")
	common.Must(err)
	pool.AppendCertsFromPEM(certBytes)

	certificate, err := tls.LoadX509KeyPair("../certs/client.pem", "../certs/client-key.pem")
	common.Must(err)
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "control.edtechstar.com",
		RootCAs:      pool,
		Certificates: []tls.Certificate{certificate},
	})

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(creds))
	common.Must(err)
	// https://github.com/kbehouse/gRPC-go-mTLS/blob/f77a12c8038c170137e1b04dc2b6ff94accb92c5/helloworld_mTLS/greeter_client/main.go#L82

	server := service.NewTrojanServerServiceClient(conn)
	stream, err := server.ListUsers(context.Background(), &service.ListUsersRequest{})
	if err != nil {
		return err
	}
	defer stream.CloseSend()
	result := []*service.ListUsersResponse{}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		result = append(result, resp)
	}
	data, err := json.Marshal(result)
	common.Must(err)
	fmt.Println(string(data))
	return nil
	// common.Must(err)
	// stream.CloseSend()
	// conn.Close()
}
