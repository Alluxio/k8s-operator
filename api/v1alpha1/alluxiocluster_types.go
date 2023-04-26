/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlluxioClusterSpec defines the desired state of AlluxioCluster
type AlluxioClusterSpec struct {
	NameOverride       string            `json:"nameOverride,omitempty" yaml:"nemOverride,omitempty"`
	Image              string            `json:"image,omitempty" yaml:"image,omitempty"`
	ImageTag           string            `json:"imageTag,omitempty" yaml:"imageTag,omitempty"`
	ImagePullPolicy    string            `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	ImagePullSecrets   []string          `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`
	User               string            `json:"user,omitempty" yaml:"user,omitempty"`
	Group              string            `json:"group,omitempty" yaml:"group,omitempty"`
	FsGroup            string            `json:"fsGroup,omitempty" yaml:"fsGroup,omitempty"`
	HostNetwork        string            `json:"hostNetwork,omitempty" yaml:"hostNetwork,omitempty"`
	DnsPolicy          string            `json:"dnsPolicy,omitempty" yaml:"dnsPolicy,omitempty"`
	ServiceAccountName string            `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`
	HostAliases        []HostAlias       `json:"hostAliases,omitempty" yaml:"hostAliases,omitempty"`
	NodeSelector       map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Tolerations        []Toleration      `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	Properties         map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
	JvmOptions         []string          `json:"jvmOptions,omitempty" yaml:"jvmOptions,omitempty"`
	PvcMounts          MountSpec         `json:"pvcMounts,omitempty" yaml:"pvcMounts,omitempty"`
	ConfigMaps         MountSpec         `json:"configMaps,omitempty" yaml:"configMaps,omitempty"`
	Secrets            MountSpec         `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	Master             MasterSpec        `json:"master,omitempty" yaml:"master,omitempty"`
	Journal            JournalSpec       `json:"journal,omitempty" yaml:"journal,omitempty"`
	Worker             WorkerSpec        `json:"worker,omitempty" yaml:"worker,omitempty"`
	Pagestore          PagestoreSpec     `json:"pagestore,omitempty" yaml:"pagestore,omitempty"`
	Metastore          MetastoreSpec     `json:"metastore,omitempty" yaml:"metastore,omitempty"`
	Proxy              ProxySpec         `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	Fuse               FuseSpec          `json:"fuse,omitempty" yaml:"fuse,omitempty"`
	Metrics            MetricsSpec       `json:"metrics,omitempty" yaml:"metrics,omitempty"`
}

type HostAlias struct {
	Ip        string   `json:"ip" yaml:"ip"`
	Hostnames []string `json:"hostnames" yaml:"hostnames"`
}

type Toleration struct {
	Key      string `json:"key" yaml:"key"`
	Operator string `json:"operator" yaml:"operator"`
	Value    string `json:"value" yaml:"value"`
	Effect   string `json:"effect" yaml:"effect"`
}

type MountSpec struct {
	Master map[string]string `json:"master,omitempty" yaml:"master,omitempty"`
	Worker map[string]string `json:"worker,omitempty" yaml:"worker,omitempty"`
	Fuse   map[string]string `json:"fuse,omitempty" yaml:"fuse,omitempty"`
	Proxy  map[string]string `json:"proxy,omitempty" yaml:"proxy,omitempty"`
}

type MasterSpec struct {
	Affinity       corev1.Affinity   `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Count          int               `json:"count,omitempty" yaml:"count,omitempty"`
	Enabled        bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Env            map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	JvmOptions     []string          `json:"jvmOptions,omitempty" yaml:"jvmOptions,omitempty"`
	LivenessProbe  ProbeSpec         `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	PodAnnotations map[string]string `json:"podAnnotations,omitempty" yaml:"podAnnotations,omitempty"`
	Ports          map[string]int    `json:"ports,omitempty" yaml:"port,omitempty"`
	ReadinessProbe ProbeSpec         `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	Resources      ResourcesSpec     `json:"resources,omitempty" yaml:"resources,omitempty"`
	StartupProbe   ProbeSpec         `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`
	Tolerations    []Toleration      `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
}

type JournalSpec struct {
	RunFormat    bool   `json:"runFormat,omitempty" yaml:"runFormat,omitempty"`
	Size         string `json:"size,omitempty" yaml:"size,omitempty"`
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}

type WorkerSpec struct {
	Affinity             corev1.Affinity   `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Count                int               `json:"count,omitempty" yaml:"count,omitempty"`
	Env                  map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	JvmOptions           []string          `json:"jvmOptions,omitempty" yaml:"jvmOptions,omitempty"`
	LimitOneWorkerPerPod bool              `json:"limitOneWorkerPerPod,omitempty" yaml:"limitOneWorkerPerPod,omitempty"`
	LivenessProbe        ProbeSpec         `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	NodeSelector         map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	PodAnnotations       map[string]string `json:"podAnnotations,omitempty" yaml:"podAnnotations,omitempty"`
	Ports                map[string]int    `json:"ports,omitempty" yaml:"ports,omitempty"`
	ReadinessProbe       ProbeSpec         `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	Resources            ResourcesSpec     `json:"resources,omitempty" yaml:"resources,omitempty"`
	StartupProbe         ProbeSpec         `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`
	Tolerations          []Toleration      `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
}

