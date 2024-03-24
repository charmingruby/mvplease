package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

func NewStorageService(logger *logrus.Logger) (*CloudflareService, error) {
	creds := loadCredentials()

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", creds.CloudflareAccountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(creds.AccessKeyID, creds.AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &CloudflareService{
		client:      client,
		credentials: creds,
		logger:      logger,
	}, nil

}

type CloudflareService struct {
	client      *s3.Client
	credentials *cloudflareCredentials
	logger      *logrus.Logger
}

type cloudflareCredentials struct {
	BucketName          string
	CloudflareAccountID string
	AccessKeyID         string
	AccessKeySecret     string
}

func loadCredentials() *cloudflareCredentials {
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return &cloudflareCredentials{
		BucketName:          bucketName,
		CloudflareAccountID: accountId,
		AccessKeyID:         accessKeyId,
		AccessKeySecret:     accessKeySecret,
	}
}

func (c *CloudflareService) Upload(file io.Reader, key string) error {
	c.logger.Info("Uploading a file to Cloudflare...")

	_, err := c.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(c.credentials.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return err
	}

	c.logger.Info("File uploaded successfully to Cloudflare.")

	return nil
}

func (c *CloudflareService) DeleteByKey(key string) error {
	c.logger.Info("Removing a file from Cloudflare...")

	_, err := c.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(c.credentials.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error removing Cloudflare file: %v", err))

		return err
	}

	c.logger.Info("Removed successfully the file from Cloudflare.")

	return nil
}
