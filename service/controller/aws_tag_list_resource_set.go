package controller

import (
	"github.com/giantswarm/aws-operator/service/controller/resource/awsclient"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"github.com/giantswarm/operatorkit/resource"
	"github.com/giantswarm/operatorkit/resource/wrapper/metricsresource"
	"github.com/giantswarm/operatorkit/resource/wrapper/retryresource"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/aws-tag-operator/client/aws"
	"github.com/giantswarm/aws-tag-operator/service/controller/resource/awstaglist"
	"github.com/giantswarm/aws-tag-operator/service/key"
)

type awsTagListResourceSetConfig struct {
	AWSConfig    aws.Config
	K8sClient    k8sclient.Interface
	K8sAWSClient kubernetes.Interface
	Logger       micrologger.Logger
}

func newAWSTagListResourceSet(config awsTagListResourceSetConfig) (*controller.ResourceSet, error) {
	var err error

	var awsClientResource resource.Interface
	{
		c := awsclient.Config{
			K8sClient:     config.K8sAWSClient,
			Logger:        config.Logger,
			ToClusterFunc: key.ToCluster,

			AWSConfig: config.AWSConfig,
		}

		awsClientResource, err = awsclient.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var awsTagListResource resource.Interface
	{
		c := awstaglist.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		awsTagListResource, err = awstaglist.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resources := []resource.Interface{
		awsClientResource,
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

		return false
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
