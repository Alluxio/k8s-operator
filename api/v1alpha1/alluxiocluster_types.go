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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlluxioClusterSpec defines the desired state of AlluxioCluster
type AlluxioClusterSpec struct {
	Image              string            `json:"image" yaml:"image"`
	ImageTag           string            `json:"imageTag" yaml:"imageTag"`
	ImagePullPolicy    string            `json:"imagePullPolicy" yaml:"imagePullPolicy"`
	ImagePullSecrets   []string          `json:"imagePullSecrets" yaml:"imagePullSecrets"`
	User               string            `json:"user" yaml:"user"`
	Group              string            `json:"group" yaml:"group"`
	FsGroup            string            `json:"fsGroup" yaml:"fsGroup"`
	HostNetwork        string            `json:"hostNetwork" yaml:"hostNetwork"`
	DnsPolicy          string            `json:"dnsPolicy" yaml:"dnsPolicy"`
	ServiceAccountName string            `json:"serviceAccountName" yaml:"serviceAccountName"`
	HostAliases        []HostAlias       `json:"hostAliases" yaml:"hostAliases"`
	NodeSelector       map[string]string `json:"nodeSelector" yaml:"nodeSelector"`
	Tolerations        []Toleration      `json:"tolerations" yaml:"tolerations"`
	JvmOptions         []string          `json:"jvmOptions" yaml:"jvmOptions"`
	PvcMounts          MountSpec         `json:"pvcMounts" yaml:"pvcMounts"`
	ConfigMaps         MountSpec         `json:"configMaps" yaml:"ConfigMaps"`
	Secrets            MountSpec         `json:"secrets" yaml:"secrets"`
	Dataset            DatasetSpec       `json:"dataset" yaml:"dataset"`
	Master             MasterSpec        `json:"master" yaml:"master"`
	Journal            JournalSpec       `json:"journal" yaml:"journal"`
	Worker             WorkerSpec        `json:"worker" yaml:"worker"`
	Pagestore          PagestoreSpec     `json:"pagestore" yaml:"pagestore"`
	Metastore          MetastoreSpec     `json:"metastore" yaml:"metastore"`
	Proxy              ProxySpec         `json:"proxy" yaml:"proxy"`
	Fuse               FuseSpec          `json:"fuse" yaml:"fuse"`
	Metrics            MetricsSpec       `json:"metrics" yaml:"metrics"`
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
	Master map[string]string `json:"master" yaml:"master"`
	Worker map[string]string `json:"worker" yaml:"worker"`
	Fuse   map[string]string `json:"fuse" yaml:"fuse"`
	Proxy  map[string]string `json:"proxy" yaml:"proxy"`
}

type DatasetSpec struct {
	Path        string            `json:"path" yaml:"path"`
	Credentials map[string]string `json:"credentials" yaml:"credentials"`
}

type MasterSpec struct {
	Affinity       corev1.Affinity   `json:"affinity" yaml:"affinity"`
	Count          int               `json:"count" yaml:"count"`
	Enabled        bool              `json:"enabled" yaml:"enabled"`
	Env            map[string]string `json:"env" yaml:"env"`
	JvmOptions     []string          `json:"jvmOptions" yaml:"jvmOptions"`
	LivenessProbe  ProbeSpec         `json:"livenessProbe" yaml:"livenessProbe"`
	NodeSelector   map[string]string `json:"nodeSelector" yaml:"nodeSelector"`
	PodAnnotations map[string]string `json:"podAnnotations" yaml:"podAnnotations"`
	Ports          map[string]int    `json:"ports" yaml:"port"`
	ReadinessProbe ProbeSpec         `json:"readinessProbe" yaml:"readinessProbe"`
	Resources      ResourcesSpec     `json:"resources" yaml:"resources"`
	StartupProbe   ProbeSpec         `json:"startupProbe" yaml:"startupProbe"`
	Tolerations    []Toleration      `json:"tolerations" yaml:"tolerations"`
}

type JournalSpec struct {
	RunFormat    bool   `json:"runFormat" yaml:"runFormat"`
	Size         string `json:"size" yaml:"size"`
	StorageClass string `json:"storageClass" yaml:"storageClass"`
}

type WorkerSpec struct {
	Affinity             corev1.Affinity   `json:"affinity" yaml:"affinity"`
	Count                int               `json:"count" yaml:"count"`
	Env                  map[string]string `json:"env" yaml:"env"`
	JvmOptions           []string          `json:"jvmOptions" yaml:"jvmOptions"`
	LimitOneWorkerPerPod bool              `json:"limitOneWorkerPerPod" yaml:"limitOneWorkerPerPod"`
	LivenessProbe        ProbeSpec         `json:"livenessProbe" yaml:"livenessProbe"`
	NodeSelector         map[string]string `json:"nodeSelector" yaml:"nodeSelector"`
	PodAnnotations       map[string]string `json:"podAnnotations" yaml:"podAnnotations"`
	Ports                map[string]int    `json:"ports" yaml:"ports"`
	ReadinessProbe       ProbeSpec         `json:"readinessProbe" yaml:"readinessProbe"`
	Resources            ResourcesSpec     `json:"resources" yaml:"resources"`
	StartupProbe         ProbeSpec         `json:"startupProbe" yaml:"startupProbe"`
	Tolerations          []Toleration      `json:"tolerations" yaml:"tolerations"`
}

type PagestoreSpec struct {
	HostPath     string `json:"hostPath" yaml:"hostPath"`
	MemoryBacked bool   `json:"memoryBacked" yaml:"memoryBacked"`
	Quota        string `json:"quota" yaml:"quota"`
	StorageClass string `json:"storageClass" yaml:"storageClass"`
	Type         string `json:"type" yaml:"type"`
}

