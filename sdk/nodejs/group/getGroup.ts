// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve a group from Okta.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = okta.group.getGroup({
 *     label: "Example App",
 * });
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/group.html.markdown.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> & GetGroupResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetGroupResult> = pulumi.runtime.invoke("okta:group/getGroup:getGroup", {
        "includeUsers": args.includeUsers,
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * whether or not to retrieve all member ids.
     */
    readonly includeUsers?: boolean;
    /**
     * name of group to retrieve.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * description of group.
     */
    readonly description: string;
    readonly includeUsers?: boolean;
    /**
     * name of group.
     */
    readonly name: string;
    /**
     * user ids that are members of this group, only included if `includeUsers` is set to `true`.
     */
    readonly users: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
