// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an SWA Application.
 * 
 * This resource allows you to create and configure an SWA Application.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 * 
 * const example = new okta.app.Swa("example", {
 *     buttonField: "btn-login",
 *     label: "example",
 *     passwordField: "txtbox-password",
 *     url: "https://example.com/login.html",
 *     usernameField: "txtbox-username",
 * });
 * ```
 *
 * > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_swa.html.markdown.
 */
export class Swa extends pulumi.CustomResource {
    /**
     * Get an existing Swa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwaState, opts?: pulumi.CustomResourceOptions): Swa {
        return new Swa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/swa:Swa';

    /**
     * Returns true if the given object is an instance of Swa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Swa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Swa.__pulumiType;
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
     * Login button field.
     */
    public readonly buttonField!: pulumi.Output<string | undefined>;
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
     * Login password field.
     */
    public readonly passwordField!: pulumi.Output<string | undefined>;
    /**
     * name of application from the Okta Integration Network, if not included a custom app will be created.
     */
    public readonly preconfiguredApp!: pulumi.Output<string | undefined>;
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
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * A regex that further restricts URL to the specified regex.
     */
    public readonly urlRegex!: pulumi.Output<string | undefined>;
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
    public readonly usernameField!: pulumi.Output<string | undefined>;
    /**
     * The users assigned to the application. See `okta.app.User` for a more flexible approach.
     */
    public readonly users!: pulumi.Output<outputs.app.SwaUser[] | undefined>;

    /**
     * Create a Swa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwaArgs | SwaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SwaState | undefined;
            inputs["accessibilityErrorRedirectUrl"] = state ? state.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = state ? state.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = state ? state.autoSubmitToolbar : undefined;
            inputs["buttonField"] = state ? state.buttonField : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["hideIos"] = state ? state.hideIos : undefined;
            inputs["hideWeb"] = state ? state.hideWeb : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["passwordField"] = state ? state.passwordField : undefined;
            inputs["preconfiguredApp"] = state ? state.preconfiguredApp : undefined;
            inputs["signOnMode"] = state ? state.signOnMode : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["urlRegex"] = state ? state.urlRegex : undefined;
            inputs["userNameTemplate"] = state ? state.userNameTemplate : undefined;
            inputs["userNameTemplateType"] = state ? state.userNameTemplateType : undefined;
            inputs["usernameField"] = state ? state.usernameField : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as SwaArgs | undefined;
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            inputs["accessibilityErrorRedirectUrl"] = args ? args.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = args ? args.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = args ? args.autoSubmitToolbar : undefined;
            inputs["buttonField"] = args ? args.buttonField : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["hideIos"] = args ? args.hideIos : undefined;
            inputs["hideWeb"] = args ? args.hideWeb : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["passwordField"] = args ? args.passwordField : undefined;
            inputs["preconfiguredApp"] = args ? args.preconfiguredApp : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["urlRegex"] = args ? args.urlRegex : undefined;
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
        super(Swa.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Swa resources.
 */
export interface SwaState {
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
     * Login button field.
     */
    readonly buttonField?: pulumi.Input<string>;
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
     * Login password field.
     */
    readonly passwordField?: pulumi.Input<string>;
    /**
     * name of application from the Okta Integration Network, if not included a custom app will be created.
     */
    readonly preconfiguredApp?: pulumi.Input<string>;
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
     * A regex that further restricts URL to the specified regex.
     */
    readonly urlRegex?: pulumi.Input<string>;
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
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.SwaUser>[]>;
}

/**
 * The set of arguments for constructing a Swa resource.
 */
export interface SwaArgs {
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
     * Login button field.
     */
    readonly buttonField?: pulumi.Input<string>;
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
     * Login password field.
     */
    readonly passwordField?: pulumi.Input<string>;
    /**
     * name of application from the Okta Integration Network, if not included a custom app will be created.
     */
    readonly preconfiguredApp?: pulumi.Input<string>;
    /**
     * Status of application. By default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * A regex that further restricts URL to the specified regex.
     */
    readonly urlRegex?: pulumi.Input<string>;
    /**
     * Login username field.
     */
    readonly usernameField?: pulumi.Input<string>;
    /**
     * The users assigned to the application. See `okta.app.User` for a more flexible approach.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.SwaUser>[]>;
}
