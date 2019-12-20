package aws

import (
	"github.com/giantswarm/aws-operator/flag/service/aws/accesskey"
)

type AWS struct {
	HostAccessKey accesskey.AccessKey
	Region        string
}
