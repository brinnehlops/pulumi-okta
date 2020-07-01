// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package group

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve a group from Okta.
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
// 		_, err := group.LookupGroup(ctx, &group.LookupGroupArgs{
// 			Name: "Example App",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("okta:group/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// whether or not to retrieve all member ids.
	IncludeUsers *bool `pulumi:"includeUsers"`
	// name of group to retrieve.
	Name string `pulumi:"name"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// description of group.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	IncludeUsers *bool  `pulumi:"includeUsers"`
	// name of group.
	Name string `pulumi:"name"`
	// user ids that are members of this group, only included if `includeUsers` is set to `true`.
	Users []string `pulumi:"users"`
}
