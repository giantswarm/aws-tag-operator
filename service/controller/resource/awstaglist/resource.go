package awstaglist

import (
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/aws-tag-operator/client/aws"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	Name = "aws_tag_list"
)

type Config struct {
	AWSClients aws.Interface
	K8sClient  k8sclient.Interface
	Logger     micrologger.Logger
}

type Resource struct {
	awsClients aws.Interface
	logger     micrologger.Logger
	k8sClient  k8sclient.Interface
}

func New(config Config) (*Resource, error) {
	if config.AWSClients == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.AWSClients must not be empty", config)
	}
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sclient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	r := &Resource{
		AWSClients: config.AWSClients,
		logger:     config.Logger,
		k8sClient:  config.K8sClient,
	}

	return r, nil
}

func (r *Resource) Name() string {
	return Name
}

func (r *Resource) toAWSTagList(v interface{}) (v1alpha1.AWSTagList, error) {
	if v == nil {
		return v1alpha1.AWSTagList{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AWSTagList{}, v)
	}

	p, ok := v.(*v1alpha1.AWSTagList)
	if !ok {
		return v1alpha1.AWSTagList{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AWSTagList{}, v)
	}

	c := p.DeepCopy()

	return *c, nil
}
