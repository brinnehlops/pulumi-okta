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
     * Allow or deny access based on the rule conditions: ALLOW or DENY.
     */
    public readonly access!: pulumi.Output<string | undefined>;
    /**
     * Authentication entrypoint: ANY or RADIUS.
     */
    public readonly authtype!: pulumi.Output<string | undefined>;
    /**
     * When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
     */
    public readonly enroll!: pulumi.Output<string | undefined>;
    /**
     * Elapsed time before the next MFA challenge
     */
    public readonly mfaLifetime!: pulumi.Output<number | undefined>;
    /**
     * Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: DEVICE, SESSION or
     * ALWAYS
     */
    public readonly mfaPrompt!: pulumi.Output<string | undefined>;
    /**
     * Remember MFA device.
     */
    public readonly mfaRememberDevice!: pulumi.Output<boolean | undefined>;
    /**
     * Require MFA.
     */
    public readonly mfaRequired!: pulumi.Output<boolean | undefined>;
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
     * Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW
     */
    public readonly passwordChange!: pulumi.Output<string | undefined>;
    /**
     * Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW
     */
    public readonly passwordReset!: pulumi.Output<string | undefined>;
    /**
     * Allow or deny a user to unlock. Default = DENY
     */
    public readonly passwordUnlock!: pulumi.Output<string | undefined>;
    /**
     * Policy ID.
     */
    public readonly policyid!: pulumi.Output<string>;
    /**
     * Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Max minutes a session can be idle.
     */
    public readonly sessionIdle!: pulumi.Output<number | undefined>;
    /**
     * Max minutes a session is active: Disable = 0.
     */
    public readonly sessionLifetime!: pulumi.Output<number | undefined>;
    /**
     * Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session
     * cookies.
     */
    public readonly sessionPersistent!: pulumi.Output<boolean | undefined>;
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
            inputs["access"] = state ? state.access : undefined;
            inputs["authtype"] = state ? state.authtype : undefined;
            inputs["enroll"] = state ? state.enroll : undefined;
            inputs["mfaLifetime"] = state ? state.mfaLifetime : undefined;
            inputs["mfaPrompt"] = state ? state.mfaPrompt : undefined;
            inputs["mfaRememberDevice"] = state ? state.mfaRememberDevice : undefined;
            inputs["mfaRequired"] = state ? state.mfaRequired : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkConnection"] = state ? state.networkConnection : undefined;
            inputs["networkExcludes"] = state ? state.networkExcludes : undefined;
            inputs["networkIncludes"] = state ? state.networkIncludes : undefined;
            inputs["passwordChange"] = state ? state.passwordChange : undefined;
            inputs["passwordReset"] = state ? state.passwordReset : undefined;
            inputs["passwordUnlock"] = state ? state.passwordUnlock : undefined;
            inputs["policyid"] = state ? state.policyid : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["sessionIdle"] = state ? state.sessionIdle : undefined;
            inputs["sessionLifetime"] = state ? state.sessionLifetime : undefined;
            inputs["sessionPersistent"] = state ? state.sessionPersistent : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["usersExcludeds"] = state ? state.usersExcludeds : undefined;
        } else {
            const args = argsOrState as RuleMfaArgs | undefined;
            if (!args || args.policyid === undefined) {
                throw new Error("Missing required property 'policyid'");
            }
            inputs["access"] = args ? args.access : undefined;
            inputs["authtype"] = args ? args.authtype : undefined;
            inputs["enroll"] = args ? args.enroll : undefined;
            inputs["mfaLifetime"] = args ? args.mfaLifetime : undefined;
            inputs["mfaPrompt"] = args ? args.mfaPrompt : undefined;
            inputs["mfaRememberDevice"] = args ? args.mfaRememberDevice : undefined;
            inputs["mfaRequired"] = args ? args.mfaRequired : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkConnection"] = args ? args.networkConnection : undefined;
            inputs["networkExcludes"] = args ? args.networkExcludes : undefined;
            inputs["networkIncludes"] = args ? args.networkIncludes : undefined;
            inputs["passwordChange"] = args ? args.passwordChange : undefined;
            inputs["passwordReset"] = args ? args.passwordReset : undefined;
            inputs["passwordUnlock"] = args ? args.passwordUnlock : undefined;
            inputs["policyid"] = args ? args.policyid : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["sessionIdle"] = args ? args.sessionIdle : undefined;
            inputs["sessionLifetime"] = args ? args.sessionLifetime : undefined;
            inputs["sessionPersistent"] = args ? args.sessionPersistent : undefined;
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
     * Allow or deny access based on the rule conditions: ALLOW or DENY.
     */
    readonly access?: pulumi.Input<string>;
    /**
     * Authentication entrypoint: ANY or RADIUS.
     */
    readonly authtype?: pulumi.Input<string>;
    /**
     * When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
     */
    readonly enroll?: pulumi.Input<string>;
    /**
     * Elapsed time before the next MFA challenge
     */
    readonly mfaLifetime?: pulumi.Input<number>;
    /**
     * Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: DEVICE, SESSION or
     * ALWAYS
     */
    readonly mfaPrompt?: pulumi.Input<string>;
    /**
     * Remember MFA device.
     */
    readonly mfaRememberDevice?: pulumi.Input<boolean>;
    /**
     * Require MFA.
     */
    readonly mfaRequired?: pulumi.Input<boolean>;
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
     * Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW
     */
    readonly passwordChange?: pulumi.Input<string>;
    /**
     * Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW
     */
    readonly passwordReset?: pulumi.Input<string>;
    /**
     * Allow or deny a user to unlock. Default = DENY
     */
    readonly passwordUnlock?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    readonly policyid?: pulumi.Input<string>;
    /**
     * Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Max minutes a session can be idle.
     */
    readonly sessionIdle?: pulumi.Input<number>;
    /**
     * Max minutes a session is active: Disable = 0.
     */
    readonly sessionLifetime?: pulumi.Input<number>;
    /**
     * Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session
     * cookies.
     */
    readonly sessionPersistent?: pulumi.Input<boolean>;
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
     * Allow or deny access based on the rule conditions: ALLOW or DENY.
     */
    readonly access?: pulumi.Input<string>;
    /**
     * Authentication entrypoint: ANY or RADIUS.
     */
    readonly authtype?: pulumi.Input<string>;
    /**
     * When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
     */
    readonly enroll?: pulumi.Input<string>;
    /**
     * Elapsed time before the next MFA challenge
     */
    readonly mfaLifetime?: pulumi.Input<number>;
    /**
     * Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: DEVICE, SESSION or
     * ALWAYS
     */
    readonly mfaPrompt?: pulumi.Input<string>;
    /**
     * Remember MFA device.
     */
    readonly mfaRememberDevice?: pulumi.Input<boolean>;
    /**
     * Require MFA.
     */
    readonly mfaRequired?: pulumi.Input<boolean>;
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
     * Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW
     */
    readonly passwordChange?: pulumi.Input<string>;
    /**
     * Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW
     */
    readonly passwordReset?: pulumi.Input<string>;
    /**
     * Allow or deny a user to unlock. Default = DENY
     */
    readonly passwordUnlock?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    readonly policyid: pulumi.Input<string>;
    /**
     * Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Max minutes a session can be idle.
     */
    readonly sessionIdle?: pulumi.Input<number>;
    /**
     * Max minutes a session is active: Disable = 0.
     */
    readonly sessionLifetime?: pulumi.Input<number>;
    /**
     * Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session
     * cookies.
     */
    readonly sessionPersistent?: pulumi.Input<boolean>;
    /**
     * Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Set of User IDs to Exclude
     */
    readonly usersExcludeds?: pulumi.Input<pulumi.Input<string>[]>;
}
