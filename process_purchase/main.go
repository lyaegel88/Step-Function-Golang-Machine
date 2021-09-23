package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Request struct {
	TransactionType string `json:"TransactionType"`
	Total int `json:"Total"`
}

type Response struct {
	TransactionType string `json:"TransactionType"`
	Timestamp string `json:"Timestamp"`
	Message string `json:"Message"`
	Total int `json:"Total"`
}

func Handler(request Request)  (Response, error){

	fmt.Print("Received message from Step Functions:")
	fmt.Print(request.TransactionType)

	return Response{
		TransactionType: request.TransactionType,
		Timestamp: time.Now().Format("2006-01-02 03:04:05"),
		Message: "Hello from lambda land inside the ProcessPurchase function",
		Total: request.Total,
	}, nil
}

func main()  {
	lambda.Start(Handler)
}