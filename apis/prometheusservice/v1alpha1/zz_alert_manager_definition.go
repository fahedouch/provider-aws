/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// AlertManagerDefinitionParameters defines the desired state of AlertManagerDefinition
type AlertManagerDefinitionParameters struct {
	// Region is which region the AlertManagerDefinition will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The alert manager definition data.
	// +kubebuilder:validation:Required
	Data                                   []byte `json:"data"`
	CustomAlertManagerDefinitionParameters `json:",inline"`
}

// AlertManagerDefinitionSpec defines the desired state of AlertManagerDefinition
type AlertManagerDefinitionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AlertManagerDefinitionParameters `json:"forProvider"`
}

// AlertManagerDefinitionObservation defines the observed state of AlertManagerDefinition
type AlertManagerDefinitionObservation struct {
	// Status code of this definition.
	StatusCode *string `json:"statusCode,omitempty"`
	// The reason for failure if any.
	StatusReason *string `json:"statusReason,omitempty"`

	CustomAlertManagerDefinitionObservation `json:",inline"`
}

// AlertManagerDefinitionStatus defines the observed state of AlertManagerDefinition.
type AlertManagerDefinitionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AlertManagerDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertManagerDefinition is the Schema for the AlertManagerDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AlertManagerDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertManagerDefinitionSpec   `json:"spec"`
	Status            AlertManagerDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertManagerDefinitionList contains a list of AlertManagerDefinitions
type AlertManagerDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertManagerDefinition `json:"items"`
}

// Repository type metadata.
var (
	AlertManagerDefinitionKind             = "AlertManagerDefinition"
	AlertManagerDefinitionGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertManagerDefinitionKind}.String()
	AlertManagerDefinitionKindAPIVersion   = AlertManagerDefinitionKind + "." + GroupVersion.String()
	AlertManagerDefinitionGroupVersionKind = GroupVersion.WithKind(AlertManagerDefinitionKind)
)

func init() {
	SchemeBuilder.Register(&AlertManagerDefinition{}, &AlertManagerDefinitionList{})
}
