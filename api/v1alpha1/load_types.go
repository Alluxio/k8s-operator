/*
 * The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
 * (the "License"). You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type LoadSpec struct {
	Dataset string `json:"dataset"`
	Path    string `json:"path"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="DatasetPhase",type="string",JSONPath=`.status.phase`,priority=0

// Load is the Schema for the loads API
type Load struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadSpec   `json:"spec,omitempty"`
	Status LoadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadList contains a list of Load. Operator wouldn't work without this list.
type LoadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Load `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Load{}, &LoadList{})
}
