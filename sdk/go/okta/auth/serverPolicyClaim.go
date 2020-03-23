// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package auth

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Authorization Server Policy Rule.
//
// This resource allows you to create and configure an Authorization Server Policy Rule.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/auth_server_policy_rule.html.markdown.
type ServerPolicyClaim struct {
	pulumi.CustomResourceState

	// Lifetime of access token. Can be set to a value between 5 and 1440.
	AccessTokenLifetimeMinutes pulumi.IntPtrOutput `pulumi:"accessTokenLifetimeMinutes"`
	// Auth Server ID.
	AuthServerId pulumi.StringOutput `pulumi:"authServerId"`
	// Accepted grant type values, `"authorizationCode"`, `"implicit"`, `"password"`
	GrantTypeWhitelists pulumi.StringArrayOutput `pulumi:"grantTypeWhitelists"`
	GroupBlacklists pulumi.StringArrayOutput `pulumi:"groupBlacklists"`
	GroupWhitelists pulumi.StringArrayOutput `pulumi:"groupWhitelists"`
	// The ID of the inline token to trigger.
	InlineHookId pulumi.StringPtrOutput `pulumi:"inlineHookId"`
	// Auth Server Policy Rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Auth Server Policy ID.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// Priority of the auth server policy rule.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Lifetime of refresh token.
	RefreshTokenLifetimeMinutes pulumi.IntPtrOutput `pulumi:"refreshTokenLifetimeMinutes"`
	RefreshTokenWindowMinutes pulumi.IntPtrOutput `pulumi:"refreshTokenWindowMinutes"`
	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
	ScopeWhitelists pulumi.StringArrayOutput `pulumi:"scopeWhitelists"`
	// The status of the Auth Server Policy Rule.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The type of the Auth Server Policy Rule.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	UserBlacklists pulumi.StringArrayOutput `pulumi:"userBlacklists"`
	UserWhitelists pulumi.StringArrayOutput `pulumi:"userWhitelists"`
}

