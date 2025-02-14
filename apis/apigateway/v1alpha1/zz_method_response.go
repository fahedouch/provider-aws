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

// MethodResponseParameters defines the desired state of MethodResponse
type MethodResponseParameters struct {
	// Region is which region the MethodResponse will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The HTTP verb of the Method resource.
	// +kubebuilder:validation:Required
	HTTPMethod *string `json:"httpMethod"`
	// Specifies the Model resources used for the response's content type. Response
	// models are represented as a key/value map, with a content type as the key
	// and a Model name as the value.
	ResponseModels map[string]*string `json:"responseModels,omitempty"`
	// A key-value map specifying required or optional response parameters that
	// API Gateway can send back to the caller. A key defines a method response
	// header name and the associated value is a Boolean flag indicating whether
	// the method response parameter is required or not. The method response header
	// names must match the pattern of method.response.header.{name}, where name
	// is a valid and unique header name. The response parameter names defined here
	// are available in the integration response to be mapped from an integration
	// response header expressed in integration.response.header.{name}, a static
	// value enclosed within a pair of single quotes (e.g., 'application/json'),
	// or a JSON expression from the back-end response payload in the form of integration.response.body.{JSON-expression},
	// where JSON-expression is a valid JSON expression without the $ prefix.)
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty"`
	// The method response's status code.
	// +kubebuilder:validation:Required
	StatusCode                     *string `json:"statusCode"`
	CustomMethodResponseParameters `json:",inline"`
}

// MethodResponseSpec defines the desired state of MethodResponse
type MethodResponseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MethodResponseParameters `json:"forProvider"`
}

// MethodResponseObservation defines the observed state of MethodResponse
type MethodResponseObservation struct {
	CustomMethodResponseObservation `json:",inline"`
}

// MethodResponseStatus defines the observed state of MethodResponse.
type MethodResponseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MethodResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponse is the Schema for the MethodResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MethodResponseSpec   `json:"spec"`
	Status            MethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponseList contains a list of MethodResponses
type MethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MethodResponse `json:"items"`
}

// Repository type metadata.
var (
	MethodResponseKind             = "MethodResponse"
	MethodResponseGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MethodResponseKind}.String()
	MethodResponseKindAPIVersion   = MethodResponseKind + "." + GroupVersion.String()
	MethodResponseGroupVersionKind = GroupVersion.WithKind(MethodResponseKind)
)

func init() {
	SchemeBuilder.Register(&MethodResponse{}, &MethodResponseList{})
}
