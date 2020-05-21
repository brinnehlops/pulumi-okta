// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Secure Password Store Application.
 *
 * This resource allows you to create and configure a Secure Password Store Application.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.app.SecurePasswordStore("example", {
 *     credentialsScheme: "ADMIN_SETS_CREDENTIALS",
 *     label: "example",
 *     passwordField: "pass",
 *     url: "http://test.com",
 *     usernameField: "user",
 * });
 * ```
 */
export class SecurePasswordStore extends pulumi.CustomResource {
    /**
     * Get an existing SecurePasswordStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurePasswordStoreState, opts?: pulumi.CustomResourceOptions): SecurePasswordStore {
        return new SecurePasswordStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/securePasswordStore:SecurePasswordStore';

    /**
     * Returns true if the given object is an instance of SecurePasswordStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurePasswordStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurePasswordStore.__pulumiType;
    }

    /**
     * Custom error page URL.
     */
    public readonly accessibilityErrorRedirectUrl!: pulumi.Output<string | undefined>;
    /**
     * Enable self service. By default it is `false`.
     */
    public readonly accessibilitySelfService!: pulumi.Output<boolean | undefined>;
    /**
     * Display auto submit toolbar.
     */
    public readonly autoSubmitToolbar!: pulumi.Output<boolean | undefined>;
    /**
     * Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.
     */
    public readonly credentialsScheme!: pulumi.Output<string | undefined>;
    /**
     * Groups associated with the application. See `okta.app.GroupAssignment` for a more flexible approach.
     */
    public readonly groups!: pulumi.Output<string[] | undefined>;
    /**
     * Do not display application icon on mobile app.
     */
    public readonly hideIos!: pulumi.Output<boolean | undefined>;
    /**
     * Do not display application icon to users.
     */
    public readonly hideWeb!: pulumi.Output<boolean | undefined>;
    /**
     * The display name of the Application.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * Name assigned to the application by Okta.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Name of optional param in the login form.
     */
    public readonly optionalField1!: pulumi.Output<string | undefined>;
    /**
     * Name of optional value in the login form.
     */
    public readonly optionalField1Value!: pulumi.Output<string | undefined>;
    /**
     * Name of optional param in the login form.
     */
    public readonly optionalField2!: pulumi.Output<string | undefined>;
    /**
     * Name of optional value in the login form.
     */
    public readonly optionalField2Value!: pulumi.Output<string | undefined>;
    /**
     * Name of optional param in the login form.
     */
    public readonly optionalField3!: pulumi.Output<string | undefined>;
    /**
     * Name of optional value in the login form.
     */
    public readonly optionalField3Value!: pulumi.Output<string | undefined>;
    /**
     * Login password field.
     */
    public readonly passwordField!: pulumi.Output<string>;
    /**
     * Allow user to reveal password.
     */
    public readonly revealPassword!: pulumi.Output<boolean | undefined>;
    /**
     * Shared password, required for certain schemes.
     */
    public readonly sharedPassword!: pulumi.Output<string | undefined>;
    /**
     * Shared username, required for certain schemes.
     */
    public readonly sharedUsername!: pulumi.Output<string | undefined>;
    /**
     * Sign on mode of application.
     */
    public /*out*/ readonly signOnMode!: pulumi.Output<string>;
    /**
     * Status of application. By default it is `"ACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Login URL.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The default username assigned to each user.
     */
    public /*out*/ readonly userNameTemplate!: pulumi.Output<string>;
    /**
     * The Username template type.
     */
    public /*out*/ readonly userNameTemplateType!: pulumi.Output<string>;
    /**
     * Login username field.
     */
    public readonly usernameField!: pulumi.Output<string>;
    /**
     * The users assigned to the application. See `okta.app.User` for a more flexible approach.
     */
    public readonly users!: pulumi.Output<outputs.app.SecurePasswordStoreUser[] | undefined>;

    /**
     * Create a SecurePasswordStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurePasswordStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurePasswordStoreArgs | SecurePasswordStoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecurePasswordStoreState | undefined;
            inputs["accessibilityErrorRedirectUrl"] = state ? state.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = state ? state.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = state ? state.autoSubmitToolbar : undefined;
            inputs["credentialsScheme"] = state ? state.credentialsScheme : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["hideIos"] = state ? state.hideIos : undefined;
            inputs["hideWeb"] = state ? state.hideWeb : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["optionalField1"] = state ? state.optionalField1 : undefined;
            inputs["optionalField1Value"] = state ? state.optionalField1Value : undefined;
            inputs["optionalField2"] = state ? state.optionalField2 : undefined;
            inputs["optionalField2Value"] = state ? state.optionalField2Value : undefined;
            inputs["optionalField3"] = state ? state.optionalField3 : undefined;
            inputs["optionalField3Value"] = state ? state.optionalField3Value : undefined;
            inputs["passwordField"] = state ? state.passwordField : undefined;
            inputs["revealPassword"] = state ? state.revealPassword : undefined;
            inputs["sharedPassword"] = state ? state.sharedPassword : undefined;
            inputs["sharedUsername"] = state ? state.sharedUsername : undefined;
            inputs["signOnMode"] = state ? state.signOnMode : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["userNameTemplate"] = state ? state.userNameTemplate : undefined;
            inputs["userNameTemplateType"] = state ? state.userNameTemplateType : undefined;
            inputs["usernameField"] = state ? state.usernameField : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as SecurePasswordStoreArgs | undefined;
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            if (!args || args.passwordField === undefined) {
                throw new Error("Missing required property 'passwordField'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            if (!args || args.usernameField === undefined) {
                throw new Error("Missing required property 'usernameField'");
            }
            inputs["accessibilityErrorRedirectUrl"] = args ? args.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = args ? args.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = args ? args.autoSubmitToolbar : undefined;
            inputs["credentialsScheme"] = args ? args.credentialsScheme : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["hideIos"] = args ? args.hideIos : undefined;
            inputs["hideWeb"] = args ? args.hideWeb : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["optionalField1"] = args ? args.optionalField1 : undefined;
            inputs["optionalField1Value"] = args ? args.optionalField1Value : undefined;
            inputs["optionalField2"] = args ? args.optionalField2 : undefined;
            inputs["optionalField2Value"] = args ? args.optionalField2Value : undefined;
            inputs["optionalField3"] = args ? args.optionalField3 : undefined;
            inputs["optionalField3Value"] = args ? args.optionalField3Value : undefined;
            inputs["passwordField"] = args ? args.passwordField : undefined;
            inputs["revealPassword"] = args ? args.revealPassword : undefined;
            inputs["sharedPassword"] = args ? args.sharedPassword : undefined;
            inputs["sharedUsername"] = args ? args.sharedUsername : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["usernameField"] = args ? args.usernameField : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["signOnMode"] = undefined /*out*/;
            inputs["userNameTemplate"] = undefined /*out*/;
            inputs["userNameTemplateType"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecurePasswordStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurePasswordStore resources.
 */
export interface SecurePasswordStoreState {
    /**
     * Custom error page URL.
     */
    readonly accessibilityErrorRedirectUrl?: pulumi.Input<string>;
    /**
     * Enable self service. By default it is `false`.
     */
    readonly accessibilitySelfService?: pulumi.Input<boolean>;
    /**
     * Display auto submit toolbar.
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
    /**
     * Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.
     */
    readonly credentialsScheme?: pulumi.Input<string>;
    /**
     * Groups associated with the application. See `okta.app.GroupAssignment` for a more flexible approach.
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Do not display application icon on mobile app.
     */
    readonly hideIos?: pulumi.Input<boolean>;
    /**
     * Do not display application icon to users.
     */
    readonly hideWeb?: pulumi.Input<boolean>;
    /**
     * The display name of the Application.
     */
    readonly label?: pulumi.Input<string>;
    /**
     * Name assigned to the application by Okta.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form.
     */
    readonly optionalField1?: pulumi.Input<string>;
    /**
     * Name of optional value in the login form.
     */
    readonly optionalField1Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form.
     */
    readonly optionalField2?: pulumi.Input<string>;
    /**
     * Name of optional value in the login form.
     */
    readonly optionalField2Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form.
     */
    readonly optionalField3?: pulumi.Input<string>;
    /**
     * Name of optional value in the login form.
     */
    readonly optionalField3Value?: pulumi.Input<string>;
    /**
     * Login password field.
     */
    readonly passwordField?: pulumi.Input<string>;
    /**
     * Allow user to reveal password.
     */
    readonly revealPassword?: pulumi.Input<boolean>;
    /**
     * Shared password, required for certain schemes.
     */
    readonly sharedPassword?: pulumi.Input<string>;
    /**
     * Shared username, required for certain schemes.
     */
    readonly sharedUsername?: pulumi.Input<string>;
    /**
     * Sign on mode of application.
     */
    readonly signOnMode?: pulumi.Input<string>;
    /**
     * Status of application. By default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * The default username assigned to each user.
     */
    readonly userNameTemplate?: pulumi.Input<string>;
    /**
     * The Username template type.
     */
    readonly userNameTemplateType?: pulumi.Input<string>;
    /**
     * Login username field.
     */
    readonly usernameField?: pulumi.Input<string>;
    /**
     * The users assigned to the application. See `okta.app.User` for a more flexible approach.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.SecurePasswordStoreUser>[]>;
}

/**
 * The set of arguments for constructing a SecurePasswordStore resource.
 */
export interface SecurePasswordStoreArgs {
    /**
     * Custom error page URL.
     */
    readonly accessibilityErrorRedirectUrl?: pulumi.Input<string>;
    /**
     * Enable self service. By default it is `false`.
     */
    readonly accessibilitySelfService?: pulumi.Input<boolean>;
    /**
     * Display auto submit toolbar.
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
    /**
     * Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.
     */
    readonly credentialsScheme?: pulumi.Input<string>;
    /**
     * Groups associated with the application. See `okta.app.GroupAssignment` for a more flexible approach.
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Do not display application icon on mobile app.
     */
    readonly hideIos?: pulumi.Input<boolean>;
    /**
     * Do not display application icon to users.
     */
    readonly hideWeb?: pulumi.Input<boolean>;
    /**
     * The display name of the Application.
     */
    readonly label: pulumi.Input<string>;
    /**
     * Name of optional param in the login form.
     */
    readonly optionalField1?: pulumi.Input<string>;
    /**
     * Name of optional value in the login form.
     */
    readonly optionalField1Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form.
     */
    readonly optionalField2?: pulumi.Input<string>;
    /**
     * Name of optional value in the login form.
     */
    readonly optionalField2Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form.
     */
    readonly optionalField3?: pulumi.Input<string>;
    /**
     * Name of optional value in the login form.
     */
    readonly optionalField3Value?: pulumi.Input<string>;
    /**
     * Login password field.
     */
    readonly passwordField: pulumi.Input<string>;
    /**
     * Allow user to reveal password.
     */
    readonly revealPassword?: pulumi.Input<boolean>;
    /**
     * Shared password, required for certain schemes.
     */
    readonly sharedPassword?: pulumi.Input<string>;
    /**
     * Shared username, required for certain schemes.
     */
    readonly sharedUsername?: pulumi.Input<string>;
    /**
     * Status of application. By default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL.
     */
    readonly url: pulumi.Input<string>;
    /**
     * Login username field.
     */
    readonly usernameField: pulumi.Input<string>;
    /**
     * The users assigned to the application. See `okta.app.User` for a more flexible approach.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.SecurePasswordStoreUser>[]>;
}
