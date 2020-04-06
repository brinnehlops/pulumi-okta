// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package app

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an SAML Application.
//
// This resource allows you to create and configure an SAML Application.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_saml.html.markdown.
type Saml struct {
	pulumi.CustomResourceState

	// Custom error page URL.
	AccessibilityErrorRedirectUrl pulumi.StringPtrOutput `pulumi:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	AccessibilityLoginRedirectUrl pulumi.StringPtrOutput `pulumi:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	AccessibilitySelfService pulumi.BoolPtrOutput `pulumi:"accessibilitySelfService"`
	// Application settings in JSON format.
	AppSettingsJson pulumi.StringPtrOutput `pulumi:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned pulumi.BoolPtrOutput `pulumi:"assertionSigned"`
	// List of SAML Attribute statements.
	AttributeStatements SamlAttributeStatementArrayOutput `pulumi:"attributeStatements"`
	// Audience restriction.
	Audience pulumi.StringPtrOutput `pulumi:"audience"`
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	AuthnContextClassRef pulumi.StringPtrOutput `pulumi:"authnContextClassRef"`
	// Display auto submit toolbar.
	AutoSubmitToolbar pulumi.BoolPtrOutput `pulumi:"autoSubmitToolbar"`
	// The raw signing certificate.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	DefaultRelayState pulumi.StringPtrOutput `pulumi:"defaultRelayState"`
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	Destination pulumi.StringPtrOutput `pulumi:"destination"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	DigestAlgorithm pulumi.StringPtrOutput `pulumi:"digestAlgorithm"`
	// Entity ID, the ID portion of the `entityUrl`.
	EntityKey pulumi.StringOutput `pulumi:"entityKey"`
	// Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8.
	EntityUrl pulumi.StringOutput `pulumi:"entityUrl"`
	// features enabled.
	Features pulumi.StringArrayOutput `pulumi:"features"`
	// Groups associated with the application
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Do not display application icon on mobile app.
	HideIos pulumi.BoolPtrOutput `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrOutput `pulumi:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn pulumi.BoolPtrOutput `pulumi:"honorForceAuthn"`
	// `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` location from the SAML metadata.
	HttpPostBinding pulumi.StringOutput `pulumi:"httpPostBinding"`
	// `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` location from the SAML metadata.
	HttpRedirectBinding pulumi.StringOutput `pulumi:"httpRedirectBinding"`
	// SAML issuer ID.
	IdpIssuer pulumi.StringPtrOutput `pulumi:"idpIssuer"`
	// Certificate key ID.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Certificate name. This modulates the rotation of keys. New name == new key.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// Number of years the certificate is valid.
	KeyYearsValid pulumi.IntPtrOutput `pulumi:"keyYearsValid"`
	// label of application.
	Label pulumi.StringOutput `pulumi:"label"`
	// The raw SAML metadata in XML.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name of the attribute statement.
	Name pulumi.StringOutput `pulumi:"name"`
	// name of application from the Okta Integration Network, if not included a custom app will be created.
	PreconfiguredApp pulumi.StringPtrOutput `pulumi:"preconfiguredApp"`
	// The location where the app may present the SAML assertion.
	Recipient pulumi.StringPtrOutput `pulumi:"recipient"`
	// Denotes whether the request is compressed or not.
	RequestCompressed pulumi.BoolPtrOutput `pulumi:"requestCompressed"`
	// Determines whether the SAML auth response message is digitally signed.
	ResponseSigned pulumi.BoolPtrOutput `pulumi:"responseSigned"`
	// Sign on mode of application.
	SignOnMode pulumi.StringOutput `pulumi:"signOnMode"`
	// Signature algorithm used ot digitally sign the assertion and response.
	SignatureAlgorithm pulumi.StringPtrOutput `pulumi:"signatureAlgorithm"`
	// SAML service provider issuer.
	SpIssuer pulumi.StringPtrOutput `pulumi:"spIssuer"`
	// Single Sign on Url.
	SsoUrl pulumi.StringPtrOutput `pulumi:"ssoUrl"`
	// status of application.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Identifies the SAML processing rules.
	SubjectNameIdFormat pulumi.StringPtrOutput `pulumi:"subjectNameIdFormat"`
	// Template for app user's username when a user is assigned to the app.
	SubjectNameIdTemplate pulumi.StringPtrOutput `pulumi:"subjectNameIdTemplate"`
	// Username template.
	UserNameTemplate pulumi.StringPtrOutput `pulumi:"userNameTemplate"`
	// Username template suffix.
	UserNameTemplateSuffix pulumi.StringPtrOutput `pulumi:"userNameTemplateSuffix"`
	// Username template type.
	UserNameTemplateType pulumi.StringPtrOutput `pulumi:"userNameTemplateType"`
	// Users associated with the application
	Users SamlUserArrayOutput `pulumi:"users"`
}

