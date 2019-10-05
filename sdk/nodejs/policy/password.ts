// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Password Policy.
 * 
 * This resource allows you to create and configure a Password Policy.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = new okta.policy.Password("example", {
 *     description: "Example",
 *     groupsIncludeds: [okta_group_everyone.id],
 *     passwordHistoryCount: 4,
 *     status: "ACTIVE",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/policy_password.html.markdown.
 */
export class Password extends pulumi.CustomResource {
    /**
     * Get an existing Password resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PasswordState, opts?: pulumi.CustomResourceOptions): Password {
        return new Password(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:policy/password:Password';

    /**
     * Returns true if the given object is an instance of Password.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Password {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Password.__pulumiType;
    }

    /**
     * Authentication Provider: `"OKTA"` or `"ACTIVE_DIRECTORY"`. Default is `"OKTA"`.
     */
    public readonly authProvider!: pulumi.Output<string | undefined>;
    /**
     * Policy Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Enable or disable email password recovery: ACTIVE or INACTIVE.
     */
    public readonly emailRecovery!: pulumi.Output<string | undefined>;
    /**
     * List of Group IDs to Include.
     */
    public readonly groupsIncludeds!: pulumi.Output<string[] | undefined>;
    /**
     * Policy Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of minutes before a locked account is unlocked: 0 = no limit.
     */
    public readonly passwordAutoUnlockMinutes!: pulumi.Output<number | undefined>;
    /**
     * Check Passwords Against Common Password Dictionary.
     */
    public readonly passwordDictionaryLookup!: pulumi.Output<boolean | undefined>;
    /**
     * User firstName attribute must be excluded from the password.
     */
    public readonly passwordExcludeFirstName!: pulumi.Output<boolean | undefined>;
    /**
     * User lastName attribute must be excluded from the password.
     */
    public readonly passwordExcludeLastName!: pulumi.Output<boolean | undefined>;
    /**
     * If the user name must be excluded from the password.
     */
    public readonly passwordExcludeUsername!: pulumi.Output<boolean | undefined>;
    /**
     * Length in days a user will be warned before password expiry: 0 = no warning.
     */
    public readonly passwordExpireWarnDays!: pulumi.Output<number | undefined>;
    /**
     * Number of distinct passwords that must be created before they can be reused: 0 = none.
     */
    public readonly passwordHistoryCount!: pulumi.Output<number | undefined>;
    /**
     * Length in days a password is valid before expiry: 0 = no limit.",
     */
    public readonly passwordMaxAgeDays!: pulumi.Output<number | undefined>;
    /**
     * Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
     */
    public readonly passwordMaxLockoutAttempts!: pulumi.Output<number | undefined>;
    /**
     * Minimum time interval in minutes between password changes: 0 = no limit.
     */
    public readonly passwordMinAgeMinutes!: pulumi.Output<number | undefined>;
    /**
     * Minimum password length. Default is 8.
     */
    public readonly passwordMinLength!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of lower case characters in password.
     */
    public readonly passwordMinLowercase!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of numbers in password.
     */
    public readonly passwordMinNumber!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of symbols in password.
     */
    public readonly passwordMinSymbol!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of upper case characters in password.
     */
    public readonly passwordMinUppercase!: pulumi.Output<number | undefined>;
    /**
     * If a user should be informed when their account is locked.
     */
    public readonly passwordShowLockoutFailures!: pulumi.Output<boolean | undefined>;
    /**
     * Priority of the policy.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Min length of the password recovery question answer.
     */
    public readonly questionMinLength!: pulumi.Output<number | undefined>;
    /**
     * Enable or disable security question password recovery: ACTIVE or INACTIVE.
     */
    public readonly questionRecovery!: pulumi.Output<string | undefined>;
    /**
     * Lifetime in minutes of the recovery email token.
     */
    public readonly recoveryEmailToken!: pulumi.Output<number | undefined>;
    /**
     * When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
     */
    public readonly skipUnlock!: pulumi.Output<boolean | undefined>;
    /**
     * Enable or disable SMS password recovery: ACTIVE or INACTIVE.
     */
    public readonly smsRecovery!: pulumi.Output<string | undefined>;
    /**
     * Policy Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a Password resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PasswordArgs | PasswordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PasswordState | undefined;
            inputs["authProvider"] = state ? state.authProvider : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["emailRecovery"] = state ? state.emailRecovery : undefined;
            inputs["groupsIncludeds"] = state ? state.groupsIncludeds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["passwordAutoUnlockMinutes"] = state ? state.passwordAutoUnlockMinutes : undefined;
            inputs["passwordDictionaryLookup"] = state ? state.passwordDictionaryLookup : undefined;
            inputs["passwordExcludeFirstName"] = state ? state.passwordExcludeFirstName : undefined;
            inputs["passwordExcludeLastName"] = state ? state.passwordExcludeLastName : undefined;
            inputs["passwordExcludeUsername"] = state ? state.passwordExcludeUsername : undefined;
            inputs["passwordExpireWarnDays"] = state ? state.passwordExpireWarnDays : undefined;
            inputs["passwordHistoryCount"] = state ? state.passwordHistoryCount : undefined;
            inputs["passwordMaxAgeDays"] = state ? state.passwordMaxAgeDays : undefined;
            inputs["passwordMaxLockoutAttempts"] = state ? state.passwordMaxLockoutAttempts : undefined;
            inputs["passwordMinAgeMinutes"] = state ? state.passwordMinAgeMinutes : undefined;
            inputs["passwordMinLength"] = state ? state.passwordMinLength : undefined;
            inputs["passwordMinLowercase"] = state ? state.passwordMinLowercase : undefined;
            inputs["passwordMinNumber"] = state ? state.passwordMinNumber : undefined;
            inputs["passwordMinSymbol"] = state ? state.passwordMinSymbol : undefined;
            inputs["passwordMinUppercase"] = state ? state.passwordMinUppercase : undefined;
            inputs["passwordShowLockoutFailures"] = state ? state.passwordShowLockoutFailures : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["questionMinLength"] = state ? state.questionMinLength : undefined;
            inputs["questionRecovery"] = state ? state.questionRecovery : undefined;
            inputs["recoveryEmailToken"] = state ? state.recoveryEmailToken : undefined;
            inputs["skipUnlock"] = state ? state.skipUnlock : undefined;
            inputs["smsRecovery"] = state ? state.smsRecovery : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as PasswordArgs | undefined;
            inputs["authProvider"] = args ? args.authProvider : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["emailRecovery"] = args ? args.emailRecovery : undefined;
            inputs["groupsIncludeds"] = args ? args.groupsIncludeds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["passwordAutoUnlockMinutes"] = args ? args.passwordAutoUnlockMinutes : undefined;
            inputs["passwordDictionaryLookup"] = args ? args.passwordDictionaryLookup : undefined;
            inputs["passwordExcludeFirstName"] = args ? args.passwordExcludeFirstName : undefined;
            inputs["passwordExcludeLastName"] = args ? args.passwordExcludeLastName : undefined;
            inputs["passwordExcludeUsername"] = args ? args.passwordExcludeUsername : undefined;
            inputs["passwordExpireWarnDays"] = args ? args.passwordExpireWarnDays : undefined;
            inputs["passwordHistoryCount"] = args ? args.passwordHistoryCount : undefined;
            inputs["passwordMaxAgeDays"] = args ? args.passwordMaxAgeDays : undefined;
            inputs["passwordMaxLockoutAttempts"] = args ? args.passwordMaxLockoutAttempts : undefined;
            inputs["passwordMinAgeMinutes"] = args ? args.passwordMinAgeMinutes : undefined;
            inputs["passwordMinLength"] = args ? args.passwordMinLength : undefined;
            inputs["passwordMinLowercase"] = args ? args.passwordMinLowercase : undefined;
            inputs["passwordMinNumber"] = args ? args.passwordMinNumber : undefined;
            inputs["passwordMinSymbol"] = args ? args.passwordMinSymbol : undefined;
            inputs["passwordMinUppercase"] = args ? args.passwordMinUppercase : undefined;
            inputs["passwordShowLockoutFailures"] = args ? args.passwordShowLockoutFailures : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["questionMinLength"] = args ? args.questionMinLength : undefined;
            inputs["questionRecovery"] = args ? args.questionRecovery : undefined;
            inputs["recoveryEmailToken"] = args ? args.recoveryEmailToken : undefined;
            inputs["skipUnlock"] = args ? args.skipUnlock : undefined;
            inputs["smsRecovery"] = args ? args.smsRecovery : undefined;
            inputs["status"] = args ? args.status : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Password.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Password resources.
 */
export interface PasswordState {
    /**
     * Authentication Provider: `"OKTA"` or `"ACTIVE_DIRECTORY"`. Default is `"OKTA"`.
     */
    readonly authProvider?: pulumi.Input<string>;
    /**
     * Policy Description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Enable or disable email password recovery: ACTIVE or INACTIVE.
     */
    readonly emailRecovery?: pulumi.Input<string>;
    /**
     * List of Group IDs to Include.
     */
    readonly groupsIncludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of minutes before a locked account is unlocked: 0 = no limit.
     */
    readonly passwordAutoUnlockMinutes?: pulumi.Input<number>;
    /**
     * Check Passwords Against Common Password Dictionary.
     */
    readonly passwordDictionaryLookup?: pulumi.Input<boolean>;
    /**
     * User firstName attribute must be excluded from the password.
     */
    readonly passwordExcludeFirstName?: pulumi.Input<boolean>;
    /**
     * User lastName attribute must be excluded from the password.
     */
    readonly passwordExcludeLastName?: pulumi.Input<boolean>;
    /**
     * If the user name must be excluded from the password.
     */
    readonly passwordExcludeUsername?: pulumi.Input<boolean>;
    /**
     * Length in days a user will be warned before password expiry: 0 = no warning.
     */
    readonly passwordExpireWarnDays?: pulumi.Input<number>;
    /**
     * Number of distinct passwords that must be created before they can be reused: 0 = none.
     */
    readonly passwordHistoryCount?: pulumi.Input<number>;
    /**
     * Length in days a password is valid before expiry: 0 = no limit.",
     */
    readonly passwordMaxAgeDays?: pulumi.Input<number>;
    /**
     * Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
     */
    readonly passwordMaxLockoutAttempts?: pulumi.Input<number>;
    /**
     * Minimum time interval in minutes between password changes: 0 = no limit.
     */
    readonly passwordMinAgeMinutes?: pulumi.Input<number>;
    /**
     * Minimum password length. Default is 8.
     */
    readonly passwordMinLength?: pulumi.Input<number>;
    /**
     * Minimum number of lower case characters in password.
     */
    readonly passwordMinLowercase?: pulumi.Input<number>;
    /**
     * Minimum number of numbers in password.
     */
    readonly passwordMinNumber?: pulumi.Input<number>;
    /**
     * Minimum number of symbols in password.
     */
    readonly passwordMinSymbol?: pulumi.Input<number>;
    /**
     * Minimum number of upper case characters in password.
     */
    readonly passwordMinUppercase?: pulumi.Input<number>;
    /**
     * If a user should be informed when their account is locked.
     */
    readonly passwordShowLockoutFailures?: pulumi.Input<boolean>;
    /**
     * Priority of the policy.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Min length of the password recovery question answer.
     */
    readonly questionMinLength?: pulumi.Input<number>;
    /**
     * Enable or disable security question password recovery: ACTIVE or INACTIVE.
     */
    readonly questionRecovery?: pulumi.Input<string>;
    /**
     * Lifetime in minutes of the recovery email token.
     */
    readonly recoveryEmailToken?: pulumi.Input<number>;
    /**
     * When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
     */
    readonly skipUnlock?: pulumi.Input<boolean>;
    /**
     * Enable or disable SMS password recovery: ACTIVE or INACTIVE.
     */
    readonly smsRecovery?: pulumi.Input<string>;
    /**
     * Policy Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Password resource.
 */
export interface PasswordArgs {
    /**
     * Authentication Provider: `"OKTA"` or `"ACTIVE_DIRECTORY"`. Default is `"OKTA"`.
     */
    readonly authProvider?: pulumi.Input<string>;
    /**
     * Policy Description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Enable or disable email password recovery: ACTIVE or INACTIVE.
     */
    readonly emailRecovery?: pulumi.Input<string>;
    /**
     * List of Group IDs to Include.
     */
    readonly groupsIncludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy Name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of minutes before a locked account is unlocked: 0 = no limit.
     */
    readonly passwordAutoUnlockMinutes?: pulumi.Input<number>;
    /**
     * Check Passwords Against Common Password Dictionary.
     */
    readonly passwordDictionaryLookup?: pulumi.Input<boolean>;
    /**
     * User firstName attribute must be excluded from the password.
     */
    readonly passwordExcludeFirstName?: pulumi.Input<boolean>;
    /**
     * User lastName attribute must be excluded from the password.
     */
    readonly passwordExcludeLastName?: pulumi.Input<boolean>;
    /**
     * If the user name must be excluded from the password.
     */
    readonly passwordExcludeUsername?: pulumi.Input<boolean>;
    /**
     * Length in days a user will be warned before password expiry: 0 = no warning.
     */
    readonly passwordExpireWarnDays?: pulumi.Input<number>;
    /**
     * Number of distinct passwords that must be created before they can be reused: 0 = none.
     */
    readonly passwordHistoryCount?: pulumi.Input<number>;
    /**
     * Length in days a password is valid before expiry: 0 = no limit.",
     */
    readonly passwordMaxAgeDays?: pulumi.Input<number>;
    /**
     * Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
     */
    readonly passwordMaxLockoutAttempts?: pulumi.Input<number>;
    /**
     * Minimum time interval in minutes between password changes: 0 = no limit.
     */
    readonly passwordMinAgeMinutes?: pulumi.Input<number>;
    /**
     * Minimum password length. Default is 8.
     */
    readonly passwordMinLength?: pulumi.Input<number>;
    /**
     * Minimum number of lower case characters in password.
     */
    readonly passwordMinLowercase?: pulumi.Input<number>;
    /**
     * Minimum number of numbers in password.
     */
    readonly passwordMinNumber?: pulumi.Input<number>;
    /**
     * Minimum number of symbols in password.
     */
    readonly passwordMinSymbol?: pulumi.Input<number>;
    /**
     * Minimum number of upper case characters in password.
     */
    readonly passwordMinUppercase?: pulumi.Input<number>;
    /**
     * If a user should be informed when their account is locked.
     */
    readonly passwordShowLockoutFailures?: pulumi.Input<boolean>;
    /**
     * Priority of the policy.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Min length of the password recovery question answer.
     */
    readonly questionMinLength?: pulumi.Input<number>;
    /**
     * Enable or disable security question password recovery: ACTIVE or INACTIVE.
     */
    readonly questionRecovery?: pulumi.Input<string>;
    /**
     * Lifetime in minutes of the recovery email token.
     */
    readonly recoveryEmailToken?: pulumi.Input<number>;
    /**
     * When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
     */
    readonly skipUnlock?: pulumi.Input<boolean>;
    /**
     * Enable or disable SMS password recovery: ACTIVE or INACTIVE.
     */
    readonly smsRecovery?: pulumi.Input<string>;
    /**
     * Policy Status: `"ACTIVE"` or `"INACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
}
