package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

//  https://docs.amplify.aws/cli/function/secrets#accessing-the-values-in-your-function
func GetSSMParameterNames() ([]string, error) {
	secretNames := []string{"capem", "clientpem", "clientkeypem"}
	ssmParameterNames := []string{}
	for _, name := range secretNames {
		ssmParameterNames = append(ssmParameterNames, os.Getenv(name))
	}
	return ssmParameterNames, nil
}

type MyResponse struct {
	Message []string `json:"Names"`
}

func HandleRequest(ctx context.Context) (MyResponse, error) {
	names, err := GetSSMParameterNames()
	if err != nil {
		log.Print(err)
	}
	return MyResponse{names}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
