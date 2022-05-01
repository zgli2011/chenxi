package gcp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewClient(accessKey string, secretKey string, region string, projectId string) (aws.Credentials, error) {
	appCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))
	return appCreds.Retrieve(context.TODO())
}
