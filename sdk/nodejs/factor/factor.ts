// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows you to manage the activation of Okta MFA methods.
 * 
 * This resource allows you to manage Okta MFA methods.
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/factor.html.markdown.
 */
export class Factor extends pulumi.CustomResource {
    /**
     * Get an existing Factor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FactorState, opts?: pulumi.CustomResourceOptions): Factor {
        return new Factor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:factor/factor:Factor';

    /**
     * Returns true if the given object is an instance of Factor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Factor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Factor.__pulumiType;
    }

    /**
     * Whether or not to activate the provider, by default it is set to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * Factor provider ID
     */
    public readonly providerId!: pulumi.Output<string>;

    /**
     * Create a Factor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FactorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FactorArgs | FactorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FactorState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["providerId"] = state ? state.providerId : undefined;
        } else {
            const args = argsOrState as FactorArgs | undefined;
            if (!args || args.providerId === undefined) {
                throw new Error("Missing required property 'providerId'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["providerId"] = args ? args.providerId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Factor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Factor resources.
 */
export interface FactorState {
    /**
     * Whether or not to activate the provider, by default it is set to `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * Factor provider ID
     */
    readonly providerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Factor resource.
 */
export interface FactorArgs {
    /**
     * Whether or not to activate the provider, by default it is set to `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * Factor provider ID
     */
    readonly providerId: pulumi.Input<string>;
}
