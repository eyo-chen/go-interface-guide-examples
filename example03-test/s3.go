package main

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// S3Manager is an interface that defines the methods for managing S3 objects.
type S3Manager interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// S3Uploader is a struct that holds a S3Manager and provides a method to upload data to S3.
type S3Uploader struct {
	client S3Manager
}

// NewS3Uploader is a function that creates a new S3Uploader with the given S3Manager.
func NewS3Uploader(client S3Manager) *S3Uploader {
	return &S3Uploader{client: client}
}

// Upload is a method that uploads data to S3.
func (s *S3Uploader) Upload(ctx context.Context, bucketName, objectKey string, data []byte) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(data),
	})

	return err
}
