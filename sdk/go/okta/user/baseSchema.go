// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a User Base Schema property.
// 
// This resource allows you to configure a base user schema property.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/user_base_schema.html.markdown.
type BaseSchema struct {
	s *pulumi.ResourceState
}

// NewBaseSchema registers a new resource with the given unique name, arguments, and options.
func NewBaseSchema(ctx *pulumi.Context,
	name string, args *BaseSchemaArgs, opts ...pulumi.ResourceOpt) (*BaseSchema, error) {
	if args == nil || args.Index == nil {
		return nil, errors.New("missing required argument 'Index'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["index"] = nil
		inputs["master"] = nil
		inputs["permissions"] = nil
		inputs["required"] = nil
		inputs["title"] = nil
		inputs["type"] = nil
	} else {
		inputs["index"] = args.Index
		inputs["master"] = args.Master
		inputs["permissions"] = args.Permissions
		inputs["required"] = args.Required
		inputs["title"] = args.Title
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("okta:user/baseSchema:BaseSchema", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BaseSchema{s: s}, nil
}

// GetBaseSchema gets an existing BaseSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaseSchema(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BaseSchemaState, opts ...pulumi.ResourceOpt) (*BaseSchema, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["index"] = state.Index
		inputs["master"] = state.Master
		inputs["permissions"] = state.Permissions
		inputs["required"] = state.Required
		inputs["title"] = state.Title
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("okta:user/baseSchema:BaseSchema", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BaseSchema{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BaseSchema) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BaseSchema) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The property name.
func (r *BaseSchema) Index() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["index"])
}

// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
func (r *BaseSchema) Master() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["master"])
}

// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
func (r *BaseSchema) Permissions() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["permissions"])
}

// Whether the property is required for this application's users.
func (r *BaseSchema) Required() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["required"])
}

// The property display name.
func (r *BaseSchema) Title() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["title"])
}

// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
func (r *BaseSchema) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering BaseSchema resources.
type BaseSchemaState struct {
	// The property name.
	Index interface{}
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master interface{}
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions interface{}
	// Whether the property is required for this application's users.
	Required interface{}
	// The property display name.
	Title interface{}
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type interface{}
}

// The set of arguments for constructing a BaseSchema resource.
type BaseSchemaArgs struct {
	// The property name.
	Index interface{}
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master interface{}
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions interface{}
	// Whether the property is required for this application's users.
	Required interface{}
	// The property display name.
	Title interface{}
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type interface{}
}
