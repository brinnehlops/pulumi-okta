// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package idp

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve a SAML IdP from Okta.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/d/idp_saml.html.markdown.
func LookupSaml(ctx *pulumi.Context, args *GetSamlArgs) (*GetSamlResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["id"] = args.Id
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("okta:idp/getSaml:getSaml", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSamlResult{
		AcsBinding: outputs["acsBinding"],
		AcsType: outputs["acsType"],
		Audience: outputs["audience"],
		Id: outputs["id"],
		Issuer: outputs["issuer"],
		IssuerMode: outputs["issuerMode"],
		Kid: outputs["kid"],
		Name: outputs["name"],
		SsoBinding: outputs["ssoBinding"],
		SsoDestination: outputs["ssoDestination"],
		SsoUrl: outputs["ssoUrl"],
		SubjectFilter: outputs["subjectFilter"],
		SubjectFormats: outputs["subjectFormats"],
		Type: outputs["type"],
	}, nil
}

// A collection of arguments for invoking getSaml.
type GetSamlArgs struct {
	// The id of the idp to retrieve, conflicts with `name`.
	Id interface{}
	// The name of the idp to retrieve, conflicts with `id`.
	Name interface{}
}

// A collection of values returned by getSaml.
type GetSamlResult struct {
	// HTTP binding used to receive a SAMLResponse message from the IdP.
	AcsBinding interface{}
	// Determines whether to publish an instance-specific (trust) or organization (shared) ACS endpoint in the SAML metadata.
	AcsType interface{}
	// URI that identifies the target Okta IdP instance (SP)
	Audience interface{}
	// id of idp.
	Id interface{}
	// URI that identifies the issuer (IdP).
	Issuer interface{}
	// indicates whether Okta uses the original Okta org domain URL, or a custom domain URL in the request to the IdP.
	IssuerMode interface{}
	// Key ID reference to the IdP's X.509 signature certificate.
	Kid interface{}
	// name of the idp.
	Name interface{}
	// single sign on binding.
	SsoBinding interface{}
	// SSO request binding, HTTP-POST or HTTP-REDIRECT.
	SsoDestination interface{}
	// single sign on url.
	SsoUrl interface{}
	// regular expression pattern used to filter untrusted IdP usernames.
	SubjectFilter interface{}
	// Expression to generate or transform a unique username for the IdP user.
	SubjectFormats interface{}
	// type of idp.
	Type interface{}
}
