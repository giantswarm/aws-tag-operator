package awstaglist

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/giantswarm/microerror"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	fmt.Println("YOLO")
	al, err := r.ToAWSTagList(obj)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, tag := range al.Spec.TagCollection {
		fmt.Println(tag)
	}

	i := &ec2.DescribeVolumesInput{}
	o, err := r.AWSClients.EC2Client().DescribeVolumes(i)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, v := range o.Volumes {
		fmt.Println(v.GoString())
	}

	return nil
}
