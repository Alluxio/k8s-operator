package v1alpha1

import v1 "k8s.io/api/apps/v1"

type ClusterPhase string

const (
	ClusterPhaseNone               ClusterPhase = ""
	ClusterPhaseCreatingOrUpdating ClusterPhase = "Creating/Updating"
	ClusterPhaseReady              ClusterPhase = "Ready"
)

// AlluxioClusterStatus defines the observed state of AlluxioCluster
type AlluxioClusterStatus struct {
	Phase  ClusterPhase    `json:"phase"`
	Master *v1.StatefulSet `json:"master"`
	Worker *v1.Deployment  `json:"worker"`
	Fuse   *v1.DaemonSet   `json:"fuse"`
	Proxy  *v1.DaemonSet   `json:"proxy"`
}
