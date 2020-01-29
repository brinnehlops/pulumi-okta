// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Sign On Policy Rule.
// 
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_signon.html.markdown.
type RuleSignon struct {
	pulumi.CustomResourceState

	// Allow or deny access based on the rule conditions: `"ALLOW"` or `"DENY"`. The default is `"ALLOW"`.
	Access pulumi.StringPtrOutput `pulumi:"access"`
	// Authentication entrypoint: `"ANY"` or `"RADIUS"`.
	Authtype pulumi.StringPtrOutput `pulumi:"authtype"`
	// Elapsed time before the next MFA challenge.
	MfaLifetime pulumi.IntPtrOutput `pulumi:"mfaLifetime"`
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: `"DEVICE"`, `"SESSION"` or `"ALWAYS"`.
	MfaPrompt pulumi.StringPtrOutput `pulumi:"mfaPrompt"`
	// Remember MFA device. The default `false`.
	MfaRememberDevice pulumi.BoolPtrOutput `pulumi:"mfaRememberDevice"`
	// Require MFA. By default is `false`.
	MfaRequired pulumi.BoolPtrOutput `pulumi:"mfaRequired"`
	// Policy Rule Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection pulumi.StringPtrOutput `pulumi:"networkConnection"`
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes pulumi.StringArrayOutput `pulumi:"networkExcludes"`
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes pulumi.StringArrayOutput `pulumi:"networkIncludes"`
	// Policy ID.
	Policyid pulumi.StringOutput `pulumi:"policyid"`
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Max minutes a session can be idle.",
	SessionIdle pulumi.IntPtrOutput `pulumi:"sessionIdle"`
	// Max minutes a session is active: Disable = 0.
	SessionLifetime pulumi.IntPtrOutput `pulumi:"sessionLifetime"`
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.
	SessionPersistent pulumi.BoolPtrOutput `pulumi:"sessionPersistent"`
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Set of User IDs to Exclude
	UsersExcludeds pulumi.StringArrayOutput `pulumi:"usersExcludeds"`
}

