// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class MfaPolicy extends pulumi.CustomResource {
    /**
     * Get an existing MfaPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaPolicyState, opts?: pulumi.CustomResourceOptions): MfaPolicy {
        return new MfaPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:deprecated/mfaPolicy:MfaPolicy';

    /**
     * Returns true if the given object is an instance of MfaPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaPolicy.__pulumiType;
    }

    /**
     * Policy Description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly duo!: pulumi.Output<outputs.deprecated.MfaPolicyDuo | undefined>;
    public readonly fidoU2f!: pulumi.Output<outputs.deprecated.MfaPolicyFidoU2f | undefined>;
    public readonly fidoWebauthn!: pulumi.Output<outputs.deprecated.MfaPolicyFidoWebauthn | undefined>;
    public readonly googleOtp!: pulumi.Output<outputs.deprecated.MfaPolicyGoogleOtp | undefined>;
    /**
     * List of Group IDs to Include
     */
    public readonly groupsIncludeds!: pulumi.Output<string[] | undefined>;
    /**
     * Policy Name
     */
    public readonly name!: pulumi.Output<string>;
    public readonly oktaCall!: pulumi.Output<outputs.deprecated.MfaPolicyOktaCall | undefined>;
    public readonly oktaOtp!: pulumi.Output<outputs.deprecated.MfaPolicyOktaOtp | undefined>;
    public readonly oktaPassword!: pulumi.Output<outputs.deprecated.MfaPolicyOktaPassword | undefined>;
    public readonly oktaPush!: pulumi.Output<outputs.deprecated.MfaPolicyOktaPush | undefined>;
    public readonly oktaQuestion!: pulumi.Output<outputs.deprecated.MfaPolicyOktaQuestion | undefined>;
    public readonly oktaSms!: pulumi.Output<outputs.deprecated.MfaPolicyOktaSms | undefined>;
    /**
     * Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an
     * invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    public readonly rsaToken!: pulumi.Output<outputs.deprecated.MfaPolicyRsaToken | undefined>;
    /**
     * Policy Status: ACTIVE or INACTIVE.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly symantecVip!: pulumi.Output<outputs.deprecated.MfaPolicySymantecVip | undefined>;
    public readonly yubikeyToken!: pulumi.Output<outputs.deprecated.MfaPolicyYubikeyToken | undefined>;

    /**
     * Create a MfaPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MfaPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaPolicyArgs | MfaPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MfaPolicyState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["duo"] = state ? state.duo : undefined;
            inputs["fidoU2f"] = state ? state.fidoU2f : undefined;
            inputs["fidoWebauthn"] = state ? state.fidoWebauthn : undefined;
            inputs["googleOtp"] = state ? state.googleOtp : undefined;
            inputs["groupsIncludeds"] = state ? state.groupsIncludeds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["oktaCall"] = state ? state.oktaCall : undefined;
            inputs["oktaOtp"] = state ? state.oktaOtp : undefined;
            inputs["oktaPassword"] = state ? state.oktaPassword : undefined;
            inputs["oktaPush"] = state ? state.oktaPush : undefined;
            inputs["oktaQuestion"] = state ? state.oktaQuestion : undefined;
            inputs["oktaSms"] = state ? state.oktaSms : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["rsaToken"] = state ? state.rsaToken : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["symantecVip"] = state ? state.symantecVip : undefined;
            inputs["yubikeyToken"] = state ? state.yubikeyToken : undefined;
        } else {
            const args = argsOrState as MfaPolicyArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["duo"] = args ? args.duo : undefined;
            inputs["fidoU2f"] = args ? args.fidoU2f : undefined;
            inputs["fidoWebauthn"] = args ? args.fidoWebauthn : undefined;
            inputs["googleOtp"] = args ? args.googleOtp : undefined;
            inputs["groupsIncludeds"] = args ? args.groupsIncludeds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["oktaCall"] = args ? args.oktaCall : undefined;
            inputs["oktaOtp"] = args ? args.oktaOtp : undefined;
            inputs["oktaPassword"] = args ? args.oktaPassword : undefined;
            inputs["oktaPush"] = args ? args.oktaPush : undefined;
            inputs["oktaQuestion"] = args ? args.oktaQuestion : undefined;
            inputs["oktaSms"] = args ? args.oktaSms : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["rsaToken"] = args ? args.rsaToken : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["symantecVip"] = args ? args.symantecVip : undefined;
            inputs["yubikeyToken"] = args ? args.yubikeyToken : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MfaPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaPolicy resources.
 */
