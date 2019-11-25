// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Idp struct {
	s *pulumi.ResourceState
}

// NewIdp registers a new resource with the given unique name, arguments, and options.
func NewIdp(ctx *pulumi.Context,
	name string, args *IdpArgs, opts ...pulumi.ResourceOpt) (*Idp, error) {
	if args == nil || args.AcsBinding == nil {
		return nil, errors.New("missing required argument 'AcsBinding'")
	}
	if args == nil || args.AuthorizationBinding == nil {
		return nil, errors.New("missing required argument 'AuthorizationBinding'")
	}
	if args == nil || args.AuthorizationUrl == nil {
		return nil, errors.New("missing required argument 'AuthorizationUrl'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.IssuerUrl == nil {
		return nil, errors.New("missing required argument 'IssuerUrl'")
	}
	if args == nil || args.JwksBinding == nil {
		return nil, errors.New("missing required argument 'JwksBinding'")
	}
	if args == nil || args.JwksUrl == nil {
		return nil, errors.New("missing required argument 'JwksUrl'")
	}
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil || args.TokenBinding == nil {
		return nil, errors.New("missing required argument 'TokenBinding'")
	}
	if args == nil || args.TokenUrl == nil {
		return nil, errors.New("missing required argument 'TokenUrl'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountLinkAction"] = nil
		inputs["accountLinkGroupIncludes"] = nil
		inputs["acsBinding"] = nil
		inputs["acsType"] = nil
		inputs["authorizationBinding"] = nil
		inputs["authorizationUrl"] = nil
		inputs["clientId"] = nil
		inputs["clientSecret"] = nil
		inputs["deprovisionedAction"] = nil
		inputs["groupsAction"] = nil
		inputs["groupsAssignments"] = nil
		inputs["groupsAttribute"] = nil
		inputs["groupsFilters"] = nil
		inputs["issuerMode"] = nil
		inputs["issuerUrl"] = nil
		inputs["jwksBinding"] = nil
		inputs["jwksUrl"] = nil
		inputs["maxClockSkew"] = nil
		inputs["name"] = nil
		inputs["profileMaster"] = nil
		inputs["protocolType"] = nil
		inputs["provisioningAction"] = nil
		inputs["requestSignatureAlgorithm"] = nil
		inputs["requestSignatureScope"] = nil
		inputs["responseSignatureAlgorithm"] = nil
		inputs["responseSignatureScope"] = nil
		inputs["scopes"] = nil
		inputs["status"] = nil
		inputs["subjectMatchAttribute"] = nil
		inputs["subjectMatchType"] = nil
		inputs["suspendedAction"] = nil
		inputs["tokenBinding"] = nil
		inputs["tokenUrl"] = nil
		inputs["userInfoBinding"] = nil
		inputs["userInfoUrl"] = nil
		inputs["usernameTemplate"] = nil
	} else {
		inputs["accountLinkAction"] = args.AccountLinkAction
		inputs["accountLinkGroupIncludes"] = args.AccountLinkGroupIncludes
		inputs["acsBinding"] = args.AcsBinding
		inputs["acsType"] = args.AcsType
		inputs["authorizationBinding"] = args.AuthorizationBinding
		inputs["authorizationUrl"] = args.AuthorizationUrl
		inputs["clientId"] = args.ClientId
		inputs["clientSecret"] = args.ClientSecret
		inputs["deprovisionedAction"] = args.DeprovisionedAction
		inputs["groupsAction"] = args.GroupsAction
		inputs["groupsAssignments"] = args.GroupsAssignments
		inputs["groupsAttribute"] = args.GroupsAttribute
		inputs["groupsFilters"] = args.GroupsFilters
		inputs["issuerMode"] = args.IssuerMode
		inputs["issuerUrl"] = args.IssuerUrl
		inputs["jwksBinding"] = args.JwksBinding
		inputs["jwksUrl"] = args.JwksUrl
		inputs["maxClockSkew"] = args.MaxClockSkew
		inputs["name"] = args.Name
		inputs["profileMaster"] = args.ProfileMaster
		inputs["protocolType"] = args.ProtocolType
		inputs["provisioningAction"] = args.ProvisioningAction
		inputs["requestSignatureAlgorithm"] = args.RequestSignatureAlgorithm
		inputs["requestSignatureScope"] = args.RequestSignatureScope
		inputs["responseSignatureAlgorithm"] = args.ResponseSignatureAlgorithm
		inputs["responseSignatureScope"] = args.ResponseSignatureScope
		inputs["scopes"] = args.Scopes
		inputs["status"] = args.Status
		inputs["subjectMatchAttribute"] = args.SubjectMatchAttribute
		inputs["subjectMatchType"] = args.SubjectMatchType
		inputs["suspendedAction"] = args.SuspendedAction
		inputs["tokenBinding"] = args.TokenBinding
		inputs["tokenUrl"] = args.TokenUrl
		inputs["userInfoBinding"] = args.UserInfoBinding
		inputs["userInfoUrl"] = args.UserInfoUrl
		inputs["usernameTemplate"] = args.UsernameTemplate
	}
	inputs["type"] = nil
	s, err := ctx.RegisterResource("okta:deprecated/idp:Idp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Idp{s: s}, nil
}

// GetIdp gets an existing Idp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IdpState, opts ...pulumi.ResourceOpt) (*Idp, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountLinkAction"] = state.AccountLinkAction
		inputs["accountLinkGroupIncludes"] = state.AccountLinkGroupIncludes
		inputs["acsBinding"] = state.AcsBinding
		inputs["acsType"] = state.AcsType
		inputs["authorizationBinding"] = state.AuthorizationBinding
		inputs["authorizationUrl"] = state.AuthorizationUrl
		inputs["clientId"] = state.ClientId
		inputs["clientSecret"] = state.ClientSecret
		inputs["deprovisionedAction"] = state.DeprovisionedAction
		inputs["groupsAction"] = state.GroupsAction
		inputs["groupsAssignments"] = state.GroupsAssignments
		inputs["groupsAttribute"] = state.GroupsAttribute
		inputs["groupsFilters"] = state.GroupsFilters
		inputs["issuerMode"] = state.IssuerMode
		inputs["issuerUrl"] = state.IssuerUrl
		inputs["jwksBinding"] = state.JwksBinding
		inputs["jwksUrl"] = state.JwksUrl
		inputs["maxClockSkew"] = state.MaxClockSkew
		inputs["name"] = state.Name
		inputs["profileMaster"] = state.ProfileMaster
		inputs["protocolType"] = state.ProtocolType
		inputs["provisioningAction"] = state.ProvisioningAction
		inputs["requestSignatureAlgorithm"] = state.RequestSignatureAlgorithm
		inputs["requestSignatureScope"] = state.RequestSignatureScope
		inputs["responseSignatureAlgorithm"] = state.ResponseSignatureAlgorithm
		inputs["responseSignatureScope"] = state.ResponseSignatureScope
		inputs["scopes"] = state.Scopes
		inputs["status"] = state.Status
		inputs["subjectMatchAttribute"] = state.SubjectMatchAttribute
		inputs["subjectMatchType"] = state.SubjectMatchType
		inputs["suspendedAction"] = state.SuspendedAction
		inputs["tokenBinding"] = state.TokenBinding
		inputs["tokenUrl"] = state.TokenUrl
		inputs["type"] = state.Type
		inputs["userInfoBinding"] = state.UserInfoBinding
		inputs["userInfoUrl"] = state.UserInfoUrl
		inputs["usernameTemplate"] = state.UsernameTemplate
	}
	s, err := ctx.ReadResource("okta:deprecated/idp:Idp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Idp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Idp) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Idp) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Idp) AccountLinkAction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accountLinkAction"])
}

