package s3Service

import (
	awsConfig "adtec/backend/src/utils/aws"
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func Upload(
	key string,
	body io.ReadSeeker,
	acl types.ObjectCannedACL,
) (*s3.PutObjectOutput, error) {

	ctx := context.TODO()
	bucket := os.Getenv("AWS_BUCKET")

	s3Client, err := awsConfig.GetS3Client()
	if err != nil {
		return nil, err
	}

	object := s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
		// Metadata: map[string]*string{
		// 	"x-amz-meta-my-key": aws.String(key),
		// },
	}

	if acl != "" {
		object.ACL = acl
	}

	data, err := s3Client.PutObject(ctx, &object)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	return data, nil
}
