package v1alpha1

type ClusterPhase string

const (
	ClusterPhaseNone     ClusterPhase = ""
	ClusterPhaseCreating ClusterPhase = "Creating"
	ClusterPhaseReady    ClusterPhase = "Ready"
)

// AlluxioClusterStatus defines the observed state of AlluxioCluster
type AlluxioClusterStatus struct {
	Phase ClusterPhase `json:"phase"`
}
