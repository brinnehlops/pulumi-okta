// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package group

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates Group level Admin Role Assignments.
//
// This resource allows you to create and configure Group level Admin Role Assignments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-okta/sdk/v2/go/okta/group"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := group.NewRoles(ctx, "example", &group.RolesArgs{
// 			AdminRoles: pulumi.StringArray{
// 				pulumi.String("SUPER_ADMIN"),
// 			},
// 			GroupId: pulumi.String("<group id>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Roles struct {
	pulumi.CustomResourceState

	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles pulumi.StringArrayOutput `pulumi:"adminRoles"`
	// The ID of group to attach admin roles to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
}

// NewRoles registers a new resource with the given unique name, arguments, and options.
func NewRoles(ctx *pulumi.Context,
	name string, args *RolesArgs, opts ...pulumi.ResourceOption) (*Roles, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil {
		args = &RolesArgs{}
	}
	var resource Roles
	err := ctx.RegisterResource("okta:group/roles:Roles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoles gets an existing Roles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolesState, opts ...pulumi.ResourceOption) (*Roles, error) {
	var resource Roles
	err := ctx.ReadResource("okta:group/roles:Roles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Roles resources.
type rolesState struct {
	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles []string `pulumi:"adminRoles"`
	// The ID of group to attach admin roles to.
	GroupId *string `pulumi:"groupId"`
}

type RolesState struct {
	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles pulumi.StringArrayInput
	// The ID of group to attach admin roles to.
	GroupId pulumi.StringPtrInput
}

func (RolesState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolesState)(nil)).Elem()
}

type rolesArgs struct {
	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles []string `pulumi:"adminRoles"`
	// The ID of group to attach admin roles to.
	GroupId string `pulumi:"groupId"`
}

// The set of arguments for constructing a Roles resource.
type RolesArgs struct {
	// Admin roles associated with the group. It can be any of the following values `"SUPER_ADMIN"`, `"ORG_ADMIN"`, `"APP_ADMIN"`, `"USER_ADMIN"`, `"HELP_DESK_ADMIN"`, `"READ_ONLY_ADMIN"`, `"MOBILE_ADMIN"`, `"API_ACCESS_MANAGEMENT_ADMIN"`, `"REPORT_ADMIN"`.
	AdminRoles pulumi.StringArrayInput
	// The ID of group to attach admin roles to.
	GroupId pulumi.StringInput
}

func (RolesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolesArgs)(nil)).Elem()
}
