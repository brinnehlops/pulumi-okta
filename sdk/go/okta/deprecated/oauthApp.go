// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type OauthApp struct {
	pulumi.CustomResourceState

	// Requested key rotation mode.
	AutoKeyRotation pulumi.BoolPtrOutput `pulumi:"autoKeyRotation"`
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrOutput `pulumi:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	ClientBasicSecret pulumi.StringPtrOutput `pulumi:"clientBasicSecret"`
	// OAuth client ID.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// OAuth client secret key. This will be in plain text in your statefile unless you set omit_secret above.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// URI to a web page providing information about the client.
	ClientUri pulumi.StringPtrOutput `pulumi:"clientUri"`
	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED.
	// Default value is TRUSTED
	ConsentMethod pulumi.StringPtrOutput `pulumi:"consentMethod"`
	// This property allows you to set your client_id.
	CustomClientId pulumi.StringPtrOutput `pulumi:"customClientId"`
	// List of OAuth 2.0 grant types. Conditional validation params found here
	// https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per
	// app type.
	GrantTypes pulumi.StringArrayOutput `pulumi:"grantTypes"`
	// Groups associated with the application
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrOutput `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrOutput `pulumi:"hideWeb"`
	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a
	// custom domain URL as the issuer of ID token for this client.
	IssuerMode pulumi.StringPtrOutput `pulumi:"issuerMode"`
	// Pretty name of app.
	Label pulumi.StringOutput `pulumi:"label"`
	// URI that initiates login.
	LoginUri pulumi.StringPtrOutput `pulumi:"loginUri"`
	// URI that references a logo for the client.
	LogoUri pulumi.StringPtrOutput `pulumi:"logoUri"`
	// name of app.
	Name pulumi.StringOutput `pulumi:"name"`
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false
	// your app will be recreated.
	OmitSecret pulumi.BoolPtrOutput `pulumi:"omitSecret"`
	// *Early Access Property*. URI to web page providing client policy document.
	PolicyUri pulumi.StringPtrOutput `pulumi:"policyUri"`
	// List of URIs for redirection after logout
	PostLogoutRedirectUris pulumi.StringArrayOutput `pulumi:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see
	// okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	RedirectUris pulumi.StringArrayOutput `pulumi:"redirectUris"`
	// List of OAuth 2.0 response type strings.
	ResponseTypes pulumi.StringArrayOutput `pulumi:"responseTypes"`
	// Sign on mode of application.
	SignOnMode pulumi.StringOutput `pulumi:"signOnMode"`
	// Status of application.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Requested authentication method for the token endpoint.
	TokenEndpointAuthMethod pulumi.StringPtrOutput `pulumi:"tokenEndpointAuthMethod"`
	// *Early Access Property*. URI to web page providing client tos (terms of service).
	TosUri pulumi.StringPtrOutput `pulumi:"tosUri"`
	// The type of client application.
	Type pulumi.StringOutput `pulumi:"type"`
	// Users associated with the application
	Users OauthAppUserArrayOutput `pulumi:"users"`
}

// NewOauthApp registers a new resource with the given unique name, arguments, and options.
func NewOauthApp(ctx *pulumi.Context,
	name string, args *OauthAppArgs, opts ...pulumi.ResourceOption) (*OauthApp, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &OauthAppArgs{}
	}
	var resource OauthApp
	err := ctx.RegisterResource("okta:deprecated/oauthApp:OauthApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOauthApp gets an existing OauthApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOauthApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OauthAppState, opts ...pulumi.ResourceOption) (*OauthApp, error) {
	var resource OauthApp
	err := ctx.ReadResource("okta:deprecated/oauthApp:OauthApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OauthApp resources.
type oauthAppState struct {
	// Requested key rotation mode.
	AutoKeyRotation *bool `pulumi:"autoKeyRotation"`
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	ClientBasicSecret *string `pulumi:"clientBasicSecret"`
	// OAuth client ID.
	ClientId *string `pulumi:"clientId"`
	// OAuth client secret key. This will be in plain text in your statefile unless you set omit_secret above.
	ClientSecret *string `pulumi:"clientSecret"`
	// URI to a web page providing information about the client.
	ClientUri *string `pulumi:"clientUri"`
	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED.
	// Default value is TRUSTED
	ConsentMethod *string `pulumi:"consentMethod"`
	// This property allows you to set your client_id.
	CustomClientId *string `pulumi:"customClientId"`
	// List of OAuth 2.0 grant types. Conditional validation params found here
	// https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per
	// app type.
	GrantTypes []string `pulumi:"grantTypes"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a
	// custom domain URL as the issuer of ID token for this client.
	IssuerMode *string `pulumi:"issuerMode"`
	// Pretty name of app.
	Label *string `pulumi:"label"`
	// URI that initiates login.
	LoginUri *string `pulumi:"loginUri"`
	// URI that references a logo for the client.
	LogoUri *string `pulumi:"logoUri"`
	// name of app.
	Name *string `pulumi:"name"`
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false
	// your app will be recreated.
	OmitSecret *bool `pulumi:"omitSecret"`
	// *Early Access Property*. URI to web page providing client policy document.
	PolicyUri *string `pulumi:"policyUri"`
	// List of URIs for redirection after logout
	PostLogoutRedirectUris []string `pulumi:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile
	Profile *string `pulumi:"profile"`
	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see
	// okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	RedirectUris []string `pulumi:"redirectUris"`
	// List of OAuth 2.0 response type strings.
	ResponseTypes []string `pulumi:"responseTypes"`
	// Sign on mode of application.
	SignOnMode *string `pulumi:"signOnMode"`
	// Status of application.
	Status *string `pulumi:"status"`
	// Requested authentication method for the token endpoint.
	TokenEndpointAuthMethod *string `pulumi:"tokenEndpointAuthMethod"`
	// *Early Access Property*. URI to web page providing client tos (terms of service).
	TosUri *string `pulumi:"tosUri"`
	// The type of client application.
	Type *string `pulumi:"type"`
	// Users associated with the application
	Users []OauthAppUser `pulumi:"users"`
}

type OauthAppState struct {
	// Requested key rotation mode.
	AutoKeyRotation pulumi.BoolPtrInput
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	ClientBasicSecret pulumi.StringPtrInput
	// OAuth client ID.
	ClientId pulumi.StringPtrInput
	// OAuth client secret key. This will be in plain text in your statefile unless you set omit_secret above.
	ClientSecret pulumi.StringPtrInput
	// URI to a web page providing information about the client.
	ClientUri pulumi.StringPtrInput
	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED.
	// Default value is TRUSTED
	ConsentMethod pulumi.StringPtrInput
	// This property allows you to set your client_id.
	CustomClientId pulumi.StringPtrInput
	// List of OAuth 2.0 grant types. Conditional validation params found here
	// https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per
	// app type.
	GrantTypes pulumi.StringArrayInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a
	// custom domain URL as the issuer of ID token for this client.
	IssuerMode pulumi.StringPtrInput
	// Pretty name of app.
	Label pulumi.StringPtrInput
	// URI that initiates login.
	LoginUri pulumi.StringPtrInput
	// URI that references a logo for the client.
	LogoUri pulumi.StringPtrInput
	// name of app.
	Name pulumi.StringPtrInput
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false
	// your app will be recreated.
	OmitSecret pulumi.BoolPtrInput
	// *Early Access Property*. URI to web page providing client policy document.
	PolicyUri pulumi.StringPtrInput
	// List of URIs for redirection after logout
	PostLogoutRedirectUris pulumi.StringArrayInput
	// Custom JSON that represents an OAuth application's profile
	Profile pulumi.StringPtrInput
	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see
	// okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	RedirectUris pulumi.StringArrayInput
	// List of OAuth 2.0 response type strings.
	ResponseTypes pulumi.StringArrayInput
	// Sign on mode of application.
	SignOnMode pulumi.StringPtrInput
	// Status of application.
	Status pulumi.StringPtrInput
	// Requested authentication method for the token endpoint.
	TokenEndpointAuthMethod pulumi.StringPtrInput
	// *Early Access Property*. URI to web page providing client tos (terms of service).
	TosUri pulumi.StringPtrInput
	// The type of client application.
	Type pulumi.StringPtrInput
	// Users associated with the application
	Users OauthAppUserArrayInput
}

func (OauthAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthAppState)(nil)).Elem()
}

type oauthAppArgs struct {
	// Requested key rotation mode.
	AutoKeyRotation *bool `pulumi:"autoKeyRotation"`
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	ClientBasicSecret *string `pulumi:"clientBasicSecret"`
	// URI to a web page providing information about the client.
	ClientUri *string `pulumi:"clientUri"`
	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED.
	// Default value is TRUSTED
	ConsentMethod *string `pulumi:"consentMethod"`
	// This property allows you to set your client_id.
	CustomClientId *string `pulumi:"customClientId"`
	// List of OAuth 2.0 grant types. Conditional validation params found here
	// https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per
	// app type.
	GrantTypes []string `pulumi:"grantTypes"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a
	// custom domain URL as the issuer of ID token for this client.
	IssuerMode *string `pulumi:"issuerMode"`
	// Pretty name of app.
	Label string `pulumi:"label"`
	// URI that initiates login.
	LoginUri *string `pulumi:"loginUri"`
	// URI that references a logo for the client.
	LogoUri *string `pulumi:"logoUri"`
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false
	// your app will be recreated.
	OmitSecret *bool `pulumi:"omitSecret"`
	// *Early Access Property*. URI to web page providing client policy document.
	PolicyUri *string `pulumi:"policyUri"`
	// List of URIs for redirection after logout
	PostLogoutRedirectUris []string `pulumi:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile
	Profile *string `pulumi:"profile"`
	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see
	// okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	RedirectUris []string `pulumi:"redirectUris"`
	// List of OAuth 2.0 response type strings.
	ResponseTypes []string `pulumi:"responseTypes"`
	// Status of application.
	Status *string `pulumi:"status"`
	// Requested authentication method for the token endpoint.
	TokenEndpointAuthMethod *string `pulumi:"tokenEndpointAuthMethod"`
	// *Early Access Property*. URI to web page providing client tos (terms of service).
	TosUri *string `pulumi:"tosUri"`
	// The type of client application.
	Type string `pulumi:"type"`
	// Users associated with the application
	Users []OauthAppUser `pulumi:"users"`
}

// The set of arguments for constructing a OauthApp resource.
type OauthAppArgs struct {
	// Requested key rotation mode.
	AutoKeyRotation pulumi.BoolPtrInput
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	ClientBasicSecret pulumi.StringPtrInput
	// URI to a web page providing information about the client.
	ClientUri pulumi.StringPtrInput
	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED.
	// Default value is TRUSTED
	ConsentMethod pulumi.StringPtrInput
	// This property allows you to set your client_id.
	CustomClientId pulumi.StringPtrInput
	// List of OAuth 2.0 grant types. Conditional validation params found here
	// https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per
	// app type.
	GrantTypes pulumi.StringArrayInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a
	// custom domain URL as the issuer of ID token for this client.
	IssuerMode pulumi.StringPtrInput
	// Pretty name of app.
	Label pulumi.StringInput
	// URI that initiates login.
	LoginUri pulumi.StringPtrInput
	// URI that references a logo for the client.
	LogoUri pulumi.StringPtrInput
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false
	// your app will be recreated.
	OmitSecret pulumi.BoolPtrInput
	// *Early Access Property*. URI to web page providing client policy document.
	PolicyUri pulumi.StringPtrInput
	// List of URIs for redirection after logout
	PostLogoutRedirectUris pulumi.StringArrayInput
	// Custom JSON that represents an OAuth application's profile
	Profile pulumi.StringPtrInput
	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see
	// okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	RedirectUris pulumi.StringArrayInput
	// List of OAuth 2.0 response type strings.
	ResponseTypes pulumi.StringArrayInput
	// Status of application.
	Status pulumi.StringPtrInput
	// Requested authentication method for the token endpoint.
	TokenEndpointAuthMethod pulumi.StringPtrInput
	// *Early Access Property*. URI to web page providing client tos (terms of service).
	TosUri pulumi.StringPtrInput
	// The type of client application.
	Type pulumi.StringInput
	// Users associated with the application
	Users OauthAppUserArrayInput
}

func (OauthAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthAppArgs)(nil)).Elem()
}