// NewRuleSignon registers a new resource with the given unique name, arguments, and options.
func NewRuleSignon(ctx *pulumi.Context,
	name string, args *RuleSignonArgs, opts ...pulumi.ResourceOption) (*RuleSignon, error) {
	if args == nil || args.Policyid == nil {
		return nil, errors.New("missing required argument 'Policyid'")
	}
	if args == nil {
		args = &RuleSignonArgs{}
	}
	var resource RuleSignon
	err := ctx.RegisterResource("okta:policy/ruleSignon:RuleSignon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleSignon gets an existing RuleSignon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleSignon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleSignonState, opts ...pulumi.ResourceOption) (*RuleSignon, error) {
	var resource RuleSignon
	err := ctx.ReadResource("okta:policy/ruleSignon:RuleSignon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleSignon resources.
type ruleSignonState struct {
	// Allow or deny access based on the rule conditions: `"ALLOW"` or `"DENY"`. The default is `"ALLOW"`.
	Access *string `pulumi:"access"`
	// Authentication entrypoint: `"ANY"` or `"RADIUS"`.
	Authtype *string `pulumi:"authtype"`
	// Elapsed time before the next MFA challenge.
	MfaLifetime *int `pulumi:"mfaLifetime"`
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: `"DEVICE"`, `"SESSION"` or `"ALWAYS"`.
	MfaPrompt *string `pulumi:"mfaPrompt"`
	// Remember MFA device. The default `false`.
	MfaRememberDevice *bool `pulumi:"mfaRememberDevice"`
	// Require MFA. By default is `false`.
	MfaRequired *bool `pulumi:"mfaRequired"`
	// Policy Rule Name.
	Name *string `pulumi:"name"`
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection *string `pulumi:"networkConnection"`
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes []string `pulumi:"networkExcludes"`
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes []string `pulumi:"networkIncludes"`
	// Policy ID.
	Policyid *string `pulumi:"policyid"`
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority *int `pulumi:"priority"`
	// Max minutes a session can be idle.",
	SessionIdle *int `pulumi:"sessionIdle"`
	// Max minutes a session is active: Disable = 0.
	SessionLifetime *int `pulumi:"sessionLifetime"`
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.
	SessionPersistent *bool `pulumi:"sessionPersistent"`
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status *string `pulumi:"status"`
	// Set of User IDs to Exclude
	UsersExcludeds []string `pulumi:"usersExcludeds"`
}

type RuleSignonState struct {
	// Allow or deny access based on the rule conditions: `"ALLOW"` or `"DENY"`. The default is `"ALLOW"`.
	Access pulumi.StringPtrInput
	// Authentication entrypoint: `"ANY"` or `"RADIUS"`.
	Authtype pulumi.StringPtrInput
	// Elapsed time before the next MFA challenge.
	MfaLifetime pulumi.IntPtrInput
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: `"DEVICE"`, `"SESSION"` or `"ALWAYS"`.
	MfaPrompt pulumi.StringPtrInput
	// Remember MFA device. The default `false`.
	MfaRememberDevice pulumi.BoolPtrInput
	// Require MFA. By default is `false`.
	MfaRequired pulumi.BoolPtrInput
	// Policy Rule Name.
	Name pulumi.StringPtrInput
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection pulumi.StringPtrInput
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes pulumi.StringArrayInput
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes pulumi.StringArrayInput
	// Policy ID.
	Policyid pulumi.StringPtrInput
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	// Max minutes a session can be idle.",
	SessionIdle pulumi.IntPtrInput
	// Max minutes a session is active: Disable = 0.
	SessionLifetime pulumi.IntPtrInput
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.
	SessionPersistent pulumi.BoolPtrInput
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status pulumi.StringPtrInput
	// Set of User IDs to Exclude
	UsersExcludeds pulumi.StringArrayInput
}

func (RuleSignonState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleSignonState)(nil)).Elem()
}

type ruleSignonArgs struct {
	// Allow or deny access based on the rule conditions: `"ALLOW"` or `"DENY"`. The default is `"ALLOW"`.
	Access *string `pulumi:"access"`
	// Authentication entrypoint: `"ANY"` or `"RADIUS"`.
	Authtype *string `pulumi:"authtype"`
	// Elapsed time before the next MFA challenge.
	MfaLifetime *int `pulumi:"mfaLifetime"`
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: `"DEVICE"`, `"SESSION"` or `"ALWAYS"`.
	MfaPrompt *string `pulumi:"mfaPrompt"`
	// Remember MFA device. The default `false`.
	MfaRememberDevice *bool `pulumi:"mfaRememberDevice"`
	// Require MFA. By default is `false`.
	MfaRequired *bool `pulumi:"mfaRequired"`
	// Policy Rule Name.
	Name *string `pulumi:"name"`
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection *string `pulumi:"networkConnection"`
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes []string `pulumi:"networkExcludes"`
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes []string `pulumi:"networkIncludes"`
	// Policy ID.
	Policyid string `pulumi:"policyid"`
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority *int `pulumi:"priority"`
	// Max minutes a session can be idle.",
	SessionIdle *int `pulumi:"sessionIdle"`
	// Max minutes a session is active: Disable = 0.
	SessionLifetime *int `pulumi:"sessionLifetime"`
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.
	SessionPersistent *bool `pulumi:"sessionPersistent"`
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status *string `pulumi:"status"`
	// Set of User IDs to Exclude
	UsersExcludeds []string `pulumi:"usersExcludeds"`
}

// The set of arguments for constructing a RuleSignon resource.
type RuleSignonArgs struct {
	// Allow or deny access based on the rule conditions: `"ALLOW"` or `"DENY"`. The default is `"ALLOW"`.
	Access pulumi.StringPtrInput
	// Authentication entrypoint: `"ANY"` or `"RADIUS"`.
	Authtype pulumi.StringPtrInput
	// Elapsed time before the next MFA challenge.
	MfaLifetime pulumi.IntPtrInput
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: `"DEVICE"`, `"SESSION"` or `"ALWAYS"`.
	MfaPrompt pulumi.StringPtrInput
	// Remember MFA device. The default `false`.
	MfaRememberDevice pulumi.BoolPtrInput
	// Require MFA. By default is `false`.
	MfaRequired pulumi.BoolPtrInput
	// Policy Rule Name.
	Name pulumi.StringPtrInput
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection pulumi.StringPtrInput
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes pulumi.StringArrayInput
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes pulumi.StringArrayInput
	// Policy ID.
	Policyid pulumi.StringInput
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	// Max minutes a session can be idle.",
	SessionIdle pulumi.IntPtrInput
	// Max minutes a session is active: Disable = 0.
	SessionLifetime pulumi.IntPtrInput
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.
	SessionPersistent pulumi.BoolPtrInput
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status pulumi.StringPtrInput
	// Set of User IDs to Exclude
	UsersExcludeds pulumi.StringArrayInput
}

func (RuleSignonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleSignonArgs)(nil)).Elem()
}

