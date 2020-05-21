// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an Application User Schema property.
 *
 * This resource allows you to create and configure a custom user schema property and associate it with an application.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.app.UserSchema("example", {
 *     appId: "<app id>",
 *     description: "My custom property name",
 *     index: "customPropertyName",
 *     master: "OKTA",
 *     scope: "SELF",
 *     title: "customPropertyName",
 *     type: "string",
 * });
 * ```
 */
export class UserSchema extends pulumi.CustomResource {
    /**
     * Get an existing UserSchema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserSchemaState, opts?: pulumi.CustomResourceOptions): UserSchema {
        return new UserSchema(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/userSchema:UserSchema';

    /**
     * Returns true if the given object is an instance of UserSchema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserSchema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserSchema.__pulumiType;
    }

    /**
     * The Application's ID the user custom schema property should be assigned to.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Array of values that an array property's items can be set to.
     */
    public readonly arrayEnums!: pulumi.Output<string[] | undefined>;
    /**
     * Display name and value an enum array can be set to.
     */
    public readonly arrayOneOfs!: pulumi.Output<outputs.app.UserSchemaArrayOneOf[] | undefined>;
    /**
     * The type of the array elements if `type` is set to `"array"`.
     */
    public readonly arrayType!: pulumi.Output<string | undefined>;
    /**
     * The description of the user schema property.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Array of values a primitive property can be set to. See `arrayEnum` for arrays.
     */
    public readonly enums!: pulumi.Output<string[] | undefined>;
    /**
     * External name of the user schema property.
     */
    public readonly externalName!: pulumi.Output<string | undefined>;
    /**
     * The property name.
     */
    public readonly index!: pulumi.Output<string>;
    /**
     * Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
     */
    public readonly master!: pulumi.Output<string | undefined>;
    /**
     * The maximum length of the user property value. Only applies to type `"string"`.
     */
    public readonly maxLength!: pulumi.Output<number | undefined>;
    /**
     * The minimum length of the user property value. Only applies to type `"string"`.
     */
    public readonly minLength!: pulumi.Output<number | undefined>;
    /**
     * Array of maps containing a mapping for display name to enum value.
     */
    public readonly oneOfs!: pulumi.Output<outputs.app.UserSchemaOneOf[] | undefined>;
    /**
     * Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
     */
    public readonly permissions!: pulumi.Output<string | undefined>;
    /**
     * Whether the property is required for this application's users.
     */
    public readonly required!: pulumi.Output<boolean | undefined>;
    /**
     * determines whether an app user attribute can be set at the Individual or Group Level.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * display name for the enum value.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a UserSchema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserSchemaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserSchemaArgs | UserSchemaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserSchemaState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["arrayEnums"] = state ? state.arrayEnums : undefined;
            inputs["arrayOneOfs"] = state ? state.arrayOneOfs : undefined;
            inputs["arrayType"] = state ? state.arrayType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enums"] = state ? state.enums : undefined;
            inputs["externalName"] = state ? state.externalName : undefined;
            inputs["index"] = state ? state.index : undefined;
            inputs["master"] = state ? state.master : undefined;
            inputs["maxLength"] = state ? state.maxLength : undefined;
            inputs["minLength"] = state ? state.minLength : undefined;
            inputs["oneOfs"] = state ? state.oneOfs : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["required"] = state ? state.required : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as UserSchemaArgs | undefined;
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
            inputs["arrayEnums"] = args ? args.arrayEnums : undefined;
            inputs["arrayOneOfs"] = args ? args.arrayOneOfs : undefined;
            inputs["arrayType"] = args ? args.arrayType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enums"] = args ? args.enums : undefined;
            inputs["externalName"] = args ? args.externalName : undefined;
            inputs["index"] = args ? args.index : undefined;
            inputs["master"] = args ? args.master : undefined;
            inputs["maxLength"] = args ? args.maxLength : undefined;
            inputs["minLength"] = args ? args.minLength : undefined;
            inputs["oneOfs"] = args ? args.oneOfs : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["required"] = args ? args.required : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserSchema.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserSchema resources.
 */
export interface UserSchemaState {
    /**
     * The Application's ID the user custom schema property should be assigned to.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * Array of values that an array property's items can be set to.
     */
    readonly arrayEnums?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Display name and value an enum array can be set to.
     */
    readonly arrayOneOfs?: pulumi.Input<pulumi.Input<inputs.app.UserSchemaArrayOneOf>[]>;
    /**
     * The type of the array elements if `type` is set to `"array"`.
     */
    readonly arrayType?: pulumi.Input<string>;
    /**
     * The description of the user schema property.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Array of values a primitive property can be set to. See `arrayEnum` for arrays.
     */
    readonly enums?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * External name of the user schema property.
     */
    readonly externalName?: pulumi.Input<string>;
    /**
     * The property name.
     */
    readonly index?: pulumi.Input<string>;
    /**
     * Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
     */
    readonly master?: pulumi.Input<string>;
    /**
     * The maximum length of the user property value. Only applies to type `"string"`.
     */
    readonly maxLength?: pulumi.Input<number>;
    /**
     * The minimum length of the user property value. Only applies to type `"string"`.
     */
    readonly minLength?: pulumi.Input<number>;
    /**
     * Array of maps containing a mapping for display name to enum value.
     */
    readonly oneOfs?: pulumi.Input<pulumi.Input<inputs.app.UserSchemaOneOf>[]>;
    /**
     * Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
     */
    readonly permissions?: pulumi.Input<string>;
    /**
     * Whether the property is required for this application's users.
     */
    readonly required?: pulumi.Input<boolean>;
    /**
     * determines whether an app user attribute can be set at the Individual or Group Level.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * display name for the enum value.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserSchema resource.
 */
export interface UserSchemaArgs {
    /**
     * The Application's ID the user custom schema property should be assigned to.
     */
    readonly appId: pulumi.Input<string>;
    /**
     * Array of values that an array property's items can be set to.
     */
    readonly arrayEnums?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Display name and value an enum array can be set to.
     */
    readonly arrayOneOfs?: pulumi.Input<pulumi.Input<inputs.app.UserSchemaArrayOneOf>[]>;
    /**
     * The type of the array elements if `type` is set to `"array"`.
     */
    readonly arrayType?: pulumi.Input<string>;
    /**
     * The description of the user schema property.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Array of values a primitive property can be set to. See `arrayEnum` for arrays.
     */
    readonly enums?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * External name of the user schema property.
     */
    readonly externalName?: pulumi.Input<string>;
    /**
     * The property name.
     */
    readonly index: pulumi.Input<string>;
    /**
     * Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
     */
    readonly master?: pulumi.Input<string>;
    /**
     * The maximum length of the user property value. Only applies to type `"string"`.
     */
    readonly maxLength?: pulumi.Input<number>;
    /**
     * The minimum length of the user property value. Only applies to type `"string"`.
     */
    readonly minLength?: pulumi.Input<number>;
    /**
     * Array of maps containing a mapping for display name to enum value.
     */
    readonly oneOfs?: pulumi.Input<pulumi.Input<inputs.app.UserSchemaOneOf>[]>;
    /**
     * Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
     */
    readonly permissions?: pulumi.Input<string>;
    /**
     * Whether the property is required for this application's users.
     */
    readonly required?: pulumi.Input<boolean>;
    /**
     * determines whether an app user attribute can be set at the Individual or Group Level.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * display name for the enum value.
     */
    readonly title: pulumi.Input<string>;
    /**
     * The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
     */
    readonly type: pulumi.Input<string>;
}
