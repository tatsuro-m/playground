package lib

const (
	ClusterRegion = "asia-northeast1"
	ClusterName   = "stg-ingress-gke-cluster"
	ProjectID     = "playground-318023"
	Host          = "asia-northeast1-docker.pkg.dev"
	Env           = "stg"
	AppName       = "ingress"
)

var Services = []string{"nginx", "frontend1"}
