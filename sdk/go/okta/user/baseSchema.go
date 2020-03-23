// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package user

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a User Base Schema property.
//
// This resource allows you to configure a base user schema property.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/user_base_schema.html.markdown.
type BaseSchema struct {
	pulumi.CustomResourceState

	// The property name.
	Index pulumi.StringOutput `pulumi:"index"`
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master pulumi.StringPtrOutput `pulumi:"master"`
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions pulumi.StringPtrOutput `pulumi:"permissions"`
	// Whether the property is required for this application's users.
	Required pulumi.BoolPtrOutput `pulumi:"required"`
	// The property display name.
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBaseSchema registers a new resource with the given unique name, arguments, and options.
func NewBaseSchema(ctx *pulumi.Context,
	name string, args *BaseSchemaArgs, opts ...pulumi.ResourceOption) (*BaseSchema, error) {
	if args == nil || args.Index == nil {
		return nil, errors.New("missing required argument 'Index'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &BaseSchemaArgs{}
	}
	var resource BaseSchema
	err := ctx.RegisterResource("okta:user/baseSchema:BaseSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBaseSchema gets an existing BaseSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaseSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BaseSchemaState, opts ...pulumi.ResourceOption) (*BaseSchema, error) {
	var resource BaseSchema
	err := ctx.ReadResource("okta:user/baseSchema:BaseSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BaseSchema resources.
type baseSchemaState struct {
	// The property name.
	Index *string `pulumi:"index"`
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master *string `pulumi:"master"`
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions *string `pulumi:"permissions"`
	// Whether the property is required for this application's users.
	Required *bool `pulumi:"required"`
	// The property display name.
	Title *string `pulumi:"title"`
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type *string `pulumi:"type"`
}

type BaseSchemaState struct {
	// The property name.
	Index pulumi.StringPtrInput
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master pulumi.StringPtrInput
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions pulumi.StringPtrInput
	// Whether the property is required for this application's users.
	Required pulumi.BoolPtrInput
	// The property display name.
	Title pulumi.StringPtrInput
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type pulumi.StringPtrInput
}

func (BaseSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*baseSchemaState)(nil)).Elem()
}

type baseSchemaArgs struct {
	// The property name.
	Index string `pulumi:"index"`
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master *string `pulumi:"master"`
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions *string `pulumi:"permissions"`
	// Whether the property is required for this application's users.
	Required *bool `pulumi:"required"`
	// The property display name.
	Title string `pulumi:"title"`
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a BaseSchema resource.
type BaseSchemaArgs struct {
	// The property name.
	Index pulumi.StringInput
	// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
	Master pulumi.StringPtrInput
	// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
	Permissions pulumi.StringPtrInput
	// Whether the property is required for this application's users.
	Required pulumi.BoolPtrInput
	// The property display name.
	Title pulumi.StringInput
	// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
	Type pulumi.StringInput
}

func (BaseSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*baseSchemaArgs)(nil)).Elem()
}

