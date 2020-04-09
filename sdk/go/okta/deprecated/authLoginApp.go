// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package deprecated

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AuthLoginApp struct {
	pulumi.CustomResourceState

	// Custom error page URL
	AccessibilityErrorRedirectUrl pulumi.StringPtrOutput `pulumi:"accessibilityErrorRedirectUrl"`
	// Enable self service
	AccessibilitySelfService pulumi.BoolPtrOutput `pulumi:"accessibilitySelfService"`
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrOutput `pulumi:"autoSubmitToolbar"`
	// Application credentials scheme
	CredentialsScheme pulumi.StringPtrOutput `pulumi:"credentialsScheme"`
	// Groups associated with the application
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrOutput `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrOutput `pulumi:"hideWeb"`
	// Pretty name of app.
	Label pulumi.StringOutput `pulumi:"label"`
	// name of app.
	Name pulumi.StringOutput `pulumi:"name"`
	// Preconfigured app name
	PreconfiguredApp pulumi.StringPtrOutput `pulumi:"preconfiguredApp"`
	// Allow user to reveal password
	RevealPassword pulumi.BoolPtrOutput `pulumi:"revealPassword"`
	// Shared password, required for certain schemes.
	SharedPassword pulumi.StringPtrOutput `pulumi:"sharedPassword"`
	// Shared username, required for certain schemes.
	SharedUsername pulumi.StringPtrOutput `pulumi:"sharedUsername"`
	// Sign on mode of application.
	SignOnMode pulumi.StringOutput `pulumi:"signOnMode"`
	// Post login redirect URL
	SignOnRedirectUrl pulumi.StringPtrOutput `pulumi:"signOnRedirectUrl"`
	// Login URL
	SignOnUrl pulumi.StringPtrOutput `pulumi:"signOnUrl"`
	// Status of application.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Username template
	UserNameTemplate pulumi.StringOutput `pulumi:"userNameTemplate"`
	// Username template type
	UserNameTemplateType pulumi.StringOutput `pulumi:"userNameTemplateType"`
	// Users associated with the application
	Users AuthLoginAppUserArrayOutput `pulumi:"users"`
}

// NewAuthLoginApp registers a new resource with the given unique name, arguments, and options.
func NewAuthLoginApp(ctx *pulumi.Context,
	name string, args *AuthLoginAppArgs, opts ...pulumi.ResourceOption) (*AuthLoginApp, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil {
		args = &AuthLoginAppArgs{}
	}
	var resource AuthLoginApp
	err := ctx.RegisterResource("okta:deprecated/authLoginApp:AuthLoginApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthLoginApp gets an existing AuthLoginApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthLoginApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthLoginAppState, opts ...pulumi.ResourceOption) (*AuthLoginApp, error) {
	var resource AuthLoginApp
	err := ctx.ReadResource("okta:deprecated/authLoginApp:AuthLoginApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthLoginApp resources.
type authLoginAppState struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Enable self service
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Application credentials scheme
	CredentialsScheme *string `pulumi:"credentialsScheme"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Pretty name of app.
	Label *string `pulumi:"label"`
	// name of app.
	Name *string `pulumi:"name"`
	// Preconfigured app name
	PreconfiguredApp *string `pulumi:"preconfiguredApp"`
	// Allow user to reveal password
	RevealPassword *bool `pulumi:"revealPassword"`
	// Shared password, required for certain schemes.
	SharedPassword *string `pulumi:"sharedPassword"`
	// Shared username, required for certain schemes.
	SharedUsername *string `pulumi:"sharedUsername"`
	// Sign on mode of application.
	SignOnMode *string `pulumi:"signOnMode"`
	// Post login redirect URL
	SignOnRedirectUrl *string `pulumi:"signOnRedirectUrl"`
	// Login URL
	SignOnUrl *string `pulumi:"signOnUrl"`
	// Status of application.
	Status *string `pulumi:"status"`
	// Username template
	UserNameTemplate *string `pulumi:"userNameTemplate"`
	// Username template type
	UserNameTemplateType *string `pulumi:"userNameTemplateType"`
	// Users associated with the application
	Users []AuthLoginAppUser `pulumi:"users"`
}

type AuthLoginAppState struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl pulumi.StringPtrInput
	// Enable self service
	AccessibilitySelfService pulumi.BoolPtrInput
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Application credentials scheme
	CredentialsScheme pulumi.StringPtrInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// Pretty name of app.
	Label pulumi.StringPtrInput
	// name of app.
	Name pulumi.StringPtrInput
	// Preconfigured app name
	PreconfiguredApp pulumi.StringPtrInput
	// Allow user to reveal password
	RevealPassword pulumi.BoolPtrInput
	// Shared password, required for certain schemes.
	SharedPassword pulumi.StringPtrInput
	// Shared username, required for certain schemes.
	SharedUsername pulumi.StringPtrInput
	// Sign on mode of application.
	SignOnMode pulumi.StringPtrInput
	// Post login redirect URL
	SignOnRedirectUrl pulumi.StringPtrInput
	// Login URL
	SignOnUrl pulumi.StringPtrInput
	// Status of application.
	Status pulumi.StringPtrInput
	// Username template
	UserNameTemplate pulumi.StringPtrInput
	// Username template type
	UserNameTemplateType pulumi.StringPtrInput
	// Users associated with the application
	Users AuthLoginAppUserArrayInput
}

func (AuthLoginAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*authLoginAppState)(nil)).Elem()
}

type authLoginAppArgs struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Enable self service
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Application credentials scheme
	CredentialsScheme *string `pulumi:"credentialsScheme"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Pretty name of app.
	Label string `pulumi:"label"`
	// Preconfigured app name
	PreconfiguredApp *string `pulumi:"preconfiguredApp"`
	// Allow user to reveal password
	RevealPassword *bool `pulumi:"revealPassword"`
	// Shared password, required for certain schemes.
	SharedPassword *string `pulumi:"sharedPassword"`
	// Shared username, required for certain schemes.
	SharedUsername *string `pulumi:"sharedUsername"`
	// Post login redirect URL
	SignOnRedirectUrl *string `pulumi:"signOnRedirectUrl"`
	// Login URL
	SignOnUrl *string `pulumi:"signOnUrl"`
	// Status of application.
	Status *string `pulumi:"status"`
	// Users associated with the application
	Users []AuthLoginAppUser `pulumi:"users"`
}

// The set of arguments for constructing a AuthLoginApp resource.
type AuthLoginAppArgs struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl pulumi.StringPtrInput
	// Enable self service
	AccessibilitySelfService pulumi.BoolPtrInput
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Application credentials scheme
	CredentialsScheme pulumi.StringPtrInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// Pretty name of app.
	Label pulumi.StringInput
	// Preconfigured app name
	PreconfiguredApp pulumi.StringPtrInput
	// Allow user to reveal password
	RevealPassword pulumi.BoolPtrInput
	// Shared password, required for certain schemes.
	SharedPassword pulumi.StringPtrInput
	// Shared username, required for certain schemes.
	SharedUsername pulumi.StringPtrInput
	// Post login redirect URL
	SignOnRedirectUrl pulumi.StringPtrInput
	// Login URL
	SignOnUrl pulumi.StringPtrInput
	// Status of application.
	Status pulumi.StringPtrInput
	// Users associated with the application
	Users AuthLoginAppUserArrayInput
}

func (AuthLoginAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authLoginAppArgs)(nil)).Elem()
}
