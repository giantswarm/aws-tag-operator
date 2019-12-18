package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/giantswarm/microerror"
)

type Config struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	RoleARN         string
	SessionToken    string
}

type Clients struct {
	ec2Client ec2iface.EC2API
	stsClient stsiface.STSAPI
}

func (c *Clients) EC2Client() ec2iface.EC2API {
	return c.ec2Client
}

func (c *Clients) STSClient() stsiface.STSAPI {
	return c.stsClient
}

func NewClients(config Config) (*Clients, error) {
	if config.AccessKeyID == "" {
		return &Clients{}, microerror.Maskf(invalidConfigError, "%T.AccessKeyID must not be empty", config)
	}
	if config.AccessKeySecret == "" {
		return &Clients{}, microerror.Maskf(invalidConfigError, "%T.AccessKeySecret must not be empty", config)
	}
	if config.Region == "" {
		return &Clients{}, microerror.Maskf(invalidConfigError, "%T.Region must not be empty", config)
	}

	var err error

	var s *session.Session
	{
		c := &aws.Config{
			Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.AccessKeySecret, config.SessionToken),
			Region:      aws.String(config.Region),
		}

		s, err = session.NewSession(c)
		if err != nil {
			return &Clients{}, microerror.Mask(err)
		}
	}

	var c Clients
	if config.RoleARN != "" {
		creds := stscreds.NewCredentials(s, config.RoleARN)
		c = newClients(s, &aws.Config{Credentials: creds})
	} else {
		c = newClients(s)
	}

	return &c, nil
}

func newClients(session *session.Session, configs ...*aws.Config) Clients {

	c := Clients{
		ec2Client: ec2.New(session, configs...),
		stsClient: sts.New(session, configs...),
	}

	return c
}
