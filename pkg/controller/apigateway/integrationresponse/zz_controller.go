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

package integrationresponse

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigateway/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an IntegrationResponse resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create IntegrationResponse in AWS"
	errUpdate        = "cannot update IntegrationResponse in AWS"
	errDescribe      = "failed to describe IntegrationResponse"
	errDelete        = "failed to delete IntegrationResponse"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.IntegrationResponse) (managed.TypedExternalClient[*svcapitypes.IntegrationResponse], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.IntegrationResponse) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetIntegrationResponseInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetIntegrationResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateIntegrationResponse(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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

func (e *external) Create(ctx context.Context, cr *svcapitypes.IntegrationResponse) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GeneratePutIntegrationResponseInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.PutIntegrationResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.ContentHandling != nil {
		cr.Spec.ForProvider.ContentHandling = resp.ContentHandling
	} else {
		cr.Spec.ForProvider.ContentHandling = nil
	}
	if resp.ResponseParameters != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range resp.ResponseParameters {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		cr.Spec.ForProvider.ResponseParameters = f1
	} else {
		cr.Spec.ForProvider.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range resp.ResponseTemplates {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		cr.Spec.ForProvider.ResponseTemplates = f2
	} else {
		cr.Spec.ForProvider.ResponseTemplates = nil
	}
	if resp.SelectionPattern != nil {
		cr.Spec.ForProvider.SelectionPattern = resp.SelectionPattern
	} else {
		cr.Spec.ForProvider.SelectionPattern = nil
	}
	if resp.StatusCode != nil {
		cr.Spec.ForProvider.StatusCode = resp.StatusCode
	} else {
		cr.Spec.ForProvider.StatusCode = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.IntegrationResponse) (managed.ExternalUpdate, error) {
	input := GenerateUpdateIntegrationResponseInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateIntegrationResponseWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.IntegrationResponse) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteIntegrationResponseInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteIntegrationResponseWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.APIGatewayAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.APIGatewayAPI
	preObserve     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseInput) error
	postObserve    func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.IntegrationResponse, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.IntegrationResponseParameters, *svcsdk.IntegrationResponse) error
	isUpToDate     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.IntegrationResponse) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.PutIntegrationResponseInput) error
	postCreate     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.IntegrationResponse, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.DeleteIntegrationResponseInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.DeleteIntegrationResponseOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.UpdateIntegrationResponseInput) error
	postUpdate     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.IntegrationResponse, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.IntegrationResponse, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.IntegrationResponseParameters, *svcsdk.IntegrationResponse) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.IntegrationResponse) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.PutIntegrationResponseInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.IntegrationResponse, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.DeleteIntegrationResponseInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.DeleteIntegrationResponseOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.UpdateIntegrationResponseInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.IntegrationResponse, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
