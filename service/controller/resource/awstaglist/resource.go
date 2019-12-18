package awstaglist

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/aws-tag-operator/client/aws"
	"github.com/giantswarm/aws-tag-operator/service/controller/controllercontext"
	"github.com/giantswarm/aws-tag-operator/service/internal/credential"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	Name = "aws_tag_list"
)

type Config struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
	AWSConfig aws.Config
}

type Resource struct {
	logger    micrologger.Logger
	k8sClient k8sclient.Interface

	AWSConfig aws.Config
}

func New(config Config) (*Resource, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sclient must not be empty", config)
	}

	r := &Resource{
		logger:    config.Logger,
		k8sClient: config.K8sClient,
		AWSConfig: config.AWSConfig,
	}

	return r, nil
}

func (r *Resource) Name() string {
	return Name
}

func (r *Resource) addAWSClientsToContext(ctx context.Context, cn string) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	{
		arn, err := credential.GetARN(r.k8sClient, cn)
		if err != nil {
			return microerror.Mask(err)
		}

		c := r.AWSConfig
		c.RoleARN = arn

		clients, err := aws.NewClients(c)
		if err != nil {
			return microerror.Mask(err)
		}

		cc.Client.AWS = clients
	}

	return nil
}

func (r *Resource) ToAWSTagList(v interface{}) (v1alpha1.AWSTagList, error) {
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
