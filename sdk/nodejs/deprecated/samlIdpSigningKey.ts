// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SamlIdpSigningKey extends pulumi.CustomResource {
    /**
     * Get an existing SamlIdpSigningKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamlIdpSigningKeyState, opts?: pulumi.CustomResourceOptions): SamlIdpSigningKey {
        return new SamlIdpSigningKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:deprecated/samlIdpSigningKey:SamlIdpSigningKey';

    /**
     * Returns true if the given object is an instance of SamlIdpSigningKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamlIdpSigningKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SamlIdpSigningKey.__pulumiType;
    }

    public /*out*/ readonly created!: pulumi.Output<string>;
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    public /*out*/ readonly kid!: pulumi.Output<string>;
    public /*out*/ readonly kty!: pulumi.Output<string>;
    public /*out*/ readonly use!: pulumi.Output<string>;
    /**
     * base64-encoded X.509 certificate chain with DER encoding
     */
    public readonly x5cs!: pulumi.Output<string[]>;
    public /*out*/ readonly x5tS256!: pulumi.Output<string>;

    /**
     * Create a SamlIdpSigningKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlIdpSigningKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamlIdpSigningKeyArgs | SamlIdpSigningKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SamlIdpSigningKeyState | undefined;
            inputs["created"] = state ? state.created : undefined;
            inputs["expiresAt"] = state ? state.expiresAt : undefined;
            inputs["kid"] = state ? state.kid : undefined;
            inputs["kty"] = state ? state.kty : undefined;
            inputs["use"] = state ? state.use : undefined;
            inputs["x5cs"] = state ? state.x5cs : undefined;
            inputs["x5tS256"] = state ? state.x5tS256 : undefined;
        } else {
            const args = argsOrState as SamlIdpSigningKeyArgs | undefined;
            if (!args || args.x5cs === undefined) {
                throw new Error("Missing required property 'x5cs'");
            }
            inputs["x5cs"] = args ? args.x5cs : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["expiresAt"] = undefined /*out*/;
            inputs["kid"] = undefined /*out*/;
            inputs["kty"] = undefined /*out*/;
            inputs["use"] = undefined /*out*/;
            inputs["x5tS256"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SamlIdpSigningKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SamlIdpSigningKey resources.
 */
export interface SamlIdpSigningKeyState {
    readonly created?: pulumi.Input<string>;
    readonly expiresAt?: pulumi.Input<string>;
    readonly kid?: pulumi.Input<string>;
    readonly kty?: pulumi.Input<string>;
    readonly use?: pulumi.Input<string>;
    /**
     * base64-encoded X.509 certificate chain with DER encoding
     */
    readonly x5cs?: pulumi.Input<pulumi.Input<string>[]>;
    readonly x5tS256?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SamlIdpSigningKey resource.
 */
export interface SamlIdpSigningKeyArgs {
    /**
     * base64-encoded X.509 certificate chain with DER encoding
     */
    readonly x5cs: pulumi.Input<pulumi.Input<string>[]>;
}
