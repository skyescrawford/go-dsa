package main

func main() {
	s3 := &S3Provider{Bucket: "my-app-assets"}
	gcs := &GCSProvider{ProjectID: "my-gcp-project-123"}

	uploader := &BaseUploader{provider: s3}
	uploader.Store("hello.txt", "hello, world! 🚀")

	secure := &SecureUploader{base: &BaseUploader{provider: gcs}}
	secure.Store("secret.txt", "super super secret 🥴")
}
