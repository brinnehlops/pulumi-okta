// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package idp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Social Identity Provider.
//
// This resource allows you to create and configure an Social Identity Provider.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/idp_social.html.markdown.
type Social struct {
	pulumi.CustomResourceState

	// Specifies the account linking action for an IdP user.
	AccountLinkAction pulumi.StringPtrOutput `pulumi:"accountLinkAction"`
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes pulumi.StringArrayOutput `pulumi:"accountLinkGroupIncludes"`
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding pulumi.StringOutput `pulumi:"authorizationBinding"`
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl pulumi.StringOutput `pulumi:"authorizationUrl"`
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction pulumi.StringPtrOutput `pulumi:"deprovisionedAction"`
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction pulumi.StringPtrOutput `pulumi:"groupsAction"`
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments pulumi.StringArrayOutput `pulumi:"groupsAssignments"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute pulumi.StringPtrOutput `pulumi:"groupsAttribute"`
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters pulumi.StringArrayOutput `pulumi:"groupsFilters"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode pulumi.StringPtrOutput `pulumi:"issuerMode"`
	MatchAttribute pulumi.StringPtrOutput `pulumi:"matchAttribute"`
	MatchType pulumi.StringPtrOutput `pulumi:"matchType"`
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew pulumi.IntPtrOutput `pulumi:"maxClockSkew"`
	// The Application's display name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster pulumi.BoolPtrOutput `pulumi:"profileMaster"`
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType pulumi.StringPtrOutput `pulumi:"protocolType"`
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction pulumi.StringPtrOutput `pulumi:"provisioningAction"`
	// The XML digital signature algorithm used when signing an AuthnRequest message.
	RequestSignatureAlgorithm pulumi.StringPtrOutput `pulumi:"requestSignatureAlgorithm"`
	// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
	RequestSignatureScope pulumi.StringPtrOutput `pulumi:"requestSignatureScope"`
	// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
	ResponseSignatureAlgorithm pulumi.StringPtrOutput `pulumi:"responseSignatureAlgorithm"`
	// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
	ResponseSignatureScope pulumi.StringPtrOutput `pulumi:"responseSignatureScope"`
	// The scopes of the IdP.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Status of the IdP.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute pulumi.StringPtrOutput `pulumi:"subjectMatchAttribute"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType pulumi.StringPtrOutput `pulumi:"subjectMatchType"`
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction pulumi.StringPtrOutput `pulumi:"suspendedAction"`
	// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	TokenBinding pulumi.StringOutput `pulumi:"tokenBinding"`
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl pulumi.StringOutput `pulumi:"tokenUrl"`
	// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate pulumi.StringPtrOutput `pulumi:"usernameTemplate"`
}

