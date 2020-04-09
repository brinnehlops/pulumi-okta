// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package deprecated

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetDefaultPolicies(ctx *pulumi.Context, args *GetDefaultPoliciesArgs, opts ...pulumi.InvokeOption) (*GetDefaultPoliciesResult, error) {
	var rv GetDefaultPoliciesResult
	err := ctx.Invoke("okta:deprecated/getDefaultPolicies:getDefaultPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultPolicies.
type GetDefaultPoliciesArgs struct {
	Type string `pulumi:"type"`
}

// A collection of values returned by getDefaultPolicies.
type GetDefaultPoliciesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Type string `pulumi:"type"`
}
