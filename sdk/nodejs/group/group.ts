// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Okta Group.
 * 
 * This resource allows you to create and configure an Okta Group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = new okta.group.Group("example", {
 *     description: "My Example Group",
 * });
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/group.html.markdown.
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:group/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * The description of the Okta Group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the Okta Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Users associated with the group. This can also be done per user.
     */
    public readonly users!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["users"] = args ? args.users : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * The description of the Okta Group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the Okta Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Users associated with the group. This can also be done per user.
     */
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * The description of the Okta Group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the Okta Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Users associated with the group. This can also be done per user.
     */
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}
