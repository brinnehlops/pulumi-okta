// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SamlIdp struct {
	pulumi.CustomResourceState

	AccountLinkAction        pulumi.StringPtrOutput   `pulumi:"accountLinkAction"`
	AccountLinkGroupIncludes pulumi.StringArrayOutput `pulumi:"accountLinkGroupIncludes"`
	AcsBinding               pulumi.StringOutput      `pulumi:"acsBinding"`
	AcsType                  pulumi.StringPtrOutput   `pulumi:"acsType"`
	Audience                 pulumi.StringOutput      `pulumi:"audience"`
	DeprovisionedAction      pulumi.StringPtrOutput   `pulumi:"deprovisionedAction"`
	GroupsAction             pulumi.StringPtrOutput   `pulumi:"groupsAction"`
	GroupsAssignments        pulumi.StringArrayOutput `pulumi:"groupsAssignments"`
	GroupsAttribute          pulumi.StringPtrOutput   `pulumi:"groupsAttribute"`
	GroupsFilters            pulumi.StringArrayOutput `pulumi:"groupsFilters"`
	Issuer                   pulumi.StringOutput      `pulumi:"issuer"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode pulumi.StringPtrOutput `pulumi:"issuerMode"`
	Kid        pulumi.StringOutput    `pulumi:"kid"`
	// name of idp
	Name               pulumi.StringOutput    `pulumi:"name"`
	NameFormat         pulumi.StringPtrOutput `pulumi:"nameFormat"`
	ProfileMaster      pulumi.BoolPtrOutput   `pulumi:"profileMaster"`
	ProvisioningAction pulumi.StringPtrOutput `pulumi:"provisioningAction"`
	// algorithm to use to sign requests
	RequestSignatureAlgorithm pulumi.StringPtrOutput `pulumi:"requestSignatureAlgorithm"`
	// algorithm to use to sign response
	RequestSignatureScope pulumi.StringPtrOutput `pulumi:"requestSignatureScope"`
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm pulumi.StringPtrOutput `pulumi:"responseSignatureAlgorithm"`
	// algorithm to use to sign response
	ResponseSignatureScope pulumi.StringPtrOutput   `pulumi:"responseSignatureScope"`
	SsoBinding             pulumi.StringPtrOutput   `pulumi:"ssoBinding"`
	SsoDestination         pulumi.StringPtrOutput   `pulumi:"ssoDestination"`
	SsoUrl                 pulumi.StringOutput      `pulumi:"ssoUrl"`
	Status                 pulumi.StringPtrOutput   `pulumi:"status"`
	SubjectFilter          pulumi.StringPtrOutput   `pulumi:"subjectFilter"`
	SubjectFormats         pulumi.StringArrayOutput `pulumi:"subjectFormats"`
	SubjectMatchAttribute  pulumi.StringPtrOutput   `pulumi:"subjectMatchAttribute"`
	SubjectMatchType       pulumi.StringPtrOutput   `pulumi:"subjectMatchType"`
	SuspendedAction        pulumi.StringPtrOutput   `pulumi:"suspendedAction"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
	UsernameTemplate       pulumi.StringPtrOutput   `pulumi:"usernameTemplate"`
}

