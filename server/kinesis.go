package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var streamName string
var kinesisSession *session.Session
var kinesisClient *kinesis.Kinesis

func _getSession() *session.Session {
	if kinesisSession == nil {
		sessionRegion, ok := os.LookupEnv("AWS_REGION")
		if !ok {
			sessionRegion = "us-east-1"
		}
		kinesisSession = session.New(&aws.Config{Region: aws.String(sessionRegion)})
	}
	return kinesisSession
}

func _getClient() *kinesis.Kinesis {
	if kinesisClient == nil {
		kinesisClient = kinesis.New(_getSession())
	}
	return kinesisClient
}

func putRecord(body []byte, partitionKey string) {
	c := _getClient()

	if streamName == "" {
		var ok bool
		streamName, ok = os.LookupEnv("STREAM")
		if !ok {
			streamName = "TestStream"
		}
	}

	_, err := c.PutRecord(&kinesis.PutRecordInput{
		Data:         body,
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(partitionKey),
	})
	if err != nil {
		log.Println("Error putting to Kinesis")
	} else {
		log.Println("Put to kinesis successful.")
	}
}
