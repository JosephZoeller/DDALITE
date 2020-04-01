package kubeutil

// Node Represents Kubernetes Nodes
type Node struct {
	Name             string
	Status           string
	Role             string
	Age              string
	Version          string
	InternalIP       string
	ExternalIP       string
	OSImage          string
	KernelVer        string
	ContainerRunTime string
}

// Pod represents K8 Pod
type Pod struct {
	Name          string
	Ready         string
	Status        string
	Restart       string
	Age           string
	IPaddr        string
	Node          string
	NominatedNode string
	ReadinessGate string
}

// Service represents K8 Service
type Service struct {
	Name       string
	Type       string
	ClusterIP  string
	ExternalIP string
	Port       string
	Age        string
	Selector   string
}
