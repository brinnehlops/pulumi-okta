// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inline

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an inline hook.
// 
// This resource allows you to create and configure an inline hook.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/inline_hook.html.markdown.
type Hook struct {
	s *pulumi.ResourceState
}

// NewHook registers a new resource with the given unique name, arguments, and options.
func NewHook(ctx *pulumi.Context,
	name string, args *HookArgs, opts ...pulumi.ResourceOpt) (*Hook, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["auth"] = nil
		inputs["channel"] = nil
		inputs["headers"] = nil
		inputs["name"] = nil
		inputs["status"] = nil
		inputs["type"] = nil
		inputs["version"] = nil
	} else {
		inputs["auth"] = args.Auth
		inputs["channel"] = args.Channel
		inputs["headers"] = args.Headers
		inputs["name"] = args.Name
		inputs["status"] = args.Status
		inputs["type"] = args.Type
		inputs["version"] = args.Version
	}
	s, err := ctx.RegisterResource("okta:inline/hook:Hook", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Hook{s: s}, nil
}

// GetHook gets an existing Hook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHook(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HookState, opts ...pulumi.ResourceOpt) (*Hook, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["auth"] = state.Auth
		inputs["channel"] = state.Channel
		inputs["headers"] = state.Headers
		inputs["name"] = state.Name
		inputs["status"] = state.Status
		inputs["type"] = state.Type
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("okta:inline/hook:Hook", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Hook{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Hook) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Hook) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Authentication required for inline hook request.
func (r *Hook) Auth() *pulumi.Output {
	return r.s.State["auth"]
}

// Details of the endpoint the inline hook will hit.
func (r *Hook) Channel() *pulumi.Output {
	return r.s.State["channel"]
}

// Map of headers to send along in inline hook request.
func (r *Hook) Headers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["headers"])
}

// The inline hook display name.
func (r *Hook) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Hook) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The type of hook to trigger. Currently only `"HTTP"` is supported.
func (r *Hook) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// The version of the endpoint.
func (r *Hook) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Hook resources.
type HookState struct {
	// Authentication required for inline hook request.
	Auth interface{}
	// Details of the endpoint the inline hook will hit.
	Channel interface{}
	// Map of headers to send along in inline hook request.
	Headers interface{}
	// The inline hook display name.
	Name interface{}
	Status interface{}
	// The type of hook to trigger. Currently only `"HTTP"` is supported.
	Type interface{}
	// The version of the endpoint.
	Version interface{}
}

// The set of arguments for constructing a Hook resource.
type HookArgs struct {
	// Authentication required for inline hook request.
	Auth interface{}
	// Details of the endpoint the inline hook will hit.
	Channel interface{}
	// Map of headers to send along in inline hook request.
	Headers interface{}
	// The inline hook display name.
	Name interface{}
	Status interface{}
	// The type of hook to trigger. Currently only `"HTTP"` is supported.
	Type interface{}
	// The version of the endpoint.
	Version interface{}
}
