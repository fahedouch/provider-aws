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

package publicdnsnamespace

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/servicediscovery"
	svcsdk "github.com/aws/aws-sdk-go/service/servicediscovery"
	svcsdkapi "github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/servicediscovery/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an PublicDNSNamespace resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create PublicDNSNamespace in AWS"
	errUpdate        = "cannot update PublicDNSNamespace in AWS"
	errDescribe      = "failed to describe PublicDNSNamespace"
	errDelete        = "failed to delete PublicDNSNamespace"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.PublicDNSNamespace) (managed.TypedExternalClient[*svcapitypes.PublicDNSNamespace], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.PublicDNSNamespace) (managed.ExternalObservation, error) {
	return e.observe(ctx, cr)
}

func (e *external) Create(ctx context.Context, cr *svcapitypes.PublicDNSNamespace) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreatePublicDnsNamespaceInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreatePublicDnsNamespaceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.OperationId != nil {
		cr.Status.AtProvider.OperationID = resp.OperationId
	} else {
		cr.Status.AtProvider.OperationID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.PublicDNSNamespace) (managed.ExternalUpdate, error) {
	input := GenerateUpdatePublicDnsNamespaceInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdatePublicDnsNamespaceWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.PublicDNSNamespace) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	return e.delete(ctx, cr)

}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ServiceDiscoveryAPI, opts []option) *external {
	e := &external{
		kube:       kube,
		client:     client,
		observe:    nopObserve,
		preCreate:  nopPreCreate,
		postCreate: nopPostCreate,
		delete:     nopDelete,
		preUpdate:  nopPreUpdate,
		postUpdate: nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube       client.Client
	client     svcsdkapi.ServiceDiscoveryAPI
	observe    func(context.Context, *svcapitypes.PublicDNSNamespace) (managed.ExternalObservation, error)
	preCreate  func(context.Context, *svcapitypes.PublicDNSNamespace, *svcsdk.CreatePublicDnsNamespaceInput) error
	postCreate func(context.Context, *svcapitypes.PublicDNSNamespace, *svcsdk.CreatePublicDnsNamespaceOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	delete     func(context.Context, *svcapitypes.PublicDNSNamespace) (managed.ExternalDelete, error)
	preUpdate  func(context.Context, *svcapitypes.PublicDNSNamespace, *svcsdk.UpdatePublicDnsNamespaceInput) error
	postUpdate func(context.Context, *svcapitypes.PublicDNSNamespace, *svcsdk.UpdatePublicDnsNamespaceOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopObserve(context.Context, *svcapitypes.PublicDNSNamespace) (managed.ExternalObservation, error) {
	return managed.ExternalObservation{}, nil
}

func nopPreCreate(context.Context, *svcapitypes.PublicDNSNamespace, *svcsdk.CreatePublicDnsNamespaceInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.PublicDNSNamespace, _ *svcsdk.CreatePublicDnsNamespaceOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopDelete(context.Context, *svcapitypes.PublicDNSNamespace) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, nil
}
func nopPreUpdate(context.Context, *svcapitypes.PublicDNSNamespace, *svcsdk.UpdatePublicDnsNamespaceInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.PublicDNSNamespace, _ *svcsdk.UpdatePublicDnsNamespaceOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
