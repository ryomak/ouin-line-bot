package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ryomak/ouin-line-bot/line-bot/src/handler"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler.LineHandler)
}
