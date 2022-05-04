package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func NewClient(accessKey string, secretKey string, region string, projectId string) (*ecs20140526.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     &accessKey,
		AccessKeySecret: &secretKey,
	}
	config.Endpoint = tea.String("ecs." + region + ".aliyuncs.com")
	client := &ecs20140526.Client{}
	client, err := ecs20140526.NewClient(config)
	return client, err
}
