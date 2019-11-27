// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AuthLoginApp struct {
	s *pulumi.ResourceState
}

// NewAuthLoginApp registers a new resource with the given unique name, arguments, and options.
func NewAuthLoginApp(ctx *pulumi.Context,
	name string, args *AuthLoginAppArgs, opts ...pulumi.ResourceOpt) (*AuthLoginApp, error) {
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
	s, err := ctx.RegisterResource("okta:deprecated/authLoginApp:AuthLoginApp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthLoginApp{s: s}, nil
}

// GetAuthLoginApp gets an existing AuthLoginApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthLoginApp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthLoginAppState, opts ...pulumi.ResourceOpt) (*AuthLoginApp, error) {
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
	s, err := ctx.ReadResource("okta:deprecated/authLoginApp:AuthLoginApp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthLoginApp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthLoginApp) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthLoginApp) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Custom error page URL
func (r *AuthLoginApp) AccessibilityErrorRedirectUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessibilityErrorRedirectUrl"])
}

// Enable self service
func (r *AuthLoginApp) AccessibilitySelfService() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["accessibilitySelfService"])
}

// Display auto submit toolbar
func (r *AuthLoginApp) AutoSubmitToolbar() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoSubmitToolbar"])
}

// Application credentials scheme
func (r *AuthLoginApp) CredentialsScheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["credentialsScheme"])
}

// Groups associated with the application
func (r *AuthLoginApp) Groups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["groups"])
}

// Do not display application icon on mobile app
func (r *AuthLoginApp) HideIos() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["hideIos"])
}

// Do not display application icon to users
func (r *AuthLoginApp) HideWeb() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["hideWeb"])
}

// Pretty name of app.
func (r *AuthLoginApp) Label() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["label"])
}

// name of app.
func (r *AuthLoginApp) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Preconfigured app name
func (r *AuthLoginApp) PreconfiguredApp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["preconfiguredApp"])
}

// Allow user to reveal password
func (r *AuthLoginApp) RevealPassword() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["revealPassword"])
}

// Shared password, required for certain schemes.
func (r *AuthLoginApp) SharedPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedPassword"])
}

// Shared username, required for certain schemes.
func (r *AuthLoginApp) SharedUsername() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sharedUsername"])
}

// Sign on mode of application.
func (r *AuthLoginApp) SignOnMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["signOnMode"])
}

// Post login redirect URL
func (r *AuthLoginApp) SignOnRedirectUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["signOnRedirectUrl"])
}

// Login URL
func (r *AuthLoginApp) SignOnUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["signOnUrl"])
}

// Status of application.
func (r *AuthLoginApp) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// Username template
func (r *AuthLoginApp) UserNameTemplate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userNameTemplate"])
}

// Username template type
func (r *AuthLoginApp) UserNameTemplateType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userNameTemplateType"])
}

// Users associated with the application
func (r *AuthLoginApp) Users() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering AuthLoginApp resources.
type AuthLoginAppState struct {
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
	// Pretty name of app.
	Label interface{}
	// name of app.
	Name interface{}
	// Preconfigured app name
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
	// Status of application.
	Status interface{}
	// Username template
	UserNameTemplate interface{}
	// Username template type
	UserNameTemplateType interface{}
	// Users associated with the application
	Users interface{}
}

// The set of arguments for constructing a AuthLoginApp resource.
type AuthLoginAppArgs struct {
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
	// Pretty name of app.
	Label interface{}
	// Preconfigured app name
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
	// Status of application.
	Status interface{}
	// Users associated with the application
	Users interface{}
}
