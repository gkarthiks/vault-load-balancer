package globals

import discovery "github.com/gkarthiks/k8s-discovery"

var (
	VaultIPList []string
	K8s         *discovery.K8s
	Namespace   string
	HttpTimeout string
)

const (
	HealthCheckPath = ":8200/v1/sys/seal-status"
	ProxyPath       = ":8200"
	DefaultTimeOut  = "1"
	DefaultBalancerPort  = 8000
)
