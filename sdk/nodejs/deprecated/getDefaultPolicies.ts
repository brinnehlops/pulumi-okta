// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDefaultPolicies(args: GetDefaultPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultPoliciesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("okta:deprecated/getDefaultPolicies:getDefaultPolicies", {
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultPolicies.
 */
export interface GetDefaultPoliciesArgs {
    readonly type: string;
}

/**
 * A collection of values returned by getDefaultPolicies.
 */
export interface GetDefaultPoliciesResult {
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
