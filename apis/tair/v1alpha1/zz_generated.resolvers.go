// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1"
	common "github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Account.
func (mg *Account) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AuditLogConfig.
func (mg *AuditLogConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Connection.
func (mg *Connection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.InstanceIDRef,
		Selector:     mg.Spec.InitProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha1.SecurityGroupList{},
			Managed: &v1alpha1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VswitchID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.VswitchIDRef,
		Selector:     mg.Spec.ForProvider.VswitchIDSelector,
		To: reference.To{
			List:    &v1alpha11.VswitchList{},
			Managed: &v1alpha11.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VswitchID")
	}
	mg.Spec.ForProvider.VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VswitchIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha1.SecurityGroupList{},
			Managed: &v1alpha1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VswitchID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.VswitchIDRef,
		Selector:     mg.Spec.InitProvider.VswitchIDSelector,
		To: reference.To{
			List:    &v1alpha11.VswitchList{},
			Managed: &v1alpha11.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VswitchID")
	}
	mg.Spec.InitProvider.VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VswitchIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TairInstance.
func (mg *TairInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha1.SecurityGroupList{},
			Managed: &v1alpha1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha11.VPCList{},
			Managed: &v1alpha11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VswitchID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.VswitchIDRef,
		Selector:     mg.Spec.ForProvider.VswitchIDSelector,
		To: reference.To{
			List:    &v1alpha11.VswitchList{},
			Managed: &v1alpha11.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VswitchID")
	}
	mg.Spec.ForProvider.VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VswitchIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha1.SecurityGroupList{},
			Managed: &v1alpha1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha11.VPCList{},
			Managed: &v1alpha11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VswitchID),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.VswitchIDRef,
		Selector:     mg.Spec.InitProvider.VswitchIDSelector,
		To: reference.To{
			List:    &v1alpha11.VswitchList{},
			Managed: &v1alpha11.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VswitchID")
	}
	mg.Spec.InitProvider.VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VswitchIDRef = rsp.ResolvedReference

	return nil
}
