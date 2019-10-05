// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Auto Login Okta Application.
// 
// This resource allows you to create and configure an Auto Login Okta Application.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/app_auto_login.html.markdown.
type AutoLogin struct {
	s *pulumi.ResourceState
}

// NewAutoLogin registers a new resource with the given unique name, arguments, and options.
func NewAutoLogin(ctx *pulumi.Context,
	name string, args *AutoLoginArgs, opts ...pulumi.ResourceOpt) (*AutoLogin, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessibilityErrorRedirectUrl"] = nil
		inputs["accessibilitySelfService"] = nil
		inputs["autoSubmitToolbar"] = nil
		inputs["credentialsScheme"] = nil
		inputs["groups"] = nil
		inputs["hideIos"] = nil
		inputs["hideWeb"] = nil
		inputs["label"] = nil
		inputs["preconfiguredApp"] = nil
		inputs["revealPassword"] = nil
		inputs["sharedPassword"] = nil
		inputs["sharedUsername"] = nil
		inputs["signOnRedirectUrl"] = nil
		inputs["signOnUrl"] = nil
		inputs["status"] = nil
		inputs["users"] = nil
	} else {
		inputs["accessibilityErrorRedirectUrl"] = args.AccessibilityErrorRedirectUrl
		inputs["accessibilitySelfService"] = args.AccessibilitySelfService
		inputs["autoSubmitToolbar"] = args.AutoSubmitToolbar
		inputs["credentialsScheme"] = args.CredentialsScheme
		inputs["groups"] = args.Groups
		inputs["hideIos"] = args.HideIos
		inputs["hideWeb"] = args.HideWeb
		inputs["label"] = args.Label
		inputs["preconfiguredApp"] = args.PreconfiguredApp
		inputs["revealPassword"] = args.RevealPassword
		inputs["sharedPassword"] = args.SharedPassword
		inputs["sharedUsername"] = args.SharedUsername
		inputs["signOnRedirectUrl"] = args.SignOnRedirectUrl
		inputs["signOnUrl"] = args.SignOnUrl
		inputs["status"] = args.Status
		inputs["users"] = args.Users
	}
	inputs["name"] = nil
	inputs["signOnMode"] = nil
	inputs["userNameTemplate"] = nil
	inputs["userNameTemplateType"] = nil
	s, err := ctx.RegisterResource("okta:app/autoLogin:AutoLogin", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AutoLogin{s: s}, nil
}

// GetAutoLogin gets an existing AutoLogin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoLogin(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AutoLoginState, opts ...pulumi.ResourceOpt) (*AutoLogin, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessibilityErrorRedirectUrl"] = state.AccessibilityErrorRedirectUrl
		inputs["accessibilitySelfService"] = state.AccessibilitySelfService
		inputs["autoSubmitToolbar"] = state.AutoSubmitToolbar
		inputs["credentialsScheme"] = state.CredentialsScheme
		inputs["groups"] = state.Groups
		inputs["hideIos"] = state.HideIos
		inputs["hideWeb"] = state.HideWeb
		inputs["label"] = state.Label
		inputs["name"] = state.Name
		inputs["preconfiguredApp"] = state.PreconfiguredApp
		inputs["revealPassword"] = state.RevealPassword
		inputs["sharedPassword"] = state.SharedPassword
		inputs["sharedUsername"] = state.SharedUsername
		inputs["signOnMode"] = state.SignOnMode
		inputs["signOnRedirectUrl"] = state.SignOnRedirectUrl
		inputs["signOnUrl"] = state.SignOnUrl
		inputs["status"] = state.Status
		inputs["userNameTemplate"] = state.UserNameTemplate
		inputs["userNameTemplateType"] = state.UserNameTemplateType
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("okta:app/autoLogin:AutoLogin", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AutoLogin{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AutoLogin) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AutoLogin) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Custom error page URL
func (r *AutoLogin) AccessibilityErrorRedirectUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessibilityErrorRedirectUrl"])
}

