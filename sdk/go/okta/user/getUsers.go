// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package user

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve a list of users from Okta.
// 
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/users.html.markdown.
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("okta:user/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// Map of search criteria to use to find users. It supports the following properties.
	Searches []GetUsersSearch `pulumi:"searches"`
	Users []GetUsersUser `pulumi:"users"`
}


// A collection of values returned by getUsers.
type GetUsersResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Searches []GetUsersSearch `pulumi:"searches"`
	// collection of users retrieved from Okta with the following properties.
	Users []GetUsersUser `pulumi:"users"`
}

