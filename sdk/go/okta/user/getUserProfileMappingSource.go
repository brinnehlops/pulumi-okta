// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package user

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve the base user Profile Mapping source or target from Okta.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/user_profile_mapping_source.html.markdown.
func GetUserProfileMappingSource(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetUserProfileMappingSourceResult, error) {
	var rv GetUserProfileMappingSourceResult
	err := ctx.Invoke("okta:user/getUserProfileMappingSource:getUserProfileMappingSource", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getUserProfileMappingSource.
type GetUserProfileMappingSourceResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// name of source.
	Name string `pulumi:"name"`
	// type of source.
	Type string `pulumi:"type"`
}
