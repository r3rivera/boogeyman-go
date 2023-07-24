package s3

import (
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const AWS_REGION = "us-east-1"

var onceAwsConfig sync.Once

func getS3Client() *s3.S3 {
	var s3c *s3.S3
	onceAwsConfig.Do(func() {
		s, err := session.NewSession(&aws.Config{
			Region: aws.String(AWS_REGION)},
		)
		if err != nil {
			panic(err)
		}
		s3c = s3.New(s)
	})
	return s3c
}

func GeneratePresignedUrl(fileKey, bucket string) (string, error) {
	svc := getS3Client()
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileKey),
	})
	str, err := req.Presign(10 * time.Minute)
	if err != nil {
		return "", err
	}
	return str, nil
}

type S3FileInfo struct {
	filePath string
	fileName string
	bucket   string
}

func NewS3FileInfo(filePath, fileName, bucket string) *S3FileInfo {
	return &S3FileInfo{
		filePath: filePath,
		fileName: fileName,
		bucket:   bucket,
	}
}

func (f *S3FileInfo) UploadFile() error {
	svc := getS3Client()

	file, err := os.Open(f.filePath + f.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(f.bucket),
		Key:    aws.String(f.filePath + f.fileName),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}
