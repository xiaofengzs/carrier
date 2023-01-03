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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CarrierSpec defines the desired state of Carrier
type CarrierSpec struct {
	Image string `json:"image"`
	Port  int32  `json:"port"`

	// +kubebuilder:validation:Maximum=5
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas"`

	// +kubebuilder:validation:Optional
	EnableService bool `json:"enableService,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceConfig ServiceConfig `json:"serviceConfig"`
}

// ServiceConfig defines the desired state of ServiceConfig
type ServiceConfig struct {

	// +kubebuilder:validation:Optional
	Type       string `json:"type,omitempty"`
	Port       int32  `json:"port"`
	TargetPort int32  `json:"targetPort"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Maximum=32767
	// +kubebuilder:validation:Minimum=30000
	NodePort int32 `json:"nodePort,omitempty"`
}

// CarrierStatus defines the observed state of Carrier
type CarrierStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Carrier is the Schema for the carriers API
type Carrier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CarrierSpec   `json:"spec,omitempty"`
	Status CarrierStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CarrierList contains a list of Carrier
type CarrierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Carrier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Carrier{}, &CarrierList{})
}
