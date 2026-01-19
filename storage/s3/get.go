package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)


func (s *Store) GetObjByKey(bucket, key string) (io.ReadCloser, error) {
	req := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	resp, err := s.Client.GetObject(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get object from S3 bucket %s with key %s: %w", bucket, key, err)
	}

	return resp.Body, nil
}

