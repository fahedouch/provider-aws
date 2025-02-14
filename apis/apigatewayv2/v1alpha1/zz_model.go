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

// ModelParameters defines the desired state of Model
type ModelParameters struct {
	// Region is which region the Model will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	ContentType *string `json:"contentType,omitempty"`

	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name"`

	// +kubebuilder:validation:Required
	Schema                *string `json:"schema"`
	CustomModelParameters `json:",inline"`
}

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ModelParameters `json:"forProvider"`
}

// ModelObservation defines the observed state of Model
type ModelObservation struct {
	ModelID *string `json:"modelID,omitempty"`

	CustomModelObservation `json:",inline"`
}

// ModelStatus defines the observed state of Model.
type ModelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Model is the Schema for the Models API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelSpec   `json:"spec"`
	Status            ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelList contains a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

// Repository type metadata.
var (
	ModelKind             = "Model"
	ModelGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ModelKind}.String()
	ModelKindAPIVersion   = ModelKind + "." + GroupVersion.String()
	ModelGroupVersionKind = GroupVersion.WithKind(ModelKind)
)

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}