// NewSaml registers a new resource with the given unique name, arguments, and options.
func NewSaml(ctx *pulumi.Context,
	name string, args *SamlArgs, opts ...pulumi.ResourceOption) (*Saml, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil {
		args = &SamlArgs{}
	}
	var resource Saml
	err := ctx.RegisterResource("okta:app/saml:Saml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSaml gets an existing Saml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlState, opts ...pulumi.ResourceOption) (*Saml, error) {
	var resource Saml
	err := ctx.ReadResource("okta:app/saml:Saml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Saml resources.
type samlState struct {
	// Custom error page URL.
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	AccessibilityLoginRedirectUrl *string `pulumi:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// Application settings in JSON format.
	AppSettingsJson *string `pulumi:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned *bool `pulumi:"assertionSigned"`
	// List of SAML Attribute statements.
	AttributeStatements []SamlAttributeStatement `pulumi:"attributeStatements"`
	// Audience restriction.
	Audience *string `pulumi:"audience"`
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	AuthnContextClassRef *string `pulumi:"authnContextClassRef"`
	// Display auto submit toolbar.
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// The raw signing certificate.
	Certificate *string `pulumi:"certificate"`
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	DefaultRelayState *string `pulumi:"defaultRelayState"`
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	Destination *string `pulumi:"destination"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	DigestAlgorithm *string `pulumi:"digestAlgorithm"`
	// Entity ID, the ID portion of the `entityUrl`.
	EntityKey *string `pulumi:"entityKey"`
	// Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8.
	EntityUrl *string `pulumi:"entityUrl"`
	// features enabled.
	Features []string `pulumi:"features"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app.
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn *bool `pulumi:"honorForceAuthn"`
	// `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` location from the SAML metadata.
	HttpPostBinding *string `pulumi:"httpPostBinding"`
	// `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` location from the SAML metadata.
	HttpRedirectBinding *string `pulumi:"httpRedirectBinding"`
	// SAML issuer ID.
	IdpIssuer *string `pulumi:"idpIssuer"`
	// Certificate key ID.
	KeyId *string `pulumi:"keyId"`
	// Certificate name. This modulates the rotation of keys. New name == new key.
	KeyName *string `pulumi:"keyName"`
	// Number of years the certificate is valid.
	KeyYearsValid *int `pulumi:"keyYearsValid"`
	// label of application.
	Label *string `pulumi:"label"`
	// The raw SAML metadata in XML.
	Metadata *string `pulumi:"metadata"`
	// The name of the attribute statement.
	Name *string `pulumi:"name"`
	// name of application from the Okta Integration Network, if not included a custom app will be created.
	PreconfiguredApp *string `pulumi:"preconfiguredApp"`
	// The location where the app may present the SAML assertion.
	Recipient *string `pulumi:"recipient"`
	// Denotes whether the request is compressed or not.
	RequestCompressed *bool `pulumi:"requestCompressed"`
	// Determines whether the SAML auth response message is digitally signed.
	ResponseSigned *bool `pulumi:"responseSigned"`
	// Sign on mode of application.
	SignOnMode *string `pulumi:"signOnMode"`
	// Signature algorithm used ot digitally sign the assertion and response.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// SAML service provider issuer.
	SpIssuer *string `pulumi:"spIssuer"`
	// Single Sign on Url.
	SsoUrl *string `pulumi:"ssoUrl"`
	// status of application.
	Status *string `pulumi:"status"`
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
	// Users associated with the application
	Users []SamlUser `pulumi:"users"`
}

type SamlState struct {
	// Custom error page URL.
	AccessibilityErrorRedirectUrl pulumi.StringPtrInput
	// Custom login page URL.
	AccessibilityLoginRedirectUrl pulumi.StringPtrInput
	// Enable self service.
	AccessibilitySelfService pulumi.BoolPtrInput
	// Application settings in JSON format.
	AppSettingsJson pulumi.StringPtrInput
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned pulumi.BoolPtrInput
	// List of SAML Attribute statements.
	AttributeStatements SamlAttributeStatementArrayInput
	// Audience restriction.
	Audience pulumi.StringPtrInput
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	AuthnContextClassRef pulumi.StringPtrInput
	// Display auto submit toolbar.
	AutoSubmitToolbar pulumi.BoolPtrInput
	// The raw signing certificate.
	Certificate pulumi.StringPtrInput
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	DefaultRelayState pulumi.StringPtrInput
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	Destination pulumi.StringPtrInput
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	DigestAlgorithm pulumi.StringPtrInput
	// Entity ID, the ID portion of the `entityUrl`.
	EntityKey pulumi.StringPtrInput
	// Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8.
	EntityUrl pulumi.StringPtrInput
	// features enabled.
	Features pulumi.StringArrayInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app.
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn pulumi.BoolPtrInput
	// `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` location from the SAML metadata.
	HttpPostBinding pulumi.StringPtrInput
	// `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` location from the SAML metadata.
	HttpRedirectBinding pulumi.StringPtrInput
	// SAML issuer ID.
	IdpIssuer pulumi.StringPtrInput
	// Certificate key ID.
	KeyId pulumi.StringPtrInput
	// Certificate name. This modulates the rotation of keys. New name == new key.
	KeyName pulumi.StringPtrInput
	// Number of years the certificate is valid.
	KeyYearsValid pulumi.IntPtrInput
	// label of application.
	Label pulumi.StringPtrInput
	// The raw SAML metadata in XML.
	Metadata pulumi.StringPtrInput
	// The name of the attribute statement.
	Name pulumi.StringPtrInput
	// name of application from the Okta Integration Network, if not included a custom app will be created.
	PreconfiguredApp pulumi.StringPtrInput
	// The location where the app may present the SAML assertion.
	Recipient pulumi.StringPtrInput
	// Denotes whether the request is compressed or not.
	RequestCompressed pulumi.BoolPtrInput
	// Determines whether the SAML auth response message is digitally signed.
	ResponseSigned pulumi.BoolPtrInput
	// Sign on mode of application.
	SignOnMode pulumi.StringPtrInput
	// Signature algorithm used ot digitally sign the assertion and response.
	SignatureAlgorithm pulumi.StringPtrInput
	// SAML service provider issuer.
	SpIssuer pulumi.StringPtrInput
	// Single Sign on Url.
	SsoUrl pulumi.StringPtrInput
	// status of application.
	Status pulumi.StringPtrInput
	// Identifies the SAML processing rules.
	SubjectNameIdFormat pulumi.StringPtrInput
	// Template for app user's username when a user is assigned to the app.
	SubjectNameIdTemplate pulumi.StringPtrInput
	// Username template.
	UserNameTemplate pulumi.StringPtrInput
	// Username template suffix.
	UserNameTemplateSuffix pulumi.StringPtrInput
	// Username template type.
	UserNameTemplateType pulumi.StringPtrInput
	// Users associated with the application
	Users SamlUserArrayInput
}

func (SamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlState)(nil)).Elem()
}

type samlArgs struct {
	// Custom error page URL.
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	AccessibilityLoginRedirectUrl *string `pulumi:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// Application settings in JSON format.
	AppSettingsJson *string `pulumi:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned *bool `pulumi:"assertionSigned"`
	// List of SAML Attribute statements.
	AttributeStatements []SamlAttributeStatement `pulumi:"attributeStatements"`
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
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app.
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn *bool `pulumi:"honorForceAuthn"`
	// SAML issuer ID.
	IdpIssuer *string `pulumi:"idpIssuer"`
	// Certificate name. This modulates the rotation of keys. New name == new key.
	KeyName *string `pulumi:"keyName"`
	// Number of years the certificate is valid.
	KeyYearsValid *int `pulumi:"keyYearsValid"`
	// label of application.
	Label string `pulumi:"label"`
	// name of application from the Okta Integration Network, if not included a custom app will be created.
	PreconfiguredApp *string `pulumi:"preconfiguredApp"`
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
	Status *string `pulumi:"status"`
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
	// Users associated with the application
	Users []SamlUser `pulumi:"users"`
}

// The set of arguments for constructing a Saml resource.
type SamlArgs struct {
	// Custom error page URL.
	AccessibilityErrorRedirectUrl pulumi.StringPtrInput
	// Custom login page URL.
	AccessibilityLoginRedirectUrl pulumi.StringPtrInput
	// Enable self service.
	AccessibilitySelfService pulumi.BoolPtrInput
	// Application settings in JSON format.
	AppSettingsJson pulumi.StringPtrInput
	// Determines whether the SAML assertion is digitally signed.
	AssertionSigned pulumi.BoolPtrInput
	// List of SAML Attribute statements.
	AttributeStatements SamlAttributeStatementArrayInput
	// Audience restriction.
	Audience pulumi.StringPtrInput
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	AuthnContextClassRef pulumi.StringPtrInput
	// Display auto submit toolbar.
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	DefaultRelayState pulumi.StringPtrInput
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	Destination pulumi.StringPtrInput
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	DigestAlgorithm pulumi.StringPtrInput
	// features enabled.
	Features pulumi.StringArrayInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app.
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// Prompt user to re-authenticate if SP asks for it.
	HonorForceAuthn pulumi.BoolPtrInput
	// SAML issuer ID.
	IdpIssuer pulumi.StringPtrInput
	// Certificate name. This modulates the rotation of keys. New name == new key.
	KeyName pulumi.StringPtrInput
	// Number of years the certificate is valid.
	KeyYearsValid pulumi.IntPtrInput
	// label of application.
	Label pulumi.StringInput
	// name of application from the Okta Integration Network, if not included a custom app will be created.
	PreconfiguredApp pulumi.StringPtrInput
	// The location where the app may present the SAML assertion.
	Recipient pulumi.StringPtrInput
	// Denotes whether the request is compressed or not.
	RequestCompressed pulumi.BoolPtrInput
	// Determines whether the SAML auth response message is digitally signed.
	ResponseSigned pulumi.BoolPtrInput
	// Signature algorithm used ot digitally sign the assertion and response.
	SignatureAlgorithm pulumi.StringPtrInput
	// SAML service provider issuer.
	SpIssuer pulumi.StringPtrInput
	// Single Sign on Url.
	SsoUrl pulumi.StringPtrInput
	// status of application.
	Status pulumi.StringPtrInput
	// Identifies the SAML processing rules.
	SubjectNameIdFormat pulumi.StringPtrInput
	// Template for app user's username when a user is assigned to the app.
	SubjectNameIdTemplate pulumi.StringPtrInput
	// Username template.
	UserNameTemplate pulumi.StringPtrInput
	// Username template suffix.
	UserNameTemplateSuffix pulumi.StringPtrInput
	// Username template type.
	UserNameTemplateType pulumi.StringPtrInput
	// Users associated with the application
	Users SamlUserArrayInput
}

func (SamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlArgs)(nil)).Elem()
}
