// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package group

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve a group from Okta.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/d/group.html.markdown.
func LookupGroup(ctx *pulumi.Context, args *GetGroupArgs) (*GetGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["includeUsers"] = args.IncludeUsers
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("okta:group/getGroup:getGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetGroupResult{
		Description: outputs["description"],
		IncludeUsers: outputs["includeUsers"],
		Name: outputs["name"],
		Users: outputs["users"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getGroup.
type GetGroupArgs struct {
	// whether or not to retrieve all member ids.
	IncludeUsers interface{}
	// name of group to retrieve.
	Name interface{}
}

// A collection of values returned by getGroup.
type GetGroupResult struct {
	// description of group.
	Description interface{}
	IncludeUsers interface{}
	// name of group.
	Name interface{}
	// user ids that are members of this group, only included if `includeUsers` is set to `true`.
	Users interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
