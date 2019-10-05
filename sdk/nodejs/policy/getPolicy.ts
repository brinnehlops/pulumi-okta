// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve a policy from Okta.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = okta.policy.getPolicy({
 *     name: "Password Policy Example",
 *     type: "PASSWORD",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/d/policy.html.markdown.
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> & GetPolicyResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetPolicyResult> = pulumi.runtime.invoke("okta:policy/getPolicy:getPolicy", {
        "name": args.name,
        "type": args.type,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * name of policy to retrieve.
     */
    readonly name: string;
    /**
     * type of policy to retrieve.
     */
    readonly type: string;
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    /**
     * name of policy.
     */
    readonly name: string;
    /**
     * type of policy.
     */
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
