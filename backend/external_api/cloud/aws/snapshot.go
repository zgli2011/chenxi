package aws

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/smithy-go"
)

type ISnapshot interface {
	List(snapshotIds []string, volumeId *string) []types.Snapshot
	Create(volumeId string, description string) (*ec2.CreateSnapshotOutput, error)
	Delete(snapshotId string) (*ec2.DeleteSnapshotOutput, error)
}

type Snapshot struct {
	client *ec2.Client
	ctx    context.Context
}

func (s *Snapshot) List(snapshotIds []string, volumeId *string) []types.Snapshot {
	input := &ec2.DescribeSnapshotsInput{}
	if len(snapshotIds) > 0 {
		input.SnapshotIds = snapshotIds
	}
	if volumeId != nil {
		input.Filters = []types.Filter{{Name: aws.String("volume-id"), Values: []string{*volumeId}}}
	}
	output, err := s.client.DescribeSnapshots(s.ctx, input)
	if err != nil {
		fmt.Println(err)
	}
	return output.Snapshots
}

func (s *Snapshot) Create(volumeId string, description string) (*ec2.CreateSnapshotOutput, error) {
	input := &ec2.CreateSnapshotInput{
		VolumeId:    &volumeId,
		Description: aws.String(description + time.Now().Format("200601021504")),
		DryRun:      aws.Bool(false),
	}
	output, err := s.client.CreateSnapshot(s.ctx, input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		return s.client.CreateSnapshot(s.ctx, input)
	}
	return output, err
}

func (s *Snapshot) Delete(snapshotId string) (*ec2.DeleteSnapshotOutput, error) {
	input := &ec2.DeleteSnapshotInput{
		SnapshotId: &snapshotId,
		DryRun:     aws.Bool(false),
	}
	output, err := s.client.DeleteSnapshot(s.ctx, input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		return s.client.DeleteSnapshot(s.ctx, input)
	}
	return output, err
}
