package aliyun

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type IInstance interface {
	List(instanceIds []string) []types.Reservation
	Create(instanceParam InstanceParam) (*ec2.RunInstancesOutput, error)
	Delete(instanceId string) error
	ShutDown(instanceIds []string) error
	Start(instanceIds []string) error
	Reboot(instanceIds []string) error
	Status() error
	ChangeIntanceType(instanceId string, instanceType string) error
}

func (i *Instance) List(instanceIds []string) []types.Reservation {
	instances := []types.Reservation{}

	input := ec2.DescribeInstancesInput{}
	if len(instanceIds) > 0 {
		input.InstanceIds = instanceIds
	}
	for {
		result, errs := i.client.DescribeInstances(i.ctx, &input)
		if errs != nil {
			fmt.Println(errs)
			break
		}
		instances = append(instances, result.Reservations...)
		if result.NextToken != nil {
			input.NextToken = result.NextToken
		} else {
			break
		}
	}
	return instances
}

type Instance struct {
	input  ec2.DescribeInstancesInput
	client *ec2.Client
	ctx    context.Context
}

type InstanceParam struct {
	ImageId              string
	SubnetId             string
	SecurityGroupIds     []string
	InstanceType         string
	Tags                 []map[string]string
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
