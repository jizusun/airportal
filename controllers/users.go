package controllers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p4gefau1t/trojan-go/api/service"
	"github.com/p4gefau1t/trojan-go/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})

type UserControllerConfig struct {
	Client *grpc.ClientConn
}

type userController struct {
	// UserControllerConfig
	Client service.TrojanServerServiceClient
}

type GrpcClientConfig struct {
	ClientCertFile string
	ClientKeyFile  string
	RootCaCertFile string
	Host           string
	Port           int
}

func NewGrpcClient(config GrpcClientConfig) (*grpc.ClientConn, error) {
	// https://github.com/p4gefau1t/trojan-go/blob/07fec5eb263658766a5f3597c98bbeb7d7843f31/api/service/server_test.go#L160-L313
	pool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(config.RootCaCertFile)
	common.Must(err)
	pool.AppendCertsFromPEM(certBytes)

	certificate, err := tls.LoadX509KeyPair(config.ClientCertFile, config.ClientKeyFile)
	common.Must(err)
	creds := credentials.NewTLS(&tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{certificate},
	})

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(creds))
	common.Must(err)
	return conn, nil
}

func NewUserController(config UserControllerConfig) userController {
	// https://github.com/kbehouse/gRPC-go-mTLS/blob/f77a12c8038c170137e1b04dc2b6ff94accb92c5/helloworld_mTLS/greeter_client/main.go#L82
	client := service.NewTrojanServerServiceClient(config.Client)
	return userController{
		Client: client,
	}
}

func (u *userController) ListUsers(c *gin.Context) {
	stream, err := u.Client.ListUsers(context.Background(), &service.ListUsersRequest{})
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}
	defer stream.CloseSend()
	result := []*service.ListUsersResponse{}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		result = append(result, resp)
	}
	c.JSON(200, result)
}
