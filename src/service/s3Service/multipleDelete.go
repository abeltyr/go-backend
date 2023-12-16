package s3Service

import (
	awsConfig "adtec/backend/src/utils/aws"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func MultipleDelete(keys []types.ObjectIdentifier) error {

	bucket := os.Getenv("AWS_BUCKET")

	ctx := context.TODO()

	s3Client, err := awsConfig.GetS3Client()
	if err != nil {
		return err
	}

	quiteData := false

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &types.Delete{
			Objects: keys,
			Quiet:   &quiteData,
		},
	}

	// Perform the DeleteObjects operation
	_, err = s3Client.DeleteObjects(ctx, input)
	if err != nil {
		log.Fatalf("Failed to delete objects, %v", err)
	}

	return nil
}
