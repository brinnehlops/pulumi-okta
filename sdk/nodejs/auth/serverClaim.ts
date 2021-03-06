// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Authorization Server Claim.
 *
 * This resource allows you to create and configure an Authorization Server Claim.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.auth.ServerClaim("example", {
 *     authServerId: "<auth server id>",
 *     claimType: "IDENTITY",
 *     scopes: [okta_auth_server_scope_example.name],
 *     value: "String.substringAfter(user.email, \"@\") == \"example.com\"",
 * });
 * ```
 */
export class ServerClaim extends pulumi.CustomResource {
    /**
     * Get an existing ServerClaim resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerClaimState, opts?: pulumi.CustomResourceOptions): ServerClaim {
        return new ServerClaim(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:auth/serverClaim:ServerClaim';

    /**
     * Returns true if the given object is an instance of ServerClaim.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerClaim {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerClaim.__pulumiType;
    }

    /**
     * Specifies whether to include claims in token, by default is is set to `true`.
     */
    public readonly alwaysIncludeInToken!: pulumi.Output<boolean | undefined>;
    /**
     * The Application's display name.
     */
    public readonly authServerId!: pulumi.Output<string>;
    /**
     * Specifies whether the claim is for an access token `"RESOURCE"` or ID token `"IDENTITY"`.
     */
    public readonly claimType!: pulumi.Output<string>;
    /**
     * Specifies the type of group filter if `valueType` is `"GROUPS"`. Can be set to one of the following `"STARTS_WITH"`, `"EQUALS"`, `"CONTAINS"`, `"REGEX"`.
     */
    public readonly groupFilterType!: pulumi.Output<string | undefined>;
    /**
     * The name of the claim.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of scopes the auth server claim is tied to.
     */
    public readonly scopes!: pulumi.Output<string[] | undefined>;
    /**
     * The status of the application. It defaults to `"ACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The value of the claim.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The type of value of the claim. It can be set to `"EXPRESSION"` or `"GROUPS"`. It defaults to `"EXPRESSION"`.
     */
    public readonly valueType!: pulumi.Output<string | undefined>;

    /**
     * Create a ServerClaim resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerClaimArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerClaimArgs | ServerClaimState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerClaimState | undefined;
            inputs["alwaysIncludeInToken"] = state ? state.alwaysIncludeInToken : undefined;
            inputs["authServerId"] = state ? state.authServerId : undefined;
            inputs["claimType"] = state ? state.claimType : undefined;
            inputs["groupFilterType"] = state ? state.groupFilterType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["value"] = state ? state.value : undefined;
            inputs["valueType"] = state ? state.valueType : undefined;
        } else {
            const args = argsOrState as ServerClaimArgs | undefined;
            if (!args || args.authServerId === undefined) {
                throw new Error("Missing required property 'authServerId'");
            }
            if (!args || args.claimType === undefined) {
                throw new Error("Missing required property 'claimType'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["alwaysIncludeInToken"] = args ? args.alwaysIncludeInToken : undefined;
            inputs["authServerId"] = args ? args.authServerId : undefined;
            inputs["claimType"] = args ? args.claimType : undefined;
            inputs["groupFilterType"] = args ? args.groupFilterType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["valueType"] = args ? args.valueType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerClaim.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerClaim resources.
 */
export interface ServerClaimState {
    /**
     * Specifies whether to include claims in token, by default is is set to `true`.
     */
    readonly alwaysIncludeInToken?: pulumi.Input<boolean>;
    /**
     * The Application's display name.
     */
    readonly authServerId?: pulumi.Input<string>;
    /**
     * Specifies whether the claim is for an access token `"RESOURCE"` or ID token `"IDENTITY"`.
     */
    readonly claimType?: pulumi.Input<string>;
    /**
     * Specifies the type of group filter if `valueType` is `"GROUPS"`. Can be set to one of the following `"STARTS_WITH"`, `"EQUALS"`, `"CONTAINS"`, `"REGEX"`.
     */
    readonly groupFilterType?: pulumi.Input<string>;
    /**
     * The name of the claim.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The list of scopes the auth server claim is tied to.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the application. It defaults to `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The value of the claim.
     */
    readonly value?: pulumi.Input<string>;
    /**
     * The type of value of the claim. It can be set to `"EXPRESSION"` or `"GROUPS"`. It defaults to `"EXPRESSION"`.
     */
    readonly valueType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerClaim resource.
 */
export interface ServerClaimArgs {
    /**
     * Specifies whether to include claims in token, by default is is set to `true`.
     */
    readonly alwaysIncludeInToken?: pulumi.Input<boolean>;
    /**
     * The Application's display name.
     */
    readonly authServerId: pulumi.Input<string>;
    /**
     * Specifies whether the claim is for an access token `"RESOURCE"` or ID token `"IDENTITY"`.
     */
    readonly claimType: pulumi.Input<string>;
    /**
     * Specifies the type of group filter if `valueType` is `"GROUPS"`. Can be set to one of the following `"STARTS_WITH"`, `"EQUALS"`, `"CONTAINS"`, `"REGEX"`.
     */
    readonly groupFilterType?: pulumi.Input<string>;
    /**
     * The name of the claim.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The list of scopes the auth server claim is tied to.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the application. It defaults to `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The value of the claim.
     */
    readonly value: pulumi.Input<string>;
    /**
     * The type of value of the claim. It can be set to `"EXPRESSION"` or `"GROUPS"`. It defaults to `"EXPRESSION"`.
     */
    readonly valueType?: pulumi.Input<string>;
}
