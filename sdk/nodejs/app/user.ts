// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Application User.
 *
 * This resource allows you to create and configure an Application User.
 *
 * __When using this resource, make sure to add the following `lifefycle` argument to the application resource you are assigning to:__
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.app.User("example", {
 *     appId: "<app_id>",
 *     userId: "<user id>",
 *     username: "example",
 * });
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * App to associate user with.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The password to use.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The JSON profile of the App User.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * User to associate the application with.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * The username to use for the app user.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["profile"] = state ? state.profile : undefined;
            inputs["userId"] = state ? state.userId : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.userId === undefined) {
                throw new Error("Missing required property 'userId'");
            }
            if (!args || args.username === undefined) {
                throw new Error("Missing required property 'username'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["profile"] = args ? args.profile : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * App to associate user with.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The password to use.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The JSON profile of the App User.
     */
    readonly profile?: pulumi.Input<string>;
    /**
     * User to associate the application with.
     */
    readonly userId?: pulumi.Input<string>;
    /**
     * The username to use for the app user.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * App to associate user with.
     */
    readonly appId: pulumi.Input<string>;
    /**
     * The password to use.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The JSON profile of the App User.
     */
    readonly profile?: pulumi.Input<string>;
    /**
     * User to associate the application with.
     */
    readonly userId: pulumi.Input<string>;
    /**
     * The username to use for the app user.
     */
    readonly username: pulumi.Input<string>;
}