type PagestoreSpec struct {
	HostPath     string `json:"hostPath,omitempty" yaml:"hostPath,omitempty"`
	MemoryBacked bool   `json:"memoryBacked,omitempty" yaml:"memoryBacked,omitempty"`
	Quota        string `json:"quota,omitempty" yaml:"quota,omitempty"`
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
	Type         string `json:"type,omitempty" yaml:"type,omitempty"`
}

type MetastoreSpec struct {
	Enabled      bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Size         string `json:"size,omitempty" yaml:"size,omitempty"`
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}

type ProxySpec struct {
	Affinity       corev1.Affinity   `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Enabled        bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Env            map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	JvmOptions     []string          `json:"jvmOptions,omitempty" yaml:"jvmOptions,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	PodAnnotations map[string]string `json:"podAnnotations,omitempty" yaml:"podAnnotations,omitempty"`
	Ports          map[string]int    `json:"ports,omitempty" yaml:"ports,omitempty"`
	Resources      ResourcesSpec     `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tolerations    []Toleration      `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
}

type FuseSpec struct {
	Affinity       corev1.Affinity   `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Enabled        bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Env            map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	JvmOptions     []string          `json:"jvmOptions,omitempty" yaml:"jvmOptions,omitempty"`
	MountOptions   []string          `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	PodAnnotations map[string]string `json:"podAnnotations,omitempty" yaml:"podAnnotations,omitempty"`
	Resources      ResourcesSpec     `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tolerations    []Toleration      `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
}

type ResourcesSpec struct {
	Limits   CpuMemSpec `json:"limits,omitempty" yaml:"limits,omitempty"`
	Requests CpuMemSpec `json:"requests,omitempty" yaml:"requests,omitempty"`
}

type CpuMemSpec struct {
	Cpu    string `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}

type ProbeSpec struct {
	FailureThreshold    int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`
	InitialDelaySeconds int `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       int `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`
	SuccessThreshold    int `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`
	TimeoutSeconds      int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`
}

type MetricsSpec struct {
	ConsoleSink              ConsoleSinkSpec              `json:"consoleSink,omitempty" yaml:"consoleSink,omitempty"`
	CsvSink                  CsvSinkSpec                  `json:"csvSink,omitempty" yaml:"csvSink,omitempty"`
	GraphiteSink             GraphiteSinkSpec             `json:"graphiteSink,omitempty" yaml:"graphiteSink,omitempty"`
	JmxSink                  JmxSinkSpec                  `json:"jmxSink,omitempty" yaml:"jmxSink,omitempty"`
	PrometheusMetricsServlet PrometheusMetricsServletSpec `json:"prometheusMetricsServlet,omitempty" yaml:"prometheusMetricsServlet,omitempty"`
	Slf4jSink                Slf4jSinkSpec                `json:"slf4jSink,omitempty" yaml:"slf4jSink,omitempty"`
}

type ConsoleSinkSpec struct {
	Enabled bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Period  int    `json:"period,omitempty" yaml:"period,omitempty"`
	Unit    string `json:"unit,omitempty" yaml:"unit,omitempty"`
}

type CsvSinkSpec struct {
	Directory string `json:"directory,omitempty" yaml:"directory,omitempty"`
	Enabled   bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Period    int    `json:"period,omitempty" yaml:"period,omitempty"`
	Unit      string `json:"unit,omitempty" yaml:"unit,omitempty"`
}

type GraphiteSinkSpec struct {
	Enabled  bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Period   int    `json:"period,omitempty" yaml:"period,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Prefix   string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Unit     string `json:"unit,omitempty" yaml:"unit,omitempty"`
}

type JmxSinkSpec struct {
	Enabled bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Domain  string `json:"domain,omitempty" yaml:"domain,omitempty"`
}

type PrometheusMetricsServletSpec struct {
	Enabled        bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	PodAnnotations map[string]string `json:"podAnnotations,omitempty" yaml:"podAnnotations,omitempty"`
}

type Slf4jSinkSpec struct {
	Enabled     bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	FilterClass string `json:"filterClass,omitempty" yaml:"filterClass,omitempty"`
	FilterRegex string `json:"filterRegex,omitempty" yaml:"filterRegex,omitempty"`
	Period      int    `json:"period,omitempty" yaml:"period,omitempty"`
	Unit        string `json:"unit,omitempty" yaml:"unit,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ClusterPhase",type="string",JSONPath=`.status.phase`,priority=0
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=`.metadata.creationTimestamp`,priority=0

// AlluxioCluster is the Schema for the alluxioclusters API
type AlluxioCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlluxioClusterSpec   `json:"spec,omitempty"`
	Status AlluxioClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlluxioClusterList contains a list of AlluxioCluster
type AlluxioClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlluxioCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlluxioCluster{}, &AlluxioClusterList{})
}
