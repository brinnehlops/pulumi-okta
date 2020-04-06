// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package app

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an OIDC Application.
//
// This resource allows you to create and configure an OIDC Application.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_oauth.html.markdown.
type OAuth struct {
	pulumi.CustomResourceState

	// Requested key rotation mode.
	AutoKeyRotation pulumi.BoolPtrOutput `pulumi:"autoKeyRotation"`
	// Display auto submit toolbar.
	AutoSubmitToolbar pulumi.BoolPtrOutput `pulumi:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
	ClientBasicSecret pulumi.StringPtrOutput `pulumi:"clientBasicSecret"`
	// The client ID of the application.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The client secret of the application.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// URI to a web page providing information about the client.
	ClientUri pulumi.StringPtrOutput `pulumi:"clientUri"`
	// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
	ConsentMethod pulumi.StringPtrOutput `pulumi:"consentMethod"`
	// This property allows you to set the application's client id.
	CustomClientId pulumi.StringPtrOutput `pulumi:"customClientId"`
	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	GrantTypes pulumi.StringArrayOutput `pulumi:"grantTypes"`
	// The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Do not display application icon on mobile app.
	HideIos pulumi.BoolPtrOutput `pulumi:"hideIos"`
	// Do not display application icon to users.
	HideWeb pulumi.BoolPtrOutput `pulumi:"hideWeb"`
	// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	IssuerMode pulumi.StringPtrOutput `pulumi:"issuerMode"`
	// The Application's display name.
	Label pulumi.StringOutput `pulumi:"label"`
	// URI that initiates login.
	LoginUri pulumi.StringPtrOutput `pulumi:"loginUri"`
	// URI that references a logo for the client.
	LogoUri pulumi.StringPtrOutput `pulumi:"logoUri"`
	// Name assigned to the application by Okta.
	Name pulumi.StringOutput `pulumi:"name"`
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	OmitSecret pulumi.BoolPtrOutput `pulumi:"omitSecret"`
	// URI to web page providing client policy document.
	PolicyUri pulumi.StringPtrOutput `pulumi:"policyUri"`
	// List of URIs for redirection after logout.
	PostLogoutRedirectUris pulumi.StringArrayOutput `pulumi:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile.
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// List of URIs for use in the redirect-based flow. This is required for all application types except service.
	RedirectUris pulumi.StringArrayOutput `pulumi:"redirectUris"`
	// List of OAuth 2.0 response type strings.
	ResponseTypes pulumi.StringArrayOutput `pulumi:"responseTypes"`
	// Sign on mode of application.
	SignOnMode pulumi.StringOutput `pulumi:"signOnMode"`
	// The status of the application, by default it is `"ACTIVE"`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
	TokenEndpointAuthMethod pulumi.StringPtrOutput `pulumi:"tokenEndpointAuthMethod"`
	// URI to web page providing client tos (terms of service).
	TosUri pulumi.StringPtrOutput `pulumi:"tosUri"`
	// The type of OAuth application.
	Type pulumi.StringOutput `pulumi:"type"`
	// The users assigned to the application. It is recommended not to use this and instead use `app.User`.
	Users OAuthUserArrayOutput `pulumi:"users"`
}

