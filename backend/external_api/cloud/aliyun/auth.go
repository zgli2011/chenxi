package aliyun

import (
	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func NewClient(accessKey string, secretKey string, region string, projectId string) (*ecs.Client, error) {
	return ecs.NewClientWithAccessKey(region, accessKey, secretKey)
}
