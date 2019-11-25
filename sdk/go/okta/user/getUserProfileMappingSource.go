// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve the base user Profile Mapping source or target from Okta.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/user_profile_mapping_source.html.markdown.
func LookupUserProfileMappingSource(ctx *pulumi.Context) (*GetUserProfileMappingSourceResult, error) {
	outputs, err := ctx.Invoke("okta:user/getUserProfileMappingSource:getUserProfileMappingSource", nil)
	if err != nil {
		return nil, err
	}
	return &GetUserProfileMappingSourceResult{
		Name: outputs["name"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of values returned by getUserProfileMappingSource.
type GetUserProfileMappingSourceResult struct {
	// name of source.
	Name interface{}
	// type of source.
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
