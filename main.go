package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type DogRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

type DogResponse struct {
	Message string `json:"message"`
}

func getHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := DogResponse{
		Message: "Post your dog's name and breed!",
	}
	jbytes, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(jbytes), StatusCode: 200}, nil
}

func postHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var dreq DogRequest
	err := json.Unmarshal([]byte(request.Body), &dreq)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
	}
	dres := DogResponse{
		Message: fmt.Sprintf("Your dog is named '%s' and it's a %s.", dreq.Name, dreq.Breed),
	}
	jbytes, err := json.Marshal(dres)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(jbytes), StatusCode: 200}, nil
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return getHandler(ctx, request)
	case "POST":
		return postHandler(ctx, request)
	default:
		return events.APIGatewayProxyResponse{Body: "Method not allowed", StatusCode: 405}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
