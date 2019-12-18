package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

type Interface interface {
	EC2Client() ec2iface.EC2API
	STSClient() stsiface.STSAPI
}
