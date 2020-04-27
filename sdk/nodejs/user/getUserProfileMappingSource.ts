// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve the base user Profile Mapping source or target from Okta.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = pulumi.output(okta.user.getUserProfileMappingSource({ async: true }));
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/user_profile_mapping_source.html.markdown.
 */
export function getUserProfileMappingSource(opts?: pulumi.InvokeOptions): Promise<GetUserProfileMappingSourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("okta:user/getUserProfileMappingSource:getUserProfileMappingSource", {
    }, opts);
}

/**
 * A collection of values returned by getUserProfileMappingSource.
 */
export interface GetUserProfileMappingSourceResult {
    /**
     * name of source.
     */
    readonly name: string;
    /**
     * type of source.
     */
    readonly type: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
