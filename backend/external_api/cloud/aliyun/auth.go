package aliyun

import (
	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type Aliyun struct {
	Region    string
	Key       string
	Secret    string
	ECSClient *ecs.Client
}

func (aliyun *Aliyun) auth() error {
	client, err := ecs.NewClientWithAccessKey(aliyun.Region, aliyun.Key, aliyun.Secret)
	aliyun.ECSClient = client
	return err
}
