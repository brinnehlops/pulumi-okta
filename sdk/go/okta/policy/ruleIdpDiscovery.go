// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an IdP Discovery Policy Rule.
//
// This resource allows you to create and configure an IdP Discovery Policy Rule.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_idp_discovery.html.markdown.
type RuleIdpDiscovery struct {
	pulumi.CustomResourceState

	// Applications to exclude in discovery rule
	AppExcludes RuleIdpDiscoveryAppExcludeArrayOutput `pulumi:"appExcludes"`
	// Applications to include in discovery rule
	AppIncludes RuleIdpDiscoveryAppIncludeArrayOutput `pulumi:"appIncludes"`
	IdpId       pulumi.StringPtrOutput                `pulumi:"idpId"`
	IdpType     pulumi.StringPtrOutput                `pulumi:"idpType"`
	// Policy Rule Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection pulumi.StringPtrOutput `pulumi:"networkConnection"`
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes pulumi.StringArrayOutput `pulumi:"networkExcludes"`
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes  pulumi.StringArrayOutput                   `pulumi:"networkIncludes"`
	PlatformIncludes RuleIdpDiscoveryPlatformIncludeArrayOutput `pulumi:"platformIncludes"`
	// Policy ID.
	Policyid pulumi.StringOutput `pulumi:"policyid"`
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status                  pulumi.StringPtrOutput                           `pulumi:"status"`
	UserIdentifierAttribute pulumi.StringPtrOutput                           `pulumi:"userIdentifierAttribute"`
	UserIdentifierPatterns  RuleIdpDiscoveryUserIdentifierPatternArrayOutput `pulumi:"userIdentifierPatterns"`
	UserIdentifierType      pulumi.StringPtrOutput                           `pulumi:"userIdentifierType"`
}

