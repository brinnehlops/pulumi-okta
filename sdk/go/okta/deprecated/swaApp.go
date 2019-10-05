// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SwaApp struct {
	s *pulumi.ResourceState
}

// NewSwaApp registers a new resource with the given unique name, arguments, and options.
func NewSwaApp(ctx *pulumi.Context,
	name string, args *SwaAppArgs, opts ...pulumi.ResourceOpt) (*SwaApp, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessibilityErrorRedirectUrl"] = nil
		inputs["accessibilitySelfService"] = nil
		inputs["autoSubmitToolbar"] = nil
		inputs["buttonField"] = nil
		inputs["groups"] = nil
		inputs["hideIos"] = nil
		inputs["hideWeb"] = nil
		inputs["label"] = nil
		inputs["passwordField"] = nil
		inputs["preconfiguredApp"] = nil
		inputs["status"] = nil
		inputs["url"] = nil
		inputs["urlRegex"] = nil
		inputs["usernameField"] = nil
		inputs["users"] = nil
	} else {
		inputs["accessibilityErrorRedirectUrl"] = args.AccessibilityErrorRedirectUrl
		inputs["accessibilitySelfService"] = args.AccessibilitySelfService
		inputs["autoSubmitToolbar"] = args.AutoSubmitToolbar
		inputs["buttonField"] = args.ButtonField
		inputs["groups"] = args.Groups
		inputs["hideIos"] = args.HideIos
		inputs["hideWeb"] = args.HideWeb
		inputs["label"] = args.Label
		inputs["passwordField"] = args.PasswordField
		inputs["preconfiguredApp"] = args.PreconfiguredApp
		inputs["status"] = args.Status
		inputs["url"] = args.Url
		inputs["urlRegex"] = args.UrlRegex
		inputs["usernameField"] = args.UsernameField
		inputs["users"] = args.Users
	}
	inputs["name"] = nil
	inputs["signOnMode"] = nil
	inputs["userNameTemplate"] = nil
	inputs["userNameTemplateType"] = nil
	s, err := ctx.RegisterResource("okta:deprecated/swaApp:SwaApp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SwaApp{s: s}, nil
}

// GetSwaApp gets an existing SwaApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwaApp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SwaAppState, opts ...pulumi.ResourceOpt) (*SwaApp, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessibilityErrorRedirectUrl"] = state.AccessibilityErrorRedirectUrl
		inputs["accessibilitySelfService"] = state.AccessibilitySelfService
		inputs["autoSubmitToolbar"] = state.AutoSubmitToolbar
		inputs["buttonField"] = state.ButtonField
		inputs["groups"] = state.Groups
		inputs["hideIos"] = state.HideIos
		inputs["hideWeb"] = state.HideWeb
		inputs["label"] = state.Label
		inputs["name"] = state.Name
		inputs["passwordField"] = state.PasswordField
		inputs["preconfiguredApp"] = state.PreconfiguredApp
		inputs["signOnMode"] = state.SignOnMode
		inputs["status"] = state.Status
		inputs["url"] = state.Url
		inputs["urlRegex"] = state.UrlRegex
		inputs["userNameTemplate"] = state.UserNameTemplate
		inputs["userNameTemplateType"] = state.UserNameTemplateType
		inputs["usernameField"] = state.UsernameField
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("okta:deprecated/swaApp:SwaApp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SwaApp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SwaApp) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SwaApp) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Custom error page URL
func (r *SwaApp) AccessibilityErrorRedirectUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessibilityErrorRedirectUrl"])
}

// Enable self service
func (r *SwaApp) AccessibilitySelfService() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["accessibilitySelfService"])
}

// Display auto submit toolbar
func (r *SwaApp) AutoSubmitToolbar() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoSubmitToolbar"])
}

// Login button field
func (r *SwaApp) ButtonField() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["buttonField"])
}

// Groups associated with the application
func (r *SwaApp) Groups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groups"])
}

// Do not display application icon on mobile app
func (r *SwaApp) HideIos() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideIos"])
}

// Do not display application icon to users
func (r *SwaApp) HideWeb() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideWeb"])
}

// Pretty name of app.
func (r *SwaApp) Label() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["label"])
}

// name of app.
func (r *SwaApp) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Login password field
func (r *SwaApp) PasswordField() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordField"])
}

// Preconfigured app name
func (r *SwaApp) PreconfiguredApp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["preconfiguredApp"])
}

// Sign on mode of application.
func (r *SwaApp) SignOnMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["signOnMode"])
}

// Status of application.
func (r *SwaApp) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Login URL
func (r *SwaApp) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// A regex that further restricts URL to the specified regex
func (r *SwaApp) UrlRegex() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["urlRegex"])
}

// Username template
func (r *SwaApp) UserNameTemplate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userNameTemplate"])
}

// Username template type
func (r *SwaApp) UserNameTemplateType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userNameTemplateType"])
}

// Login username field
func (r *SwaApp) UsernameField() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["usernameField"])
}

// Users associated with the application
func (r *SwaApp) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering SwaApp resources.
type SwaAppState struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl interface{}
	// Enable self service
	AccessibilitySelfService interface{}
	// Display auto submit toolbar
	AutoSubmitToolbar interface{}
	// Login button field
	ButtonField interface{}
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
	// Login password field
	PasswordField interface{}
	// Preconfigured app name
	PreconfiguredApp interface{}
	// Sign on mode of application.
	SignOnMode interface{}
	// Status of application.
	Status interface{}
	// Login URL
	Url interface{}
	// A regex that further restricts URL to the specified regex
	UrlRegex interface{}
	// Username template
	UserNameTemplate interface{}
	// Username template type
	UserNameTemplateType interface{}
	// Login username field
	UsernameField interface{}
	// Users associated with the application
	Users interface{}
}

// The set of arguments for constructing a SwaApp resource.
type SwaAppArgs struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl interface{}
	// Enable self service
	AccessibilitySelfService interface{}
	// Display auto submit toolbar
	AutoSubmitToolbar interface{}
	// Login button field
	ButtonField interface{}
	// Groups associated with the application
	Groups interface{}
	// Do not display application icon on mobile app
	HideIos interface{}
	// Do not display application icon to users
	HideWeb interface{}
	// Pretty name of app.
	Label interface{}
	// Login password field
	PasswordField interface{}
	// Preconfigured app name
	PreconfiguredApp interface{}
	// Status of application.
	Status interface{}
	// Login URL
	Url interface{}
	// A regex that further restricts URL to the specified regex
	UrlRegex interface{}
	// Login username field
	UsernameField interface{}
	// Users associated with the application
	Users interface{}
}
