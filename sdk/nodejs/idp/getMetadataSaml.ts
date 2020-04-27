// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve SAML IdP metadata from Okta.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = pulumi.output(okta.idp.getMetadataSaml({
 *     id: "<idp id>",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/idp_metadata_saml.html.markdown.
 */
export function getMetadataSaml(args?: GetMetadataSamlArgs, opts?: pulumi.InvokeOptions): Promise<GetMetadataSamlResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("okta:idp/getMetadataSaml:getMetadataSaml", {
        "idpId": args.idpId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMetadataSaml.
 */
export interface GetMetadataSamlArgs {
    /**
     * The id of the IdP to retrieve metadata for.
     */
    readonly idpId?: string;
}

/**
 * A collection of values returned by getMetadataSaml.
 */
export interface GetMetadataSamlResult {
    /**
     * whether assertions are signed.
     */
    readonly assertionsSigned: boolean;
    /**
     * whether authn requests are signed.
     */
    readonly authnRequestSigned: boolean;
    /**
     * SAML request encryption certificate.
     */
    readonly encryptionCertificate: string;
    /**
     * Entity URL for instance `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf`.
     */
    readonly entityId: string;
    /**
     * urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.
     */
    readonly httpPostBinding: string;
    /**
     * urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.
     */
    readonly httpRedirectBinding: string;
    readonly idpId?: string;
    /**
     * raw IdP metadata.
     */
    readonly metadata: string;
    /**
     * SAML request signing certificate.
     */
    readonly signingCertificate: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
