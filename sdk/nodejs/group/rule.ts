// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Okta Group Rule.
 *
 * This resource allows you to create and configure an Okta Group Rule.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.group.Rule("example", {
 *     expressionType: "urn:okta:expression:1.0",
 *     expressionValue: "String.startsWith(user.firstName,\"andy\")",
 *     groupAssignments: ["<group id>"],
 *     status: "ACTIVE",
 * });
 * ```
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:group/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
     */
    public readonly expressionType!: pulumi.Output<string | undefined>;
    /**
     * The expression value.
     */
    public readonly expressionValue!: pulumi.Output<string>;
    /**
     * The list of group ids to assign the users to.
     */
    public readonly groupAssignments!: pulumi.Output<string[]>;
    /**
     * The name of the Group Rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The status of the group rule.
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RuleState | undefined;
            inputs["expressionType"] = state ? state.expressionType : undefined;
            inputs["expressionValue"] = state ? state.expressionValue : undefined;
            inputs["groupAssignments"] = state ? state.groupAssignments : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if (!args || args.expressionValue === undefined) {
                throw new Error("Missing required property 'expressionValue'");
            }
            if (!args || args.groupAssignments === undefined) {
                throw new Error("Missing required property 'groupAssignments'");
            }
            inputs["expressionType"] = args ? args.expressionType : undefined;
            inputs["expressionValue"] = args ? args.expressionValue : undefined;
            inputs["groupAssignments"] = args ? args.groupAssignments : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["status"] = args ? args.status : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
     */
    readonly expressionType?: pulumi.Input<string>;
    /**
     * The expression value.
     */
    readonly expressionValue?: pulumi.Input<string>;
    /**
     * The list of group ids to assign the users to.
     */
    readonly groupAssignments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Group Rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The status of the group rule.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
     */
    readonly expressionType?: pulumi.Input<string>;
    /**
     * The expression value.
     */
    readonly expressionValue: pulumi.Input<string>;
    /**
     * The list of group ids to assign the users to.
     */
    readonly groupAssignments: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Group Rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The status of the group rule.
     */
    readonly status?: pulumi.Input<string>;
}
