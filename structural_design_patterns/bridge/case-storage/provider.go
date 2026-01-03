package main

import "fmt"

// implementation interface
type StorageProvider interface {
	UploadFile(name string, data []byte) error
}

// concrete implementation
type S3Provider struct {
	Bucket string
}

func (s *S3Provider) UploadFile(name string, data []byte) error {
	fmt.Printf("Uploading %s to AWS S3 Bucket: %s\n", name, s.Bucket)
	return nil
}

type GCSProvider struct {
	ProjectID string
}

func (g *GCSProvider) UploadFile(name string, data []byte) error {
	fmt.Printf("Uploading %s to Google Cloud Project: %s\n", name, g.ProjectID)
	return nil
}
