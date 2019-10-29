// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Authorization Server Policy.
// 
// This resource allows you to create and configure an Authorization Server Policy.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/auth_server_policy.html.markdown.
type ServerPolicy struct {
	s *pulumi.ResourceState
}

// NewServerPolicy registers a new resource with the given unique name, arguments, and options.
func NewServerPolicy(ctx *pulumi.Context,
	name string, args *ServerPolicyArgs, opts ...pulumi.ResourceOpt) (*ServerPolicy, error) {
	if args == nil || args.AuthServerId == nil {
		return nil, errors.New("missing required argument 'AuthServerId'")
	}
	if args == nil || args.ClientWhitelists == nil {
		return nil, errors.New("missing required argument 'ClientWhitelists'")
	}
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["authServerId"] = nil
		inputs["clientWhitelists"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["priority"] = nil
		inputs["status"] = nil
		inputs["type"] = nil
	} else {
		inputs["authServerId"] = args.AuthServerId
		inputs["clientWhitelists"] = args.ClientWhitelists
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["priority"] = args.Priority
		inputs["status"] = args.Status
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("okta:auth/serverPolicy:ServerPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServerPolicy{s: s}, nil
}

// GetServerPolicy gets an existing ServerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServerPolicyState, opts ...pulumi.ResourceOpt) (*ServerPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["authServerId"] = state.AuthServerId
		inputs["clientWhitelists"] = state.ClientWhitelists
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["priority"] = state.Priority
		inputs["status"] = state.Status
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("okta:auth/serverPolicy:ServerPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ServerPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ServerPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ServerPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the Auth Server.
func (r *ServerPolicy) AuthServerId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authServerId"])
}

// The clients to whitelist the policy for. `["ALL_CLIENTS"]` is a special value that can be used to whitelist for all clients. Otherwise it is a list of client ids.
func (r *ServerPolicy) ClientWhitelists() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["clientWhitelists"])
}

// The description of the Auth Server Policy.
func (r *ServerPolicy) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The name of the Auth Server Policy.
func (r *ServerPolicy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The priority of the Auth Server Policy.
func (r *ServerPolicy) Priority() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["priority"])
}

// The status of the Auth Server Policy.
func (r *ServerPolicy) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The type of the Auth Server Policy.
func (r *ServerPolicy) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering ServerPolicy resources.
type ServerPolicyState struct {
	// The ID of the Auth Server.
	AuthServerId interface{}
	// The clients to whitelist the policy for. `["ALL_CLIENTS"]` is a special value that can be used to whitelist for all clients. Otherwise it is a list of client ids.
	ClientWhitelists interface{}
	// The description of the Auth Server Policy.
	Description interface{}
	// The name of the Auth Server Policy.
	Name interface{}
	// The priority of the Auth Server Policy.
	Priority interface{}
	// The status of the Auth Server Policy.
	Status interface{}
	// The type of the Auth Server Policy.
	Type interface{}
}

// The set of arguments for constructing a ServerPolicy resource.
type ServerPolicyArgs struct {
	// The ID of the Auth Server.
	AuthServerId interface{}
	// The clients to whitelist the policy for. `["ALL_CLIENTS"]` is a special value that can be used to whitelist for all clients. Otherwise it is a list of client ids.
	ClientWhitelists interface{}
	// The description of the Auth Server Policy.
	Description interface{}
	// The name of the Auth Server Policy.
	Name interface{}
	// The priority of the Auth Server Policy.
	Priority interface{}
	// The status of the Auth Server Policy.
	Status interface{}
	// The type of the Auth Server Policy.
	Type interface{}
}
