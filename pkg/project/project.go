package project

var (
	bundleVersion = "0.0.1"
	description   = "The aws-tag-operator does something."
	gitSHA        = "n/a"
	name          = "aws-tag-operator"
	source        = "https://github.com/giantswarm/aws-tag-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
