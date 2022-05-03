package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"
)

type IDisk interface {
	List(VolumeIds []string) []types.Volume
	Create()
	Delete()
	Update(diskId string, diskSize int, diskType string, performanceLevel int) (*ec2.ModifyVolumeOutput, error)
}

type Disk struct {
	client *ec2.Client
	ctx    context.Context
}

func (d *Disk) List(volumeIds []string, instanceId *string) []types.Volume {
	disks := []types.Volume{}
	input := &ec2.DescribeVolumesInput{}
	if instanceId != nil {
		input.Filters = append(input.Filters, types.Filter{
			Name:   aws.String("attachment.instance-id"),
			Values: []string{*instanceId}},
		)
	} else {
		if len(volumeIds) > 0 {
			input.VolumeIds = volumeIds
		}
	}

	for {
		result, err := d.client.DescribeVolumes(d.ctx, input)
		if err != nil {
			fmt.Println(err)
			break
		}
		disks = append(disks, result.Volumes...)
		if result.NextToken != nil {
			input.NextToken = result.NextToken
		} else {
			break
		}
	}
	return disks
}

func (d *Disk) Update(diskId string, diskSize int, diskType string, performanceLevel int) (*ec2.ModifyVolumeOutput, error) {
	input := &ec2.ModifyVolumeInput{
		VolumeId: &diskId,
		DryRun:   aws.Bool(true),
	}
	if diskSize != 0 {
		input.Size = aws.Int32(int32(diskSize))
	}
	if diskType != "" {
		input.VolumeType = types.VolumeType(diskType)
	}
	if performanceLevel != 0 {
		input.Iops = aws.Int32(int32(performanceLevel))
	}
	output, err := d.client.ModifyVolume(d.ctx, input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		return d.client.ModifyVolume(d.ctx, input)
	}
	return output, err
}
