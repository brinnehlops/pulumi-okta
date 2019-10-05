// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package idp

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an OIDC Identity Provider.
// 
// This resource allows you to create and configure an OIDC Identity Provider.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/idp_oidc.html.markdown.
type Oidc struct {
	s *pulumi.ResourceState
}

// NewOidc registers a new resource with the given unique name, arguments, and options.
func NewOidc(ctx *pulumi.Context,
	name string, args *OidcArgs, opts ...pulumi.ResourceOpt) (*Oidc, error) {
	if args == nil || args.AcsBinding == nil {
		return nil, errors.New("missing required argument 'AcsBinding'")
	}
	if args == nil || args.AuthorizationBinding == nil {
		return nil, errors.New("missing required argument 'AuthorizationBinding'")
	}
	if args == nil || args.AuthorizationUrl == nil {
		return nil, errors.New("missing required argument 'AuthorizationUrl'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.IssuerUrl == nil {
		return nil, errors.New("missing required argument 'IssuerUrl'")
	}
	if args == nil || args.JwksBinding == nil {
		return nil, errors.New("missing required argument 'JwksBinding'")
	}
	if args == nil || args.JwksUrl == nil {
		return nil, errors.New("missing required argument 'JwksUrl'")
	}
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil || args.TokenBinding == nil {
		return nil, errors.New("missing required argument 'TokenBinding'")
	}
	if args == nil || args.TokenUrl == nil {
		return nil, errors.New("missing required argument 'TokenUrl'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountLinkAction"] = nil
		inputs["accountLinkGroupIncludes"] = nil
		inputs["acsBinding"] = nil
		inputs["acsType"] = nil
		inputs["authorizationBinding"] = nil
		inputs["authorizationUrl"] = nil
		inputs["clientId"] = nil
		inputs["clientSecret"] = nil
		inputs["deprovisionedAction"] = nil
		inputs["groupsAction"] = nil
		inputs["groupsAssignments"] = nil
		inputs["groupsAttribute"] = nil
		inputs["groupsFilters"] = nil
		inputs["issuerMode"] = nil
		inputs["issuerUrl"] = nil
		inputs["jwksBinding"] = nil
		inputs["jwksUrl"] = nil
		inputs["maxClockSkew"] = nil
		inputs["name"] = nil
		inputs["profileMaster"] = nil
		inputs["protocolType"] = nil
		inputs["provisioningAction"] = nil
		inputs["requestSignatureAlgorithm"] = nil
		inputs["requestSignatureScope"] = nil
		inputs["responseSignatureAlgorithm"] = nil
		inputs["responseSignatureScope"] = nil
		inputs["scopes"] = nil
		inputs["status"] = nil
		inputs["subjectMatchAttribute"] = nil
		inputs["subjectMatchType"] = nil
		inputs["suspendedAction"] = nil
		inputs["tokenBinding"] = nil
		inputs["tokenUrl"] = nil
		inputs["userInfoBinding"] = nil
		inputs["userInfoUrl"] = nil
		inputs["usernameTemplate"] = nil
	} else {
		inputs["accountLinkAction"] = args.AccountLinkAction
		inputs["accountLinkGroupIncludes"] = args.AccountLinkGroupIncludes
		inputs["acsBinding"] = args.AcsBinding
		inputs["acsType"] = args.AcsType
		inputs["authorizationBinding"] = args.AuthorizationBinding
		inputs["authorizationUrl"] = args.AuthorizationUrl
		inputs["clientId"] = args.ClientId
		inputs["clientSecret"] = args.ClientSecret
		inputs["deprovisionedAction"] = args.DeprovisionedAction
		inputs["groupsAction"] = args.GroupsAction
		inputs["groupsAssignments"] = args.GroupsAssignments
		inputs["groupsAttribute"] = args.GroupsAttribute
		inputs["groupsFilters"] = args.GroupsFilters
		inputs["issuerMode"] = args.IssuerMode
		inputs["issuerUrl"] = args.IssuerUrl
		inputs["jwksBinding"] = args.JwksBinding
		inputs["jwksUrl"] = args.JwksUrl
		inputs["maxClockSkew"] = args.MaxClockSkew
		inputs["name"] = args.Name
		inputs["profileMaster"] = args.ProfileMaster
		inputs["protocolType"] = args.ProtocolType
		inputs["provisioningAction"] = args.ProvisioningAction
		inputs["requestSignatureAlgorithm"] = args.RequestSignatureAlgorithm
		inputs["requestSignatureScope"] = args.RequestSignatureScope
		inputs["responseSignatureAlgorithm"] = args.ResponseSignatureAlgorithm
		inputs["responseSignatureScope"] = args.ResponseSignatureScope
		inputs["scopes"] = args.Scopes
		inputs["status"] = args.Status
		inputs["subjectMatchAttribute"] = args.SubjectMatchAttribute
		inputs["subjectMatchType"] = args.SubjectMatchType
		inputs["suspendedAction"] = args.SuspendedAction
		inputs["tokenBinding"] = args.TokenBinding
		inputs["tokenUrl"] = args.TokenUrl
		inputs["userInfoBinding"] = args.UserInfoBinding
		inputs["userInfoUrl"] = args.UserInfoUrl
		inputs["usernameTemplate"] = args.UsernameTemplate
	}
	inputs["type"] = nil
	s, err := ctx.RegisterResource("okta:idp/oidc:Oidc", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Oidc{s: s}, nil
}

// GetOidc gets an existing Oidc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidc(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OidcState, opts ...pulumi.ResourceOpt) (*Oidc, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountLinkAction"] = state.AccountLinkAction
		inputs["accountLinkGroupIncludes"] = state.AccountLinkGroupIncludes
		inputs["acsBinding"] = state.AcsBinding
		inputs["acsType"] = state.AcsType
		inputs["authorizationBinding"] = state.AuthorizationBinding
		inputs["authorizationUrl"] = state.AuthorizationUrl
		inputs["clientId"] = state.ClientId
		inputs["clientSecret"] = state.ClientSecret
		inputs["deprovisionedAction"] = state.DeprovisionedAction
		inputs["groupsAction"] = state.GroupsAction
		inputs["groupsAssignments"] = state.GroupsAssignments
		inputs["groupsAttribute"] = state.GroupsAttribute
		inputs["groupsFilters"] = state.GroupsFilters
		inputs["issuerMode"] = state.IssuerMode
		inputs["issuerUrl"] = state.IssuerUrl
		inputs["jwksBinding"] = state.JwksBinding
		inputs["jwksUrl"] = state.JwksUrl
		inputs["maxClockSkew"] = state.MaxClockSkew
		inputs["name"] = state.Name
		inputs["profileMaster"] = state.ProfileMaster
		inputs["protocolType"] = state.ProtocolType
		inputs["provisioningAction"] = state.ProvisioningAction
		inputs["requestSignatureAlgorithm"] = state.RequestSignatureAlgorithm
		inputs["requestSignatureScope"] = state.RequestSignatureScope
		inputs["responseSignatureAlgorithm"] = state.ResponseSignatureAlgorithm
		inputs["responseSignatureScope"] = state.ResponseSignatureScope
		inputs["scopes"] = state.Scopes
		inputs["status"] = state.Status
		inputs["subjectMatchAttribute"] = state.SubjectMatchAttribute
		inputs["subjectMatchType"] = state.SubjectMatchType
		inputs["suspendedAction"] = state.SuspendedAction
		inputs["tokenBinding"] = state.TokenBinding
		inputs["tokenUrl"] = state.TokenUrl
		inputs["type"] = state.Type
		inputs["userInfoBinding"] = state.UserInfoBinding
		inputs["userInfoUrl"] = state.UserInfoUrl
		inputs["usernameTemplate"] = state.UsernameTemplate
	}
	s, err := ctx.ReadResource("okta:idp/oidc:Oidc", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Oidc{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Oidc) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Oidc) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the account linking action for an IdP user.
func (r *Oidc) AccountLinkAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accountLinkAction"])
}

