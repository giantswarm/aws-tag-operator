package controller

import (
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/aws-tag-operator/client/aws"
	"github.com/giantswarm/aws-tag-operator/pkg/project"
)

type AWSTagListConfig struct {
	AWSConfig    aws.Config
	K8sAWSClient kubernetes.Interface
	K8sClient    k8sclient.Interface
	Logger       micrologger.Logger
}

type AWSTagList struct {
	*controller.Controller
}

func NewAWSTagList(config AWSTagListConfig) (*AWSTagList, error) {
	var err error

	resourceSets, err := newAWSTagListResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			CRD:          v1alpha1.NewAWSTagListCRD(),
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(v1alpha1.AWSTagList)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/aws-tag-operator.
			Name: project.Name(),
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &AWSTagList{
		Controller: operatorkitController,
	}

	return c, nil
}

func newAWSTagListResourceSets(config AWSTagListConfig) ([]*controller.ResourceSet, error) {
	var err error

	var resourceSet *controller.ResourceSet
	{
		c := awsTagListResourceSetConfig{
			AWSConfig:    config.AWSConfig,
			K8sAWSClient: config.K8sAWSClient,
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
		}

		resourceSet, err = newAWSTagListResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resourceSets := []*controller.ResourceSet{
		resourceSet,
	}

	return resourceSets, nil
}
