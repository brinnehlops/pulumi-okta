// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package group

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve the Everyone group from Okta. The same can be achieved with the `group.Group` data source with `name = "Everyone"`. This is simply a shortcut.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/everyone_group.html.markdown.
func GetEveryoneGroup(ctx *pulumi.Context, args *GetEveryoneGroupArgs, opts ...pulumi.InvokeOption) (*GetEveryoneGroupResult, error) {
	var rv GetEveryoneGroupResult
	err := ctx.Invoke("okta:group/getEveryoneGroup:getEveryoneGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEveryoneGroup.
type GetEveryoneGroupArgs struct {
	IncludeUsers *bool `pulumi:"includeUsers"`
}

// A collection of values returned by getEveryoneGroup.
type GetEveryoneGroupResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	IncludeUsers *bool  `pulumi:"includeUsers"`
}
