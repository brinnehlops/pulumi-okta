// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an MFA Policy.
//
// This resource allows you to create and configure an MFA Policy.
type Mfa struct {
	pulumi.CustomResourceState

	// Policy Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// DUO MFA policy settings.
	Duo MfaDuoPtrOutput `pulumi:"duo"`
	// Fido U2F MFA policy settings.
	FidoU2f MfaFidoU2fPtrOutput `pulumi:"fidoU2f"`
	// Fido Web Authn MFA policy settings.
	FidoWebauthn MfaFidoWebauthnPtrOutput `pulumi:"fidoWebauthn"`
	// Google OTP MFA policy settings.
	GoogleOtp MfaGoogleOtpPtrOutput `pulumi:"googleOtp"`
	// List of Group IDs to Include.
	GroupsIncludeds pulumi.StringArrayOutput `pulumi:"groupsIncludeds"`
	// Policy Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Okta Call MFA policy settings.
	OktaCall MfaOktaCallPtrOutput `pulumi:"oktaCall"`
	// Okta OTP MFA policy settings.
	OktaOtp MfaOktaOtpPtrOutput `pulumi:"oktaOtp"`
	// Okta Password MFA policy settings.
	OktaPassword MfaOktaPasswordPtrOutput `pulumi:"oktaPassword"`
	// Okta Push MFA policy settings.
	OktaPush MfaOktaPushPtrOutput `pulumi:"oktaPush"`
	// Okta Question MFA policy settings.
	OktaQuestion MfaOktaQuestionPtrOutput `pulumi:"oktaQuestion"`
	// Okta SMS MFA policy settings.
	OktaSms MfaOktaSmsPtrOutput `pulumi:"oktaSms"`
	// Priority of the policy.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// RSA Token MFA policy settings.
	RsaToken MfaRsaTokenPtrOutput `pulumi:"rsaToken"`
	// Policy Status: `"ACTIVE"` or `"INACTIVE"`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Symantec VIP MFA policy settings.
	SymantecVip MfaSymantecVipPtrOutput `pulumi:"symantecVip"`
	// Yubikey Token MFA policy settings.
	YubikeyToken MfaYubikeyTokenPtrOutput `pulumi:"yubikeyToken"`
}

