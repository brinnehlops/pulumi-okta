// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package auth

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve an auth server from Okta.
//
//
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/auth_server.html.markdown.
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("okta:auth/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type LookupServerArgs struct {
	// The name of the auth server to retrieve.
	Name string `pulumi:"name"`
}

// A collection of values returned by getServer.
type LookupServerResult struct {
	// array of audiences,
	Audiences []string `pulumi:"audiences"`
	// last time credentials were rotated.
	CredentialsLastRotated string `pulumi:"credentialsLastRotated"`
	// next time credentials will be rotated
	CredentialsNextRotation string `pulumi:"credentialsNextRotation"`
	// mode of credential rotation, auto or manual.
	CredentialsRotationMode string `pulumi:"credentialsRotationMode"`
	// description of Authorization server.
	Description string `pulumi:"description"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// auth server key id.
	Kid string `pulumi:"kid"`
	// The name of the auth server.
	Name string `pulumi:"name"`
	// the activation status of the authorization server.
	Status string `pulumi:"status"`
}
