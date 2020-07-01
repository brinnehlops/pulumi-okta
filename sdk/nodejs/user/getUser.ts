// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve a users from Okta.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = pulumi.output(okta.user.getUser({
 *     searches: [
 *         {
 *             name: "profile.firstName",
 *             value: "John",
 *         },
 *         {
 *             name: "profile.lastName",
 *             value: "Doe",
 *         },
 *     ],
 * }, { async: true }));
 * ```
 */
export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("okta:user/getUser:getUser", {
        "searches": args.searches,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * Map of search criteria. It supports the following properties.
     */
    readonly searches: inputs.user.GetUserSearch[];
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * Administrator roles assigned to user.
     */
    readonly adminRoles: string[];
    /**
     * user profile property.
     */
    readonly city: string;
    /**
     * user profile property.
     */
    readonly costCenter: string;
    /**
     * user profile property.
     */
    readonly countryCode: string;
    /**
     * raw JSON containing all custom profile attributes.
     */
    readonly customProfileAttributes: string;
    /**
     * user profile property.
     */
    readonly department: string;
    /**
     * user profile property.
     */
    readonly displayName: string;
    /**
     * user profile property.
     */
    readonly division: string;
    /**
     * user profile property.
     */
    readonly email: string;
    /**
     * user profile property.
     */
    readonly employeeNumber: string;
    /**
     * user profile property.
     */
    readonly firstName: string;
    /**
     * user profile property.
     */
    readonly groupMemberships: string[];
    /**
     * user profile property.
     */
    readonly honorificPrefix: string;
    /**
     * user profile property.
     */
    readonly honorificSuffix: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * user profile property.
     */
    readonly lastName: string;
    /**
     * user profile property.
     */
    readonly locale: string;
    /**
     * user profile property.
     */
    readonly login: string;
    /**
     * user profile property.
     */
    readonly manager: string;
    /**
     * user profile property.
     */
    readonly managerId: string;
    /**
     * user profile property.
     */
    readonly middleName: string;
    /**
     * user profile property.
     */
    readonly mobilePhone: string;
    /**
     * user profile property.
     */
    readonly nickName: string;
    /**
     * user profile property.
     */
    readonly organization: string;
    /**
     * user profile property.
     */
    readonly postalAddress: string;
    /**
     * user profile property.
     */
    readonly preferredLanguage: string;
    /**
     * user profile property.
     */
    readonly primaryPhone: string;
    /**
     * user profile property.
     */
    readonly profileUrl: string;
    readonly searches: outputs.user.GetUserSearch[];
    /**
     * user profile property.
     */
    readonly secondEmail: string;
    /**
     * user profile property.
     */
    readonly state: string;
    /**
     * user profile property.
     */
    readonly status: string;
    /**
     * user profile property.
     */
    readonly streetAddress: string;
    /**
     * user profile property.
     */
    readonly timezone: string;
    /**
     * user profile property.
     */
    readonly title: string;
    /**
     * user profile property.
     */
    readonly userType: string;
    /**
     * user profile property.
     */
    readonly zipCode: string;
}
