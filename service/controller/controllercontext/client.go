package controllercontext

import (
	"github.com/giantswarm/aws-tag-operator/client/aws"
)

type ContextClient struct {
	Cluster ContextClientCluster
}

type ContextClientCluster struct {
	AWS aws.Clients
}
