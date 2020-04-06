// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package app

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve the collaborators for a given repository.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/app_metadata_saml.html.markdown.
func GetMetadataSaml(ctx *pulumi.Context, args *GetMetadataSamlArgs, opts ...pulumi.InvokeOption) (*GetMetadataSamlResult, error) {
	var rv GetMetadataSamlResult
	err := ctx.Invoke("okta:app/getMetadataSaml:getMetadataSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMetadataSaml.
type GetMetadataSamlArgs struct {
	// The application ID.
	AppId string `pulumi:"appId"`
	// Certificate Key ID.
	KeyId string `pulumi:"keyId"`
}

// A collection of values returned by getMetadataSaml.
type GetMetadataSamlResult struct {
	AppId string `pulumi:"appId"`
	// public certificate from application metadata.
	Certificate string `pulumi:"certificate"`
	// Entity URL for instance `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf`.
	EntityId string `pulumi:"entityId"`
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.
	HttpPostBinding string `pulumi:"httpPostBinding"`
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.
	HttpRedirectBinding string `pulumi:"httpRedirectBinding"`
	// id is the provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	KeyId string `pulumi:"keyId"`
	// raw metadata of application.
	Metadata string `pulumi:"metadata"`
	// Whether authn requests are signed.
	WantAuthnRequestsSigned bool `pulumi:"wantAuthnRequestsSigned"`
}
