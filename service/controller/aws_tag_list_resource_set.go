package controller

import (
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"github.com/giantswarm/operatorkit/resource"
	"github.com/giantswarm/operatorkit/resource/wrapper/metricsresource"
	"github.com/giantswarm/operatorkit/resource/wrapper/retryresource"

	"github.com/giantswarm/aws-tag-operator/client/aws"
	"github.com/giantswarm/aws-tag-operator/service/controller/resource/awstaglist"
)

type awsTagListResourceSetConfig struct {
	AWSClients aws.Interface
	K8sClient  k8sclient.Interface
	Logger     micrologger.Logger
}

func newAWSTagListResourceSet(config awsTagListResourceSetConfig) (*controller.ResourceSet, error) {
	var err error

	var awsTagListResource resource.Interface
	{
		c := awstaglist.Config{
			AWSClients: config.AWSClients,
			K8sClient:  config.K8sClient,
			Logger:     config.Logger,
		}

		awsTagListResource, err = awstaglist.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resources := []resource.Interface{
		awsTagListResource,
	}

	{
		c := retryresource.WrapConfig{
			Logger: config.Logger,
		}

		resources, err = retryresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	{
		c := metricsresource.WrapConfig{}

		resources, err = metricsresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	handlesFunc := func(obj interface{}) bool {
		//For now reconcile all CRs
		return true
	}

	var resourceSet *controller.ResourceSet
	{
		c := controller.ResourceSetConfig{
			Handles:   handlesFunc,
			Logger:    config.Logger,
			Resources: resources,
		}

		resourceSet, err = controller.NewResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return resourceSet, nil
}
