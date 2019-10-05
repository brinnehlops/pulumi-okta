// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SwaApp extends pulumi.CustomResource {
    /**
     * Get an existing SwaApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwaAppState, opts?: pulumi.CustomResourceOptions): SwaApp {
        return new SwaApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:deprecated/swaApp:SwaApp';

    /**
     * Returns true if the given object is an instance of SwaApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwaApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwaApp.__pulumiType;
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
     * Login button field
     */
    public readonly buttonField!: pulumi.Output<string | undefined>;
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
     * Login password field
     */
    public readonly passwordField!: pulumi.Output<string | undefined>;
    /**
     * Preconfigured app name
     */
    public readonly preconfiguredApp!: pulumi.Output<string | undefined>;
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
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * A regex that further restricts URL to the specified regex
     */
    public readonly urlRegex!: pulumi.Output<string | undefined>;
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
    public readonly usernameField!: pulumi.Output<string | undefined>;
    /**
     * Users associated with the application
     */
    public readonly users!: pulumi.Output<outputs.deprecated.SwaAppUser[] | undefined>;

    /**
     * Create a SwaApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwaAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwaAppArgs | SwaAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SwaAppState | undefined;
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
            const args = argsOrState as SwaAppArgs | undefined;
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
        super(SwaApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwaApp resources.
 */
export interface SwaAppState {
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
     * Login button field
     */
    readonly buttonField?: pulumi.Input<string>;
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
     * Login password field
     */
    readonly passwordField?: pulumi.Input<string>;
    /**
     * Preconfigured app name
     */
    readonly preconfiguredApp?: pulumi.Input<string>;
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
     * A regex that further restricts URL to the specified regex
     */
    readonly urlRegex?: pulumi.Input<string>;
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
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.SwaAppUser>[]>;
}

/**
 * The set of arguments for constructing a SwaApp resource.
 */
export interface SwaAppArgs {
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
     * Login button field
     */
    readonly buttonField?: pulumi.Input<string>;
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
     * Login password field
     */
    readonly passwordField?: pulumi.Input<string>;
    /**
     * Preconfigured app name
     */
    readonly preconfiguredApp?: pulumi.Input<string>;
    /**
     * Status of application.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL
     */
    readonly url?: pulumi.Input<string>;
    /**
     * A regex that further restricts URL to the specified regex
     */
    readonly urlRegex?: pulumi.Input<string>;
    /**
     * Login username field
     */
    readonly usernameField?: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.SwaAppUser>[]>;
}
