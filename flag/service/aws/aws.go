package aws

import (
	"github.com/giantswarm/aws-tag-operator/flag/service/aws/accesskey"
)

type AWS struct {
	HostAccessKey accesskey.AccessKey
	Region        string
}
