package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go-storage-s3/common/log"
	"go-storage-s3/configs"
)

func Connect(cf *configs.Config) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(cf.S3.AccessKey, cf.S3.SecretKey, cf.S3.Token),
		Endpoint:    aws.String(cf.S3.Endpoint),
		Region:      aws.String(cf.S3.Region)},
	)
	if err != nil {
		log.Fatalf(err, "init session s3 fail")
	}

	return s3.New(sess)
}
