package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DynamoDb() (*dynamodb.Client, error) {

	cfg, err := AwsConfig()
	if err != nil {
		log.Println("unable to load SDK config, %v", err)
		return nil, err
	}

	// Create a DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return svc, nil
}
