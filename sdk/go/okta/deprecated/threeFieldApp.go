// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package deprecated

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ThreeFieldApp struct {
	pulumi.CustomResourceState

	// Custom error page URL
	AccessibilityErrorRedirectUrl pulumi.StringPtrOutput `pulumi:"accessibilityErrorRedirectUrl"`
	// Enable self service
	AccessibilitySelfService pulumi.BoolPtrOutput `pulumi:"accessibilitySelfService"`
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrOutput `pulumi:"autoSubmitToolbar"`
	// Login button field CSS selector
	ButtonSelector pulumi.StringOutput `pulumi:"buttonSelector"`
	// Extra field CSS selector
	ExtraFieldSelector pulumi.StringOutput `pulumi:"extraFieldSelector"`
	// Value for extra form field
	ExtraFieldValue pulumi.StringOutput `pulumi:"extraFieldValue"`
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
	// Login password field CSS selector
	PasswordSelector pulumi.StringOutput `pulumi:"passwordSelector"`
	// Sign on mode of application.
	SignOnMode pulumi.StringOutput `pulumi:"signOnMode"`
	// Status of application.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Login URL
	Url pulumi.StringOutput `pulumi:"url"`
	// A regex that further restricts URL to the specified regex
	UrlRegex pulumi.StringPtrOutput `pulumi:"urlRegex"`
	// Username template
	UserNameTemplate pulumi.StringOutput `pulumi:"userNameTemplate"`
	// Username template type
	UserNameTemplateType pulumi.StringOutput `pulumi:"userNameTemplateType"`
	// Login username field CSS selector
	UsernameSelector pulumi.StringOutput `pulumi:"usernameSelector"`
	// Users associated with the application
	Users ThreeFieldAppUserArrayOutput `pulumi:"users"`
}

// NewThreeFieldApp registers a new resource with the given unique name, arguments, and options.
func NewThreeFieldApp(ctx *pulumi.Context,
	name string, args *ThreeFieldAppArgs, opts ...pulumi.ResourceOption) (*ThreeFieldApp, error) {
	if args == nil || args.ButtonSelector == nil {
		return nil, errors.New("missing required argument 'ButtonSelector'")
	}
	if args == nil || args.ExtraFieldSelector == nil {
		return nil, errors.New("missing required argument 'ExtraFieldSelector'")
	}
	if args == nil || args.ExtraFieldValue == nil {
		return nil, errors.New("missing required argument 'ExtraFieldValue'")
	}
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil || args.PasswordSelector == nil {
		return nil, errors.New("missing required argument 'PasswordSelector'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil || args.UsernameSelector == nil {
		return nil, errors.New("missing required argument 'UsernameSelector'")
	}
	if args == nil {
		args = &ThreeFieldAppArgs{}
	}
	var resource ThreeFieldApp
	err := ctx.RegisterResource("okta:deprecated/threeFieldApp:ThreeFieldApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThreeFieldApp gets an existing ThreeFieldApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreeFieldApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreeFieldAppState, opts ...pulumi.ResourceOption) (*ThreeFieldApp, error) {
	var resource ThreeFieldApp
	err := ctx.ReadResource("okta:deprecated/threeFieldApp:ThreeFieldApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThreeFieldApp resources.
type threeFieldAppState struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Enable self service
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Login button field CSS selector
	ButtonSelector *string `pulumi:"buttonSelector"`
	// Extra field CSS selector
	ExtraFieldSelector *string `pulumi:"extraFieldSelector"`
	// Value for extra form field
	ExtraFieldValue *string `pulumi:"extraFieldValue"`
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
	// Login password field CSS selector
	PasswordSelector *string `pulumi:"passwordSelector"`
	// Sign on mode of application.
	SignOnMode *string `pulumi:"signOnMode"`
	// Status of application.
	Status *string `pulumi:"status"`
	// Login URL
	Url *string `pulumi:"url"`
	// A regex that further restricts URL to the specified regex
	UrlRegex *string `pulumi:"urlRegex"`
	// Username template
	UserNameTemplate *string `pulumi:"userNameTemplate"`
	// Username template type
	UserNameTemplateType *string `pulumi:"userNameTemplateType"`
	// Login username field CSS selector
	UsernameSelector *string `pulumi:"usernameSelector"`
	// Users associated with the application
	Users []ThreeFieldAppUser `pulumi:"users"`
}

type ThreeFieldAppState struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl pulumi.StringPtrInput
	// Enable self service
	AccessibilitySelfService pulumi.BoolPtrInput
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Login button field CSS selector
	ButtonSelector pulumi.StringPtrInput
	// Extra field CSS selector
	ExtraFieldSelector pulumi.StringPtrInput
	// Value for extra form field
	ExtraFieldValue pulumi.StringPtrInput
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
	// Login password field CSS selector
	PasswordSelector pulumi.StringPtrInput
	// Sign on mode of application.
	SignOnMode pulumi.StringPtrInput
	// Status of application.
	Status pulumi.StringPtrInput
	// Login URL
	Url pulumi.StringPtrInput
	// A regex that further restricts URL to the specified regex
	UrlRegex pulumi.StringPtrInput
	// Username template
	UserNameTemplate pulumi.StringPtrInput
	// Username template type
	UserNameTemplateType pulumi.StringPtrInput
	// Login username field CSS selector
	UsernameSelector pulumi.StringPtrInput
	// Users associated with the application
	Users ThreeFieldAppUserArrayInput
}

func (ThreeFieldAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*threeFieldAppState)(nil)).Elem()
}

type threeFieldAppArgs struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl *string `pulumi:"accessibilityErrorRedirectUrl"`
	// Enable self service
	AccessibilitySelfService *bool `pulumi:"accessibilitySelfService"`
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Login button field CSS selector
	ButtonSelector string `pulumi:"buttonSelector"`
	// Extra field CSS selector
	ExtraFieldSelector string `pulumi:"extraFieldSelector"`
	// Value for extra form field
	ExtraFieldValue string `pulumi:"extraFieldValue"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// Pretty name of app.
	Label string `pulumi:"label"`
	// Login password field CSS selector
	PasswordSelector string `pulumi:"passwordSelector"`
	// Status of application.
	Status *string `pulumi:"status"`
	// Login URL
	Url string `pulumi:"url"`
	// A regex that further restricts URL to the specified regex
	UrlRegex *string `pulumi:"urlRegex"`
	// Login username field CSS selector
	UsernameSelector string `pulumi:"usernameSelector"`
	// Users associated with the application
	Users []ThreeFieldAppUser `pulumi:"users"`
}

// The set of arguments for constructing a ThreeFieldApp resource.
type ThreeFieldAppArgs struct {
	// Custom error page URL
	AccessibilityErrorRedirectUrl pulumi.StringPtrInput
	// Enable self service
	AccessibilitySelfService pulumi.BoolPtrInput
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Login button field CSS selector
	ButtonSelector pulumi.StringInput
	// Extra field CSS selector
	ExtraFieldSelector pulumi.StringInput
	// Value for extra form field
	ExtraFieldValue pulumi.StringInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// Pretty name of app.
	Label pulumi.StringInput
	// Login password field CSS selector
	PasswordSelector pulumi.StringInput
	// Status of application.
	Status pulumi.StringPtrInput
	// Login URL
	Url pulumi.StringInput
	// A regex that further restricts URL to the specified regex
	UrlRegex pulumi.StringPtrInput
	// Login username field CSS selector
	UsernameSelector pulumi.StringInput
	// Users associated with the application
	Users ThreeFieldAppUserArrayInput
}

func (ThreeFieldAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threeFieldAppArgs)(nil)).Elem()
}
