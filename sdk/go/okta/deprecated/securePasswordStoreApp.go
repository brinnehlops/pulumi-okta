// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecurePasswordStoreApp struct {
	s *pulumi.ResourceState
}

// NewSecurePasswordStoreApp registers a new resource with the given unique name, arguments, and options.
func NewSecurePasswordStoreApp(ctx *pulumi.Context,
	name string, args *SecurePasswordStoreAppArgs, opts ...pulumi.ResourceOpt) (*SecurePasswordStoreApp, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil || args.PasswordField == nil {
		return nil, errors.New("missing required argument 'PasswordField'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil || args.UsernameField == nil {
		return nil, errors.New("missing required argument 'UsernameField'")
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
		inputs["optionalField1"] = nil
		inputs["optionalField1Value"] = nil
		inputs["optionalField2"] = nil
		inputs["optionalField2Value"] = nil
		inputs["optionalField3"] = nil
		inputs["optionalField3Value"] = nil
		inputs["passwordField"] = nil
		inputs["revealPassword"] = nil
		inputs["sharedPassword"] = nil
		inputs["sharedUsername"] = nil
		inputs["status"] = nil
		inputs["url"] = nil
		inputs["usernameField"] = nil
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
		inputs["optionalField1"] = args.OptionalField1
		inputs["optionalField1Value"] = args.OptionalField1Value
		inputs["optionalField2"] = args.OptionalField2
		inputs["optionalField2Value"] = args.OptionalField2Value
		inputs["optionalField3"] = args.OptionalField3
		inputs["optionalField3Value"] = args.OptionalField3Value
		inputs["passwordField"] = args.PasswordField
		inputs["revealPassword"] = args.RevealPassword
		inputs["sharedPassword"] = args.SharedPassword
		inputs["sharedUsername"] = args.SharedUsername
		inputs["status"] = args.Status
		inputs["url"] = args.Url
		inputs["usernameField"] = args.UsernameField
		inputs["users"] = args.Users
	}
	inputs["name"] = nil
	inputs["signOnMode"] = nil
	inputs["userNameTemplate"] = nil
	inputs["userNameTemplateType"] = nil
	s, err := ctx.RegisterResource("okta:deprecated/securePasswordStoreApp:SecurePasswordStoreApp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurePasswordStoreApp{s: s}, nil
}

// GetSecurePasswordStoreApp gets an existing SecurePasswordStoreApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurePasswordStoreApp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurePasswordStoreAppState, opts ...pulumi.ResourceOpt) (*SecurePasswordStoreApp, error) {
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
		inputs["optionalField1"] = state.OptionalField1
		inputs["optionalField1Value"] = state.OptionalField1Value
		inputs["optionalField2"] = state.OptionalField2
		inputs["optionalField2Value"] = state.OptionalField2Value
		inputs["optionalField3"] = state.OptionalField3
		inputs["optionalField3Value"] = state.OptionalField3Value
		inputs["passwordField"] = state.PasswordField
		inputs["revealPassword"] = state.RevealPassword
		inputs["sharedPassword"] = state.SharedPassword
		inputs["sharedUsername"] = state.SharedUsername
		inputs["signOnMode"] = state.SignOnMode
		inputs["status"] = state.Status
		inputs["url"] = state.Url
		inputs["userNameTemplate"] = state.UserNameTemplate
		inputs["userNameTemplateType"] = state.UserNameTemplateType
		inputs["usernameField"] = state.UsernameField
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("okta:deprecated/securePasswordStoreApp:SecurePasswordStoreApp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurePasswordStoreApp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecurePasswordStoreApp) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecurePasswordStoreApp) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Custom error page URL
func (r *SecurePasswordStoreApp) AccessibilityErrorRedirectUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessibilityErrorRedirectUrl"])
}

// Enable self service
func (r *SecurePasswordStoreApp) AccessibilitySelfService() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["accessibilitySelfService"])
}

// Display auto submit toolbar
func (r *SecurePasswordStoreApp) AutoSubmitToolbar() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoSubmitToolbar"])
}

// Application credentials scheme
func (r *SecurePasswordStoreApp) CredentialsScheme() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["credentialsScheme"])
}

// Groups associated with the application
func (r *SecurePasswordStoreApp) Groups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groups"])
}

