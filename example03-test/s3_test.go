package main

import (
	"context"
	"testing"
)

func TestUpload(t *testing.T) {
	mockUploader := NewS3Uploader(&MockS3Manager{})

	err := mockUploader.Upload(context.Background(), "test-bucket", "test-object", []byte("test-data"))
	if err != nil {
		t.Errorf("Upload failed: %v", err)
	}
}