func (r *Idp) AccountLinkGroupIncludes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["accountLinkGroupIncludes"])
}

func (r *Idp) AcsBinding() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["acsBinding"])
}

func (r *Idp) AcsType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["acsType"])
}

func (r *Idp) AuthorizationBinding() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["authorizationBinding"])
}

func (r *Idp) AuthorizationUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["authorizationUrl"])
}

func (r *Idp) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

func (r *Idp) ClientSecret() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientSecret"])
}

func (r *Idp) DeprovisionedAction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["deprovisionedAction"])
}

func (r *Idp) GroupsAction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["groupsAction"])
}

func (r *Idp) GroupsAssignments() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["groupsAssignments"])
}

func (r *Idp) GroupsAttribute() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["groupsAttribute"])
}

func (r *Idp) GroupsFilters() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["groupsFilters"])
}

// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
func (r *Idp) IssuerMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["issuerMode"])
}

func (r *Idp) IssuerUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["issuerUrl"])
}

func (r *Idp) JwksBinding() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["jwksBinding"])
}

func (r *Idp) JwksUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["jwksUrl"])
}

func (r *Idp) MaxClockSkew() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxClockSkew"])
}

// name of idp
func (r *Idp) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Idp) ProfileMaster() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["profileMaster"])
}

func (r *Idp) ProtocolType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocolType"])
}

