package repositories

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go-storage-s3/configs"
	"go-storage-s3/core/ports"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

type FileRepository struct {
	bucket           string
	svc              *s3.S3
	presignedExpired time.Duration
}

func NewFileRepository(config *configs.Config, svc *s3.S3) ports.FileS3Repository {
	if !checkBucketExisted(svc, config.S3.Bucket) {
		log.Fatalf("Bucket in '%s' s3 not available\n", config.S3.Bucket)
	}
	return &FileRepository{
		bucket:           config.S3.Bucket,
		svc:              svc,
		presignedExpired: 15 * time.Minute,
	}
}

func checkBucketExisted(svc *s3.S3, bucketCheck string) bool {
	buckets, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatalf("Cannot get list bucket **** %v", err)
	}
	for _, bucket := range buckets.Buckets {
		if bucketCheck == *bucket.Name {
			return true
		}
	}
	return false
}

func (u FileRepository) GetPresignedUrl(_ context.Context, path string) (string, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(path),
		//ContentLength: aws.Int64(1234),
	}
	request, _ := u.svc.PutObjectRequest(input)
	return request.Presign(u.presignedExpired)
}

func (u *FileRepository) Upload() {
	cf := configs.Get()
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(cf.S3.AccessKey, cf.S3.SecretKey, cf.S3.Token),
		Endpoint:    aws.String(cf.S3.Endpoint),
		Region:      aws.String(cf.S3.Region)},
	)
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open("upload/a.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(u.bucket + "/a.pdf"),
		Body:   f,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}
