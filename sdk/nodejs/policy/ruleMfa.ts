// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an MFA Policy Rule.
 * 
 * This resource allows you to create and configure an MFA Policy Rule.
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_mfa.html.markdown.
 */
export class RuleMfa extends pulumi.CustomResource {
    /**
     * Get an existing RuleMfa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleMfaState, opts?: pulumi.CustomResourceOptions): RuleMfa {
        return new RuleMfa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:policy/ruleMfa:RuleMfa';

    /**
     * Returns true if the given object is an instance of RuleMfa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleMfa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleMfa.__pulumiType;
    }

    /**
     * When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
     */
    public readonly enroll!: pulumi.Output<string | undefined>;
    /**
     * Policy Rule Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
     */
    public readonly networkConnection!: pulumi.Output<string | undefined>;
    /**
     * The network zones to exclude. Conflicts with `networkIncludes`.
     */
    public readonly networkExcludes!: pulumi.Output<string[] | undefined>;
    /**
     * The network zones to include. Conflicts with `networkExcludes`.
     */
    public readonly networkIncludes!: pulumi.Output<string[] | undefined>;
    /**
     * Policy ID.
     */
    public readonly policyid!: pulumi.Output<string>;
    /**
     * Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Set of User IDs to Exclude
     */
    public readonly usersExcludeds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RuleMfa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleMfaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleMfaArgs | RuleMfaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RuleMfaState | undefined;
            inputs["enroll"] = state ? state.enroll : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkConnection"] = state ? state.networkConnection : undefined;
            inputs["networkExcludes"] = state ? state.networkExcludes : undefined;
            inputs["networkIncludes"] = state ? state.networkIncludes : undefined;
            inputs["policyid"] = state ? state.policyid : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["usersExcludeds"] = state ? state.usersExcludeds : undefined;
        } else {
            const args = argsOrState as RuleMfaArgs | undefined;
            if (!args || args.policyid === undefined) {
                throw new Error("Missing required property 'policyid'");
            }
            inputs["enroll"] = args ? args.enroll : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkConnection"] = args ? args.networkConnection : undefined;
            inputs["networkExcludes"] = args ? args.networkExcludes : undefined;
            inputs["networkIncludes"] = args ? args.networkIncludes : undefined;
            inputs["policyid"] = args ? args.policyid : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["usersExcludeds"] = args ? args.usersExcludeds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RuleMfa.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleMfa resources.
 */
export interface RuleMfaState {
    /**
     * When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
     */
    readonly enroll?: pulumi.Input<string>;
    /**
     * Policy Rule Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
     */
    readonly networkConnection?: pulumi.Input<string>;
    /**
     * The network zones to exclude. Conflicts with `networkIncludes`.
     */
    readonly networkExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network zones to include. Conflicts with `networkExcludes`.
     */
    readonly networkIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy ID.
     */
    readonly policyid?: pulumi.Input<string>;
    /**
     * Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Set of User IDs to Exclude
     */
    readonly usersExcludeds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RuleMfa resource.
 */
export interface RuleMfaArgs {
    /**
     * When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
     */
    readonly enroll?: pulumi.Input<string>;
    /**
     * Policy Rule Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
     */
    readonly networkConnection?: pulumi.Input<string>;
    /**
     * The network zones to exclude. Conflicts with `networkIncludes`.
     */
    readonly networkExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network zones to include. Conflicts with `networkExcludes`.
     */
    readonly networkIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy ID.
     */
    readonly policyid: pulumi.Input<string>;
    /**
     * Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Set of User IDs to Exclude
     */
    readonly usersExcludeds?: pulumi.Input<pulumi.Input<string>[]>;
}