// NewRuleIdpDiscovery registers a new resource with the given unique name, arguments, and options.
func NewRuleIdpDiscovery(ctx *pulumi.Context,
	name string, args *RuleIdpDiscoveryArgs, opts ...pulumi.ResourceOption) (*RuleIdpDiscovery, error) {
	if args == nil || args.Policyid == nil {
		return nil, errors.New("missing required argument 'Policyid'")
	}
	if args == nil {
		args = &RuleIdpDiscoveryArgs{}
	}
	var resource RuleIdpDiscovery
	err := ctx.RegisterResource("okta:policy/ruleIdpDiscovery:RuleIdpDiscovery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleIdpDiscovery gets an existing RuleIdpDiscovery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleIdpDiscovery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleIdpDiscoveryState, opts ...pulumi.ResourceOption) (*RuleIdpDiscovery, error) {
	var resource RuleIdpDiscovery
	err := ctx.ReadResource("okta:policy/ruleIdpDiscovery:RuleIdpDiscovery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleIdpDiscovery resources.
type ruleIdpDiscoveryState struct {
	// Applications to exclude in discovery rule
	AppExcludes []RuleIdpDiscoveryAppExclude `pulumi:"appExcludes"`
	// Applications to include in discovery rule
	AppIncludes []RuleIdpDiscoveryAppInclude `pulumi:"appIncludes"`
	IdpId       *string                      `pulumi:"idpId"`
	IdpType     *string                      `pulumi:"idpType"`
	// Policy Rule Name.
	Name *string `pulumi:"name"`
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection *string `pulumi:"networkConnection"`
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes []string `pulumi:"networkExcludes"`
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes  []string                          `pulumi:"networkIncludes"`
	PlatformIncludes []RuleIdpDiscoveryPlatformInclude `pulumi:"platformIncludes"`
	// Policy ID.
	Policyid *string `pulumi:"policyid"`
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority *int `pulumi:"priority"`
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status                  *string                                 `pulumi:"status"`
	UserIdentifierAttribute *string                                 `pulumi:"userIdentifierAttribute"`
	UserIdentifierPatterns  []RuleIdpDiscoveryUserIdentifierPattern `pulumi:"userIdentifierPatterns"`
	UserIdentifierType      *string                                 `pulumi:"userIdentifierType"`
}

type RuleIdpDiscoveryState struct {
	// Applications to exclude in discovery rule
	AppExcludes RuleIdpDiscoveryAppExcludeArrayInput
	// Applications to include in discovery rule
	AppIncludes RuleIdpDiscoveryAppIncludeArrayInput
	IdpId       pulumi.StringPtrInput
	IdpType     pulumi.StringPtrInput
	// Policy Rule Name.
	Name pulumi.StringPtrInput
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection pulumi.StringPtrInput
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes pulumi.StringArrayInput
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes  pulumi.StringArrayInput
	PlatformIncludes RuleIdpDiscoveryPlatformIncludeArrayInput
	// Policy ID.
	Policyid pulumi.StringPtrInput
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status                  pulumi.StringPtrInput
	UserIdentifierAttribute pulumi.StringPtrInput
	UserIdentifierPatterns  RuleIdpDiscoveryUserIdentifierPatternArrayInput
	UserIdentifierType      pulumi.StringPtrInput
}

func (RuleIdpDiscoveryState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleIdpDiscoveryState)(nil)).Elem()
}

type ruleIdpDiscoveryArgs struct {
	// Applications to exclude in discovery rule
	AppExcludes []RuleIdpDiscoveryAppExclude `pulumi:"appExcludes"`
	// Applications to include in discovery rule
	AppIncludes []RuleIdpDiscoveryAppInclude `pulumi:"appIncludes"`
	IdpId       *string                      `pulumi:"idpId"`
	IdpType     *string                      `pulumi:"idpType"`
	// Policy Rule Name.
	Name *string `pulumi:"name"`
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection *string `pulumi:"networkConnection"`
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes []string `pulumi:"networkExcludes"`
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes  []string                          `pulumi:"networkIncludes"`
	PlatformIncludes []RuleIdpDiscoveryPlatformInclude `pulumi:"platformIncludes"`
	// Policy ID.
	Policyid string `pulumi:"policyid"`
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority *int `pulumi:"priority"`
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status                  *string                                 `pulumi:"status"`
	UserIdentifierAttribute *string                                 `pulumi:"userIdentifierAttribute"`
	UserIdentifierPatterns  []RuleIdpDiscoveryUserIdentifierPattern `pulumi:"userIdentifierPatterns"`
	UserIdentifierType      *string                                 `pulumi:"userIdentifierType"`
}

// The set of arguments for constructing a RuleIdpDiscovery resource.
type RuleIdpDiscoveryArgs struct {
	// Applications to exclude in discovery rule
	AppExcludes RuleIdpDiscoveryAppExcludeArrayInput
	// Applications to include in discovery rule
	AppIncludes RuleIdpDiscoveryAppIncludeArrayInput
	IdpId       pulumi.StringPtrInput
	IdpType     pulumi.StringPtrInput
	// Policy Rule Name.
	Name pulumi.StringPtrInput
	// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
	NetworkConnection pulumi.StringPtrInput
	// The network zones to exclude. Conflicts with `networkIncludes`.
	NetworkExcludes pulumi.StringArrayInput
	// The network zones to include. Conflicts with `networkExcludes`.
	NetworkIncludes  pulumi.StringArrayInput
	PlatformIncludes RuleIdpDiscoveryPlatformIncludeArrayInput
	// Policy ID.
	Policyid pulumi.StringInput
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
	Status                  pulumi.StringPtrInput
	UserIdentifierAttribute pulumi.StringPtrInput
	UserIdentifierPatterns  RuleIdpDiscoveryUserIdentifierPatternArrayInput
	UserIdentifierType      pulumi.StringPtrInput
}

func (RuleIdpDiscoveryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleIdpDiscoveryArgs)(nil)).Elem()
}