func (r *Idp) ProvisioningAction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["provisioningAction"])
}

// algorithm to use to sign requests
func (r *Idp) RequestSignatureAlgorithm() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["requestSignatureAlgorithm"])
}

// algorithm to use to sign response
func (r *Idp) RequestSignatureScope() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["requestSignatureScope"])
}

// algorithm to use to sign requests
func (r *Idp) ResponseSignatureAlgorithm() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["responseSignatureAlgorithm"])
}

// algorithm to use to sign response
func (r *Idp) ResponseSignatureScope() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["responseSignatureScope"])
}

func (r *Idp) Scopes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["scopes"])
}

func (r *Idp) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

func (r *Idp) SubjectMatchAttribute() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["subjectMatchAttribute"])
}

func (r *Idp) SubjectMatchType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["subjectMatchType"])
}

func (r *Idp) SuspendedAction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["suspendedAction"])
}

func (r *Idp) TokenBinding() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tokenBinding"])
}

func (r *Idp) TokenUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tokenUrl"])
}

func (r *Idp) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

func (r *Idp) UserInfoBinding() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userInfoBinding"])
}

func (r *Idp) UserInfoUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userInfoUrl"])
}

func (r *Idp) UsernameTemplate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["usernameTemplate"])
}

// Input properties used for looking up and filtering Idp resources.
type IdpState struct {
	AccountLinkAction interface{}
	AccountLinkGroupIncludes interface{}
	AcsBinding interface{}
	AcsType interface{}
	AuthorizationBinding interface{}
	AuthorizationUrl interface{}
	ClientId interface{}
	ClientSecret interface{}
	DeprovisionedAction interface{}
	GroupsAction interface{}
	GroupsAssignments interface{}
	GroupsAttribute interface{}
	GroupsFilters interface{}
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode interface{}
	IssuerUrl interface{}
	JwksBinding interface{}
	JwksUrl interface{}
	MaxClockSkew interface{}
	// name of idp
	Name interface{}
	ProfileMaster interface{}
	ProtocolType interface{}
	ProvisioningAction interface{}
	// algorithm to use to sign requests
	RequestSignatureAlgorithm interface{}
	// algorithm to use to sign response
	RequestSignatureScope interface{}
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm interface{}
	// algorithm to use to sign response
	ResponseSignatureScope interface{}
	Scopes interface{}
	Status interface{}
	SubjectMatchAttribute interface{}
	SubjectMatchType interface{}
	SuspendedAction interface{}
	TokenBinding interface{}
	TokenUrl interface{}
	Type interface{}
	UserInfoBinding interface{}
	UserInfoUrl interface{}
	UsernameTemplate interface{}
}

// The set of arguments for constructing a Idp resource.
type IdpArgs struct {
	AccountLinkAction interface{}
	AccountLinkGroupIncludes interface{}
	AcsBinding interface{}
	AcsType interface{}
	AuthorizationBinding interface{}
	AuthorizationUrl interface{}
	ClientId interface{}
	ClientSecret interface{}
	DeprovisionedAction interface{}
	GroupsAction interface{}
	GroupsAssignments interface{}
	GroupsAttribute interface{}
	GroupsFilters interface{}
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
	IssuerMode interface{}
	IssuerUrl interface{}
	JwksBinding interface{}
	JwksUrl interface{}
	MaxClockSkew interface{}
	// name of idp
	Name interface{}
	ProfileMaster interface{}
	ProtocolType interface{}
	ProvisioningAction interface{}
	// algorithm to use to sign requests
	RequestSignatureAlgorithm interface{}
	// algorithm to use to sign response
	RequestSignatureScope interface{}
	// algorithm to use to sign requests
	ResponseSignatureAlgorithm interface{}
	// algorithm to use to sign response
	ResponseSignatureScope interface{}
	Scopes interface{}
	Status interface{}
	SubjectMatchAttribute interface{}
	SubjectMatchType interface{}
	SuspendedAction interface{}
	TokenBinding interface{}
	TokenUrl interface{}
	UserInfoBinding interface{}
	UserInfoUrl interface{}
	UsernameTemplate interface{}
}
