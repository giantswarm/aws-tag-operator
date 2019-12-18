package key

import (
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/microerror"
)

func CredentialName(clusterName string) string {
	return clusterName
}

func CredentialNamespace() string {
	return "giantswarm"
}

func ToCluster(v interface{}) (v1alpha1.AWSTagList, error) {
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
