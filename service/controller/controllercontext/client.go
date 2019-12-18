package controllercontext

import (
	"github.com/giantswarm/aws-tag-operator/client/aws"
)

type ContextClient struct {
	AWS aws.Clients
}
