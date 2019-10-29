// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an inline hook.
 * 
 * This resource allows you to create and configure an inline hook.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = new okta.inline.Hook("example", {
 *     auth: {
 *         key: "Authorization",
 *         type: "HEADER",
 *         value: "secret",
 *     },
 *     channel: {
 *         method: "POST",
 *         uri: "https://example.com/test",
 *         version: "1.0.0",
 *     },
 *     type: "com.okta.oauth2.tokens.transform",
 *     version: "1.0.1",
 * });
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/inline_hook.html.markdown.
 */
export class Hook extends pulumi.CustomResource {
    /**
     * Get an existing Hook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HookState, opts?: pulumi.CustomResourceOptions): Hook {
        return new Hook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:inline/hook:Hook';

    /**
     * Returns true if the given object is an instance of Hook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hook.__pulumiType;
    }

    /**
     * Authentication required for inline hook request.
     */
    public readonly auth!: pulumi.Output<outputs.inline.HookAuth | undefined>;
    /**
     * Details of the endpoint the inline hook will hit.
     */
    public readonly channel!: pulumi.Output<outputs.inline.HookChannel | undefined>;
    /**
     * Map of headers to send along in inline hook request.
     */
    public readonly headers!: pulumi.Output<outputs.inline.HookHeader[] | undefined>;
    /**
     * The inline hook display name.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The type of hook to trigger. Currently only `"HTTP"` is supported.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The version of the endpoint.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Hook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HookArgs | HookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HookState | undefined;
            inputs["auth"] = state ? state.auth : undefined;
            inputs["channel"] = state ? state.channel : undefined;
            inputs["headers"] = state ? state.headers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as HookArgs | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["auth"] = args ? args.auth : undefined;
            inputs["channel"] = args ? args.channel : undefined;
            inputs["headers"] = args ? args.headers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Hook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Hook resources.
 */
export interface HookState {
    /**
     * Authentication required for inline hook request.
     */
    readonly auth?: pulumi.Input<inputs.inline.HookAuth>;
    /**
     * Details of the endpoint the inline hook will hit.
     */
    readonly channel?: pulumi.Input<inputs.inline.HookChannel>;
    /**
     * Map of headers to send along in inline hook request.
     */
    readonly headers?: pulumi.Input<pulumi.Input<inputs.inline.HookHeader>[]>;
    /**
     * The inline hook display name.
     */
    readonly name?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    /**
     * The type of hook to trigger. Currently only `"HTTP"` is supported.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The version of the endpoint.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Hook resource.
 */
export interface HookArgs {
    /**
     * Authentication required for inline hook request.
     */
    readonly auth?: pulumi.Input<inputs.inline.HookAuth>;
    /**
     * Details of the endpoint the inline hook will hit.
     */
    readonly channel?: pulumi.Input<inputs.inline.HookChannel>;
    /**
     * Map of headers to send along in inline hook request.
     */
    readonly headers?: pulumi.Input<pulumi.Input<inputs.inline.HookHeader>[]>;
    /**
     * The inline hook display name.
     */
    readonly name?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    /**
     * The type of hook to trigger. Currently only `"HTTP"` is supported.
     */
    readonly type: pulumi.Input<string>;
    /**
     * The version of the endpoint.
     */
    readonly version: pulumi.Input<string>;
}
