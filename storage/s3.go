package storage

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	Bucket  string
	Client  *s3.S3
	BaseURL string
}

// NewS3Storage initializes a new S3-compatible storage instance
func NewS3Storage(bucket, region, baseURL, accessKey, secretKey string) *S3Storage {
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey, secretKey, "",
		),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	return &S3Storage{
		Bucket:  bucket,
		Client:  s3.New(sess),
		BaseURL: baseURL,
	}
}

// Upload uploads a file to the S3 bucket
func (s *S3Storage) Upload(fileName string, fileContent io.Reader, contentType string) (string, error) {
	// Read the file content into a buffer
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(fileContent)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %v", err)
	}

	// Upload the file to S3
	_, err = s.Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"), // Adjust permissions as needed
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Return the file path (relative path)
	return fmt.Sprintf("%s/%s", s.BaseURL, fileName), nil
}

// Get generates a public URL for the file
func (s *S3Storage) Get(fileName string) (string, error) {
	return fmt.Sprintf("%s/%s", s.BaseURL, fileName), nil
}

// Delete removes a file from the S3 bucket
func (s *S3Storage) Delete(fileName string) error {
	_, err := s.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}

	return nil
}
