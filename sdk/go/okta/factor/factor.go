// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package factor

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows you to manage the activation of Okta MFA methods.
// 
// This resource allows you to manage Okta MFA methods.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/factor.html.markdown.
type Factor struct {
	s *pulumi.ResourceState
}

// NewFactor registers a new resource with the given unique name, arguments, and options.
func NewFactor(ctx *pulumi.Context,
	name string, args *FactorArgs, opts ...pulumi.ResourceOpt) (*Factor, error) {
	if args == nil || args.ProviderId == nil {
		return nil, errors.New("missing required argument 'ProviderId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["active"] = nil
		inputs["providerId"] = nil
	} else {
		inputs["active"] = args.Active
		inputs["providerId"] = args.ProviderId
	}
	s, err := ctx.RegisterResource("okta:factor/factor:Factor", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Factor{s: s}, nil
}

// GetFactor gets an existing Factor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFactor(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FactorState, opts ...pulumi.ResourceOpt) (*Factor, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["active"] = state.Active
		inputs["providerId"] = state.ProviderId
	}
	s, err := ctx.ReadResource("okta:factor/factor:Factor", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Factor{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Factor) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Factor) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Whether or not to activate the provider, by default it is set to `true`.
func (r *Factor) Active() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["active"])
}

// Factor provider ID
func (r *Factor) ProviderId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["providerId"])
}

// Input properties used for looking up and filtering Factor resources.
type FactorState struct {
	// Whether or not to activate the provider, by default it is set to `true`.
	Active interface{}
	// Factor provider ID
	ProviderId interface{}
}

// The set of arguments for constructing a Factor resource.
type FactorArgs struct {
	// Whether or not to activate the provider, by default it is set to `true`.
	Active interface{}
	// Factor provider ID
	ProviderId interface{}
}
