// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package idp

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve SAML IdP metadata from Okta.
//
//
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/idp_metadata_saml.html.markdown.
func GetMetadataSaml(ctx *pulumi.Context, args *GetMetadataSamlArgs, opts ...pulumi.InvokeOption) (*GetMetadataSamlResult, error) {
	var rv GetMetadataSamlResult
	err := ctx.Invoke("okta:idp/getMetadataSaml:getMetadataSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMetadataSaml.
type GetMetadataSamlArgs struct {
	// The id of the IdP to retrieve metadata for.
	IdpId *string `pulumi:"idpId"`
}

// A collection of values returned by getMetadataSaml.
type GetMetadataSamlResult struct {
	// whether assertions are signed.
	AssertionsSigned bool `pulumi:"assertionsSigned"`
	// whether authn requests are signed.
	AuthnRequestSigned bool `pulumi:"authnRequestSigned"`
	// SAML request encryption certificate.
	EncryptionCertificate string `pulumi:"encryptionCertificate"`
	// Entity URL for instance `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf`.
	EntityId string `pulumi:"entityId"`
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.
	HttpPostBinding string `pulumi:"httpPostBinding"`
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.
	HttpRedirectBinding string `pulumi:"httpRedirectBinding"`
	// id is the provider-assigned unique ID for this managed resource.
	Id    string  `pulumi:"id"`
	IdpId *string `pulumi:"idpId"`
	// raw IdP metadata.
	Metadata string `pulumi:"metadata"`
	// SAML request signing certificate.
	SigningCertificate string `pulumi:"signingCertificate"`
}
