// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an Auto Login Okta Application.
 * 
 * This resource allows you to create and configure an Auto Login Okta Application.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = new okta.app.AutoLogin("example", {
 *     credentialsScheme: "EDIT_USERNAME_AND_PASSWORD",
 *     label: "Example App",
 *     revealPassword: true,
 *     signOnRedirectUrl: "https://example.com",
 *     signOnUrl: "https://example.com/login.html",
 * });
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_auto_login.html.markdown.
 */
export class AutoLogin extends pulumi.CustomResource {
    /**
     * Get an existing AutoLogin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoLoginState, opts?: pulumi.CustomResourceOptions): AutoLogin {
        return new AutoLogin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/autoLogin:AutoLogin';

    /**
     * Returns true if the given object is an instance of AutoLogin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoLogin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoLogin.__pulumiType;
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
     * The Application's display name.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * Name assigned to the application by Okta.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
     */
    public readonly preconfiguredApp!: pulumi.Output<string | undefined>;
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
     * Post login redirect URL
     */
    public readonly signOnRedirectUrl!: pulumi.Output<string | undefined>;
    /**
     * Login URL
     */
    public readonly signOnUrl!: pulumi.Output<string | undefined>;
    /**
     * The status of the application, by default it is `"ACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Username template
     */
    public /*out*/ readonly userNameTemplate!: pulumi.Output<string>;
    /**
     * Username template type
     */
    public /*out*/ readonly userNameTemplateType!: pulumi.Output<string>;
    /**
     * Users associated with the application
     */
    public readonly users!: pulumi.Output<outputs.app.AutoLoginUser[] | undefined>;

    /**
     * Create a AutoLogin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoLoginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoLoginArgs | AutoLoginState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AutoLoginState | undefined;
            inputs["accessibilityErrorRedirectUrl"] = state ? state.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = state ? state.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = state ? state.autoSubmitToolbar : undefined;
            inputs["credentialsScheme"] = state ? state.credentialsScheme : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["hideIos"] = state ? state.hideIos : undefined;
            inputs["hideWeb"] = state ? state.hideWeb : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["preconfiguredApp"] = state ? state.preconfiguredApp : undefined;
            inputs["revealPassword"] = state ? state.revealPassword : undefined;
            inputs["sharedPassword"] = state ? state.sharedPassword : undefined;
            inputs["sharedUsername"] = state ? state.sharedUsername : undefined;
            inputs["signOnMode"] = state ? state.signOnMode : undefined;
            inputs["signOnRedirectUrl"] = state ? state.signOnRedirectUrl : undefined;
            inputs["signOnUrl"] = state ? state.signOnUrl : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["userNameTemplate"] = state ? state.userNameTemplate : undefined;
            inputs["userNameTemplateType"] = state ? state.userNameTemplateType : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as AutoLoginArgs | undefined;
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            inputs["accessibilityErrorRedirectUrl"] = args ? args.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = args ? args.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = args ? args.autoSubmitToolbar : undefined;
            inputs["credentialsScheme"] = args ? args.credentialsScheme : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["hideIos"] = args ? args.hideIos : undefined;
            inputs["hideWeb"] = args ? args.hideWeb : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["preconfiguredApp"] = args ? args.preconfiguredApp : undefined;
            inputs["revealPassword"] = args ? args.revealPassword : undefined;
            inputs["sharedPassword"] = args ? args.sharedPassword : undefined;
            inputs["sharedUsername"] = args ? args.sharedUsername : undefined;
            inputs["signOnRedirectUrl"] = args ? args.signOnRedirectUrl : undefined;
            inputs["signOnUrl"] = args ? args.signOnUrl : undefined;
            inputs["status"] = args ? args.status : undefined;
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
        super(AutoLogin.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoLogin resources.
 */
export interface AutoLoginState {
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
     * The Application's display name.
     */
    readonly label?: pulumi.Input<string>;
    /**
     * Name assigned to the application by Okta.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
     */
    readonly preconfiguredApp?: pulumi.Input<string>;
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
     * Post login redirect URL
     */
    readonly signOnRedirectUrl?: pulumi.Input<string>;
    /**
     * Login URL
     */
    readonly signOnUrl?: pulumi.Input<string>;
    /**
     * The status of the application, by default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Username template
     */
    readonly userNameTemplate?: pulumi.Input<string>;
    /**
     * Username template type
     */
    readonly userNameTemplateType?: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.AutoLoginUser>[]>;
}

/**
 * The set of arguments for constructing a AutoLogin resource.
 */
export interface AutoLoginArgs {
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
     * The Application's display name.
     */
    readonly label: pulumi.Input<string>;
    /**
     * Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
     */
    readonly preconfiguredApp?: pulumi.Input<string>;
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
     * Post login redirect URL
     */
    readonly signOnRedirectUrl?: pulumi.Input<string>;
    /**
     * Login URL
     */
    readonly signOnUrl?: pulumi.Input<string>;
    /**
     * The status of the application, by default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.AutoLoginUser>[]>;
}
