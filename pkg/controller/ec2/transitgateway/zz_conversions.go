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

package transitgateway

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ec2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeTransitGatewaysInput returns input for read
// operation.
func GenerateDescribeTransitGatewaysInput(cr *svcapitypes.TransitGateway) *svcsdk.DescribeTransitGatewaysInput {
	res := &svcsdk.DescribeTransitGatewaysInput{}

	if cr.Status.AtProvider.TransitGatewayID != nil {
		f3 := []*string{}
		f3 = append(f3, cr.Status.AtProvider.TransitGatewayID)
		res.SetTransitGatewayIds(f3)
	}

	return res
}

// GenerateTransitGateway returns the current state in the form of *svcapitypes.TransitGateway.
func GenerateTransitGateway(resp *svcsdk.DescribeTransitGatewaysOutput) *svcapitypes.TransitGateway {
	cr := &svcapitypes.TransitGateway{}

	found := false
	for _, elem := range resp.TransitGateways {
		if elem.CreationTime != nil {
			cr.Status.AtProvider.CreationTime = &metav1.Time{*elem.CreationTime}
		} else {
			cr.Status.AtProvider.CreationTime = nil
		}
		if elem.Description != nil {
			cr.Spec.ForProvider.Description = elem.Description
		} else {
			cr.Spec.ForProvider.Description = nil
		}
		if elem.Options != nil {
			f2 := &svcapitypes.TransitGatewayRequestOptions{}
			if elem.Options.AmazonSideAsn != nil {
				f2.AmazonSideASN = elem.Options.AmazonSideAsn
			}
			if elem.Options.AutoAcceptSharedAttachments != nil {
				f2.AutoAcceptSharedAttachments = elem.Options.AutoAcceptSharedAttachments
			}
			if elem.Options.DefaultRouteTableAssociation != nil {
				f2.DefaultRouteTableAssociation = elem.Options.DefaultRouteTableAssociation
			}
			if elem.Options.DefaultRouteTablePropagation != nil {
				f2.DefaultRouteTablePropagation = elem.Options.DefaultRouteTablePropagation
			}
			if elem.Options.DnsSupport != nil {
				f2.DNSSupport = elem.Options.DnsSupport
			}
			if elem.Options.MulticastSupport != nil {
				f2.MulticastSupport = elem.Options.MulticastSupport
			}
			if elem.Options.TransitGatewayCidrBlocks != nil {
				f2f8 := []*string{}
				for _, f2f8iter := range elem.Options.TransitGatewayCidrBlocks {
					var f2f8elem string
					f2f8elem = *f2f8iter
					f2f8 = append(f2f8, &f2f8elem)
				}
				f2.TransitGatewayCIDRBlocks = f2f8
			}
			if elem.Options.VpnEcmpSupport != nil {
				f2.VPNECMPSupport = elem.Options.VpnEcmpSupport
			}
			cr.Spec.ForProvider.Options = f2
		} else {
			cr.Spec.ForProvider.Options = nil
		}
		if elem.OwnerId != nil {
			cr.Status.AtProvider.OwnerID = elem.OwnerId
		} else {
			cr.Status.AtProvider.OwnerID = nil
		}
		if elem.State != nil {
			cr.Status.AtProvider.State = elem.State
		} else {
			cr.Status.AtProvider.State = nil
		}
		if elem.Tags != nil {
			f5 := []*svcapitypes.Tag{}
			for _, f5iter := range elem.Tags {
				f5elem := &svcapitypes.Tag{}
				if f5iter.Key != nil {
					f5elem.Key = f5iter.Key
				}
				if f5iter.Value != nil {
					f5elem.Value = f5iter.Value
				}
				f5 = append(f5, f5elem)
			}
			cr.Status.AtProvider.Tags = f5
		} else {
			cr.Status.AtProvider.Tags = nil
		}
		if elem.TransitGatewayArn != nil {
			cr.Status.AtProvider.TransitGatewayARN = elem.TransitGatewayArn
		} else {
			cr.Status.AtProvider.TransitGatewayARN = nil
		}
		if elem.TransitGatewayId != nil {
			cr.Status.AtProvider.TransitGatewayID = elem.TransitGatewayId
		} else {
			cr.Status.AtProvider.TransitGatewayID = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateTransitGatewayInput returns a create input.
func GenerateCreateTransitGatewayInput(cr *svcapitypes.TransitGateway) *svcsdk.CreateTransitGatewayInput {
	res := &svcsdk.CreateTransitGatewayInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.Options != nil {
		f1 := &svcsdk.TransitGatewayRequestOptions{}
		if cr.Spec.ForProvider.Options.AmazonSideASN != nil {
			f1.SetAmazonSideAsn(*cr.Spec.ForProvider.Options.AmazonSideASN)
		}
		if cr.Spec.ForProvider.Options.AutoAcceptSharedAttachments != nil {
			f1.SetAutoAcceptSharedAttachments(*cr.Spec.ForProvider.Options.AutoAcceptSharedAttachments)
		}
		if cr.Spec.ForProvider.Options.DefaultRouteTableAssociation != nil {
			f1.SetDefaultRouteTableAssociation(*cr.Spec.ForProvider.Options.DefaultRouteTableAssociation)
		}
		if cr.Spec.ForProvider.Options.DefaultRouteTablePropagation != nil {
			f1.SetDefaultRouteTablePropagation(*cr.Spec.ForProvider.Options.DefaultRouteTablePropagation)
		}
		if cr.Spec.ForProvider.Options.DNSSupport != nil {
			f1.SetDnsSupport(*cr.Spec.ForProvider.Options.DNSSupport)
		}
		if cr.Spec.ForProvider.Options.MulticastSupport != nil {
			f1.SetMulticastSupport(*cr.Spec.ForProvider.Options.MulticastSupport)
		}
		if cr.Spec.ForProvider.Options.TransitGatewayCIDRBlocks != nil {
			f1f6 := []*string{}
			for _, f1f6iter := range cr.Spec.ForProvider.Options.TransitGatewayCIDRBlocks {
				var f1f6elem string
				f1f6elem = *f1f6iter
				f1f6 = append(f1f6, &f1f6elem)
			}
			f1.SetTransitGatewayCidrBlocks(f1f6)
		}
		if cr.Spec.ForProvider.Options.VPNECMPSupport != nil {
			f1.SetVpnEcmpSupport(*cr.Spec.ForProvider.Options.VPNECMPSupport)
		}
		res.SetOptions(f1)
	}
	if cr.Spec.ForProvider.TagSpecifications != nil {
		f2 := []*svcsdk.TagSpecification{}
		for _, f2iter := range cr.Spec.ForProvider.TagSpecifications {
			f2elem := &svcsdk.TagSpecification{}
			if f2iter.ResourceType != nil {
				f2elem.SetResourceType(*f2iter.ResourceType)
			}
			if f2iter.Tags != nil {
				f2elemf1 := []*svcsdk.Tag{}
				for _, f2elemf1iter := range f2iter.Tags {
					f2elemf1elem := &svcsdk.Tag{}
					if f2elemf1iter.Key != nil {
						f2elemf1elem.SetKey(*f2elemf1iter.Key)
					}
					if f2elemf1iter.Value != nil {
						f2elemf1elem.SetValue(*f2elemf1iter.Value)
					}
					f2elemf1 = append(f2elemf1, f2elemf1elem)
				}
				f2elem.SetTags(f2elemf1)
			}
			f2 = append(f2, f2elem)
		}
		res.SetTagSpecifications(f2)
	}

	return res
}

// GenerateModifyTransitGatewayInput returns an update input.
func GenerateModifyTransitGatewayInput(cr *svcapitypes.TransitGateway) *svcsdk.ModifyTransitGatewayInput {
	res := &svcsdk.ModifyTransitGatewayInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.Options != nil {
		f2 := &svcsdk.ModifyTransitGatewayOptions{}
		if cr.Spec.ForProvider.Options.AmazonSideASN != nil {
			f2.SetAmazonSideAsn(*cr.Spec.ForProvider.Options.AmazonSideASN)
		}
		if cr.Spec.ForProvider.Options.AutoAcceptSharedAttachments != nil {
			f2.SetAutoAcceptSharedAttachments(*cr.Spec.ForProvider.Options.AutoAcceptSharedAttachments)
		}
		if cr.Spec.ForProvider.Options.DefaultRouteTableAssociation != nil {
			f2.SetDefaultRouteTableAssociation(*cr.Spec.ForProvider.Options.DefaultRouteTableAssociation)
		}
		if cr.Spec.ForProvider.Options.DefaultRouteTablePropagation != nil {
			f2.SetDefaultRouteTablePropagation(*cr.Spec.ForProvider.Options.DefaultRouteTablePropagation)
		}
		if cr.Spec.ForProvider.Options.DNSSupport != nil {
			f2.SetDnsSupport(*cr.Spec.ForProvider.Options.DNSSupport)
		}
		if cr.Spec.ForProvider.Options.VPNECMPSupport != nil {
			f2.SetVpnEcmpSupport(*cr.Spec.ForProvider.Options.VPNECMPSupport)
		}
		res.SetOptions(f2)
	}
	if cr.Status.AtProvider.TransitGatewayID != nil {
		res.SetTransitGatewayId(*cr.Status.AtProvider.TransitGatewayID)
	}

	return res
}

// GenerateDeleteTransitGatewayInput returns a deletion input.
func GenerateDeleteTransitGatewayInput(cr *svcapitypes.TransitGateway) *svcsdk.DeleteTransitGatewayInput {
	res := &svcsdk.DeleteTransitGatewayInput{}

	if cr.Status.AtProvider.TransitGatewayID != nil {
		res.SetTransitGatewayId(*cr.Status.AtProvider.TransitGatewayID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
