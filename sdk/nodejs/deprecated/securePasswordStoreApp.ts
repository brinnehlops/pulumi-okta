// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SecurePasswordStoreApp extends pulumi.CustomResource {
    /**
     * Get an existing SecurePasswordStoreApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurePasswordStoreAppState, opts?: pulumi.CustomResourceOptions): SecurePasswordStoreApp {
        return new SecurePasswordStoreApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:deprecated/securePasswordStoreApp:SecurePasswordStoreApp';

    /**
     * Returns true if the given object is an instance of SecurePasswordStoreApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurePasswordStoreApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurePasswordStoreApp.__pulumiType;
    }

    /**
     * Custom error page URL
     */
    public readonly accessibilityErrorRedirectUrl!: pulumi.Output<string | undefined>;
    /**
     * Enable self service
     */
    public readonly accessibilitySelfService!: pulumi.Output<boolean | undefined>;
    /**
     * Display auto submit toolbar
     */
    public readonly autoSubmitToolbar!: pulumi.Output<boolean | undefined>;
    /**
     * Application credentials scheme
     */
    public readonly credentialsScheme!: pulumi.Output<string | undefined>;
    /**
     * Groups associated with the application
     */
    public readonly groups!: pulumi.Output<string[] | undefined>;
    /**
     * Do not display application icon on mobile app
     */
    public readonly hideIos!: pulumi.Output<boolean | undefined>;
    /**
     * Do not display application icon to users
     */
    public readonly hideWeb!: pulumi.Output<boolean | undefined>;
    /**
     * Pretty name of app.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * name of app.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Name of optional param in the login form
     */
    public readonly optionalField1!: pulumi.Output<string | undefined>;
    /**
     * Name of optional value in login form
     */
    public readonly optionalField1Value!: pulumi.Output<string | undefined>;
    /**
     * Name of optional param in the login form
     */
    public readonly optionalField2!: pulumi.Output<string | undefined>;
    /**
     * Name of optional value in login form
     */
    public readonly optionalField2Value!: pulumi.Output<string | undefined>;
    /**
     * Name of optional param in the login form
     */
    public readonly optionalField3!: pulumi.Output<string | undefined>;
    /**
     * Name of optional value in login form
     */
    public readonly optionalField3Value!: pulumi.Output<string | undefined>;
    /**
     * Login password field
     */
    public readonly passwordField!: pulumi.Output<string>;
    /**
     * Allow user to reveal password
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
     * Status of application.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Login URL
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Username template
     */
    public /*out*/ readonly userNameTemplate!: pulumi.Output<string>;
    /**
     * Username template type
     */
    public /*out*/ readonly userNameTemplateType!: pulumi.Output<string>;
    /**
     * Login username field
     */
    public readonly usernameField!: pulumi.Output<string>;
    /**
     * Users associated with the application
     */
    public readonly users!: pulumi.Output<outputs.deprecated.SecurePasswordStoreAppUser[] | undefined>;

    /**
     * Create a SecurePasswordStoreApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurePasswordStoreAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurePasswordStoreAppArgs | SecurePasswordStoreAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecurePasswordStoreAppState | undefined;
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
            const args = argsOrState as SecurePasswordStoreAppArgs | undefined;
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
        super(SecurePasswordStoreApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurePasswordStoreApp resources.
 */
export interface SecurePasswordStoreAppState {
    /**
     * Custom error page URL
     */
    readonly accessibilityErrorRedirectUrl?: pulumi.Input<string>;
    /**
     * Enable self service
     */
    readonly accessibilitySelfService?: pulumi.Input<boolean>;
    /**
     * Display auto submit toolbar
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
    /**
     * Application credentials scheme
     */
    readonly credentialsScheme?: pulumi.Input<string>;
    /**
     * Groups associated with the application
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Do not display application icon on mobile app
     */
    readonly hideIos?: pulumi.Input<boolean>;
    /**
     * Do not display application icon to users
     */
    readonly hideWeb?: pulumi.Input<boolean>;
    /**
     * Pretty name of app.
     */
    readonly label?: pulumi.Input<string>;
    /**
     * name of app.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form
     */
    readonly optionalField1?: pulumi.Input<string>;
    /**
     * Name of optional value in login form
     */
    readonly optionalField1Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form
     */
    readonly optionalField2?: pulumi.Input<string>;
    /**
     * Name of optional value in login form
     */
    readonly optionalField2Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form
     */
    readonly optionalField3?: pulumi.Input<string>;
    /**
     * Name of optional value in login form
     */
    readonly optionalField3Value?: pulumi.Input<string>;
    /**
     * Login password field
     */
    readonly passwordField?: pulumi.Input<string>;
    /**
     * Allow user to reveal password
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
     * Status of application.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Username template
     */
    readonly userNameTemplate?: pulumi.Input<string>;
    /**
     * Username template type
     */
    readonly userNameTemplateType?: pulumi.Input<string>;
    /**
     * Login username field
     */
    readonly usernameField?: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.SecurePasswordStoreAppUser>[]>;
}

/**
 * The set of arguments for constructing a SecurePasswordStoreApp resource.
 */
export interface SecurePasswordStoreAppArgs {
    /**
     * Custom error page URL
     */
    readonly accessibilityErrorRedirectUrl?: pulumi.Input<string>;
    /**
     * Enable self service
     */
    readonly accessibilitySelfService?: pulumi.Input<boolean>;
    /**
     * Display auto submit toolbar
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
    /**
     * Application credentials scheme
     */
    readonly credentialsScheme?: pulumi.Input<string>;
    /**
     * Groups associated with the application
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Do not display application icon on mobile app
     */
    readonly hideIos?: pulumi.Input<boolean>;
    /**
     * Do not display application icon to users
     */
    readonly hideWeb?: pulumi.Input<boolean>;
    /**
     * Pretty name of app.
     */
    readonly label: pulumi.Input<string>;
    /**
     * Name of optional param in the login form
     */
    readonly optionalField1?: pulumi.Input<string>;
    /**
     * Name of optional value in login form
     */
    readonly optionalField1Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form
     */
    readonly optionalField2?: pulumi.Input<string>;
    /**
     * Name of optional value in login form
     */
    readonly optionalField2Value?: pulumi.Input<string>;
    /**
     * Name of optional param in the login form
     */
    readonly optionalField3?: pulumi.Input<string>;
    /**
     * Name of optional value in login form
     */
    readonly optionalField3Value?: pulumi.Input<string>;
    /**
     * Login password field
     */
    readonly passwordField: pulumi.Input<string>;
    /**
     * Allow user to reveal password
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
     * Status of application.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL
     */
    readonly url: pulumi.Input<string>;
    /**
     * Login username field
     */
    readonly usernameField: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.SecurePasswordStoreAppUser>[]>;
}
