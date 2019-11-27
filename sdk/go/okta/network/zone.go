// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Okta Network Zone.
// 
// This resource allows you to create and configure an Okta Network Zone.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/network_zone.html.markdown.
type Zone struct {
	s *pulumi.ResourceState
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOpt) (*Zone, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dynamicLocations"] = nil
		inputs["gateways"] = nil
		inputs["name"] = nil
		inputs["proxies"] = nil
		inputs["type"] = nil
	} else {
		inputs["dynamicLocations"] = args.DynamicLocations
		inputs["gateways"] = args.Gateways
		inputs["name"] = args.Name
		inputs["proxies"] = args.Proxies
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("okta:network/zone:Zone", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Zone{s: s}, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ZoneState, opts ...pulumi.ResourceOpt) (*Zone, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dynamicLocations"] = state.DynamicLocations
		inputs["gateways"] = state.Gateways
		inputs["name"] = state.Name
		inputs["proxies"] = state.Proxies
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("okta:network/zone:Zone", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Zone{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Zone) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Zone) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
func (r *Zone) DynamicLocations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["dynamicLocations"])
}

// Array of values in CIDR/range form.
func (r *Zone) Gateways() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["gateways"])
}

// Name of the Network Zone Resource.
func (r *Zone) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Array of values in CIDR/range form.
func (r *Zone) Proxies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["proxies"])
}

// Type of the Network Zone - can either be IP or DYNAMIC only.
func (r *Zone) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering Zone resources.
type ZoneState struct {
	// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
	DynamicLocations interface{}
	// Array of values in CIDR/range form.
	Gateways interface{}
	// Name of the Network Zone Resource.
	Name interface{}
	// Array of values in CIDR/range form.
	Proxies interface{}
	// Type of the Network Zone - can either be IP or DYNAMIC only.
	Type interface{}
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
	DynamicLocations interface{}
	// Array of values in CIDR/range form.
	Gateways interface{}
	// Name of the Network Zone Resource.
	Name interface{}
	// Array of values in CIDR/range form.
	Proxies interface{}
	// Type of the Network Zone - can either be IP or DYNAMIC only.
	Type interface{}
}
