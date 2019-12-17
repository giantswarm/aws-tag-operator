package controllercontext

type ContextStatus struct {
	Cluster ContextStatusCluster
}

type ContextStatusCluster struct {
	AWS             ContextStatusClusterAWS
	OperatorVersion string
}

type ContextStatusClusterAWS struct {
	AccountID string
	Region    string
}
