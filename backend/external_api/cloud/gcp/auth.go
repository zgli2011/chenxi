package google

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GCP struct {
	Region    string
	Key       string
	Secret    string
	ECSClient *ecs.Client
}

func (gcp *GCP) auth() error {
	config := &oauth2.Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURL:  "",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	// Dummy authorization flow to read auth code from stdin.
	authURL := config.AuthCodeURL("your state")
	// client, err := ecs.NewClientWithAccessKey(aws.Region, aws.Key, aws.Secret)
	// aws.ECSClient = client
	// return err
}
