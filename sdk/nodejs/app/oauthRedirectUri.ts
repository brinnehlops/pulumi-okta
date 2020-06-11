// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class OAuthRedirectUri extends pulumi.CustomResource {
    /**
     * Get an existing OAuthRedirectUri resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OAuthRedirectUriState, opts?: pulumi.CustomResourceOptions): OAuthRedirectUri {
        return new OAuthRedirectUri(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/oAuthRedirectUri:OAuthRedirectUri';

    /**
     * Returns true if the given object is an instance of OAuthRedirectUri.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OAuthRedirectUri {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OAuthRedirectUri.__pulumiType;
    }

    public readonly appId!: pulumi.Output<string>;
    /**
     * Redirect URI to append to Okta OIDC application.
     */
    public readonly uri!: pulumi.Output<string>;

    /**
     * Create a OAuthRedirectUri resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OAuthRedirectUriArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OAuthRedirectUriArgs | OAuthRedirectUriState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OAuthRedirectUriState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as OAuthRedirectUriArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.uri === undefined) {
                throw new Error("Missing required property 'uri'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["uri"] = args ? args.uri : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OAuthRedirectUri.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OAuthRedirectUri resources.
 */
export interface OAuthRedirectUriState {
    readonly appId?: pulumi.Input<string>;
    /**
     * Redirect URI to append to Okta OIDC application.
     */
    readonly uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OAuthRedirectUri resource.
 */
export interface OAuthRedirectUriArgs {
    readonly appId: pulumi.Input<string>;
    /**
     * Redirect URI to append to Okta OIDC application.
     */
    readonly uri: pulumi.Input<string>;
}