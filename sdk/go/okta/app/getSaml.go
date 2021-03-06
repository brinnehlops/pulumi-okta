// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve the collaborators for a given repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-okta/sdk/v2/go/okta/app"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Example App"
// 		_, err := app.LookupSaml(ctx, &app.LookupSamlArgs{
// 			Label: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSaml(ctx *pulumi.Context, args *LookupSamlArgs, opts ...pulumi.InvokeOption) (*LookupSamlResult, error) {
	var rv LookupSamlResult
	err := ctx.Invoke("okta:app/getSaml:getSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSaml.
type LookupSamlArgs struct {
	// Custom error page URL.
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	AccessibilityLoginRedirectUrl *string `pulumi:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// tells the provider to query for only `ACTIVE` applications.
	ActiveOnly *bool `pulumi:"activeOnly"`
	// Application settings in JSON format.
	AppSettingsJson *string `pulumi:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned *bool `pulumi:"assertionSigned"`
	// SAML Attribute statements.
	AttributeStatements []GetSamlAttributeStatement `pulumi:"attributeStatements"`
	// Audience restriction.
	Audience *string `pulumi:"audience"`
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	AuthnContextClassRef *string `pulumi:"authnContextClassRef"`
	// Display auto submit toolbar.
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	DefaultRelayState *string `pulumi:"defaultRelayState"`
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	Destination *string `pulumi:"destination"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	DigestAlgorithm *string `pulumi:"digestAlgorithm"`
	// features enabled.
	Features []string `pulumi:"features"`
	// Do not display application icon on mobile app.
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn *bool `pulumi:"honorForceAuthn"`
	// `id` of application to retrieve, conflicts with `label` and `labelPrefix`.
	Id *string `pulumi:"id"`
	// SAML issuer ID.
	IdpIssuer *string `pulumi:"idpIssuer"`
	// The label of the app to retrieve, conflicts with `labelPrefix` and `id`.
	Label *string `pulumi:"label"`
	// Label prefix of the app to retrieve, conflicts with `label` and `id`. This will tell the provider to do a `starts with` query as opposed to an `equals` query.
	LabelPrefix *string `pulumi:"labelPrefix"`
	// The location where the app may present the SAML assertion.
	Recipient *string `pulumi:"recipient"`
	// Denotes whether the request is compressed or not.
	RequestCompressed *bool `pulumi:"requestCompressed"`
	// Determines whether the SAML auth response message is digitally signed.
	ResponseSigned *bool `pulumi:"responseSigned"`
	// Signature algorithm used ot digitally sign the assertion and response.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// SAML service provider issuer.
	SpIssuer *string `pulumi:"spIssuer"`
	// Single Sign on Url.
	SsoUrl *string `pulumi:"ssoUrl"`
	// Identifies the SAML processing rules.
	SubjectNameIdFormat *string `pulumi:"subjectNameIdFormat"`
	// Template for app user's username when a user is assigned to the app.
	SubjectNameIdTemplate *string `pulumi:"subjectNameIdTemplate"`
	// Username template.
	UserNameTemplate *string `pulumi:"userNameTemplate"`
	// Username template suffix.
	UserNameTemplateSuffix *string `pulumi:"userNameTemplateSuffix"`
	// Username template type.
	UserNameTemplateType *string `pulumi:"userNameTemplateType"`
}

// A collection of values returned by getSaml.
type LookupSamlResult struct {
	// Custom error page URL.
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	AccessibilityLoginRedirectUrl *string `pulumi:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	ActiveOnly               *bool `pulumi:"activeOnly"`
	// Application settings in JSON format.
	AppSettingsJson *string `pulumi:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned *bool `pulumi:"assertionSigned"`
	// SAML Attribute statements.
	AttributeStatements []GetSamlAttributeStatement `pulumi:"attributeStatements"`
	// Audience restriction.
	Audience *string `pulumi:"audience"`
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	AuthnContextClassRef *string `pulumi:"authnContextClassRef"`
	// Display auto submit toolbar.
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	DefaultRelayState *string `pulumi:"defaultRelayState"`
	// description of application.
	Description string `pulumi:"description"`
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	Destination *string `pulumi:"destination"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	DigestAlgorithm *string `pulumi:"digestAlgorithm"`
	// features enabled.
	Features []string `pulumi:"features"`
	// Do not display application icon on mobile app.
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn *bool `pulumi:"honorForceAuthn"`
	// id of application.
	Id *string `pulumi:"id"`
	// SAML issuer ID.
	IdpIssuer *string `pulumi:"idpIssuer"`
	// Certificate key ID.
	KeyId string `pulumi:"keyId"`
	// label of application.
	Label       *string `pulumi:"label"`
	LabelPrefix *string `pulumi:"labelPrefix"`
	// name of application.
	Name string `pulumi:"name"`
	// The location where the app may present the SAML assertion.
	Recipient *string `pulumi:"recipient"`
	// Denotes whether the request is compressed or not.
	RequestCompressed *bool `pulumi:"requestCompressed"`
	// Determines whether the SAML auth response message is digitally signed.
	ResponseSigned *bool `pulumi:"responseSigned"`
	// Signature algorithm used ot digitally sign the assertion and response.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// SAML service provider issuer.
	SpIssuer *string `pulumi:"spIssuer"`
	// Single Sign on Url.
	SsoUrl *string `pulumi:"ssoUrl"`
	// status of application.
	Status string `pulumi:"status"`
	// Identifies the SAML processing rules.
	SubjectNameIdFormat *string `pulumi:"subjectNameIdFormat"`
	// Template for app user's username when a user is assigned to the app.
	SubjectNameIdTemplate *string `pulumi:"subjectNameIdTemplate"`
	// Username template.
	UserNameTemplate *string `pulumi:"userNameTemplate"`
	// Username template suffix.
	UserNameTemplateSuffix *string `pulumi:"userNameTemplateSuffix"`
	// Username template type.
	UserNameTemplateType *string `pulumi:"userNameTemplateType"`
}
