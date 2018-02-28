package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//type Response struct {
//	Message string `json:"message"`
//}
//
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Recieved body: ", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

//func Handler() (Response, error) {
//	return Response{
//		Message: "Go Serverless v1.0! Your function executed successfully!",
//	}, nil
//}

func main() {
	lambda.Start(Handler)
}
