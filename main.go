package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

var (
	err          error
	isJSON       bool
	jsonstr      []byte
	responseBody string
	statusCode   int
)

// The API Gateway handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	spew.Config.Indent = "    "
	statusCode = int(200)
	cacheFrom := time.Now().Format(http.TimeFormat)

	if isJSON, err = strconv.ParseBool(os.Getenv("JSON")); err != nil {
		log.Fatal("Lambda app error: ", err)
	}

	if isJSON {
		jsonstr, err = json.MarshalIndent(request, "", "    ")
		if err != nil {
			log.Fatal("Lambda app error: ", err)
		}

		responseBody = string(jsonstr)
	} else {
		responseBody = spew.Sdump(request)
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":  "text/plain; charset=utf-8",
			"Last-Modified": cacheFrom,
			"Expires":       cacheFrom,
		},
		Body:       responseBody,
		StatusCode: statusCode,
	}, nil
}

// The core function
func main() {
	lambda.Start(Handler)
}
