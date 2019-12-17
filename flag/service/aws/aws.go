package aws

import (
	"github.com/giantswarm/aws-operator/flag/service/aws/accesskey"
)

type AWS struct {
	AccessKey accesskey.AccessKey
	Region    string
}
