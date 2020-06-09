// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Okta Network Zone.
 *
 * This resource allows you to create and configure an Okta Network Zone.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.network.Zone("example", {
 *     gateways: [
 *         "1.2.3.4/24",
 *         "2.3.4.5-2.3.4.15",
 *     ],
 *     proxies: [
 *         "2.2.3.4/24",
 *         "3.3.4.5-3.3.4.15",
 *     ],
 *     type: "IP",
 * });
 * ```
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:network/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
     */
    public readonly dynamicLocations!: pulumi.Output<string[] | undefined>;
    /**
     * Array of values in CIDR/range form.
     */
    public readonly gateways!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the Network Zone Resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Array of values in CIDR/range form.
     */
    public readonly proxies!: pulumi.Output<string[] | undefined>;
    /**
     * Type of the Network Zone - can either be IP or DYNAMIC only.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ZoneState | undefined;
            inputs["dynamicLocations"] = state ? state.dynamicLocations : undefined;
            inputs["gateways"] = state ? state.gateways : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["proxies"] = state ? state.proxies : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["dynamicLocations"] = args ? args.dynamicLocations : undefined;
            inputs["gateways"] = args ? args.gateways : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["proxies"] = args ? args.proxies : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Zone.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
     */
    readonly dynamicLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of values in CIDR/range form.
     */
    readonly gateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Network Zone Resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Array of values in CIDR/range form.
     */
    readonly proxies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the Network Zone - can either be IP or DYNAMIC only.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
     */
    readonly dynamicLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of values in CIDR/range form.
     */
    readonly gateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Network Zone Resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Array of values in CIDR/range form.
     */
    readonly proxies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the Network Zone - can either be IP or DYNAMIC only.
     */
    readonly type: pulumi.Input<string>;
}