// NewOAuth registers a new resource with the given unique name, arguments, and options.
func NewOAuth(ctx *pulumi.Context,
	name string, args *OAuthArgs, opts ...pulumi.ResourceOption) (*OAuth, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &OAuthArgs{}
	}
	var resource OAuth
	err := ctx.RegisterResource("okta:app/oAuth:OAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOAuth gets an existing OAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OAuthState, opts ...pulumi.ResourceOption) (*OAuth, error) {
	var resource OAuth
	err := ctx.ReadResource("okta:app/oAuth:OAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OAuth resources.
type oauthState struct {
	// Requested key rotation mode.
	AutoKeyRotation *bool `pulumi:"autoKeyRotation"`
	// Display auto submit toolbar.
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
	ClientBasicSecret *string `pulumi:"clientBasicSecret"`
	// The client ID of the application.
	ClientId *string `pulumi:"clientId"`
	// The client secret of the application.
	ClientSecret *string `pulumi:"clientSecret"`
	// URI to a web page providing information about the client.
	ClientUri *string `pulumi:"clientUri"`
	// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
	ConsentMethod *string `pulumi:"consentMethod"`
	// This property allows you to set the application's client id.
	CustomClientId *string `pulumi:"customClientId"`
	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	GrantTypes []string `pulumi:"grantTypes"`
	// The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app.
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users.
	HideWeb *bool `pulumi:"hideWeb"`
	// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	IssuerMode *string `pulumi:"issuerMode"`
	// The Application's display name.
	Label *string `pulumi:"label"`
	// URI that initiates login.
	LoginUri *string `pulumi:"loginUri"`
	// URI that references a logo for the client.
	LogoUri *string `pulumi:"logoUri"`
	// Name assigned to the application by Okta.
	Name *string `pulumi:"name"`
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	OmitSecret *bool `pulumi:"omitSecret"`
	// URI to web page providing client policy document.
	PolicyUri *string `pulumi:"policyUri"`
	// List of URIs for redirection after logout.
	PostLogoutRedirectUris []string `pulumi:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile.
	Profile *string `pulumi:"profile"`
	// List of URIs for use in the redirect-based flow. This is required for all application types except service.
	RedirectUris []string `pulumi:"redirectUris"`
	// List of OAuth 2.0 response type strings.
	ResponseTypes []string `pulumi:"responseTypes"`
	// Sign on mode of application.
	SignOnMode *string `pulumi:"signOnMode"`
	// The status of the application, by default it is `"ACTIVE"`.
	Status *string `pulumi:"status"`
	// Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
	TokenEndpointAuthMethod *string `pulumi:"tokenEndpointAuthMethod"`
	// URI to web page providing client tos (terms of service).
	TosUri *string `pulumi:"tosUri"`
	// The type of OAuth application.
	Type *string `pulumi:"type"`
	// The users assigned to the application. It is recommended not to use this and instead use `app.User`.
	Users []OAuthUser `pulumi:"users"`
}

type OAuthState struct {
	// Requested key rotation mode.
	AutoKeyRotation pulumi.BoolPtrInput
	// Display auto submit toolbar.
	AutoSubmitToolbar pulumi.BoolPtrInput
	// OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
	ClientBasicSecret pulumi.StringPtrInput
	// The client ID of the application.
	ClientId pulumi.StringPtrInput
	// The client secret of the application.
	ClientSecret pulumi.StringPtrInput
	// URI to a web page providing information about the client.
	ClientUri pulumi.StringPtrInput
	// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
	ConsentMethod pulumi.StringPtrInput
	// This property allows you to set the application's client id.
	CustomClientId pulumi.StringPtrInput
	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	GrantTypes pulumi.StringArrayInput
	// The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app.
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users.
	HideWeb pulumi.BoolPtrInput
	// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	IssuerMode pulumi.StringPtrInput
	// The Application's display name.
	Label pulumi.StringPtrInput
	// URI that initiates login.
	LoginUri pulumi.StringPtrInput
	// URI that references a logo for the client.
	LogoUri pulumi.StringPtrInput
	// Name assigned to the application by Okta.
	Name pulumi.StringPtrInput
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	OmitSecret pulumi.BoolPtrInput
	// URI to web page providing client policy document.
	PolicyUri pulumi.StringPtrInput
	// List of URIs for redirection after logout.
	PostLogoutRedirectUris pulumi.StringArrayInput
	// Custom JSON that represents an OAuth application's profile.
	Profile pulumi.StringPtrInput
	// List of URIs for use in the redirect-based flow. This is required for all application types except service.
	RedirectUris pulumi.StringArrayInput
	// List of OAuth 2.0 response type strings.
	ResponseTypes pulumi.StringArrayInput
	// Sign on mode of application.
	SignOnMode pulumi.StringPtrInput
	// The status of the application, by default it is `"ACTIVE"`.
	Status pulumi.StringPtrInput
	// Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
	TokenEndpointAuthMethod pulumi.StringPtrInput
	// URI to web page providing client tos (terms of service).
	TosUri pulumi.StringPtrInput
	// The type of OAuth application.
	Type pulumi.StringPtrInput
	// The users assigned to the application. It is recommended not to use this and instead use `app.User`.
	Users OAuthUserArrayInput
}

func (OAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthState)(nil)).Elem()
}

type oauthArgs struct {
	// Requested key rotation mode.
	AutoKeyRotation *bool `pulumi:"autoKeyRotation"`
	// Display auto submit toolbar.
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
	ClientBasicSecret *string `pulumi:"clientBasicSecret"`
	// URI to a web page providing information about the client.
	ClientUri *string `pulumi:"clientUri"`
	// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
	ConsentMethod *string `pulumi:"consentMethod"`
	// This property allows you to set the application's client id.
	CustomClientId *string `pulumi:"customClientId"`
	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	GrantTypes []string `pulumi:"grantTypes"`
	// The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app.
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users.
	HideWeb *bool `pulumi:"hideWeb"`
	// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	IssuerMode *string `pulumi:"issuerMode"`
	// The Application's display name.
	Label string `pulumi:"label"`
	// URI that initiates login.
	LoginUri *string `pulumi:"loginUri"`
	// URI that references a logo for the client.
	LogoUri *string `pulumi:"logoUri"`
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	OmitSecret *bool `pulumi:"omitSecret"`
	// URI to web page providing client policy document.
	PolicyUri *string `pulumi:"policyUri"`
	// List of URIs for redirection after logout.
	PostLogoutRedirectUris []string `pulumi:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile.
	Profile *string `pulumi:"profile"`
	// List of URIs for use in the redirect-based flow. This is required for all application types except service.
	RedirectUris []string `pulumi:"redirectUris"`
	// List of OAuth 2.0 response type strings.
	ResponseTypes []string `pulumi:"responseTypes"`
	// The status of the application, by default it is `"ACTIVE"`.
	Status *string `pulumi:"status"`
	// Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
	TokenEndpointAuthMethod *string `pulumi:"tokenEndpointAuthMethod"`
	// URI to web page providing client tos (terms of service).
	TosUri *string `pulumi:"tosUri"`
	// The type of OAuth application.
	Type string `pulumi:"type"`
	// The users assigned to the application. It is recommended not to use this and instead use `app.User`.
	Users []OAuthUser `pulumi:"users"`
}

// The set of arguments for constructing a OAuth resource.
type OAuthArgs struct {
	// Requested key rotation mode.
	AutoKeyRotation pulumi.BoolPtrInput
	// Display auto submit toolbar.
	AutoSubmitToolbar pulumi.BoolPtrInput
	// OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
	ClientBasicSecret pulumi.StringPtrInput
	// URI to a web page providing information about the client.
	ClientUri pulumi.StringPtrInput
	// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
	ConsentMethod pulumi.StringPtrInput
	// This property allows you to set the application's client id.
	CustomClientId pulumi.StringPtrInput
	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	GrantTypes pulumi.StringArrayInput
	// The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app.
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users.
	HideWeb pulumi.BoolPtrInput
	// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	IssuerMode pulumi.StringPtrInput
	// The Application's display name.
	Label pulumi.StringInput
	// URI that initiates login.
	LoginUri pulumi.StringPtrInput
	// URI that references a logo for the client.
	LogoUri pulumi.StringPtrInput
	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	OmitSecret pulumi.BoolPtrInput
	// URI to web page providing client policy document.
	PolicyUri pulumi.StringPtrInput
	// List of URIs for redirection after logout.
	PostLogoutRedirectUris pulumi.StringArrayInput
	// Custom JSON that represents an OAuth application's profile.
	Profile pulumi.StringPtrInput
	// List of URIs for use in the redirect-based flow. This is required for all application types except service.
	RedirectUris pulumi.StringArrayInput
	// List of OAuth 2.0 response type strings.
	ResponseTypes pulumi.StringArrayInput
	// The status of the application, by default it is `"ACTIVE"`.
	Status pulumi.StringPtrInput
	// Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
	TokenEndpointAuthMethod pulumi.StringPtrInput
	// URI to web page providing client tos (terms of service).
	TosUri pulumi.StringPtrInput
	// The type of OAuth application.
	Type pulumi.StringInput
	// The users assigned to the application. It is recommended not to use this and instead use `app.User`.
	Users OAuthUserArrayInput
}

func (OAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthArgs)(nil)).Elem()
}
