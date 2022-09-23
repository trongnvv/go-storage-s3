package repositories

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"go-storage-s3/configs"
	"go-storage-s3/core/ports"
	"golang.org/x/net/context"
	"log"
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
	u.test()
	input := &s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(path),
		//ContentLength: aws.Int64(1234),
	}
	request, _ := u.svc.PutObjectRequest(input)
	return request.Presign(u.presignedExpired)
}

func (u *FileRepository) test() {
	resp, err := u.svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(u.bucket)})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", u.bucket)
	fmt.Println("")

}
