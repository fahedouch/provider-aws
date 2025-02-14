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

// TransitGatewayRouteParameters defines the desired state of TransitGatewayRoute
type TransitGatewayRouteParameters struct {
	// Region is which region the TransitGatewayRoute will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Indicates whether to drop traffic that matches this route.
	Blackhole *bool `json:"blackhole,omitempty"`
	// The CIDR range used for destination matches. Routing decisions are based
	// on the most specific match.
	// +kubebuilder:validation:Required
	DestinationCIDRBlock                *string `json:"destinationCIDRBlock"`
	CustomTransitGatewayRouteParameters `json:",inline"`
}

// TransitGatewayRouteSpec defines the desired state of TransitGatewayRoute
type TransitGatewayRouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TransitGatewayRouteParameters `json:"forProvider"`
}

// TransitGatewayRouteObservation defines the observed state of TransitGatewayRoute
type TransitGatewayRouteObservation struct {
	// The ID of the prefix list used for destination matches.
	PrefixListID *string `json:"prefixListID,omitempty"`
	// The state of the route.
	State *string `json:"state,omitempty"`
	// The attachments.
	TransitGatewayAttachments []*TransitGatewayRouteAttachment `json:"transitGatewayAttachments,omitempty"`
	// The ID of the transit gateway route table announcement.
	TransitGatewayRouteTableAnnouncementID *string `json:"transitGatewayRouteTableAnnouncementID,omitempty"`
	// The route type.
	Type *string `json:"type_,omitempty"`

	CustomTransitGatewayRouteObservation `json:",inline"`
}

// TransitGatewayRouteStatus defines the observed state of TransitGatewayRoute.
type TransitGatewayRouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TransitGatewayRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRoute is the Schema for the TransitGatewayRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteSpec   `json:"spec"`
	Status            TransitGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteList contains a list of TransitGatewayRoutes
type TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRoute `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteKind             = "TransitGatewayRoute"
	TransitGatewayRouteGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRouteKind}.String()
	TransitGatewayRouteKindAPIVersion   = TransitGatewayRouteKind + "." + GroupVersion.String()
	TransitGatewayRouteGroupVersionKind = GroupVersion.WithKind(TransitGatewayRouteKind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRoute{}, &TransitGatewayRouteList{})
}
