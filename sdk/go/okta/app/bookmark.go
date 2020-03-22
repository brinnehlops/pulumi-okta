// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package app

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Bookmark Application.
//
// This resource allows you to create and configure a Bookmark Application.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_bookmark.html.markdown.
type Bookmark struct {
	pulumi.CustomResourceState

	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrOutput `pulumi:"autoSubmitToolbar"`
	// Groups associated with the application
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrOutput `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrOutput `pulumi:"hideWeb"`
	// The Application's display name.
	Label pulumi.StringOutput `pulumi:"label"`
	// name of app.
	Name pulumi.StringOutput `pulumi:"name"`
	// Would you like Okta to add an integration for this app?
	RequestIntegration pulumi.BoolPtrOutput `pulumi:"requestIntegration"`
	// Sign on mode of application.
	SignOnMode pulumi.StringOutput `pulumi:"signOnMode"`
	// Status of application.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The URL of the bookmark.
	Url pulumi.StringOutput `pulumi:"url"`
	// Users associated with the application
	Users BookmarkUserArrayOutput `pulumi:"users"`
}

// NewBookmark registers a new resource with the given unique name, arguments, and options.
func NewBookmark(ctx *pulumi.Context,
	name string, args *BookmarkArgs, opts ...pulumi.ResourceOption) (*Bookmark, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &BookmarkArgs{}
	}
	var resource Bookmark
	err := ctx.RegisterResource("okta:app/bookmark:Bookmark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBookmark gets an existing Bookmark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBookmark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkState, opts ...pulumi.ResourceOption) (*Bookmark, error) {
	var resource Bookmark
	err := ctx.ReadResource("okta:app/bookmark:Bookmark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bookmark resources.
type bookmarkState struct {
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// The Application's display name.
	Label *string `pulumi:"label"`
	// name of app.
	Name *string `pulumi:"name"`
	// Would you like Okta to add an integration for this app?
	RequestIntegration *bool `pulumi:"requestIntegration"`
	// Sign on mode of application.
	SignOnMode *string `pulumi:"signOnMode"`
	// Status of application.
	Status *string `pulumi:"status"`
	// The URL of the bookmark.
	Url *string `pulumi:"url"`
	// Users associated with the application
	Users []BookmarkUser `pulumi:"users"`
}

type BookmarkState struct {
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// The Application's display name.
	Label pulumi.StringPtrInput
	// name of app.
	Name pulumi.StringPtrInput
	// Would you like Okta to add an integration for this app?
	RequestIntegration pulumi.BoolPtrInput
	// Sign on mode of application.
	SignOnMode pulumi.StringPtrInput
	// Status of application.
	Status pulumi.StringPtrInput
	// The URL of the bookmark.
	Url pulumi.StringPtrInput
	// Users associated with the application
	Users BookmarkUserArrayInput
}

func (BookmarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkState)(nil)).Elem()
}

type bookmarkArgs struct {
	// Display auto submit toolbar
	AutoSubmitToolbar *bool `pulumi:"autoSubmitToolbar"`
	// Groups associated with the application
	Groups []string `pulumi:"groups"`
	// Do not display application icon on mobile app
	HideIos *bool `pulumi:"hideIos"`
	// Do not display application icon to users
	HideWeb *bool `pulumi:"hideWeb"`
	// The Application's display name.
	Label string `pulumi:"label"`
	// Would you like Okta to add an integration for this app?
	RequestIntegration *bool `pulumi:"requestIntegration"`
	// Status of application.
	Status *string `pulumi:"status"`
	// The URL of the bookmark.
	Url string `pulumi:"url"`
	// Users associated with the application
	Users []BookmarkUser `pulumi:"users"`
}

// The set of arguments for constructing a Bookmark resource.
type BookmarkArgs struct {
	// Display auto submit toolbar
	AutoSubmitToolbar pulumi.BoolPtrInput
	// Groups associated with the application
	Groups pulumi.StringArrayInput
	// Do not display application icon on mobile app
	HideIos pulumi.BoolPtrInput
	// Do not display application icon to users
	HideWeb pulumi.BoolPtrInput
	// The Application's display name.
	Label pulumi.StringInput
	// Would you like Okta to add an integration for this app?
	RequestIntegration pulumi.BoolPtrInput
	// Status of application.
	Status pulumi.StringPtrInput
	// The URL of the bookmark.
	Url pulumi.StringInput
	// Users associated with the application
	Users BookmarkUserArrayInput
}

func (BookmarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkArgs)(nil)).Elem()
}

