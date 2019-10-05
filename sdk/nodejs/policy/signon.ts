// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Sign On Policy.
 * 
 * This resource allows you to create and configure a Sign On Policy.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = new okta.policy.Signon("example", {
 *     description: "Example",
 *     groupsIncludeds: [okta_group_everyone.id],
 *     status: "ACTIVE",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/policy_signon.html.markdown.
 */
export class Signon extends pulumi.CustomResource {
    /**
     * Get an existing Signon resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SignonState, opts?: pulumi.CustomResourceOptions): Signon {
        return new Signon(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:policy/signon:Signon';

    /**
     * Returns true if the given object is an instance of Signon.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Signon {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Signon.__pulumiType;
    }

    /**
     * Policy Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of Group IDs to Include.
     */
    public readonly groupsIncludeds!: pulumi.Output<string[] | undefined>;
    /**
     * Policy Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Priority of the policy.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Policy Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a Signon resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SignonArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SignonArgs | SignonState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SignonState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["groupsIncludeds"] = state ? state.groupsIncludeds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SignonArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["groupsIncludeds"] = args ? args.groupsIncludeds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["status"] = args ? args.status : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Signon.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Signon resources.
 */
export interface SignonState {
    /**
     * Policy Description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of Group IDs to Include.
     */
    readonly groupsIncludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Priority of the policy.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Policy Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Signon resource.
 */
export interface SignonArgs {
    /**
     * Policy Description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of Group IDs to Include.
     */
    readonly groupsIncludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Priority of the policy.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Policy Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
}
