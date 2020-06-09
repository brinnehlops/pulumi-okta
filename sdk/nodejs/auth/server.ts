// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Authorization Server.
 *
 * This resource allows you to create and configure an Authorization Server.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.auth.Server("example", {
 *     audiences: ["api://example"],
 *     description: "My Example Auth Server",
 *     issuerMode: "CUSTOM_URL",
 *     status: "ACTIVE",
 * });
 * ```
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:auth/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * The recipients that the tokens are intended for. This becomes the `aud` claim in an access token.
     */
    public readonly audiences!: pulumi.Output<string[]>;
    /**
     * The timestamp when the authorization server started to use the `kid` for signing tokens.
     */
    public /*out*/ readonly credentialsLastRotated!: pulumi.Output<string>;
    /**
     * The timestamp when the authorization server changes the key for signing tokens. Only returned when `credentialsRotationMode` is `"AUTO"`.
     */
    public /*out*/ readonly credentialsNextRotation!: pulumi.Output<string>;
    /**
     * The key rotation mode for the authorization server. Can be `"AUTO"` or `"MANUAL"`.
     */
    public readonly credentialsRotationMode!: pulumi.Output<string | undefined>;
    /**
     * The description of the authorization server.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The complete URL for a Custom Authorization Server. This becomes the `iss` claim in an access token.
     */
    public /*out*/ readonly issuer!: pulumi.Output<string>;
    /**
     * Allows you to use a custom issuer URL. It can be set to `"CUSTOM_URL"` or `"ORG_URL"`
     */
    public readonly issuerMode!: pulumi.Output<string | undefined>;
    /**
     * The ID of the JSON Web Key used for signing tokens issued by the authorization server.
     */
    public /*out*/ readonly kid!: pulumi.Output<string>;
    /**
     * The name of the authorization server.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The status of the auth server. It defaults to `"ACTIVE"`
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerState | undefined;
            inputs["audiences"] = state ? state.audiences : undefined;
            inputs["credentialsLastRotated"] = state ? state.credentialsLastRotated : undefined;
            inputs["credentialsNextRotation"] = state ? state.credentialsNextRotation : undefined;
            inputs["credentialsRotationMode"] = state ? state.credentialsRotationMode : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["issuer"] = state ? state.issuer : undefined;
            inputs["issuerMode"] = state ? state.issuerMode : undefined;
            inputs["kid"] = state ? state.kid : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if (!args || args.audiences === undefined) {
                throw new Error("Missing required property 'audiences'");
            }
            inputs["audiences"] = args ? args.audiences : undefined;
            inputs["credentialsRotationMode"] = args ? args.credentialsRotationMode : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["issuerMode"] = args ? args.issuerMode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["credentialsLastRotated"] = undefined /*out*/;
            inputs["credentialsNextRotation"] = undefined /*out*/;
            inputs["issuer"] = undefined /*out*/;
            inputs["kid"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * The recipients that the tokens are intended for. This becomes the `aud` claim in an access token.
     */
    readonly audiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The timestamp when the authorization server started to use the `kid` for signing tokens.
     */
    readonly credentialsLastRotated?: pulumi.Input<string>;
    /**
     * The timestamp when the authorization server changes the key for signing tokens. Only returned when `credentialsRotationMode` is `"AUTO"`.
     */
    readonly credentialsNextRotation?: pulumi.Input<string>;
    /**
     * The key rotation mode for the authorization server. Can be `"AUTO"` or `"MANUAL"`.
     */
    readonly credentialsRotationMode?: pulumi.Input<string>;
    /**
     * The description of the authorization server.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The complete URL for a Custom Authorization Server. This becomes the `iss` claim in an access token.
     */
    readonly issuer?: pulumi.Input<string>;
    /**
     * Allows you to use a custom issuer URL. It can be set to `"CUSTOM_URL"` or `"ORG_URL"`
     */
    readonly issuerMode?: pulumi.Input<string>;
    /**
     * The ID of the JSON Web Key used for signing tokens issued by the authorization server.
     */
    readonly kid?: pulumi.Input<string>;
    /**
     * The name of the authorization server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The status of the auth server. It defaults to `"ACTIVE"`
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The recipients that the tokens are intended for. This becomes the `aud` claim in an access token.
     */
    readonly audiences: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The key rotation mode for the authorization server. Can be `"AUTO"` or `"MANUAL"`.
     */
    readonly credentialsRotationMode?: pulumi.Input<string>;
    /**
     * The description of the authorization server.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Allows you to use a custom issuer URL. It can be set to `"CUSTOM_URL"` or `"ORG_URL"`
     */
    readonly issuerMode?: pulumi.Input<string>;
    /**
     * The name of the authorization server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The status of the auth server. It defaults to `"ACTIVE"`
     */
    readonly status?: pulumi.Input<string>;
}
