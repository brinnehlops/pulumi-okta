// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package idp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an OIDC Identity Provider.
//
// This resource allows you to create and configure an OIDC Identity Provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-okta/sdk/v2/go/okta/idp"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := idp.NewOidc(ctx, "example", &idp.OidcArgs{
// 			AcsBinding:           pulumi.String("HTTP-POST"),
// 			AcsType:              pulumi.String("INSTANCE"),
// 			AuthorizationBinding: pulumi.String("HTTP-REDIRECT"),
// 			AuthorizationUrl:     pulumi.String("https://idp.example.com/authorize"),
// 			ClientId:             pulumi.String("efg456"),
// 			ClientSecret:         pulumi.String("efg456"),
// 			IssuerUrl:            pulumi.String("https://id.example.com"),
// 			JwksBinding:          pulumi.String("HTTP-REDIRECT"),
// 			JwksUrl:              pulumi.String("https://idp.example.com/keys"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("openid"),
// 			},
// 			TokenBinding:     pulumi.String("HTTP-POST"),
// 			TokenUrl:         pulumi.String("https://idp.example.com/token"),
// 			UserInfoBinding:  pulumi.String("HTTP-REDIRECT"),
// 			UserInfoUrl:      pulumi.String("https://idp.example.com/userinfo"),
// 			UsernameTemplate: pulumi.String("idpuser.email"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Oidc struct {
	pulumi.CustomResourceState

	// Specifies the account linking action for an IdP user.
	AccountLinkAction pulumi.StringPtrOutput `pulumi:"accountLinkAction"`
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes pulumi.StringArrayOutput `pulumi:"accountLinkGroupIncludes"`
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding pulumi.StringOutput `pulumi:"acsBinding"`
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType pulumi.StringPtrOutput `pulumi:"acsType"`
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding pulumi.StringOutput `pulumi:"authorizationBinding"`
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl pulumi.StringOutput `pulumi:"authorizationUrl"`
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
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
	// URI that identifies the issuer.
	IssuerUrl pulumi.StringOutput `pulumi:"issuerUrl"`
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding pulumi.StringOutput `pulumi:"jwksBinding"`
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl pulumi.StringOutput `pulumi:"jwksUrl"`
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
	// algorithm to use to sign requests
	RequestSignatureAlgorithm pulumi.StringPtrOutput `pulumi:"requestSignatureAlgorithm"`
	// algorithm to use to sign response
	RequestSignatureScope pulumi.StringPtrOutput `pulumi:"requestSignatureScope"`
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm pulumi.StringPtrOutput `pulumi:"responseSignatureAlgorithm"`
	// algorithm to use to sign response
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
	// Type of OIDC IdP.
	Type            pulumi.StringOutput    `pulumi:"type"`
	UserInfoBinding pulumi.StringPtrOutput `pulumi:"userInfoBinding"`
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl pulumi.StringPtrOutput `pulumi:"userInfoUrl"`
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate pulumi.StringPtrOutput `pulumi:"usernameTemplate"`
}

