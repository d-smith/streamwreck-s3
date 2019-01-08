package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bradfitz/gomemcache/memcache"
)

var memcachedEndpoint = fmt.Sprintf("%s:%s", os.Getenv("MEMCACHED_ENDPOINT"), os.Getenv("MEMCACHED_PORT"))

var client = memcache.New(memcachedEndpoint)

//TODO - error handling at the lambda level - return error and retry? Log and continue?

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	fmt.Println(memcachedEndpoint)
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

	}
}

func main() {
	lambda.Start(handler)
}
