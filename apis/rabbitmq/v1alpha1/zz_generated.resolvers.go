// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	common "github.com/evaneos/provider-rabbitmq/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Binding.
func (mg *Binding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Source),
		Extract:      common.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.SourceRef,
		Selector:     mg.Spec.ForProvider.SourceSelector,
		To: reference.To{
			List:    &ExchangeList{},
			Managed: &Exchange{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Source")
	}
	mg.Spec.ForProvider.Source = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Source),
		Extract:      common.ExtractResourceName(),
		Reference:    mg.Spec.InitProvider.SourceRef,
		Selector:     mg.Spec.InitProvider.SourceSelector,
		To: reference.To{
			List:    &ExchangeList{},
			Managed: &Exchange{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Source")
	}
	mg.Spec.InitProvider.Source = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SourceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Exchange.
func (mg *Exchange) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FederationUpstream.
func (mg *FederationUpstream) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OperatorPolicy.
func (mg *OperatorPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Permissions.
func (mg *Permissions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserRef,
		Selector:     mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.User),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.UserRef,
		Selector:     mg.Spec.InitProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.User")
	}
	mg.Spec.InitProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Shovel.
func (mg *Shovel) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicPermissions.
func (mg *TopicPermissions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VhostRef,
		Selector:     mg.Spec.ForProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Vhost")
	}
	mg.Spec.ForProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VhostRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Vhost),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VhostRef,
		Selector:     mg.Spec.InitProvider.VhostSelector,
		To: reference.To{
			List:    &VhostList{},
			Managed: &Vhost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Vhost")
	}
	mg.Spec.InitProvider.Vhost = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VhostRef = rsp.ResolvedReference

	return nil
}
