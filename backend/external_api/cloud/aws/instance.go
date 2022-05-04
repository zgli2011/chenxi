package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/smithy-go"
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

type Instance struct {
	client *ec2.Client
	ctx    context.Context
}

func (i *Instance) List(instanceIds []string) []types.Reservation {
	instances := []types.Reservation{}

	input := ec2.DescribeInstancesInput{}
	if len(instanceIds) > 0 {
		input.InstanceIds = instanceIds
	} else {
		input.MaxResults = aws.Int32(1000)
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

func (i *Instance) Create(instanceParam InstanceParam) (*ec2.RunInstancesOutput, error) {
	input := &ec2.RunInstancesInput{
		ImageId:               aws.String(instanceParam.ImageId),
		SubnetId:              aws.String(instanceParam.SubnetId),
		SecurityGroupIds:      instanceParam.SecurityGroupIds,
		MinCount:              aws.Int32(1),
		MaxCount:              aws.Int32(1),
		InstanceType:          types.InstanceType(instanceParam.InstanceType),
		Monitoring:            &types.RunInstancesMonitoringEnabled{Enabled: aws.Bool(false)},
		DisableApiTermination: aws.Bool(false),
		DryRun:                aws.Bool(true),
	}
	// 标签
	if len(instanceParam.Tags) > 0 {
		tag := []types.Tag{}
		for key, value := range instanceParam.Tags {
			tag = append(tag, types.Tag{Key: aws.String(key), Value: aws.String(value)})
		}
		input.TagSpecifications = []types.TagSpecification{
			{ResourceType: types.ResourceTypeInstance, Tags: tag},
			{ResourceType: types.ResourceTypeVolume, Tags: tag},
		}
	}
	// 磁盘
	if len(instanceParam.VolumeDeviceMappings) > 0 {
		blockDeviceMappings := []types.BlockDeviceMapping{}
		for _, volume := range instanceParam.VolumeDeviceMappings {
			esb := types.EbsBlockDevice{
				DeleteOnTermination: aws.Bool(true),
				Encrypted:           aws.Bool(false),
				VolumeSize:          aws.Int32(volume.VolumeSize),
				VolumeType:          types.VolumeType(volume.VolumeType),
			}
			if volume.SnapshotId != "" {
				esb.SnapshotId = aws.String(volume.SnapshotId)
			}
			if volume.VolumeType == "io1" || volume.VolumeType == "io2" {
				esb.Iops = &volume.PerformanceLevel
			}

			blockDeviceMapping := types.BlockDeviceMapping{
				DeviceName: aws.String(volume.Device),
				Ebs:        &esb,
			}
			blockDeviceMappings = append(blockDeviceMappings, blockDeviceMapping)
		}
		input.BlockDeviceMappings = blockDeviceMappings
	}
	// 手动指定ip地址
	if instanceParam.PrivateIP != "" {
		input.PrivateIpAddress = &instanceParam.PrivateIP
	}
	// 指定iam
	if instanceParam.Iam != "" {
		input.IamInstanceProfile.Arn = &instanceParam.Iam
	}
	// 用于初始化脚本执行
	if instanceParam.UserData != "" {
		input.UserData = &instanceParam.UserData
	}

	output, err := i.client.RunInstances(i.ctx, input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		return i.client.RunInstances(i.ctx, input)
	}
	return output, err
}

func (i *Instance) Delete(instanceId string) error {
	modify_input := &ec2.ModifyInstanceAttributeInput{
		InstanceId:            &instanceId,
		DisableApiTermination: &types.AttributeBooleanValue{Value: aws.Bool(false)},
	}
	_, err := i.client.ModifyInstanceAttribute(i.ctx, modify_input)
	if err != nil {
		return err
	}

	terminate_input := &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceId},
		DryRun:      aws.Bool(true),
	}
	_, err = i.client.TerminateInstances(i.ctx, terminate_input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		terminate_input.DryRun = aws.Bool(false)
		_, err = i.client.TerminateInstances(i.ctx, terminate_input)
		return err
	}
	return err
}

func (i *Instance) ShutDown(instanceIds []string) error {
	shutdown_input := &ec2.StopInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(true),
		Force:       aws.Bool(true),
	}
	_, err := i.client.StopInstances(i.ctx, shutdown_input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		shutdown_input.DryRun = aws.Bool(false)
		_, err = i.client.StopInstances(i.ctx, shutdown_input)
		return err
	}
	return err
}

func (i *Instance) Start(instanceIds []string) error {
	start_input := &ec2.StartInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(false),
	}
	_, err := i.client.StartInstances(i.ctx, start_input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		start_input.DryRun = aws.Bool(false)
		_, err = i.client.StartInstances(i.ctx, start_input)
		return err
	}
	return err
}

func (i *Instance) Reboot(instanceIds []string) error {
	reboot_input := &ec2.RebootInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(false),
	}
	_, err := i.client.RebootInstances(i.ctx, reboot_input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		reboot_input.DryRun = aws.Bool(false)
		_, err = i.client.RebootInstances(i.ctx, reboot_input)
		return err
	}
	return err
}

func (i *Instance) Status(instanceIds []string) (*ec2.DescribeInstanceStatusOutput, error) {
	input := &ec2.DescribeInstanceStatusInput{
		InstanceIds: instanceIds,
	}
	return i.client.DescribeInstanceStatus(i.ctx, input)
}

func (i *Instance) ChangeIntanceType(instanceId string, instanceType string) error {
	input := &ec2.ModifyInstanceAttributeInput{
		InstanceId:   &instanceId,
		InstanceType: &types.AttributeValue{Value: aws.String(instanceType)},
		DryRun:       aws.Bool(false),
	}
	_, err := i.client.ModifyInstanceAttribute(i.ctx, input)
	return err
}