type MetastoreSpec struct {
	Enabled      bool   `json:"enabled" yaml:"enabled"`
	Size         string `json:"size" yaml:"size"`
	StorageClass string `json:"storageClass" yaml:"storageClass"`
}

type ProxySpec struct {
	Affinity       corev1.Affinity   `json:"affinity" yaml:"affinity"`
	Enabled        bool              `json:"enabled" yaml:"enabled"`
	Env            map[string]string `json:"env" yaml:"env"`
	JvmOptions     []string          `json:"jvmOptions" yaml:"jvmOptions"`
	NodeSelector   map[string]string `json:"nodeSelector" yaml:"nodeSelector"`
	PodAnnotations map[string]string `json:"podAnnotations" yaml:"podAnnotations"`
	Ports          map[string]int    `json:"ports" yaml:"ports"`
	Resources      ResourcesSpec     `json:"resources" yaml:"resources"`
	Tolerations    []Toleration      `json:"tolerations" yaml:"tolerations"`
}

type FuseSpec struct {
	Affinity       corev1.Affinity   `json:"affinity" yaml:"affinity"`
	Enabled        bool              `json:"enabled" yaml:"enabled"`
	Env            map[string]string `json:"env" yaml:"env"`
	JvmOptions     []string          `json:"jvmOptions" yaml:"jvmOptions"`
	MountOptions   []string          `json:"mountOptions" yaml:"mountOptions"`
	NodeSelector   map[string]string `json:"nodeSelector" yaml:"nodeSelector"`
	PodAnnotations map[string]string `json:"podAnnotations" yaml:"podAnnotations"`
	Resources      ResourcesSpec     `json:"resources" yaml:"resources"`
	Tolerations    []Toleration      `json:"tolerations" yaml:"tolerations"`
}

type ResourcesSpec struct {
	Limits   CpuMemSpec `json:"limits" yaml:"limits"`
	Requests CpuMemSpec `json:"requests" yaml:"requests"`
}

type CpuMemSpec struct {
	Cpu    string `json:"cpu" yaml:"cpu"`
	Memory string `json:"memory" yaml:"memory"`
}

type ProbeSpec struct {
	FailureThreshold    int `json:"failureThreshold" yaml:"failureThreshold"`
	InitialDelaySeconds int `json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds" yaml:"periodSeconds"`
	SuccessThreshold    int `json:"successThreshold" yaml:"successThreshold"`
	TimeoutSeconds      int `json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

type MetricsSpec struct {
	ConsoleSink              ConsoleSinkSpec              `json:"consoleSink" yaml:"consoleSink"`
	CsvSink                  CsvSinkSpec                  `json:"csvSink" yaml:"csvSink"`
	GraphiteSink             GraphiteSinkSpec             `json:"graphiteSink" yaml:"graphiteSink"`
	JmxSink                  JmxSinkSpec                  `json:"jmxSink" yaml:"jmxSink"`
	PrometheusMetricsServlet PrometheusMetricsServletSpec `json:"prometheusMetricsServlet" yaml:"prometheusMetricsServlet"`
	Slf4jSink                Slf4jSinkSpec                `json:"Slf4jSink" yaml:"Slf4jSink"`
}

type ConsoleSinkSpec struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Period  int    `json:"period" yaml:"period"`
	Unit    string `json:"unit" yaml:"unit"`
}

type CsvSinkSpec struct {
	Directory string `json:"directory" yaml:"directory"`
	Enabled   bool   `json:"enabled" yaml:"enabled"`
	Period    int    `json:"period" yaml:"period"`
	Unit      string `json:"unit" yaml:"unit"`
}

type GraphiteSinkSpec struct {
	Enabled  bool   `json:"enabled" yaml:"enabled"`
	Hostname string `json:"hostname" yaml:"hostname"`
	Period   int    `json:"period" yaml:"period"`
	Port     int    `json:"port" yaml:"port"`
	Prefix   string `json:"prefix" yaml:"prefix"`
	Unit     string `json:"unit" yaml:"unit"`
}

type JmxSinkSpec struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Domain  string `json:"domain" yaml:"domain"`
}

type PrometheusMetricsServletSpec struct {
	Enabled        bool              `json:"enabled" yaml:"enabled"`
	PodAnnotations map[string]string `json:"podAnnotations" yaml:"podAnnotations"`
}

type Slf4jSinkSpec struct {
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	FilterClass string `json:"filterClass" yaml:"filterClass"`
	FilterRegex string `json:"filterRegex" yaml:"filterRegex"`
	Period      int    `json:"period" yaml:"period"`
	Unit        string `json:"unit" yaml:"unit"`
}

// AlluxioClusterStatus defines the observed state of AlluxioCluster
type AlluxioClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Nodes []NodeStatus `json:"nodes" yaml:"nodes"`
}

type NodeStatus struct {
	Name  string          `json:"name" yaml:"name"`
	Phase corev1.PodPhase `json:"phase,omitempty" yaml:"phase,omitempty"`
	PodIP string          `json:"podIP,omitempty" yaml:"podIP,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AlluxioCluster is the Schema for the alluxioclusters API
type AlluxioCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlluxioClusterSpec   `json:"spec,omitempty"`
	Status AlluxioClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AlluxioClusterList contains a list of AlluxioCluster
type AlluxioClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Spec            AlluxioClusterSpec   `json:"spec,omitempty"`
	Status          AlluxioClusterStatus `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&AlluxioCluster{}, &AlluxioClusterList{})
}
