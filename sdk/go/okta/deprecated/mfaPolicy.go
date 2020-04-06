// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package deprecated

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type MfaPolicy struct {
	pulumi.CustomResourceState

	// Policy Description
	Description  pulumi.StringPtrOutput         `pulumi:"description"`
	Duo          MfaPolicyDuoPtrOutput          `pulumi:"duo"`
	FidoU2f      MfaPolicyFidoU2fPtrOutput      `pulumi:"fidoU2f"`
	FidoWebauthn MfaPolicyFidoWebauthnPtrOutput `pulumi:"fidoWebauthn"`
	GoogleOtp    MfaPolicyGoogleOtpPtrOutput    `pulumi:"googleOtp"`
	// List of Group IDs to Include
	GroupsIncludeds pulumi.StringArrayOutput `pulumi:"groupsIncludeds"`
	// Policy Name
	Name         pulumi.StringOutput            `pulumi:"name"`
	OktaCall     MfaPolicyOktaCallPtrOutput     `pulumi:"oktaCall"`
	OktaOtp      MfaPolicyOktaOtpPtrOutput      `pulumi:"oktaOtp"`
	OktaPassword MfaPolicyOktaPasswordPtrOutput `pulumi:"oktaPassword"`
	OktaPush     MfaPolicyOktaPushPtrOutput     `pulumi:"oktaPush"`
	OktaQuestion MfaPolicyOktaQuestionPtrOutput `pulumi:"oktaQuestion"`
	OktaSms      MfaPolicyOktaSmsPtrOutput      `pulumi:"oktaSms"`
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrOutput        `pulumi:"priority"`
	RsaToken MfaPolicyRsaTokenPtrOutput `pulumi:"rsaToken"`
	// Policy Status: ACTIVE or INACTIVE.
	Status       pulumi.StringPtrOutput         `pulumi:"status"`
	SymantecVip  MfaPolicySymantecVipPtrOutput  `pulumi:"symantecVip"`
	YubikeyToken MfaPolicyYubikeyTokenPtrOutput `pulumi:"yubikeyToken"`
}

