// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package group

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates Group level Admin Role Assignments.
// 
// This resource allows you to create and configure Group level Admin Role Assignments.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/group_roles.html.markdown.
type Roles struct {
	s *pulumi.ResourceState
}

// NewRoles registers a new resource with the given unique name, arguments, and options.
func NewRoles(ctx *pulumi.Context,
	name string, args *RolesArgs, opts ...pulumi.ResourceOpt) (*Roles, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminRoles"] = nil
		inputs["groupId"] = nil
	} else {
		inputs["adminRoles"] = args.AdminRoles
		inputs["groupId"] = args.GroupId
	}
	s, err := ctx.RegisterResource("okta:group/roles:Roles", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Roles{s: s}, nil
}

// GetRoles gets an existing Roles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoles(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RolesState, opts ...pulumi.ResourceOpt) (*Roles, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminRoles"] = state.AdminRoles
		inputs["groupId"] = state.GroupId
	}
	s, err := ctx.ReadResource("okta:group/roles:Roles", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Roles{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Roles) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Roles) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
func (r *Roles) AdminRoles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["adminRoles"])
}

// The ID of group to attach admin roles to.
func (r *Roles) GroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupId"])
}

// Input properties used for looking up and filtering Roles resources.
type RolesState struct {
	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles interface{}
	// The ID of group to attach admin roles to.
	GroupId interface{}
}

// The set of arguments for constructing a Roles resource.
type RolesArgs struct {
	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles interface{}
	// The ID of group to attach admin roles to.
	GroupId interface{}
}