// NewSocial registers a new resource with the given unique name, arguments, and options.
func NewSocial(ctx *pulumi.Context,
	name string, args *SocialArgs, opts ...pulumi.ResourceOption) (*Social, error) {
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &SocialArgs{}
	}
	var resource Social
	err := ctx.RegisterResource("okta:idp/social:Social", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSocial gets an existing Social resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSocial(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SocialState, opts ...pulumi.ResourceOption) (*Social, error) {
	var resource Social
	err := ctx.ReadResource("okta:idp/social:Social", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Social resources.
type socialState struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction *string `pulumi:"accountLinkAction"`
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes []string `pulumi:"accountLinkGroupIncludes"`
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding *string `pulumi:"authorizationBinding"`
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl *string `pulumi:"authorizationUrl"`
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId *string `pulumi:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret *string `pulumi:"clientSecret"`
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction *string `pulumi:"deprovisionedAction"`
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction *string `pulumi:"groupsAction"`
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments []string `pulumi:"groupsAssignments"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute *string `pulumi:"groupsAttribute"`
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters []string `pulumi:"groupsFilters"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode *string `pulumi:"issuerMode"`
	MatchAttribute *string `pulumi:"matchAttribute"`
	MatchType *string `pulumi:"matchType"`
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew *int `pulumi:"maxClockSkew"`
	// The Application's display name.
	Name *string `pulumi:"name"`
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster *bool `pulumi:"profileMaster"`
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType *string `pulumi:"protocolType"`
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// The XML digital signature algorithm used when signing an AuthnRequest message.
	RequestSignatureAlgorithm *string `pulumi:"requestSignatureAlgorithm"`
	// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
	RequestSignatureScope *string `pulumi:"requestSignatureScope"`
	// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
	ResponseSignatureScope *string `pulumi:"responseSignatureScope"`
	// The scopes of the IdP.
	Scopes []string `pulumi:"scopes"`
	// Status of the IdP.
	Status *string `pulumi:"status"`
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute *string `pulumi:"subjectMatchAttribute"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType *string `pulumi:"subjectMatchType"`
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction *string `pulumi:"suspendedAction"`
	// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	TokenBinding *string `pulumi:"tokenBinding"`
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl *string `pulumi:"tokenUrl"`
	// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
	Type *string `pulumi:"type"`
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

type SocialState struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction pulumi.StringPtrInput
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes pulumi.StringArrayInput
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding pulumi.StringPtrInput
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl pulumi.StringPtrInput
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId pulumi.StringPtrInput
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret pulumi.StringPtrInput
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction pulumi.StringPtrInput
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction pulumi.StringPtrInput
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments pulumi.StringArrayInput
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute pulumi.StringPtrInput
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters pulumi.StringArrayInput
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode pulumi.StringPtrInput
	MatchAttribute pulumi.StringPtrInput
	MatchType pulumi.StringPtrInput
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew pulumi.IntPtrInput
	// The Application's display name.
	Name pulumi.StringPtrInput
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster pulumi.BoolPtrInput
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType pulumi.StringPtrInput
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction pulumi.StringPtrInput
	// The XML digital signature algorithm used when signing an AuthnRequest message.
	RequestSignatureAlgorithm pulumi.StringPtrInput
	// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
	RequestSignatureScope pulumi.StringPtrInput
	// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
	ResponseSignatureScope pulumi.StringPtrInput
	// The scopes of the IdP.
	Scopes pulumi.StringArrayInput
	// Status of the IdP.
	Status pulumi.StringPtrInput
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute pulumi.StringPtrInput
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType pulumi.StringPtrInput
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction pulumi.StringPtrInput
	// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	TokenBinding pulumi.StringPtrInput
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl pulumi.StringPtrInput
	// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
	Type pulumi.StringPtrInput
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate pulumi.StringPtrInput
}

func (SocialState) ElementType() reflect.Type {
	return reflect.TypeOf((*socialState)(nil)).Elem()
}

type socialArgs struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction *string `pulumi:"accountLinkAction"`
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes []string `pulumi:"accountLinkGroupIncludes"`
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId *string `pulumi:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret *string `pulumi:"clientSecret"`
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction *string `pulumi:"deprovisionedAction"`
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction *string `pulumi:"groupsAction"`
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments []string `pulumi:"groupsAssignments"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute *string `pulumi:"groupsAttribute"`
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters []string `pulumi:"groupsFilters"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode *string `pulumi:"issuerMode"`
	MatchAttribute *string `pulumi:"matchAttribute"`
	MatchType *string `pulumi:"matchType"`
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew *int `pulumi:"maxClockSkew"`
	// The Application's display name.
	Name *string `pulumi:"name"`
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster *bool `pulumi:"profileMaster"`
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType *string `pulumi:"protocolType"`
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// The XML digital signature algorithm used when signing an AuthnRequest message.
	RequestSignatureAlgorithm *string `pulumi:"requestSignatureAlgorithm"`
	// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
	RequestSignatureScope *string `pulumi:"requestSignatureScope"`
	// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
	ResponseSignatureScope *string `pulumi:"responseSignatureScope"`
	// The scopes of the IdP.
	Scopes []string `pulumi:"scopes"`
	// Status of the IdP.
	Status *string `pulumi:"status"`
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute *string `pulumi:"subjectMatchAttribute"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType *string `pulumi:"subjectMatchType"`
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction *string `pulumi:"suspendedAction"`
	// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
	Type string `pulumi:"type"`
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

// The set of arguments for constructing a Social resource.
type SocialArgs struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction pulumi.StringPtrInput
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes pulumi.StringArrayInput
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId pulumi.StringPtrInput
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret pulumi.StringPtrInput
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction pulumi.StringPtrInput
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction pulumi.StringPtrInput
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments pulumi.StringArrayInput
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute pulumi.StringPtrInput
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters pulumi.StringArrayInput
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode pulumi.StringPtrInput
	MatchAttribute pulumi.StringPtrInput
	MatchType pulumi.StringPtrInput
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew pulumi.IntPtrInput
	// The Application's display name.
	Name pulumi.StringPtrInput
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster pulumi.BoolPtrInput
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType pulumi.StringPtrInput
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction pulumi.StringPtrInput
	// The XML digital signature algorithm used when signing an AuthnRequest message.
	RequestSignatureAlgorithm pulumi.StringPtrInput
	// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
	RequestSignatureScope pulumi.StringPtrInput
	// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
	ResponseSignatureScope pulumi.StringPtrInput
	// The scopes of the IdP.
	Scopes pulumi.StringArrayInput
	// Status of the IdP.
	Status pulumi.StringPtrInput
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute pulumi.StringPtrInput
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType pulumi.StringPtrInput
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction pulumi.StringPtrInput
	// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
	Type pulumi.StringInput
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate pulumi.StringPtrInput
}

func (SocialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*socialArgs)(nil)).Elem()
}