// NewSamlIdp registers a new resource with the given unique name, arguments, and options.
func NewSamlIdp(ctx *pulumi.Context,
	name string, args *SamlIdpArgs, opts ...pulumi.ResourceOption) (*SamlIdp, error) {
	if args == nil || args.AcsBinding == nil {
		return nil, errors.New("missing required argument 'AcsBinding'")
	}
	if args == nil || args.Issuer == nil {
		return nil, errors.New("missing required argument 'Issuer'")
	}
	if args == nil || args.Kid == nil {
		return nil, errors.New("missing required argument 'Kid'")
	}
	if args == nil || args.SsoUrl == nil {
		return nil, errors.New("missing required argument 'SsoUrl'")
	}
	if args == nil {
		args = &SamlIdpArgs{}
	}
	var resource SamlIdp
	err := ctx.RegisterResource("okta:deprecated/samlIdp:SamlIdp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlIdp gets an existing SamlIdp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlIdp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlIdpState, opts ...pulumi.ResourceOption) (*SamlIdp, error) {
	var resource SamlIdp
	err := ctx.ReadResource("okta:deprecated/samlIdp:SamlIdp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlIdp resources.
type samlIdpState struct {
	AccountLinkAction        *string  `pulumi:"accountLinkAction"`
	AccountLinkGroupIncludes []string `pulumi:"accountLinkGroupIncludes"`
	AcsBinding               *string  `pulumi:"acsBinding"`
	AcsType                  *string  `pulumi:"acsType"`
	Audience                 *string  `pulumi:"audience"`
	DeprovisionedAction      *string  `pulumi:"deprovisionedAction"`
	GroupsAction             *string  `pulumi:"groupsAction"`
	GroupsAssignments        []string `pulumi:"groupsAssignments"`
	GroupsAttribute          *string  `pulumi:"groupsAttribute"`
	GroupsFilters            []string `pulumi:"groupsFilters"`
	Issuer                   *string  `pulumi:"issuer"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode *string `pulumi:"issuerMode"`
	Kid        *string `pulumi:"kid"`
	// name of idp
	Name               *string `pulumi:"name"`
	NameFormat         *string `pulumi:"nameFormat"`
	ProfileMaster      *bool   `pulumi:"profileMaster"`
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// algorithm to use to sign requests
	RequestSignatureAlgorithm *string `pulumi:"requestSignatureAlgorithm"`
	// algorithm to use to sign response
	RequestSignatureScope *string `pulumi:"requestSignatureScope"`
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// algorithm to use to sign response
	ResponseSignatureScope *string  `pulumi:"responseSignatureScope"`
	SsoBinding             *string  `pulumi:"ssoBinding"`
	SsoDestination         *string  `pulumi:"ssoDestination"`
	SsoUrl                 *string  `pulumi:"ssoUrl"`
	Status                 *string  `pulumi:"status"`
	SubjectFilter          *string  `pulumi:"subjectFilter"`
	SubjectFormats         []string `pulumi:"subjectFormats"`
	SubjectMatchAttribute  *string  `pulumi:"subjectMatchAttribute"`
	SubjectMatchType       *string  `pulumi:"subjectMatchType"`
	SuspendedAction        *string  `pulumi:"suspendedAction"`
	Type                   *string  `pulumi:"type"`
	UsernameTemplate       *string  `pulumi:"usernameTemplate"`
}

type SamlIdpState struct {
	AccountLinkAction        pulumi.StringPtrInput
	AccountLinkGroupIncludes pulumi.StringArrayInput
	AcsBinding               pulumi.StringPtrInput
	AcsType                  pulumi.StringPtrInput
	Audience                 pulumi.StringPtrInput
	DeprovisionedAction      pulumi.StringPtrInput
	GroupsAction             pulumi.StringPtrInput
	GroupsAssignments        pulumi.StringArrayInput
	GroupsAttribute          pulumi.StringPtrInput
	GroupsFilters            pulumi.StringArrayInput
	Issuer                   pulumi.StringPtrInput
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode pulumi.StringPtrInput
	Kid        pulumi.StringPtrInput
	// name of idp
	Name               pulumi.StringPtrInput
	NameFormat         pulumi.StringPtrInput
	ProfileMaster      pulumi.BoolPtrInput
	ProvisioningAction pulumi.StringPtrInput
	// algorithm to use to sign requests
	RequestSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
	RequestSignatureScope pulumi.StringPtrInput
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
	ResponseSignatureScope pulumi.StringPtrInput
	SsoBinding             pulumi.StringPtrInput
	SsoDestination         pulumi.StringPtrInput
	SsoUrl                 pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	SubjectFilter          pulumi.StringPtrInput
	SubjectFormats         pulumi.StringArrayInput
	SubjectMatchAttribute  pulumi.StringPtrInput
	SubjectMatchType       pulumi.StringPtrInput
	SuspendedAction        pulumi.StringPtrInput
	Type                   pulumi.StringPtrInput
	UsernameTemplate       pulumi.StringPtrInput
}

func (SamlIdpState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlIdpState)(nil)).Elem()
}

type samlIdpArgs struct {
	AccountLinkAction        *string  `pulumi:"accountLinkAction"`
	AccountLinkGroupIncludes []string `pulumi:"accountLinkGroupIncludes"`
	AcsBinding               string   `pulumi:"acsBinding"`
	AcsType                  *string  `pulumi:"acsType"`
	DeprovisionedAction      *string  `pulumi:"deprovisionedAction"`
	GroupsAction             *string  `pulumi:"groupsAction"`
	GroupsAssignments        []string `pulumi:"groupsAssignments"`
	GroupsAttribute          *string  `pulumi:"groupsAttribute"`
	GroupsFilters            []string `pulumi:"groupsFilters"`
	Issuer                   string   `pulumi:"issuer"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode *string `pulumi:"issuerMode"`
	Kid        string  `pulumi:"kid"`
	// name of idp
	Name               *string `pulumi:"name"`
	NameFormat         *string `pulumi:"nameFormat"`
	ProfileMaster      *bool   `pulumi:"profileMaster"`
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// algorithm to use to sign requests
	RequestSignatureAlgorithm *string `pulumi:"requestSignatureAlgorithm"`
	// algorithm to use to sign response
	RequestSignatureScope *string `pulumi:"requestSignatureScope"`
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// algorithm to use to sign response
	ResponseSignatureScope *string  `pulumi:"responseSignatureScope"`
	SsoBinding             *string  `pulumi:"ssoBinding"`
	SsoDestination         *string  `pulumi:"ssoDestination"`
	SsoUrl                 string   `pulumi:"ssoUrl"`
	Status                 *string  `pulumi:"status"`
	SubjectFilter          *string  `pulumi:"subjectFilter"`
	SubjectFormats         []string `pulumi:"subjectFormats"`
	SubjectMatchAttribute  *string  `pulumi:"subjectMatchAttribute"`
	SubjectMatchType       *string  `pulumi:"subjectMatchType"`
	SuspendedAction        *string  `pulumi:"suspendedAction"`
	UsernameTemplate       *string  `pulumi:"usernameTemplate"`
}

// The set of arguments for constructing a SamlIdp resource.
type SamlIdpArgs struct {
	AccountLinkAction        pulumi.StringPtrInput
	AccountLinkGroupIncludes pulumi.StringArrayInput
	AcsBinding               pulumi.StringInput
	AcsType                  pulumi.StringPtrInput
	DeprovisionedAction      pulumi.StringPtrInput
	GroupsAction             pulumi.StringPtrInput
	GroupsAssignments        pulumi.StringArrayInput
	GroupsAttribute          pulumi.StringPtrInput
	GroupsFilters            pulumi.StringArrayInput
	Issuer                   pulumi.StringInput
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode pulumi.StringPtrInput
	Kid        pulumi.StringInput
	// name of idp
	Name               pulumi.StringPtrInput
	NameFormat         pulumi.StringPtrInput
	ProfileMaster      pulumi.BoolPtrInput
	ProvisioningAction pulumi.StringPtrInput
	// algorithm to use to sign requests
	RequestSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
	RequestSignatureScope pulumi.StringPtrInput
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// algorithm to use to sign response
	ResponseSignatureScope pulumi.StringPtrInput
	SsoBinding             pulumi.StringPtrInput
	SsoDestination         pulumi.StringPtrInput
	SsoUrl                 pulumi.StringInput
	Status                 pulumi.StringPtrInput
	SubjectFilter          pulumi.StringPtrInput
	SubjectFormats         pulumi.StringArrayInput
	SubjectMatchAttribute  pulumi.StringPtrInput
	SubjectMatchType       pulumi.StringPtrInput
	SuspendedAction        pulumi.StringPtrInput
	UsernameTemplate       pulumi.StringPtrInput
}

func (SamlIdpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlIdpArgs)(nil)).Elem()
}
