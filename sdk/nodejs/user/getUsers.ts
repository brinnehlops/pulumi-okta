// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve a list of users from Okta.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = pulumi.output(okta.user.getUsers({
 *     searches: [{
 *         comparison: "sw",
 *         name: "profile.company",
 *         value: "Articulate",
 *     }],
 * }, { async: true }));
 * ```
 */
export function getUsers(args: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("okta:user/getUsers:getUsers", {
        "searches": args.searches,
        "users": args.users,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * Map of search criteria to use to find users. It supports the following properties.
     */
    readonly searches: inputs.user.GetUsersSearch[];
    /**
     * collection of users retrieved from Okta with the following properties.
     */
    readonly users?: inputs.user.GetUsersUser[];
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly searches: outputs.user.GetUsersSearch[];
    /**
     * collection of users retrieved from Okta with the following properties.
     */
    readonly users?: outputs.user.GetUsersUser[];
}
