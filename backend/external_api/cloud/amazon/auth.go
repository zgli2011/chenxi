package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type Amazon struct {
	Region string
	Key    string
	Secret string
}

func (a *Amazon) auth(accessKey string, secretKey string) error {
	appCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))

	_, err := appCreds.Retrieve(context.TODO())

	// creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	// sess := session.Must(session.NewSession(&aws.))
	return err
}
