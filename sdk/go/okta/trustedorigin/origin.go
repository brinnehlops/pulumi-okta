// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trustedorigin

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Trusted Origin.
// 
// This resource allows you to create and configure an Trusted Origin.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/trusted_origin.html.markdown.
type Origin struct {
	s *pulumi.ResourceState
}

// NewOrigin registers a new resource with the given unique name, arguments, and options.
func NewOrigin(ctx *pulumi.Context,
	name string, args *OriginArgs, opts ...pulumi.ResourceOpt) (*Origin, error) {
	if args == nil || args.Origin == nil {
		return nil, errors.New("missing required argument 'Origin'")
	}
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["active"] = nil
		inputs["name"] = nil
		inputs["origin"] = nil
		inputs["scopes"] = nil
	} else {
		inputs["active"] = args.Active
		inputs["name"] = args.Name
		inputs["origin"] = args.Origin
		inputs["scopes"] = args.Scopes
	}
	s, err := ctx.RegisterResource("okta:trustedorigin/origin:Origin", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Origin{s: s}, nil
}

// GetOrigin gets an existing Origin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrigin(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OriginState, opts ...pulumi.ResourceOpt) (*Origin, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["active"] = state.Active
		inputs["name"] = state.Name
		inputs["origin"] = state.Origin
		inputs["scopes"] = state.Scopes
	}
	s, err := ctx.ReadResource("okta:trustedorigin/origin:Origin", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Origin{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Origin) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Origin) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Whether the Trusted Origin is active or not - can only be issued post-creation.
func (r *Origin) Active() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["active"])
}

// Name of the Trusted Origin Resource.
func (r *Origin) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The origin to trust.
func (r *Origin) Origin() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["origin"])
}

// Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
func (r *Origin) Scopes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["scopes"])
}

// Input properties used for looking up and filtering Origin resources.
type OriginState struct {
	// Whether the Trusted Origin is active or not - can only be issued post-creation.
	Active interface{}
	// Name of the Trusted Origin Resource.
	Name interface{}
	// The origin to trust.
	Origin interface{}
	// Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
	Scopes interface{}
}

// The set of arguments for constructing a Origin resource.
type OriginArgs struct {
	// Whether the Trusted Origin is active or not - can only be issued post-creation.
	Active interface{}
	// Name of the Trusted Origin Resource.
	Name interface{}
	// The origin to trust.
	Origin interface{}
	// Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
	Scopes interface{}
}