// NewOidc registers a new resource with the given unique name, arguments, and options.
func NewOidc(ctx *pulumi.Context,
	name string, args *OidcArgs, opts ...pulumi.ResourceOption) (*Oidc, error) {
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
	if args == nil {
		args = &OidcArgs{}
	}
	var resource Oidc
	err := ctx.RegisterResource("okta:idp/oidc:Oidc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidc gets an existing Oidc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcState, opts ...pulumi.ResourceOption) (*Oidc, error) {
	var resource Oidc
	err := ctx.ReadResource("okta:idp/oidc:Oidc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Oidc resources.
type oidcState struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction *string `pulumi:"accountLinkAction"`
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes []string `pulumi:"accountLinkGroupIncludes"`
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding *string `pulumi:"acsBinding"`
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType *string `pulumi:"acsType"`
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
	// URI that identifies the issuer.
	IssuerUrl *string `pulumi:"issuerUrl"`
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding *string `pulumi:"jwksBinding"`
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl *string `pulumi:"jwksUrl"`
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
	// algorithm to use to sign requests
	RequestSignatureAlgorithm *string `pulumi:"requestSignatureAlgorithm"`
	// algorithm to use to sign response
	RequestSignatureScope *string `pulumi:"requestSignatureScope"`
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// algorithm to use to sign response
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
	// Type of OIDC IdP.
	Type            *string `pulumi:"type"`
	UserInfoBinding *string `pulumi:"userInfoBinding"`
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl *string `pulumi:"userInfoUrl"`
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

type OidcState struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction pulumi.StringPtrInput
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes pulumi.StringArrayInput
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding pulumi.StringPtrInput
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType pulumi.StringPtrInput
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
	// URI that identifies the issuer.
	IssuerUrl pulumi.StringPtrInput
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding pulumi.StringPtrInput
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl pulumi.StringPtrInput
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
	// algorithm to use to sign requests
	RequestSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
	RequestSignatureScope pulumi.StringPtrInput
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
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
	// Type of OIDC IdP.
	Type            pulumi.StringPtrInput
	UserInfoBinding pulumi.StringPtrInput
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl pulumi.StringPtrInput
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate pulumi.StringPtrInput
}

func (OidcState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcState)(nil)).Elem()
}

type oidcArgs struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction *string `pulumi:"accountLinkAction"`
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes []string `pulumi:"accountLinkGroupIncludes"`
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding string `pulumi:"acsBinding"`
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType *string `pulumi:"acsType"`
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding string `pulumi:"authorizationBinding"`
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl string `pulumi:"authorizationUrl"`
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId string `pulumi:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret string `pulumi:"clientSecret"`
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
	// URI that identifies the issuer.
	IssuerUrl string `pulumi:"issuerUrl"`
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding string `pulumi:"jwksBinding"`
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl string `pulumi:"jwksUrl"`
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
	// algorithm to use to sign requests
	RequestSignatureAlgorithm *string `pulumi:"requestSignatureAlgorithm"`
	// algorithm to use to sign response
	RequestSignatureScope *string `pulumi:"requestSignatureScope"`
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// algorithm to use to sign response
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
	TokenBinding string `pulumi:"tokenBinding"`
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl        string  `pulumi:"tokenUrl"`
	UserInfoBinding *string `pulumi:"userInfoBinding"`
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl *string `pulumi:"userInfoUrl"`
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate *string `pulumi:"usernameTemplate"`
}

// The set of arguments for constructing a Oidc resource.
type OidcArgs struct {
	// Specifies the account linking action for an IdP user.
	AccountLinkAction pulumi.StringPtrInput
	// Group memberships to determine link candidates.
	AccountLinkGroupIncludes pulumi.StringArrayInput
	// The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AcsBinding pulumi.StringInput
	// The type of ACS. Default is `"INSTANCE"`.
	AcsType pulumi.StringPtrInput
	// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	AuthorizationBinding pulumi.StringInput
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	AuthorizationUrl pulumi.StringInput
	// Unique identifier issued by AS for the Okta IdP instance.
	ClientId pulumi.StringInput
	// Client secret issued by AS for the Okta IdP instance.
	ClientSecret pulumi.StringInput
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
	// URI that identifies the issuer.
	IssuerUrl pulumi.StringInput
	// The method of making a request for the OIDC JWKS. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
	JwksBinding pulumi.StringInput
	// Endpoint where the signer of the keys publishes its keys in a JWK Set.
	JwksUrl pulumi.StringInput
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
	// algorithm to use to sign requests
	RequestSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
	RequestSignatureScope pulumi.StringPtrInput
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
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
	TokenBinding pulumi.StringInput
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	TokenUrl        pulumi.StringInput
	UserInfoBinding pulumi.StringPtrInput
	// Protected resource endpoint that returns claims about the authenticated user.
	UserInfoUrl pulumi.StringPtrInput
	// Okta EL Expression to generate or transform a unique username for the IdP user.
	UsernameTemplate pulumi.StringPtrInput
}

func (OidcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcArgs)(nil)).Elem()
}
