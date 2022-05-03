package aws

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"
)

type IImage interface {
	List(imageIds []string) []types.Image
	Create(instanceId string, description string) (*ec2.CreateImageOutput, error)
	Delete(imageId string) (*ec2.DeregisterImageOutput, error)
}

type Image struct {
	client *ec2.Client
	ctx    context.Context
}

func (i *Image) List(imageIds []string) []types.Image {
	images := []types.Image{}
	input := &ec2.DescribeImagesInput{
		// Owners:  []string{"self"},
		Filters: []types.Filter{{Name: aws.String("is-public"), Values: []string{"false"}}},
	}
	if len(imageIds) > 0 {
		input.ImageIds = imageIds
	}
	result, err := i.client.DescribeImages(i.ctx, input)
	if err != nil {
		fmt.Println(err)
	}
	images = append(images, result.Images...)
	return images
}

func (i *Image) Create(instanceId string, description string) (*ec2.CreateImageOutput, error) {
	input := &ec2.CreateImageInput{
		InstanceId: &instanceId,
		NoReboot:   aws.Bool(true),
		Name:       aws.String(description + time.Now().Format("200601021504")),
		DryRun:     aws.Bool(false),
	}
	output, err := i.client.CreateImage(i.ctx, input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		return i.client.CreateImage(i.ctx, input)
	}
	return output, err
}

func (i *Image) Delete(imageId string) (*ec2.DeregisterImageOutput, error) {
	input := &ec2.DeregisterImageInput{
		ImageId: &imageId,
		DryRun:  aws.Bool(false),
	}
	output, err := i.client.DeregisterImage(i.ctx, input)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		return i.client.DeregisterImage(i.ctx, input)
	}
	return output, err
}
