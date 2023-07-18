package s3

import (
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
