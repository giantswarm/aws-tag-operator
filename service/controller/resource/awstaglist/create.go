package awstaglist

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	fmt.Println("YOLO")
	al, err := r.ToAWSTagList(obj)
	if err != nil {
		return microerror.Mask(err)
	}

	for cn := range al.Spec.ClusterIDCollection {

		err = r.addAWSClientsToContext(ctx, "")
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
