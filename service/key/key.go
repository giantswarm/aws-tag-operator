package key

import (
	"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	"github.com/giantswarm/microerror"
)

func CredentialName(cluster v1alpha1.AWSClusterConfig) string {
	return cluster.Spec.Guest.CredentialSecret.Name
}

func CredentialNamespace(cluster v1alpha1.AWSClusterConfig) string {
	return cluster.Spec.Guest.CredentialSecret.Namespace
}

func ToCluster(v interface{}) (v1alpha1.AWSClusterConfig, error) {
	if v == nil {
		return v1alpha1.AWSClusterConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AWSClusterConfig{}, v)
	}

	p, ok := v.(*v1alpha1.AWSClusterConfig)
	if !ok {
		return v1alpha1.AWSClusterConfig{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.AWSClusterConfig{}, v)
	}

	c := p.DeepCopy()

	return *c, nil
}
