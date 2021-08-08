package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

//  https://docs.amplify.aws/cli/function/secrets#accessing-the-values-in-your-function
func GetSSMParameterNames() ([]string, error) {
	secretNames := []string{"capem", "clientpem", "clientkeypem"}
	ssmParameterNames := []string{}
	for _, name := range secretNames {
		ssmParameterNames = append(ssmParameterNames, os.Getenv(name))
	}
	return ssmParameterNames, nil
}

type BodyResponse struct {
	Message []string `json:"names"`
}

// https://github.com/serverless/examples/blob/master/aws-golang-http-get-post/postFolder/postExample.go
func HandleRequest(ctx context.Context) (Response, error) {
	names, err := GetSSMParameterNames()
	if err != nil {
		log.Print(err)
	}
	bodyResponse := BodyResponse{names}
	response, err := json.Marshal(&bodyResponse)
	headers := map[string]string{
		"Content-Type":                 "application/json",
		"X-MyCompany-Func-Reply":       "hello-handler",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
		"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}
	if err != nil {
		return Response{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
		}, nil
	}
	return Response{
		Body:       string(response),
		StatusCode: 200,
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
