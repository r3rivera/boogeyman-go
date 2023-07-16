package dba

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var awsConfig aws.Config
var onceAwsConfig sync.Once
var dynamodbClient *dynamodb.Client
var onceDdbClient sync.Once

const AWS_REGION = "us-east-1"

func getAwsConfig() aws.Config {

	onceAwsConfig.Do(func() {
		var err error
		awsConfig, err = config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic(err)
		}
	})
	return awsConfig
}

func getDynamodbClient() *dynamodb.Client {
	onceDdbClient.Do(func() {
		awsConfig = getAwsConfig()
		dynamodbClient = dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
			opt.Region = AWS_REGION
		})
	})
	return dynamodbClient
}
