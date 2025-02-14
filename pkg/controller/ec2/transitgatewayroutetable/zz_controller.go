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

package transitgatewayroutetable

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/ec2"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ec2/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an TransitGatewayRouteTable resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create TransitGatewayRouteTable in AWS"
	errUpdate        = "cannot update TransitGatewayRouteTable in AWS"
	errDescribe      = "failed to describe TransitGatewayRouteTable"
	errDelete        = "failed to delete TransitGatewayRouteTable"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.TransitGatewayRouteTable) (managed.TypedExternalClient[*svcapitypes.TransitGatewayRouteTable], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.TransitGatewayRouteTable) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeTransitGatewayRouteTablesInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeTransitGatewayRouteTablesWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.TransitGatewayRouteTables) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateTransitGatewayRouteTable(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, cr *svcapitypes.TransitGatewayRouteTable) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateTransitGatewayRouteTableInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateTransitGatewayRouteTableWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.TransitGatewayRouteTable.CreationTime != nil {
		cr.Status.AtProvider.CreationTime = &metav1.Time{*resp.TransitGatewayRouteTable.CreationTime}
	} else {
		cr.Status.AtProvider.CreationTime = nil
	}
	if resp.TransitGatewayRouteTable.DefaultAssociationRouteTable != nil {
		cr.Status.AtProvider.DefaultAssociationRouteTable = resp.TransitGatewayRouteTable.DefaultAssociationRouteTable
	} else {
		cr.Status.AtProvider.DefaultAssociationRouteTable = nil
	}
	if resp.TransitGatewayRouteTable.DefaultPropagationRouteTable != nil {
		cr.Status.AtProvider.DefaultPropagationRouteTable = resp.TransitGatewayRouteTable.DefaultPropagationRouteTable
	} else {
		cr.Status.AtProvider.DefaultPropagationRouteTable = nil
	}
	if resp.TransitGatewayRouteTable.State != nil {
		cr.Status.AtProvider.State = resp.TransitGatewayRouteTable.State
	} else {
		cr.Status.AtProvider.State = nil
	}
	if resp.TransitGatewayRouteTable.Tags != nil {
		f4 := []*svcapitypes.Tag{}
		for _, f4iter := range resp.TransitGatewayRouteTable.Tags {
			f4elem := &svcapitypes.Tag{}
			if f4iter.Key != nil {
				f4elem.Key = f4iter.Key
			}
			if f4iter.Value != nil {
				f4elem.Value = f4iter.Value
			}
			f4 = append(f4, f4elem)
		}
		cr.Status.AtProvider.Tags = f4
	} else {
		cr.Status.AtProvider.Tags = nil
	}
	if resp.TransitGatewayRouteTable.TransitGatewayId != nil {
		cr.Status.AtProvider.TransitGatewayID = resp.TransitGatewayRouteTable.TransitGatewayId
	} else {
		cr.Status.AtProvider.TransitGatewayID = nil
	}
	if resp.TransitGatewayRouteTable.TransitGatewayRouteTableId != nil {
		cr.Status.AtProvider.TransitGatewayRouteTableID = resp.TransitGatewayRouteTable.TransitGatewayRouteTableId
	} else {
		cr.Status.AtProvider.TransitGatewayRouteTableID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.TransitGatewayRouteTable) (managed.ExternalUpdate, error) {
	return e.update(ctx, cr)

}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.TransitGatewayRouteTable) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteTransitGatewayRouteTableInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteTransitGatewayRouteTableWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EC2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EC2API
	preObserve     func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DescribeTransitGatewayRouteTablesInput) error
	postObserve    func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DescribeTransitGatewayRouteTablesOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.TransitGatewayRouteTable, *svcsdk.DescribeTransitGatewayRouteTablesOutput) *svcsdk.DescribeTransitGatewayRouteTablesOutput
	lateInitialize func(*svcapitypes.TransitGatewayRouteTableParameters, *svcsdk.DescribeTransitGatewayRouteTablesOutput) error
	isUpToDate     func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DescribeTransitGatewayRouteTablesOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.CreateTransitGatewayRouteTableInput) error
	postCreate     func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.CreateTransitGatewayRouteTableOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DeleteTransitGatewayRouteTableInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DeleteTransitGatewayRouteTableOutput, error) (managed.ExternalDelete, error)
	update         func(context.Context, *svcapitypes.TransitGatewayRouteTable) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DescribeTransitGatewayRouteTablesInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.TransitGatewayRouteTable, _ *svcsdk.DescribeTransitGatewayRouteTablesOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.TransitGatewayRouteTable, list *svcsdk.DescribeTransitGatewayRouteTablesOutput) *svcsdk.DescribeTransitGatewayRouteTablesOutput {
	return list
}

func nopLateInitialize(*svcapitypes.TransitGatewayRouteTableParameters, *svcsdk.DescribeTransitGatewayRouteTablesOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DescribeTransitGatewayRouteTablesOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.CreateTransitGatewayRouteTableInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.TransitGatewayRouteTable, _ *svcsdk.CreateTransitGatewayRouteTableOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.TransitGatewayRouteTable, *svcsdk.DeleteTransitGatewayRouteTableInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.TransitGatewayRouteTable, _ *svcsdk.DeleteTransitGatewayRouteTableOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopUpdate(context.Context, *svcapitypes.TransitGatewayRouteTable) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
