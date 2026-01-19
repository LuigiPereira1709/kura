package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// PutObject uploads a new object to the specified S3 bucket.
func (s *Store) PutObject(bucket, key, contentType string, body io.Reader) error {
	req := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        body,
	}

	_, err := s.Client.PutObject(context.Background(), req)
	return err
}
