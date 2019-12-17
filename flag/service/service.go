package service

import (
	"github.com/giantswarm/aws-tag-operator/flag/service/aws"
	"github.com/giantswarm/operatorkit/flag/service/kubernetes"
)

// Service is an intermediate data structure for command line configuration flags.
type Service struct {
	Kubernetes kubernetes.Kubernetes
	AWS        aws.AWS
}
