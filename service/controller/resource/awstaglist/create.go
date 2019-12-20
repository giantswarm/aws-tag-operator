package awstaglist

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	al, err := r.toAWSTagList(obj)
	if err != nil {
		return microerror.Mask(err)
	}

	tags := []*ec2.Tag{}
	for _, t := range al.Spec.TagCollection {
		tag := ec2.Tag{
			Key:   aws.String(t.Key),
			Value: aws.String(t.Value),
		}
		tags = append(tags, &tag)
		fmt.Printf("Key: %s, Value: %s\n", t.Key, t.Value)
	}

	if len(tags) <= 0 {
		return nil
	}

	pvList, err := r.k8sClient.K8sClient().CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	fmt.Printf("Volumes in cluster:\n")
	volumes := []*string{}
	for _, pv := range pvList.Items {
		v := pv.Spec.AWSElasticBlockStore.VolumeID
		vc := v[len(v)-21:]
		fmt.Printf("vol: %s\n", vc)
		volumes = append(volumes, &vc)
	}

	input := &ec2.CreateTagsInput{
		Resources: volumes,
		Tags:      tags,
	}

	result, err := r.awsClients.EC2Client().CreateTags(input)
	if err != nil {
		return microerror.Mask(err)
	}

	fmt.Println(result)

	return nil
}
