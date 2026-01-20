package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *Store) ListObjectsWithoutPrefix(bucket, notPrefix string) ([]string, error) {
	req := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(notPrefix),
	}

	resp, err := s.Client.ListObjectsV2(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects in S3 bucket %s with prefix %s: %w", bucket, notPrefix, err)
	}

	var keys []string
	for _, obj := range resp.Contents {
		if *obj.Key != notPrefix {
			keys = append(keys, *obj.Key)
		}
	}

	return keys, nil
}

// ListObjectsForPrefix lists all objects in the specified S3 bucket that match the given prefix.
func (s *Store) ListObjectsForPrefix(bucket, prefix string) ([]string, error) {
	req := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}

	resp, err := s.Client.ListObjectsV2(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects in S3 bucket %s with prefix %s: %w", bucket, prefix, err)
	}

	var keys []string
	for _, obj := range resp.Contents {
		if obj.Key != nil {
			keys = append(keys, *obj.Key)
		}
	}

	return keys, nil
}