// Do not display application icon on mobile app
func (r *SecurePasswordStoreApp) HideIos() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideIos"])
}

// Do not display application icon to users
func (r *SecurePasswordStoreApp) HideWeb() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideWeb"])
}

// Pretty name of app.
func (r *SecurePasswordStoreApp) Label() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["label"])
}

// name of app.
func (r *SecurePasswordStoreApp) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Name of optional param in the login form
func (r *SecurePasswordStoreApp) OptionalField1() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["optionalField1"])
}

// Name of optional value in login form
func (r *SecurePasswordStoreApp) OptionalField1Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["optionalField1Value"])
}

// Name of optional param in the login form
func (r *SecurePasswordStoreApp) OptionalField2() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["optionalField2"])
}

// Name of optional value in login form
func (r *SecurePasswordStoreApp) OptionalField2Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["optionalField2Value"])
}

// Name of optional param in the login form
func (r *SecurePasswordStoreApp) OptionalField3() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["optionalField3"])
}

// Name of optional value in login form
func (r *SecurePasswordStoreApp) OptionalField3Value() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["optionalField3Value"])
}

// Login password field
func (r *SecurePasswordStoreApp) PasswordField() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordField"])
}

// Allow user to reveal password
func (r *SecurePasswordStoreApp) RevealPassword() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["revealPassword"])
}

// Shared password, required for certain schemes.
func (r *SecurePasswordStoreApp) SharedPassword() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedPassword"])
}

// Shared username, required for certain schemes.
func (r *SecurePasswordStoreApp) SharedUsername() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedUsername"])
}

// Sign on mode of application.
func (r *SecurePasswordStoreApp) SignOnMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["signOnMode"])
}

// Status of application.
func (r *SecurePasswordStoreApp) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Login URL
func (r *SecurePasswordStoreApp) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Username template
func (r *SecurePasswordStoreApp) UserNameTemplate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userNameTemplate"])
}

// Username template type
func (r *SecurePasswordStoreApp) UserNameTemplateType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userNameTemplateType"])
}

// Login username field
func (r *SecurePasswordStoreApp) UsernameField() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["usernameField"])
}

// Users associated with the application
func (r *SecurePasswordStoreApp) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering SecurePasswordStoreApp resources.
type SecurePasswordStoreAppState struct {
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
	// Name of optional param in the login form
	OptionalField1 interface{}
	// Name of optional value in login form
	OptionalField1Value interface{}
	// Name of optional param in the login form
	OptionalField2 interface{}
	// Name of optional value in login form
	OptionalField2Value interface{}
	// Name of optional param in the login form
	OptionalField3 interface{}
	// Name of optional value in login form
	OptionalField3Value interface{}
	// Login password field
	PasswordField interface{}
	// Allow user to reveal password
	RevealPassword interface{}
	// Shared password, required for certain schemes.
	SharedPassword interface{}
	// Shared username, required for certain schemes.
	SharedUsername interface{}
	// Sign on mode of application.
	SignOnMode interface{}
	// Status of application.
	Status interface{}
	// Login URL
	Url interface{}
	// Username template
	UserNameTemplate interface{}
	// Username template type
	UserNameTemplateType interface{}
	// Login username field
	UsernameField interface{}
	// Users associated with the application
	Users interface{}
}

// The set of arguments for constructing a SecurePasswordStoreApp resource.
type SecurePasswordStoreAppArgs struct {
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
	// Name of optional param in the login form
	OptionalField1 interface{}
	// Name of optional value in login form
	OptionalField1Value interface{}
	// Name of optional param in the login form
	OptionalField2 interface{}
	// Name of optional value in login form
	OptionalField2Value interface{}
	// Name of optional param in the login form
	OptionalField3 interface{}
	// Name of optional value in login form
	OptionalField3Value interface{}
	// Login password field
	PasswordField interface{}
	// Allow user to reveal password
	RevealPassword interface{}
	// Shared password, required for certain schemes.
	SharedPassword interface{}
	// Shared username, required for certain schemes.
	SharedUsername interface{}
	// Status of application.
	Status interface{}
	// Login URL
	Url interface{}
	// Login username field
	UsernameField interface{}
	// Users associated with the application
	Users interface{}
}
