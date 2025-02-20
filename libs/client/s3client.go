package client

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"path"
	"time"
)

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

func NewS3Client(ctx context.Context, awsConfig AWSConfig) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(awsConfig.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				awsConfig.AccessKeyID,
				awsConfig.SecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	return s3.NewFromConfig(cfg), nil
}

func GenerateObjectKey(objectPath, schemaName string) string {
	format := "2006_01_02_15_04_05"
	timestamp := time.Now().Format(format)
	fileName := fmt.Sprintf("%s_%s.json", schemaName, timestamp)
	return path.Join(objectPath, schemaName, fileName)
}
