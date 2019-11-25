// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Bsaic Auth Application.
// 
// This resource allows you to create and configure a Basic Auth Application.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_basic_auth.html.markdown.
type BasicAuth struct {
	s *pulumi.ResourceState
}

// NewBasicAuth registers a new resource with the given unique name, arguments, and options.
func NewBasicAuth(ctx *pulumi.Context,
	name string, args *BasicAuthArgs, opts ...pulumi.ResourceOpt) (*BasicAuth, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["authUrl"] = nil
		inputs["autoSubmitToolbar"] = nil
		inputs["groups"] = nil
		inputs["hideIos"] = nil
		inputs["hideWeb"] = nil
		inputs["label"] = nil
		inputs["status"] = nil
		inputs["url"] = nil
		inputs["users"] = nil
	} else {
		inputs["authUrl"] = args.AuthUrl
		inputs["autoSubmitToolbar"] = args.AutoSubmitToolbar
		inputs["groups"] = args.Groups
		inputs["hideIos"] = args.HideIos
		inputs["hideWeb"] = args.HideWeb
		inputs["label"] = args.Label
		inputs["status"] = args.Status
		inputs["url"] = args.Url
		inputs["users"] = args.Users
	}
	inputs["name"] = nil
	inputs["signOnMode"] = nil
	s, err := ctx.RegisterResource("okta:app/basicAuth:BasicAuth", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BasicAuth{s: s}, nil
}

// GetBasicAuth gets an existing BasicAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicAuth(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BasicAuthState, opts ...pulumi.ResourceOpt) (*BasicAuth, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["authUrl"] = state.AuthUrl
		inputs["autoSubmitToolbar"] = state.AutoSubmitToolbar
		inputs["groups"] = state.Groups
		inputs["hideIos"] = state.HideIos
		inputs["hideWeb"] = state.HideWeb
		inputs["label"] = state.Label
		inputs["name"] = state.Name
		inputs["signOnMode"] = state.SignOnMode
		inputs["status"] = state.Status
		inputs["url"] = state.Url
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("okta:app/basicAuth:BasicAuth", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BasicAuth{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BasicAuth) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BasicAuth) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The URL of the authenticating site for this app.
func (r *BasicAuth) AuthUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authUrl"])
}

// Display auto submit toolbar
func (r *BasicAuth) AutoSubmitToolbar() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoSubmitToolbar"])
}

// Groups associated with the application
func (r *BasicAuth) Groups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groups"])
}

// Do not display application icon on mobile app
func (r *BasicAuth) HideIos() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideIos"])
}

// Do not display application icon to users
func (r *BasicAuth) HideWeb() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["hideWeb"])
}

// The Application's display name.
func (r *BasicAuth) Label() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["label"])
}

// name of app.
func (r *BasicAuth) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Sign on mode of application.
func (r *BasicAuth) SignOnMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["signOnMode"])
}

// Status of application.
func (r *BasicAuth) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The URL of the sign-in page for this app.
func (r *BasicAuth) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Users associated with the application
func (r *BasicAuth) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering BasicAuth resources.
type BasicAuthState struct {
	// The URL of the authenticating site for this app.
	AuthUrl interface{}
	// Display auto submit toolbar
	AutoSubmitToolbar interface{}
	// Groups associated with the application
	Groups interface{}
	// Do not display application icon on mobile app
	HideIos interface{}
	// Do not display application icon to users
	HideWeb interface{}
	// The Application's display name.
	Label interface{}
	// name of app.
	Name interface{}
	// Sign on mode of application.
	SignOnMode interface{}
	// Status of application.
	Status interface{}
	// The URL of the sign-in page for this app.
	Url interface{}
	// Users associated with the application
	Users interface{}
}

// The set of arguments for constructing a BasicAuth resource.
type BasicAuthArgs struct {
	// The URL of the authenticating site for this app.
	AuthUrl interface{}
	// Display auto submit toolbar
	AutoSubmitToolbar interface{}
	// Groups associated with the application
	Groups interface{}
	// Do not display application icon on mobile app
	HideIos interface{}
	// Do not display application icon to users
	HideWeb interface{}
	// The Application's display name.
	Label interface{}
	// Status of application.
	Status interface{}
	// The URL of the sign-in page for this app.
	Url interface{}
	// Users associated with the application
	Users interface{}
}
