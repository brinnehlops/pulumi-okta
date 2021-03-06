// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an Authorization Server Scope.
//
// This resource allows you to create and configure an Authorization Server Scope.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-okta/sdk/v2/go/okta/auth"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := auth.NewServerScope(ctx, "example", &auth.ServerScopeArgs{
// 			AuthServerId:    pulumi.String("<auth server id>"),
// 			Consent:         pulumi.String("IMPLICIT"),
// 			MetadataPublish: pulumi.String("NO_CLIENTS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ServerScope struct {
	pulumi.CustomResourceState

	// Auth Server ID.
	AuthServerId pulumi.StringOutput `pulumi:"authServerId"`
	// Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
	Consent pulumi.StringPtrOutput `pulumi:"consent"`
	// A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
	Default pulumi.BoolPtrOutput `pulumi:"default"`
	// Description of the Auth Server Scope.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
	MetadataPublish pulumi.StringPtrOutput `pulumi:"metadataPublish"`
	// Auth Server scope name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewServerScope registers a new resource with the given unique name, arguments, and options.
func NewServerScope(ctx *pulumi.Context,
	name string, args *ServerScopeArgs, opts ...pulumi.ResourceOption) (*ServerScope, error) {
	if args == nil || args.AuthServerId == nil {
		return nil, errors.New("missing required argument 'AuthServerId'")
	}
	if args == nil {
		args = &ServerScopeArgs{}
	}
	var resource ServerScope
	err := ctx.RegisterResource("okta:auth/serverScope:ServerScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerScope gets an existing ServerScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerScopeState, opts ...pulumi.ResourceOption) (*ServerScope, error) {
	var resource ServerScope
	err := ctx.ReadResource("okta:auth/serverScope:ServerScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerScope resources.
type serverScopeState struct {
	// Auth Server ID.
	AuthServerId *string `pulumi:"authServerId"`
	// Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
	Consent *string `pulumi:"consent"`
	// A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
	Default *bool `pulumi:"default"`
	// Description of the Auth Server Scope.
	Description *string `pulumi:"description"`
	// Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
	MetadataPublish *string `pulumi:"metadataPublish"`
	// Auth Server scope name.
	Name *string `pulumi:"name"`
}

type ServerScopeState struct {
	// Auth Server ID.
	AuthServerId pulumi.StringPtrInput
	// Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
	Consent pulumi.StringPtrInput
	// A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
	Default pulumi.BoolPtrInput
	// Description of the Auth Server Scope.
	Description pulumi.StringPtrInput
	// Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
	MetadataPublish pulumi.StringPtrInput
	// Auth Server scope name.
	Name pulumi.StringPtrInput
}

func (ServerScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverScopeState)(nil)).Elem()
}

type serverScopeArgs struct {
	// Auth Server ID.
	AuthServerId string `pulumi:"authServerId"`
	// Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
	Consent *string `pulumi:"consent"`
	// A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
	Default *bool `pulumi:"default"`
	// Description of the Auth Server Scope.
	Description *string `pulumi:"description"`
	// Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
	MetadataPublish *string `pulumi:"metadataPublish"`
	// Auth Server scope name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ServerScope resource.
type ServerScopeArgs struct {
	// Auth Server ID.
	AuthServerId pulumi.StringInput
	// Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
	Consent pulumi.StringPtrInput
	// A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
	Default pulumi.BoolPtrInput
	// Description of the Auth Server Scope.
	Description pulumi.StringPtrInput
	// Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
	MetadataPublish pulumi.StringPtrInput
	// Auth Server scope name.
	Name pulumi.StringPtrInput
}

func (ServerScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverScopeArgs)(nil)).Elem()
}
