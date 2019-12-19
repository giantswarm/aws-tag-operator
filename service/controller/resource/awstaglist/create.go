package awstaglist

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	al, err := r.toAWSTagList(obj)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, tag := range al.Spec.TagCollection {
		fmt.Printf("Key: %s, Value: %s", tag.Key, tag.Value)
	}

	pvList, err := r.k8sClient.K8sClient().CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	for _, pv := range pvList.Items {
		fmt.Println(pv.Spec.AWSElasticBlockStore.VolumeID)
	}
	// i := &ec2.DescribeVolumesInput{}
	// o, err := r.awsClients.EC2Client().DescribeVolumes(i)
	// if err != nil {
	// 	return microerror.Mask(err)
	// }

	// for _, v := range o.Volumes {
	// 	fmt.Println(v.GoString())
	// }

	return nil
}