// NewMfaPolicy registers a new resource with the given unique name, arguments, and options.
func NewMfaPolicy(ctx *pulumi.Context,
	name string, args *MfaPolicyArgs, opts ...pulumi.ResourceOption) (*MfaPolicy, error) {
	if args == nil {
		args = &MfaPolicyArgs{}
	}
	var resource MfaPolicy
	err := ctx.RegisterResource("okta:deprecated/mfaPolicy:MfaPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaPolicy gets an existing MfaPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaPolicyState, opts ...pulumi.ResourceOption) (*MfaPolicy, error) {
	var resource MfaPolicy
	err := ctx.ReadResource("okta:deprecated/mfaPolicy:MfaPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaPolicy resources.
type mfaPolicyState struct {
	// Policy Description
	Description  *string                `pulumi:"description"`
	Duo          *MfaPolicyDuo          `pulumi:"duo"`
	FidoU2f      *MfaPolicyFidoU2f      `pulumi:"fidoU2f"`
	FidoWebauthn *MfaPolicyFidoWebauthn `pulumi:"fidoWebauthn"`
	GoogleOtp    *MfaPolicyGoogleOtp    `pulumi:"googleOtp"`
	// List of Group IDs to Include
	GroupsIncludeds []string `pulumi:"groupsIncludeds"`
	// Policy Name
	Name         *string                `pulumi:"name"`
	OktaCall     *MfaPolicyOktaCall     `pulumi:"oktaCall"`
	OktaOtp      *MfaPolicyOktaOtp      `pulumi:"oktaOtp"`
	OktaPassword *MfaPolicyOktaPassword `pulumi:"oktaPassword"`
	OktaPush     *MfaPolicyOktaPush     `pulumi:"oktaPush"`
	OktaQuestion *MfaPolicyOktaQuestion `pulumi:"oktaQuestion"`
	OktaSms      *MfaPolicyOktaSms      `pulumi:"oktaSms"`
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority *int               `pulumi:"priority"`
	RsaToken *MfaPolicyRsaToken `pulumi:"rsaToken"`
	// Policy Status: ACTIVE or INACTIVE.
	Status       *string                `pulumi:"status"`
	SymantecVip  *MfaPolicySymantecVip  `pulumi:"symantecVip"`
	YubikeyToken *MfaPolicyYubikeyToken `pulumi:"yubikeyToken"`
}

type MfaPolicyState struct {
	// Policy Description
	Description  pulumi.StringPtrInput
	Duo          MfaPolicyDuoPtrInput
	FidoU2f      MfaPolicyFidoU2fPtrInput
	FidoWebauthn MfaPolicyFidoWebauthnPtrInput
	GoogleOtp    MfaPolicyGoogleOtpPtrInput
	// List of Group IDs to Include
	GroupsIncludeds pulumi.StringArrayInput
	// Policy Name
	Name         pulumi.StringPtrInput
	OktaCall     MfaPolicyOktaCallPtrInput
	OktaOtp      MfaPolicyOktaOtpPtrInput
	OktaPassword MfaPolicyOktaPasswordPtrInput
	OktaPush     MfaPolicyOktaPushPtrInput
	OktaQuestion MfaPolicyOktaQuestionPtrInput
	OktaSms      MfaPolicyOktaSmsPtrInput
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	RsaToken MfaPolicyRsaTokenPtrInput
	// Policy Status: ACTIVE or INACTIVE.
	Status       pulumi.StringPtrInput
	SymantecVip  MfaPolicySymantecVipPtrInput
	YubikeyToken MfaPolicyYubikeyTokenPtrInput
}

func (MfaPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaPolicyState)(nil)).Elem()
}

type mfaPolicyArgs struct {
	// Policy Description
	Description  *string                `pulumi:"description"`
	Duo          *MfaPolicyDuo          `pulumi:"duo"`
	FidoU2f      *MfaPolicyFidoU2f      `pulumi:"fidoU2f"`
	FidoWebauthn *MfaPolicyFidoWebauthn `pulumi:"fidoWebauthn"`
	GoogleOtp    *MfaPolicyGoogleOtp    `pulumi:"googleOtp"`
	// List of Group IDs to Include
	GroupsIncludeds []string `pulumi:"groupsIncludeds"`
	// Policy Name
	Name         *string                `pulumi:"name"`
	OktaCall     *MfaPolicyOktaCall     `pulumi:"oktaCall"`
	OktaOtp      *MfaPolicyOktaOtp      `pulumi:"oktaOtp"`
	OktaPassword *MfaPolicyOktaPassword `pulumi:"oktaPassword"`
	OktaPush     *MfaPolicyOktaPush     `pulumi:"oktaPush"`
	OktaQuestion *MfaPolicyOktaQuestion `pulumi:"oktaQuestion"`
	OktaSms      *MfaPolicyOktaSms      `pulumi:"oktaSms"`
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority *int               `pulumi:"priority"`
	RsaToken *MfaPolicyRsaToken `pulumi:"rsaToken"`
	// Policy Status: ACTIVE or INACTIVE.
	Status       *string                `pulumi:"status"`
	SymantecVip  *MfaPolicySymantecVip  `pulumi:"symantecVip"`
	YubikeyToken *MfaPolicyYubikeyToken `pulumi:"yubikeyToken"`
}

// The set of arguments for constructing a MfaPolicy resource.
type MfaPolicyArgs struct {
	// Policy Description
	Description  pulumi.StringPtrInput
	Duo          MfaPolicyDuoPtrInput
	FidoU2f      MfaPolicyFidoU2fPtrInput
	FidoWebauthn MfaPolicyFidoWebauthnPtrInput
	GoogleOtp    MfaPolicyGoogleOtpPtrInput
	// List of Group IDs to Include
	GroupsIncludeds pulumi.StringArrayInput
	// Policy Name
	Name         pulumi.StringPtrInput
	OktaCall     MfaPolicyOktaCallPtrInput
	OktaOtp      MfaPolicyOktaOtpPtrInput
	OktaPassword MfaPolicyOktaPasswordPtrInput
	OktaPush     MfaPolicyOktaPushPtrInput
	OktaQuestion MfaPolicyOktaQuestionPtrInput
	OktaSms      MfaPolicyOktaSmsPtrInput
	// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
	// priority is provided. API defaults it to the last/lowest if not there.
	Priority pulumi.IntPtrInput
	RsaToken MfaPolicyRsaTokenPtrInput
	// Policy Status: ACTIVE or INACTIVE.
	Status       pulumi.StringPtrInput
	SymantecVip  MfaPolicySymantecVipPtrInput
	YubikeyToken MfaPolicyYubikeyTokenPtrInput
}

func (MfaPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaPolicyArgs)(nil)).Elem()
}