// Enable self service
func (r *AutoLogin) AccessibilitySelfService() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["accessibilitySelfService"])
}

// Display auto submit toolbar
func (r *AutoLogin) AutoSubmitToolbar() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoSubmitToolbar"])
}

// Application credentials scheme
func (r *AutoLogin) CredentialsScheme() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["credentialsScheme"])
}

// Groups associated with the application
func (r *AutoLogin) Groups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groups"])
}

// Do not display application icon on mobile app
func (r *AutoLogin) HideIos() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideIos"])
}

// Do not display application icon to users
func (r *AutoLogin) HideWeb() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideWeb"])
}

// The Application's display name.
func (r *AutoLogin) Label() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["label"])
}

// Name assigned to the application by Okta.
func (r *AutoLogin) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
func (r *AutoLogin) PreconfiguredApp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["preconfiguredApp"])
}

// Allow user to reveal password
func (r *AutoLogin) RevealPassword() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["revealPassword"])
}

// Shared password, required for certain schemes.
func (r *AutoLogin) SharedPassword() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedPassword"])
}

// Shared username, required for certain schemes.
func (r *AutoLogin) SharedUsername() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedUsername"])
}

// Sign on mode of application.
func (r *AutoLogin) SignOnMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["signOnMode"])
}

// Post login redirect URL
func (r *AutoLogin) SignOnRedirectUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["signOnRedirectUrl"])
}

// Login URL
func (r *AutoLogin) SignOnUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["signOnUrl"])
}

// The status of the application, by default it is `"ACTIVE"`.
func (r *AutoLogin) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Username template
func (r *AutoLogin) UserNameTemplate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userNameTemplate"])
}

// Username template type
func (r *AutoLogin) UserNameTemplateType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userNameTemplateType"])
}

// Users associated with the application
func (r *AutoLogin) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering AutoLogin resources.
type AutoLoginState struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl interface{}
	// Enable self service
	AccessibilitySelfService interface{}
	// Display auto submit toolbar
	AutoSubmitToolbar interface{}
	// Application credentials scheme
	CredentialsScheme interface{}
	// Groups associated with the application
	Groups interface{}
	// Do not display application icon on mobile app
	HideIos interface{}
	// Do not display application icon to users
	HideWeb interface{}
	// The Application's display name.
	Label interface{}
	// Name assigned to the application by Okta.
	Name interface{}
	// Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
	PreconfiguredApp interface{}
	// Allow user to reveal password
	RevealPassword interface{}
	// Shared password, required for certain schemes.
	SharedPassword interface{}
	// Shared username, required for certain schemes.
	SharedUsername interface{}
	// Sign on mode of application.
	SignOnMode interface{}
	// Post login redirect URL
	SignOnRedirectUrl interface{}
	// Login URL
	SignOnUrl interface{}
	// The status of the application, by default it is `"ACTIVE"`.
	Status interface{}
	// Username template
	UserNameTemplate interface{}
	// Username template type
	UserNameTemplateType interface{}
	// Users associated with the application
	Users interface{}
}

// The set of arguments for constructing a AutoLogin resource.
type AutoLoginArgs struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl interface{}
	// Enable self service
	AccessibilitySelfService interface{}
	// Display auto submit toolbar
	AutoSubmitToolbar interface{}
	// Application credentials scheme
	CredentialsScheme interface{}
	// Groups associated with the application
	Groups interface{}
	// Do not display application icon on mobile app
	HideIos interface{}
	// Do not display application icon to users
	HideWeb interface{}
	// The Application's display name.
	Label interface{}
	// Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
	PreconfiguredApp interface{}
	// Allow user to reveal password
	RevealPassword interface{}
	// Shared password, required for certain schemes.
	SharedPassword interface{}
	// Shared username, required for certain schemes.
	SharedUsername interface{}
	// Post login redirect URL
	SignOnRedirectUrl interface{}
	// Login URL
	SignOnUrl interface{}
	// The status of the application, by default it is `"ACTIVE"`.
	Status interface{}
	// Users associated with the application
	Users interface{}
}