// NewMfa registers a new resource with the given unique name, arguments, and options.
func NewMfa(ctx *pulumi.Context,
	name string, args *MfaArgs, opts ...pulumi.ResourceOption) (*Mfa, error) {
	if args == nil {
		args = &MfaArgs{}
	}
	var resource Mfa
	err := ctx.RegisterResource("okta:policy/mfa:Mfa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfa gets an existing Mfa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaState, opts ...pulumi.ResourceOption) (*Mfa, error) {
	var resource Mfa
	err := ctx.ReadResource("okta:policy/mfa:Mfa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mfa resources.
type mfaState struct {
	// Policy Description.
	Description *string `pulumi:"description"`
	// DUO MFA policy settings.
	Duo *MfaDuo `pulumi:"duo"`
	// Fido U2F MFA policy settings.
	FidoU2f *MfaFidoU2f `pulumi:"fidoU2f"`
	// Fido Web Authn MFA policy settings.
	FidoWebauthn *MfaFidoWebauthn `pulumi:"fidoWebauthn"`
	// Google OTP MFA policy settings.
	GoogleOtp *MfaGoogleOtp `pulumi:"googleOtp"`
	// List of Group IDs to Include.
	GroupsIncludeds []string `pulumi:"groupsIncludeds"`
	// Policy Name.
	Name *string `pulumi:"name"`
	// Okta Call MFA policy settings.
	OktaCall *MfaOktaCall `pulumi:"oktaCall"`
	// Okta OTP MFA policy settings.
	OktaOtp *MfaOktaOtp `pulumi:"oktaOtp"`
	// Okta Password MFA policy settings.
	OktaPassword *MfaOktaPassword `pulumi:"oktaPassword"`
	// Okta Push MFA policy settings.
	OktaPush *MfaOktaPush `pulumi:"oktaPush"`
	// Okta Question MFA policy settings.
	OktaQuestion *MfaOktaQuestion `pulumi:"oktaQuestion"`
	// Okta SMS MFA policy settings.
	OktaSms *MfaOktaSms `pulumi:"oktaSms"`
	// Priority of the policy.
	Priority *int `pulumi:"priority"`
	// RSA Token MFA policy settings.
	RsaToken *MfaRsaToken `pulumi:"rsaToken"`
	// Policy Status: `"ACTIVE"` or `"INACTIVE"`.
	Status *string `pulumi:"status"`
	// Symantec VIP MFA policy settings.
	SymantecVip *MfaSymantecVip `pulumi:"symantecVip"`
	// Yubikey Token MFA policy settings.
	YubikeyToken *MfaYubikeyToken `pulumi:"yubikeyToken"`
}

type MfaState struct {
	// Policy Description.
	Description pulumi.StringPtrInput
	// DUO MFA policy settings.
	Duo MfaDuoPtrInput
	// Fido U2F MFA policy settings.
	FidoU2f MfaFidoU2fPtrInput
	// Fido Web Authn MFA policy settings.
	FidoWebauthn MfaFidoWebauthnPtrInput
	// Google OTP MFA policy settings.
	GoogleOtp MfaGoogleOtpPtrInput
	// List of Group IDs to Include.
	GroupsIncludeds pulumi.StringArrayInput
	// Policy Name.
	Name pulumi.StringPtrInput
	// Okta Call MFA policy settings.
	OktaCall MfaOktaCallPtrInput
	// Okta OTP MFA policy settings.
	OktaOtp MfaOktaOtpPtrInput
	// Okta Password MFA policy settings.
	OktaPassword MfaOktaPasswordPtrInput
	// Okta Push MFA policy settings.
	OktaPush MfaOktaPushPtrInput
	// Okta Question MFA policy settings.
	OktaQuestion MfaOktaQuestionPtrInput
	// Okta SMS MFA policy settings.
	OktaSms MfaOktaSmsPtrInput
	// Priority of the policy.
	Priority pulumi.IntPtrInput
	// RSA Token MFA policy settings.
	RsaToken MfaRsaTokenPtrInput
	// Policy Status: `"ACTIVE"` or `"INACTIVE"`.
	Status pulumi.StringPtrInput
	// Symantec VIP MFA policy settings.
	SymantecVip MfaSymantecVipPtrInput
	// Yubikey Token MFA policy settings.
	YubikeyToken MfaYubikeyTokenPtrInput
}

func (MfaState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaState)(nil)).Elem()
}

type mfaArgs struct {
	// Policy Description.
	Description *string `pulumi:"description"`
	// DUO MFA policy settings.
	Duo *MfaDuo `pulumi:"duo"`
	// Fido U2F MFA policy settings.
	FidoU2f *MfaFidoU2f `pulumi:"fidoU2f"`
	// Fido Web Authn MFA policy settings.
	FidoWebauthn *MfaFidoWebauthn `pulumi:"fidoWebauthn"`
	// Google OTP MFA policy settings.
	GoogleOtp *MfaGoogleOtp `pulumi:"googleOtp"`
	// List of Group IDs to Include.
	GroupsIncludeds []string `pulumi:"groupsIncludeds"`
	// Policy Name.
	Name *string `pulumi:"name"`
	// Okta Call MFA policy settings.
	OktaCall *MfaOktaCall `pulumi:"oktaCall"`
	// Okta OTP MFA policy settings.
	OktaOtp *MfaOktaOtp `pulumi:"oktaOtp"`
	// Okta Password MFA policy settings.
	OktaPassword *MfaOktaPassword `pulumi:"oktaPassword"`
	// Okta Push MFA policy settings.
	OktaPush *MfaOktaPush `pulumi:"oktaPush"`
	// Okta Question MFA policy settings.
	OktaQuestion *MfaOktaQuestion `pulumi:"oktaQuestion"`
	// Okta SMS MFA policy settings.
	OktaSms *MfaOktaSms `pulumi:"oktaSms"`
	// Priority of the policy.
	Priority *int `pulumi:"priority"`
	// RSA Token MFA policy settings.
	RsaToken *MfaRsaToken `pulumi:"rsaToken"`
	// Policy Status: `"ACTIVE"` or `"INACTIVE"`.
	Status *string `pulumi:"status"`
	// Symantec VIP MFA policy settings.
	SymantecVip *MfaSymantecVip `pulumi:"symantecVip"`
	// Yubikey Token MFA policy settings.
	YubikeyToken *MfaYubikeyToken `pulumi:"yubikeyToken"`
}

// The set of arguments for constructing a Mfa resource.
type MfaArgs struct {
	// Policy Description.
	Description pulumi.StringPtrInput
	// DUO MFA policy settings.
	Duo MfaDuoPtrInput
	// Fido U2F MFA policy settings.
	FidoU2f MfaFidoU2fPtrInput
	// Fido Web Authn MFA policy settings.
	FidoWebauthn MfaFidoWebauthnPtrInput
	// Google OTP MFA policy settings.
	GoogleOtp MfaGoogleOtpPtrInput
	// List of Group IDs to Include.
	GroupsIncludeds pulumi.StringArrayInput
	// Policy Name.
	Name pulumi.StringPtrInput
	// Okta Call MFA policy settings.
	OktaCall MfaOktaCallPtrInput
	// Okta OTP MFA policy settings.
	OktaOtp MfaOktaOtpPtrInput
	// Okta Password MFA policy settings.
	OktaPassword MfaOktaPasswordPtrInput
	// Okta Push MFA policy settings.
	OktaPush MfaOktaPushPtrInput
	// Okta Question MFA policy settings.
	OktaQuestion MfaOktaQuestionPtrInput
	// Okta SMS MFA policy settings.
	OktaSms MfaOktaSmsPtrInput
	// Priority of the policy.
	Priority pulumi.IntPtrInput
	// RSA Token MFA policy settings.
	RsaToken MfaRsaTokenPtrInput
	// Policy Status: `"ACTIVE"` or `"INACTIVE"`.
	Status pulumi.StringPtrInput
	// Symantec VIP MFA policy settings.
	SymantecVip MfaSymantecVipPtrInput
	// Yubikey Token MFA policy settings.
	YubikeyToken MfaYubikeyTokenPtrInput
}

func (MfaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaArgs)(nil)).Elem()
}
