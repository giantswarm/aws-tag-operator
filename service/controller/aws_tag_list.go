package controller

import (
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/giantswarm/aws-tag-operator/pkg/project"
)

type AWSTagListConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
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
				return new(corev1.Pod)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/aws-tag-operator-todo-controller.
			Name: project.Name() + "-todo-controller",
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
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
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
