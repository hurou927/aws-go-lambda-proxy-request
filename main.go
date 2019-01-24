package main

import (
	"bytes"
	"fmt"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"./lib"
	// "github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}


type Obj struct {
    Key1 string `json:"key1"`
	Key2 int    `json:"key2"`
	Kye3 bool	`json:"key3"`
	Key4 string	`json:"key4"`
	Key5 int	`json:"key5"`
} 

func main() {
	

	body := map[string] interface{} {
		"key1" : "v1",
		"key2" : 1,
		"key3" : true,
	}

	queryParams := map[string] string {
		"query": "q1",
	}
	pathParams := map[string] string {
		"path": "p1",
	}

	req, err := request.CreateProxyRequest(body,queryParams, pathParams)
	if err != nil  {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n",req)

	var ctx context.Context
	res, err := Handler(ctx, req)
	if err != nil  {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", res)
}
