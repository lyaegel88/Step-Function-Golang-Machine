package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Request struct {
	TransactionType string `json:"TransactionType"`
	Timestamp string `json:"Timestamp"`
	Message string `json:"Message"`
	Total int `json:"Total"`
}

type Response struct {
	TransactionType string `json:"TransactionType"`
	Timestamp string `json:"Timestamp"`
	Message string `json:"Message"`
	Total int `json:"Total"`
	Email string `json:"Email"`
}

func Handler(request Request)  (Response, error){

	fmt.Print("Sending Confirmation Email")
	fmt.Print(request.TransactionType)

	return Response{
		TransactionType: request.TransactionType,
		Timestamp: time.Now().Format("2006-01-02 03:04:05"),
		Message: "Your purchase has been processed!",
		Total: request.Total,
		Email: "customer@example.com",
	}, nil
}

func main()  {
	lambda.Start(Handler)
}