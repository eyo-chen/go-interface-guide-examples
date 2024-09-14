package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// MockS3Manager is a mock implementation of the S3Manager interface
type MockS3Manager struct{}

// PutObject is a method that mocks the PutObject method of the S3Manager interface
func (m *MockS3Manager) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	// mocking logic
	return &s3.PutObjectOutput{}, nil
}
