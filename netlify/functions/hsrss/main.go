package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	rssFile, err := os.ReadFile("../../../public/output.xml")
	if err != nil {
		log.Fatal(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			// https://stackoverflow.com/a/57173617/1171790
			"Content-Type":  "application/atom+xml;charset=UTF-8",
			"Pragma":        "no-cache",
			"Cache-Control": "no-cache, must-revalidate",
		},
		Body: string(rssFile),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
