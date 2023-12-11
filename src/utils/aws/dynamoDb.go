package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DynamoDb() (*dynamodb.Client, error) {

	region := os.Getenv("AWS_REGION")

	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKeyId,
				secretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return nil, err
	}

	// Create a DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return svc, nil
}
