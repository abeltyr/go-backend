package s3Service

import (
	awsConfig "adtec/backend/src/utils/aws"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Delete(key string) error {

	bucket := os.Getenv("AWS_BUCKET")

	ctx := context.TODO()

	s3Client, err := awsConfig.GetS3Client()
	if err != nil {
		return err
	}

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	// Perform the DeleteObjects operation
	_, err = s3Client.DeleteObject(ctx, input)
	if err != nil {
		log.Fatalf("Failed to delete objects, %v", err)
	}

	return nil
}
