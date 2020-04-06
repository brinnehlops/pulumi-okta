// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package idp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a SAML Identity Provider Signing Key.
//
// This resource allows you to create and configure a SAML Identity Provider Signing Key.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/idp_saml_signing_key.html.markdown.
type SamlKey struct {
	pulumi.CustomResourceState

	// Date created.
	Created pulumi.StringOutput `pulumi:"created"`
	// Date the cert expires.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// Key ID.
	Kid pulumi.StringOutput `pulumi:"kid"`
	// Identifies the cryptographic algorithm family used with the key.
	Kty pulumi.StringOutput `pulumi:"kty"`
	// Intended use of the public key.
	Use pulumi.StringOutput `pulumi:"use"`
	// base64-encoded X.509 certificate chain with DER encoding.
	X5cs pulumi.StringArrayOutput `pulumi:"x5cs"`
	// base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate.
	X5tS256 pulumi.StringOutput `pulumi:"x5tS256"`
}

// NewSamlKey registers a new resource with the given unique name, arguments, and options.
func NewSamlKey(ctx *pulumi.Context,
	name string, args *SamlKeyArgs, opts ...pulumi.ResourceOption) (*SamlKey, error) {
	if args == nil || args.X5cs == nil {
		return nil, errors.New("missing required argument 'X5cs'")
	}
	if args == nil {
		args = &SamlKeyArgs{}
	}
	var resource SamlKey
	err := ctx.RegisterResource("okta:idp/samlKey:SamlKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlKey gets an existing SamlKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlKeyState, opts ...pulumi.ResourceOption) (*SamlKey, error) {
	var resource SamlKey
	err := ctx.ReadResource("okta:idp/samlKey:SamlKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlKey resources.
type samlKeyState struct {
	// Date created.
	Created *string `pulumi:"created"`
	// Date the cert expires.
	ExpiresAt *string `pulumi:"expiresAt"`
	// Key ID.
	Kid *string `pulumi:"kid"`
	// Identifies the cryptographic algorithm family used with the key.
	Kty *string `pulumi:"kty"`
	// Intended use of the public key.
	Use *string `pulumi:"use"`
	// base64-encoded X.509 certificate chain with DER encoding.
	X5cs []string `pulumi:"x5cs"`
	// base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate.
	X5tS256 *string `pulumi:"x5tS256"`
}

type SamlKeyState struct {
	// Date created.
	Created pulumi.StringPtrInput
	// Date the cert expires.
	ExpiresAt pulumi.StringPtrInput
	// Key ID.
	Kid pulumi.StringPtrInput
	// Identifies the cryptographic algorithm family used with the key.
	Kty pulumi.StringPtrInput
	// Intended use of the public key.
	Use pulumi.StringPtrInput
	// base64-encoded X.509 certificate chain with DER encoding.
	X5cs pulumi.StringArrayInput
	// base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate.
	X5tS256 pulumi.StringPtrInput
}

func (SamlKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlKeyState)(nil)).Elem()
}

type samlKeyArgs struct {
	// base64-encoded X.509 certificate chain with DER encoding.
	X5cs []string `pulumi:"x5cs"`
}

// The set of arguments for constructing a SamlKey resource.
type SamlKeyArgs struct {
	// base64-encoded X.509 certificate chain with DER encoding.
	X5cs pulumi.StringArrayInput
}

func (SamlKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlKeyArgs)(nil)).Elem()
}
