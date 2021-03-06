// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Assigns a group to an application.
//
// This resource allows you to create an App Group assignment.
//
// __When using this resource, make sure to add the following `lifefycle` argument to the application resource you are assigning to:__
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-okta/sdk/v2/go/okta/app"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := app.NewGroupAssignment(ctx, "example", &app.GroupAssignmentArgs{
// 			AppId:   pulumi.String("<app id>"),
// 			GroupId: pulumi.String("<group id>"),
// 			Profile: pulumi.String(fmt.Sprintf("%v%v%v%v", "{\n", "  \"<app_profile_field>\": \"<value>\"\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GroupAssignment struct {
	pulumi.CustomResourceState

	// The ID of the application to assign a group to.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The ID of the group to assign the app to.
	GroupId  pulumi.StringOutput `pulumi:"groupId"`
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
}

// NewGroupAssignment registers a new resource with the given unique name, arguments, and options.
func NewGroupAssignment(ctx *pulumi.Context,
	name string, args *GroupAssignmentArgs, opts ...pulumi.ResourceOption) (*GroupAssignment, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil {
		args = &GroupAssignmentArgs{}
	}
	var resource GroupAssignment
	err := ctx.RegisterResource("okta:app/groupAssignment:GroupAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupAssignment gets an existing GroupAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupAssignmentState, opts ...pulumi.ResourceOption) (*GroupAssignment, error) {
	var resource GroupAssignment
	err := ctx.ReadResource("okta:app/groupAssignment:GroupAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupAssignment resources.
type groupAssignmentState struct {
	// The ID of the application to assign a group to.
	AppId *string `pulumi:"appId"`
	// The ID of the group to assign the app to.
	GroupId  *string `pulumi:"groupId"`
	Priority *int    `pulumi:"priority"`
	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	Profile *string `pulumi:"profile"`
}

type GroupAssignmentState struct {
	// The ID of the application to assign a group to.
	AppId pulumi.StringPtrInput
	// The ID of the group to assign the app to.
	GroupId  pulumi.StringPtrInput
	Priority pulumi.IntPtrInput
	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	Profile pulumi.StringPtrInput
}

func (GroupAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupAssignmentState)(nil)).Elem()
}

type groupAssignmentArgs struct {
	// The ID of the application to assign a group to.
	AppId string `pulumi:"appId"`
	// The ID of the group to assign the app to.
	GroupId  string `pulumi:"groupId"`
	Priority *int   `pulumi:"priority"`
	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	Profile *string `pulumi:"profile"`
}

// The set of arguments for constructing a GroupAssignment resource.
type GroupAssignmentArgs struct {
	// The ID of the application to assign a group to.
	AppId pulumi.StringInput
	// The ID of the group to assign the app to.
	GroupId  pulumi.StringInput
	Priority pulumi.IntPtrInput
	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	Profile pulumi.StringPtrInput
}

func (GroupAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupAssignmentArgs)(nil)).Elem()
}
