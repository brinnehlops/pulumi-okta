// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Assigns a group to an application.
 *
 * This resource allows you to create an App Group assignment.
 *
 * __When using this resource, make sure to add the following `lifefycle` argument to the application resource you are assigning to:__
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.app.GroupAssignment("example", {
 *     appId: "<app id>",
 *     groupId: "<group id>",
 *     profile: `{
 *   "<app_profile_field>": "<value>"
 * }
 * `,
 * });
 * ```
 */
export class GroupAssignment extends pulumi.CustomResource {
    /**
     * Get an existing GroupAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupAssignmentState, opts?: pulumi.CustomResourceOptions): GroupAssignment {
        return new GroupAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/groupAssignment:GroupAssignment';

    /**
     * Returns true if the given object is an instance of GroupAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupAssignment.__pulumiType;
    }

    /**
     * The ID of the application to assign a group to.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The ID of the group to assign the app to.
     */
    public readonly groupId!: pulumi.Output<string>;
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
     */
    public readonly profile!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupAssignmentArgs | GroupAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupAssignmentState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["profile"] = state ? state.profile : undefined;
        } else {
            const args = argsOrState as GroupAssignmentArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.groupId === undefined) {
                throw new Error("Missing required property 'groupId'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["profile"] = args ? args.profile : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupAssignment resources.
 */
export interface GroupAssignmentState {
    /**
     * The ID of the application to assign a group to.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The ID of the group to assign the app to.
     */
    readonly groupId?: pulumi.Input<string>;
    readonly priority?: pulumi.Input<number>;
    /**
     * JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
     */
    readonly profile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupAssignment resource.
 */
export interface GroupAssignmentArgs {
    /**
     * The ID of the application to assign a group to.
     */
    readonly appId: pulumi.Input<string>;
    /**
     * The ID of the group to assign the app to.
     */
    readonly groupId: pulumi.Input<string>;
    readonly priority?: pulumi.Input<number>;
    /**
     * JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
     */
    readonly profile?: pulumi.Input<string>;
}
