package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Store provides methods to interact with an S3-compatible storage service.
type Store struct {
	Client *s3.Client
}

// NewService creates a new S3Service instance with the provided S3 client.
func New(cfg aws.Config) *Store {
	return &Store{
		Client: s3.NewFromConfig(cfg),
	}
}
