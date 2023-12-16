package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Client() (*s3.Client, error) {

	cfg, err := AwsConfig()
	if err != nil {
		log.Println("unable to load SDK config, %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return client, nil

}
