// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve the collaborators for a given repository.
func GetApp(ctx *pulumi.Context, args *GetAppArgs, opts ...pulumi.InvokeOption) (*GetAppResult, error) {
	var rv GetAppResult
	err := ctx.Invoke("okta:app/getApp:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApp.
type GetAppArgs struct {
	// tells the provider to query for only `ACTIVE` applications.
	ActiveOnly *bool `pulumi:"activeOnly"`
	// `id` of application to retrieve, conflicts with `label` and `labelPrefix`.
	Id *string `pulumi:"id"`
	// The label of the app to retrieve, conflicts with `labelPrefix` and `id`.
	Label *string `pulumi:"label"`
	// Label prefix of the app to retrieve, conflicts with `label` and `id`. This will tell the provider to do a `starts with` query as opposed to an `equals` query.
	LabelPrefix *string `pulumi:"labelPrefix"`
}

// A collection of values returned by getApp.
type GetAppResult struct {
	ActiveOnly *bool `pulumi:"activeOnly"`
	// `description` of application.
	Description string `pulumi:"description"`
	// `id` of application.
	Id *string `pulumi:"id"`
	// `label` of application.
	Label       *string `pulumi:"label"`
	LabelPrefix *string `pulumi:"labelPrefix"`
	// `name` of application.
	Name string `pulumi:"name"`
	// `status` of application.
	Status string `pulumi:"status"`
}