// Group memberships to determine link candidates.
func (r *Oidc) AccountLinkGroupIncludes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["accountLinkGroupIncludes"])
}

// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
func (r *Oidc) AcsBinding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["acsBinding"])
}

// The type of ACS. Default is `"INSTANCE"`.
func (r *Oidc) AcsType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["acsType"])
}

// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
func (r *Oidc) AuthorizationBinding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authorizationBinding"])
}

// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
func (r *Oidc) AuthorizationUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authorizationUrl"])
}

// Unique identifier issued by AS for the Okta IdP instance.
func (r *Oidc) ClientId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clientId"])
}

// Client secret issued by AS for the Okta IdP instance.
func (r *Oidc) ClientSecret() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clientSecret"])
}

// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
func (r *Oidc) DeprovisionedAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deprovisionedAction"])
}

// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
func (r *Oidc) GroupsAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupsAction"])
}

// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
func (r *Oidc) GroupsAssignments() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groupsAssignments"])
}

// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
func (r *Oidc) GroupsAttribute() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupsAttribute"])
}

// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
func (r *Oidc) GroupsFilters() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groupsFilters"])
}

// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
func (r *Oidc) IssuerMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["issuerMode"])
}

// URI that identifies the issuer.
func (r *Oidc) IssuerUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["issuerUrl"])
}

// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
func (r *Oidc) JwksBinding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["jwksBinding"])
}

// Endpoint where the signer of the keys publishes its keys in a JWK Set.
func (r *Oidc) JwksUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["jwksUrl"])
}

