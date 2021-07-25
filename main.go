package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jizusun/trojan-manager/controllers"
)

func main() {
	router := gin.Default()

	grpcClientConn, err := controllers.NewGrpcClient(controllers.GrpcClientConfig{
		ClientCertFile: "certs/client.pem",
		ClientKeyFile:  "certs/client-key.pem",
		RootCaCertFile: "certs/ca.pem",
		Host:           "virmach-go.edtechstar.com",
		Port:           10000,
	})
	if err != nil {
		log.Fatal(err)
	}
	userController := controllers.NewUserController(controllers.UserControllerConfig{
		Client: grpcClientConn,
	})
	router.GET("/users", userController.ListUsers)
	router.Run()
}
