package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"go-storage-s3/configs"
	"go-storage-s3/core/ports"
	"golang.org/x/net/context"
	"time"
)

type FileRepository struct {
	bucket string
	svc    *s3.S3
}

func NewFileRepository(config *configs.Config, svc *s3.S3) ports.FileS3Repository {
	return &FileRepository{bucket: config.S3.Bucket, svc: svc}
}

func (u FileRepository) GetPresignedUrl(_ context.Context, path string) (string, error) {
	r, _ := u.svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(path),
	})
	return r.Presign(15 * time.Minute)
}