// Maximum allowable clock-skew when processing messages from the IdP.
func (r *Oidc) MaxClockSkew() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxClockSkew"])
}

// The Application's display name.
func (r *Oidc) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Determines if the IdP should act as a source of truth for user profile attributes.
func (r *Oidc) ProfileMaster() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["profileMaster"])
}

// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
func (r *Oidc) ProtocolType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocolType"])
}

// Provisioning action for an IdP user during authentication.
func (r *Oidc) ProvisioningAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["provisioningAction"])
}

// algorithm to use to sign requests
func (r *Oidc) RequestSignatureAlgorithm() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["requestSignatureAlgorithm"])
}

// algorithm to use to sign response
func (r *Oidc) RequestSignatureScope() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["requestSignatureScope"])
}

// algorithm to use to sign requests
func (r *Oidc) ResponseSignatureAlgorithm() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["responseSignatureAlgorithm"])
}

// algorithm to use to sign response
func (r *Oidc) ResponseSignatureScope() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["responseSignatureScope"])
}

// The scopes of the IdP.
func (r *Oidc) Scopes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["scopes"])
}

// Status of the IdP.
func (r *Oidc) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
func (r *Oidc) SubjectMatchAttribute() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subjectMatchAttribute"])
}

// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
func (r *Oidc) SubjectMatchType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subjectMatchType"])
}

// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
func (r *Oidc) SuspendedAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["suspendedAction"])
}

// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
func (r *Oidc) TokenBinding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tokenBinding"])
}

// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
func (r *Oidc) TokenUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tokenUrl"])
}

// Type of OIDC IdP.
func (r *Oidc) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

func (r *Oidc) UserInfoBinding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userInfoBinding"])
}

// Protected resource endpoint that returns claims about the authenticated user.
func (r *Oidc) UserInfoUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userInfoUrl"])
}

// Okta EL Expression to generate or transform a unique username for the IdP user.
func (r *Oidc) UsernameTemplate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["usernameTemplate"])
}

// Input properties used for looking up and filtering Oidc resources.
type OidcState struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction interface{}
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes interface{}
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding interface{}
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType interface{}
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding interface{}
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl interface{}
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId interface{}
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret interface{}
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction interface{}
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction interface{}
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments interface{}
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute interface{}
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters interface{}
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode interface{}
	// URI that identifies the issuer.
	IssuerUrl interface{}
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding interface{}
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl interface{}
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew interface{}
	// The Application's display name.
	Name interface{}
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster interface{}
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType interface{}
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction interface{}
	// algorithm to use to sign requests
	RequestSignatureAlgorithm interface{}
	// algorithm to use to sign response
	RequestSignatureScope interface{}
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm interface{}
	// algorithm to use to sign response
	ResponseSignatureScope interface{}
	// The scopes of the IdP.
	Scopes interface{}
	// Status of the IdP.
	Status interface{}
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute interface{}
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType interface{}
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction interface{}
	// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	TokenBinding interface{}
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl interface{}
	// Type of OIDC IdP.
	Type interface{}
	UserInfoBinding interface{}
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl interface{}
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate interface{}
}

// The set of arguments for constructing a Oidc resource.
type OidcArgs struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction interface{}
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes interface{}
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding interface{}
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType interface{}
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding interface{}
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl interface{}
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId interface{}
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret interface{}
	// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
	DeprovisionedAction interface{}
	// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
	GroupsAction interface{}
	// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
	GroupsAssignments interface{}
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	GroupsAttribute interface{}
	// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
	GroupsFilters interface{}
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
	IssuerMode interface{}
	// URI that identifies the issuer.
	IssuerUrl interface{}
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding interface{}
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl interface{}
	// Maximum allowable clock-skew when processing messages from the IdP.
	MaxClockSkew interface{}
	// The Application's display name.
	Name interface{}
	// Determines if the IdP should act as a source of truth for user profile attributes.
	ProfileMaster interface{}
	// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
	ProtocolType interface{}
	// Provisioning action for an IdP user during authentication.
	ProvisioningAction interface{}
	// algorithm to use to sign requests
	RequestSignatureAlgorithm interface{}
	// algorithm to use to sign response
	RequestSignatureScope interface{}
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm interface{}
	// algorithm to use to sign response
	ResponseSignatureScope interface{}
	// The scopes of the IdP.
	Scopes interface{}
	// Status of the IdP.
	Status interface{}
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchAttribute interface{}
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
	SubjectMatchType interface{}
	// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
	SuspendedAction interface{}
	// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	TokenBinding interface{}
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl interface{}
	UserInfoBinding interface{}
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl interface{}
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate interface{}
}