// NewServerPolicyClaim registers a new resource with the given unique name, arguments, and options.
func NewServerPolicyClaim(ctx *pulumi.Context,
	name string, args *ServerPolicyClaimArgs, opts ...pulumi.ResourceOption) (*ServerPolicyClaim, error) {
	if args == nil || args.AuthServerId == nil {
		return nil, errors.New("missing required argument 'AuthServerId'")
	}
	if args == nil || args.GrantTypeWhitelists == nil {
		return nil, errors.New("missing required argument 'GrantTypeWhitelists'")
	}
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil {
		args = &ServerPolicyClaimArgs{}
	}
	var resource ServerPolicyClaim
	err := ctx.RegisterResource("okta:auth/serverPolicyClaim:ServerPolicyClaim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerPolicyClaim gets an existing ServerPolicyClaim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerPolicyClaim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerPolicyClaimState, opts ...pulumi.ResourceOption) (*ServerPolicyClaim, error) {
	var resource ServerPolicyClaim
	err := ctx.ReadResource("okta:auth/serverPolicyClaim:ServerPolicyClaim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerPolicyClaim resources.
type serverPolicyClaimState struct {
	// Lifetime of access token. Can be set to a value between 5 and 1440.
	AccessTokenLifetimeMinutes *int `pulumi:"accessTokenLifetimeMinutes"`
	// Auth Server ID.
	AuthServerId *string `pulumi:"authServerId"`
	// Accepted grant type values, `"authorizationCode"`, `"implicit"`, `"password"`
	GrantTypeWhitelists []string `pulumi:"grantTypeWhitelists"`
	GroupBlacklists []string `pulumi:"groupBlacklists"`
	GroupWhitelists []string `pulumi:"groupWhitelists"`
	// The ID of the inline token to trigger.
	InlineHookId *string `pulumi:"inlineHookId"`
	// Auth Server Policy Rule name.
	Name *string `pulumi:"name"`
	// Auth Server Policy ID.
	PolicyId *string `pulumi:"policyId"`
	// Priority of the auth server policy rule.
	Priority *int `pulumi:"priority"`
	// Lifetime of refresh token.
	RefreshTokenLifetimeMinutes *int `pulumi:"refreshTokenLifetimeMinutes"`
	RefreshTokenWindowMinutes *int `pulumi:"refreshTokenWindowMinutes"`
	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
	ScopeWhitelists []string `pulumi:"scopeWhitelists"`
	// The status of the Auth Server Policy Rule.
	Status *string `pulumi:"status"`
	// The type of the Auth Server Policy Rule.
	Type *string `pulumi:"type"`
	UserBlacklists []string `pulumi:"userBlacklists"`
	UserWhitelists []string `pulumi:"userWhitelists"`
}

type ServerPolicyClaimState struct {
	// Lifetime of access token. Can be set to a value between 5 and 1440.
	AccessTokenLifetimeMinutes pulumi.IntPtrInput
	// Auth Server ID.
	AuthServerId pulumi.StringPtrInput
	// Accepted grant type values, `"authorizationCode"`, `"implicit"`, `"password"`
	GrantTypeWhitelists pulumi.StringArrayInput
	GroupBlacklists pulumi.StringArrayInput
	GroupWhitelists pulumi.StringArrayInput
	// The ID of the inline token to trigger.
	InlineHookId pulumi.StringPtrInput
	// Auth Server Policy Rule name.
	Name pulumi.StringPtrInput
	// Auth Server Policy ID.
	PolicyId pulumi.StringPtrInput
	// Priority of the auth server policy rule.
	Priority pulumi.IntPtrInput
	// Lifetime of refresh token.
	RefreshTokenLifetimeMinutes pulumi.IntPtrInput
	RefreshTokenWindowMinutes pulumi.IntPtrInput
	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
	ScopeWhitelists pulumi.StringArrayInput
	// The status of the Auth Server Policy Rule.
	Status pulumi.StringPtrInput
	// The type of the Auth Server Policy Rule.
	Type pulumi.StringPtrInput
	UserBlacklists pulumi.StringArrayInput
	UserWhitelists pulumi.StringArrayInput
}

func (ServerPolicyClaimState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverPolicyClaimState)(nil)).Elem()
}

type serverPolicyClaimArgs struct {
	// Lifetime of access token. Can be set to a value between 5 and 1440.
	AccessTokenLifetimeMinutes *int `pulumi:"accessTokenLifetimeMinutes"`
	// Auth Server ID.
	AuthServerId string `pulumi:"authServerId"`
	// Accepted grant type values, `"authorizationCode"`, `"implicit"`, `"password"`
	GrantTypeWhitelists []string `pulumi:"grantTypeWhitelists"`
	GroupBlacklists []string `pulumi:"groupBlacklists"`
	GroupWhitelists []string `pulumi:"groupWhitelists"`
	// The ID of the inline token to trigger.
	InlineHookId *string `pulumi:"inlineHookId"`
	// Auth Server Policy Rule name.
	Name *string `pulumi:"name"`
	// Auth Server Policy ID.
	PolicyId string `pulumi:"policyId"`
	// Priority of the auth server policy rule.
	Priority int `pulumi:"priority"`
	// Lifetime of refresh token.
	RefreshTokenLifetimeMinutes *int `pulumi:"refreshTokenLifetimeMinutes"`
	RefreshTokenWindowMinutes *int `pulumi:"refreshTokenWindowMinutes"`
	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
	ScopeWhitelists []string `pulumi:"scopeWhitelists"`
	// The status of the Auth Server Policy Rule.
	Status *string `pulumi:"status"`
	// The type of the Auth Server Policy Rule.
	Type *string `pulumi:"type"`
	UserBlacklists []string `pulumi:"userBlacklists"`
	UserWhitelists []string `pulumi:"userWhitelists"`
}

// The set of arguments for constructing a ServerPolicyClaim resource.
type ServerPolicyClaimArgs struct {
	// Lifetime of access token. Can be set to a value between 5 and 1440.
	AccessTokenLifetimeMinutes pulumi.IntPtrInput
	// Auth Server ID.
	AuthServerId pulumi.StringInput
	// Accepted grant type values, `"authorizationCode"`, `"implicit"`, `"password"`
	GrantTypeWhitelists pulumi.StringArrayInput
	GroupBlacklists pulumi.StringArrayInput
	GroupWhitelists pulumi.StringArrayInput
	// The ID of the inline token to trigger.
	InlineHookId pulumi.StringPtrInput
	// Auth Server Policy Rule name.
	Name pulumi.StringPtrInput
	// Auth Server Policy ID.
	PolicyId pulumi.StringInput
	// Priority of the auth server policy rule.
	Priority pulumi.IntInput
	// Lifetime of refresh token.
	RefreshTokenLifetimeMinutes pulumi.IntPtrInput
	RefreshTokenWindowMinutes pulumi.IntPtrInput
	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
	ScopeWhitelists pulumi.StringArrayInput
	// The status of the Auth Server Policy Rule.
	Status pulumi.StringPtrInput
	// The type of the Auth Server Policy Rule.
	Type pulumi.StringPtrInput
	UserBlacklists pulumi.StringArrayInput
	UserWhitelists pulumi.StringArrayInput
}

func (ServerPolicyClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverPolicyClaimArgs)(nil)).Elem()
}

