package s3Service

import (
	"context"
	"io"
	"os"

	awsConfig "adtec/backend/src/utils/aws"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetFile(
	key string,
) ([]byte, error) {

	ctx := context.TODO()
	bucket := os.Getenv("BUCKET")

	s3Client, err := awsConfig.GetS3Client()
	if err != nil {
		return nil, err
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := s3Client.GetObject(ctx, input)
	if err != nil {
		return nil, err
	}

	fileByte, err := io.ReadAll(result.Body)

	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return fileByte, nil
}