export interface MfaPolicyState {
    /**
     * Policy Description
     */
    readonly description?: pulumi.Input<string>;
    readonly duo?: pulumi.Input<inputs.deprecated.MfaPolicyDuo>;
    readonly fidoU2f?: pulumi.Input<inputs.deprecated.MfaPolicyFidoU2f>;
    readonly fidoWebauthn?: pulumi.Input<inputs.deprecated.MfaPolicyFidoWebauthn>;
    readonly googleOtp?: pulumi.Input<inputs.deprecated.MfaPolicyGoogleOtp>;
    /**
     * List of Group IDs to Include
     */
    readonly groupsIncludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy Name
     */
    readonly name?: pulumi.Input<string>;
    readonly oktaCall?: pulumi.Input<inputs.deprecated.MfaPolicyOktaCall>;
    readonly oktaOtp?: pulumi.Input<inputs.deprecated.MfaPolicyOktaOtp>;
    readonly oktaPassword?: pulumi.Input<inputs.deprecated.MfaPolicyOktaPassword>;
    readonly oktaPush?: pulumi.Input<inputs.deprecated.MfaPolicyOktaPush>;
    readonly oktaQuestion?: pulumi.Input<inputs.deprecated.MfaPolicyOktaQuestion>;
    readonly oktaSms?: pulumi.Input<inputs.deprecated.MfaPolicyOktaSms>;
    /**
     * Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an
     * invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    readonly priority?: pulumi.Input<number>;
    readonly rsaToken?: pulumi.Input<inputs.deprecated.MfaPolicyRsaToken>;
    /**
     * Policy Status: ACTIVE or INACTIVE.
     */
    readonly status?: pulumi.Input<string>;
    readonly symantecVip?: pulumi.Input<inputs.deprecated.MfaPolicySymantecVip>;
    readonly yubikeyToken?: pulumi.Input<inputs.deprecated.MfaPolicyYubikeyToken>;
}

/**
 * The set of arguments for constructing a MfaPolicy resource.
 */
export interface MfaPolicyArgs {
    /**
     * Policy Description
     */
    readonly description?: pulumi.Input<string>;
    readonly duo?: pulumi.Input<inputs.deprecated.MfaPolicyDuo>;
    readonly fidoU2f?: pulumi.Input<inputs.deprecated.MfaPolicyFidoU2f>;
    readonly fidoWebauthn?: pulumi.Input<inputs.deprecated.MfaPolicyFidoWebauthn>;
    readonly googleOtp?: pulumi.Input<inputs.deprecated.MfaPolicyGoogleOtp>;
    /**
     * List of Group IDs to Include
     */
    readonly groupsIncludeds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policy Name
     */
    readonly name?: pulumi.Input<string>;
    readonly oktaCall?: pulumi.Input<inputs.deprecated.MfaPolicyOktaCall>;
    readonly oktaOtp?: pulumi.Input<inputs.deprecated.MfaPolicyOktaOtp>;
    readonly oktaPassword?: pulumi.Input<inputs.deprecated.MfaPolicyOktaPassword>;
    readonly oktaPush?: pulumi.Input<inputs.deprecated.MfaPolicyOktaPush>;
    readonly oktaQuestion?: pulumi.Input<inputs.deprecated.MfaPolicyOktaQuestion>;
    readonly oktaSms?: pulumi.Input<inputs.deprecated.MfaPolicyOktaSms>;
    /**
     * Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an
     * invalid priority is provided. API defaults it to the last/lowest if not there.
     */
    readonly priority?: pulumi.Input<number>;
    readonly rsaToken?: pulumi.Input<inputs.deprecated.MfaPolicyRsaToken>;
    /**
     * Policy Status: ACTIVE or INACTIVE.
     */
    readonly status?: pulumi.Input<string>;
    readonly symantecVip?: pulumi.Input<inputs.deprecated.MfaPolicySymantecVip>;
    readonly yubikeyToken?: pulumi.Input<inputs.deprecated.MfaPolicyYubikeyToken>;
}
