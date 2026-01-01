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

// FileUploader is the "Abstraction"
type FileUploader interface {
	Store(filename string, content string)
}

// BaseUploader is refined abstraction
type BaseUploader struct {
	provider StorageProvider // this is the bridge
}

func (b *BaseUploader) Store(filename string, content string) {
	data := []byte(content)
	b.provider.UploadFile(filename, data)
}

type SecureUploader struct {
	BaseUploader
}

func (s *SecureUploader) Store(filename string, content string) {
	fmt.Println("Encrypting file before upload...")
	s.BaseUploader.Store(filename, content)
}

func main() {
	s3 := &S3Provider{Bucket: "my-app-assets"}
	gcs := &GCSProvider{ProjectID: "my-gcp-project-123"}

	uploader := &BaseUploader{provider: s3}
	uploader.Store("hello.txt", "hello, world! 🚀")

	secure := &SecureUploader{BaseUploader: BaseUploader{provider: gcs}}
	secure.Store("secret.txt", "super super secret 🥴")
}
