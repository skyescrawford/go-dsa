package main

import "fmt"

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

// wrapper for base uploader with additional feature
type SecureUploader struct {
	base *BaseUploader
}

func (s *SecureUploader) Store(filename string, content string) {
	fmt.Println("Encrypting file before upload...") // additional step
	s.base.Store(filename, content)
}
