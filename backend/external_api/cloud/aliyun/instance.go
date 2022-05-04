package aliyun

import (
	"fmt"
	"strings"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type IInstance interface {
	List(instanceIds []string) []*ecs20140526.DescribeInstancesResponseBodyInstancesInstance
	Create(instanceParam InstanceParam)
	Delete(instanceId string) error
	ShutDown(instanceIds []string) error
	Start(instanceIds []string) error
	Reboot(instanceIds []string) error
	Status() error
	ChangeIntanceType(instanceId string, instanceType string) error
}

func (i *Instance) List(instanceIds []string) []*ecs20140526.DescribeInstancesResponseBodyInstancesInstance {
	instances := []*ecs20140526.DescribeInstancesResponseBodyInstancesInstance{}
	request := &ecs20140526.DescribeInstancesRequest{}
	if len(instanceIds) > 0 {
		request.InstanceIds = tea.String(strings.Join(instanceIds, ","))
	}
	request.MaxResults = tea.Int32(100)
	for {
		response, err := i.client.DescribeInstances(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		instances = append(instances, response.Body.Instances.Instance...)
		if response.Body.NextToken != nil {
			request.NextToken = response.Body.NextToken
		} else {
			break
		}
	}
	return instances
}

type Instance struct {
	client *ecs20140526.Client
}

type InstanceParam struct {
	ImageId              string
	SubnetId             string
	SecurityGroupIds     []string
	InstanceType         string
	Tags                 map[string]string
	PrivateIP            string
	Iam                  string
	UserData             string
	VolumeDeviceMappings []struct {
		Device           string
		VolumeSize       int32
		VolumeType       string
		SnapshotId       string
		PerformanceLevel int32
	}
}

func (i *Instance) Create(instanceParam InstanceParam) {
	request := &ecs20140526.RunInstancesRequest{
		ImageId:          &instanceParam.ImageId,
		InstanceType:     &instanceParam.InstanceType,
		SecurityGroupIds: instanceParam.SecurityGroupIds,
		VSwitchId:        &instanceParam.SubnetId,
	}
	hostname := ""
	for key, value := range instanceParam.Tags {
		if key == "Name" {
			hostname = value
			break
		}
	}
	request.InstanceName = &hostname
	request.HostName = &hostname
	request.Amount = tea.Int32(1)
	request.PasswordInherit = tea.Bool(true)
	request.InstanceChargeType = tea.String("PrePaid")
	request.PeriodUnit = tea.String("Month")
	request.Period = tea.Int32(1)
	request.Amount = tea.Int32(1)
}
