// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve the Everyone group from Okta. The same can be achieved with the `okta.group.Group` data source with `name = "Everyone"`. This is simply a shortcut.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = pulumi.output(okta.group.getEveryoneGroup({ async: true }));
 * ```
 */
export function getEveryoneGroup(args?: GetEveryoneGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetEveryoneGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("okta:group/getEveryoneGroup:getEveryoneGroup", {
        "includeUsers": args.includeUsers,
    }, opts);
}

/**
 * A collection of arguments for invoking getEveryoneGroup.
 */
export interface GetEveryoneGroupArgs {
    readonly includeUsers?: boolean;
}

/**
 * A collection of values returned by getEveryoneGroup.
 */
export interface GetEveryoneGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeUsers?: boolean;
}
