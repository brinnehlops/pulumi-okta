// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SamlIdpSigningKey struct {
	pulumi.CustomResourceState

	Created   pulumi.StringOutput `pulumi:"created"`
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	Kid       pulumi.StringOutput `pulumi:"kid"`
	Kty       pulumi.StringOutput `pulumi:"kty"`
	Use       pulumi.StringOutput `pulumi:"use"`
	// base64-encoded X.509 certificate chain with DER encoding
	X5cs    pulumi.StringArrayOutput `pulumi:"x5cs"`
	X5tS256 pulumi.StringOutput      `pulumi:"x5tS256"`
}

// NewSamlIdpSigningKey registers a new resource with the given unique name, arguments, and options.
func NewSamlIdpSigningKey(ctx *pulumi.Context,
	name string, args *SamlIdpSigningKeyArgs, opts ...pulumi.ResourceOption) (*SamlIdpSigningKey, error) {
	if args == nil || args.X5cs == nil {
		return nil, errors.New("missing required argument 'X5cs'")
	}
	if args == nil {
		args = &SamlIdpSigningKeyArgs{}
	}
	var resource SamlIdpSigningKey
	err := ctx.RegisterResource("okta:deprecated/samlIdpSigningKey:SamlIdpSigningKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlIdpSigningKey gets an existing SamlIdpSigningKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlIdpSigningKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlIdpSigningKeyState, opts ...pulumi.ResourceOption) (*SamlIdpSigningKey, error) {
	var resource SamlIdpSigningKey
	err := ctx.ReadResource("okta:deprecated/samlIdpSigningKey:SamlIdpSigningKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlIdpSigningKey resources.
type samlIdpSigningKeyState struct {
	Created   *string `pulumi:"created"`
	ExpiresAt *string `pulumi:"expiresAt"`
	Kid       *string `pulumi:"kid"`
	Kty       *string `pulumi:"kty"`
	Use       *string `pulumi:"use"`
	// base64-encoded X.509 certificate chain with DER encoding
	X5cs    []string `pulumi:"x5cs"`
	X5tS256 *string  `pulumi:"x5tS256"`
}

type SamlIdpSigningKeyState struct {
	Created   pulumi.StringPtrInput
	ExpiresAt pulumi.StringPtrInput
	Kid       pulumi.StringPtrInput
	Kty       pulumi.StringPtrInput
	Use       pulumi.StringPtrInput
	// base64-encoded X.509 certificate chain with DER encoding
	X5cs    pulumi.StringArrayInput
	X5tS256 pulumi.StringPtrInput
}

func (SamlIdpSigningKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlIdpSigningKeyState)(nil)).Elem()
}

type samlIdpSigningKeyArgs struct {
	// base64-encoded X.509 certificate chain with DER encoding
	X5cs []string `pulumi:"x5cs"`
}

// The set of arguments for constructing a SamlIdpSigningKey resource.
type SamlIdpSigningKeyArgs struct {
	// base64-encoded X.509 certificate chain with DER encoding
	X5cs pulumi.StringArrayInput
}

func (SamlIdpSigningKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlIdpSigningKeyArgs)(nil)).Elem()
}
