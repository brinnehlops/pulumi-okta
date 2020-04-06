// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package template

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Okta Email Template.
//
// This resource allows you to create and configure an Okta Email Template.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/template_email.html.markdown.
type Email struct {
	pulumi.CustomResourceState

	// The default language, by default is set to `"en"`.
	DefaultLanguage pulumi.StringPtrOutput `pulumi:"defaultLanguage"`
	// Set of translations for particular template.
	Translations EmailTranslationArrayOutput `pulumi:"translations"`
	// Email template type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEmail registers a new resource with the given unique name, arguments, and options.
func NewEmail(ctx *pulumi.Context,
	name string, args *EmailArgs, opts ...pulumi.ResourceOption) (*Email, error) {
	if args == nil || args.Translations == nil {
		return nil, errors.New("missing required argument 'Translations'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &EmailArgs{}
	}
	var resource Email
	err := ctx.RegisterResource("okta:template/email:Email", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmail gets an existing Email resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailState, opts ...pulumi.ResourceOption) (*Email, error) {
	var resource Email
	err := ctx.ReadResource("okta:template/email:Email", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Email resources.
type emailState struct {
	// The default language, by default is set to `"en"`.
	DefaultLanguage *string `pulumi:"defaultLanguage"`
	// Set of translations for particular template.
	Translations []EmailTranslation `pulumi:"translations"`
	// Email template type
	Type *string `pulumi:"type"`
}

type EmailState struct {
	// The default language, by default is set to `"en"`.
	DefaultLanguage pulumi.StringPtrInput
	// Set of translations for particular template.
	Translations EmailTranslationArrayInput
	// Email template type
	Type pulumi.StringPtrInput
}

func (EmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailState)(nil)).Elem()
}

type emailArgs struct {
	// The default language, by default is set to `"en"`.
	DefaultLanguage *string `pulumi:"defaultLanguage"`
	// Set of translations for particular template.
	Translations []EmailTranslation `pulumi:"translations"`
	// Email template type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Email resource.
type EmailArgs struct {
	// The default language, by default is set to `"en"`.
	DefaultLanguage pulumi.StringPtrInput
	// Set of translations for particular template.
	Translations EmailTranslationArrayInput
	// Email template type
	Type pulumi.StringInput
}

func (EmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailArgs)(nil)).Elem()
}
