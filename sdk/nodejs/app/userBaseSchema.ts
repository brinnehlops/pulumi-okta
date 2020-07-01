// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Application User Base Schema property.
 *
 * This resource allows you to configure a base app user schema property.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.app.UserBaseSchema("example", {
 *     appId: "<app id>",
 *     index: "customPropertyName",
 *     master: "OKTA",
 *     title: "customPropertyName",
 *     type: "string",
 * });
 * ```
 */
export class UserBaseSchema extends pulumi.CustomResource {
    /**
     * Get an existing UserBaseSchema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserBaseSchemaState, opts?: pulumi.CustomResourceOptions): UserBaseSchema {
        return new UserBaseSchema(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/userBaseSchema:UserBaseSchema';

    /**
     * Returns true if the given object is an instance of UserBaseSchema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserBaseSchema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserBaseSchema.__pulumiType;
    }

    /**
     * The Application's ID the user schema property should be assigned to.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The property name.
     */
    public readonly index!: pulumi.Output<string>;
    /**
     * Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
     */
    public readonly master!: pulumi.Output<string | undefined>;
    /**
     * Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
     */
    public readonly permissions!: pulumi.Output<string | undefined>;
    /**
     * Whether the property is required for this application's users.
     */
    public readonly required!: pulumi.Output<boolean | undefined>;
    /**
     * The property display name.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a UserBaseSchema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserBaseSchemaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserBaseSchemaArgs | UserBaseSchemaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserBaseSchemaState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["index"] = state ? state.index : undefined;
            inputs["master"] = state ? state.master : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["required"] = state ? state.required : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as UserBaseSchemaArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.index === undefined) {
                throw new Error("Missing required property 'index'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["index"] = args ? args.index : undefined;
            inputs["master"] = args ? args.master : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["required"] = args ? args.required : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserBaseSchema.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserBaseSchema resources.
 */
export interface UserBaseSchemaState {
    /**
     * The Application's ID the user schema property should be assigned to.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The property name.
     */
    readonly index?: pulumi.Input<string>;
    /**
     * Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
     */
    readonly master?: pulumi.Input<string>;
    /**
     * Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
     */
    readonly permissions?: pulumi.Input<string>;
    /**
     * Whether the property is required for this application's users.
     */
    readonly required?: pulumi.Input<boolean>;
    /**
     * The property display name.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserBaseSchema resource.
 */
export interface UserBaseSchemaArgs {
    /**
     * The Application's ID the user schema property should be assigned to.
     */
    readonly appId: pulumi.Input<string>;
    /**
     * The property name.
     */
    readonly index: pulumi.Input<string>;
    /**
     * Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
     */
    readonly master?: pulumi.Input<string>;
    /**
     * Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
     */
    readonly permissions?: pulumi.Input<string>;
    /**
     * Whether the property is required for this application's users.
     */
    readonly required?: pulumi.Input<boolean>;
    /**
     * The property display name.
     */
    readonly title: pulumi.Input<string>;
    /**
     * The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
     */
    readonly type: pulumi.Input<string>;
}
