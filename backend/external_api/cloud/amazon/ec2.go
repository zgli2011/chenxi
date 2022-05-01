package amazon

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var PageSize int32 = 1000

type IInstance interface {
	List()
	Create()
	Update()
	Delete()
	ShutDown()
	Start()
	Reboot()
	Status()
	ChangeIntanceType()
}

type Instance struct {
	input  ec2.DescribeInstancesInput
	client *ec2.Client
}

func (i *Instance) List(instanceIds []string) []types.Reservation {
	instances := []types.Reservation{}

	input := ec2.DescribeInstancesInput{}
	if len(instanceIds) > 0 {
		input.InstanceIds = instanceIds
	}
	for {
		result, errs := i.client.DescribeInstances(context.TODO(), &input)
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
	ImageId string
}

func (i *Instance) Create(instanceIds []string) {
	// input := &ec2.RunInstancesInput{
	// 	ImageId:      aws.String("ami-e7527ed7"),
	// 	InstanceType: types.InstanceTypeT2Micro,
	// 	MinCount:     1,
	// 	MaxCount:     1,
	// }
}

func (i *Instance) Update(instanceIds []string) {
}

func (i *Instance) Delete(instanceIds []string) {
}

func (i *Instance) ShutDown(instanceIds []string) {
}

func (i *Instance) Start(instanceIds []string) {
}

func (i *Instance) Reboot(instanceIds []string) {
}

func (i *Instance) ChangeIntanceType(instanceId string, instanceType string) {
}
