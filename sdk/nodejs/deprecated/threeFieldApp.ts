// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ThreeFieldApp extends pulumi.CustomResource {
    /**
     * Get an existing ThreeFieldApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ThreeFieldAppState, opts?: pulumi.CustomResourceOptions): ThreeFieldApp {
        return new ThreeFieldApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:deprecated/threeFieldApp:ThreeFieldApp';

    /**
     * Returns true if the given object is an instance of ThreeFieldApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ThreeFieldApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ThreeFieldApp.__pulumiType;
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
     * Login button field CSS selector
     */
    public readonly buttonSelector!: pulumi.Output<string>;
    /**
     * Extra field CSS selector
     */
    public readonly extraFieldSelector!: pulumi.Output<string>;
    /**
     * Value for extra form field
     */
    public readonly extraFieldValue!: pulumi.Output<string>;
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
     * Login password field CSS selector
     */
    public readonly passwordSelector!: pulumi.Output<string>;
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
     * Login username field CSS selector
     */
    public readonly usernameSelector!: pulumi.Output<string>;
    /**
     * Users associated with the application
     */
    public readonly users!: pulumi.Output<outputs.deprecated.ThreeFieldAppUser[] | undefined>;

    /**
     * Create a ThreeFieldApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ThreeFieldAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ThreeFieldAppArgs | ThreeFieldAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ThreeFieldAppState | undefined;
            inputs["accessibilityErrorRedirectUrl"] = state ? state.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = state ? state.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = state ? state.autoSubmitToolbar : undefined;
            inputs["buttonSelector"] = state ? state.buttonSelector : undefined;
            inputs["extraFieldSelector"] = state ? state.extraFieldSelector : undefined;
            inputs["extraFieldValue"] = state ? state.extraFieldValue : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["hideIos"] = state ? state.hideIos : undefined;
            inputs["hideWeb"] = state ? state.hideWeb : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["passwordSelector"] = state ? state.passwordSelector : undefined;
            inputs["signOnMode"] = state ? state.signOnMode : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["urlRegex"] = state ? state.urlRegex : undefined;
            inputs["userNameTemplate"] = state ? state.userNameTemplate : undefined;
            inputs["userNameTemplateType"] = state ? state.userNameTemplateType : undefined;
            inputs["usernameSelector"] = state ? state.usernameSelector : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as ThreeFieldAppArgs | undefined;
            if (!args || args.buttonSelector === undefined) {
                throw new Error("Missing required property 'buttonSelector'");
            }
            if (!args || args.extraFieldSelector === undefined) {
                throw new Error("Missing required property 'extraFieldSelector'");
            }
            if (!args || args.extraFieldValue === undefined) {
                throw new Error("Missing required property 'extraFieldValue'");
            }
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            if (!args || args.passwordSelector === undefined) {
                throw new Error("Missing required property 'passwordSelector'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            if (!args || args.usernameSelector === undefined) {
                throw new Error("Missing required property 'usernameSelector'");
            }
            inputs["accessibilityErrorRedirectUrl"] = args ? args.accessibilityErrorRedirectUrl : undefined;
            inputs["accessibilitySelfService"] = args ? args.accessibilitySelfService : undefined;
            inputs["autoSubmitToolbar"] = args ? args.autoSubmitToolbar : undefined;
            inputs["buttonSelector"] = args ? args.buttonSelector : undefined;
            inputs["extraFieldSelector"] = args ? args.extraFieldSelector : undefined;
            inputs["extraFieldValue"] = args ? args.extraFieldValue : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["hideIos"] = args ? args.hideIos : undefined;
            inputs["hideWeb"] = args ? args.hideWeb : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["passwordSelector"] = args ? args.passwordSelector : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["urlRegex"] = args ? args.urlRegex : undefined;
            inputs["usernameSelector"] = args ? args.usernameSelector : undefined;
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
        super(ThreeFieldApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ThreeFieldApp resources.
 */
export interface ThreeFieldAppState {
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
     * Login button field CSS selector
     */
    readonly buttonSelector?: pulumi.Input<string>;
    /**
     * Extra field CSS selector
     */
    readonly extraFieldSelector?: pulumi.Input<string>;
    /**
     * Value for extra form field
     */
    readonly extraFieldValue?: pulumi.Input<string>;
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
     * Login password field CSS selector
     */
    readonly passwordSelector?: pulumi.Input<string>;
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
     * Login username field CSS selector
     */
    readonly usernameSelector?: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.ThreeFieldAppUser>[]>;
}

/**
 * The set of arguments for constructing a ThreeFieldApp resource.
 */
export interface ThreeFieldAppArgs {
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
     * Login button field CSS selector
     */
    readonly buttonSelector: pulumi.Input<string>;
    /**
     * Extra field CSS selector
     */
    readonly extraFieldSelector: pulumi.Input<string>;
    /**
     * Value for extra form field
     */
    readonly extraFieldValue: pulumi.Input<string>;
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
     * Login password field CSS selector
     */
    readonly passwordSelector: pulumi.Input<string>;
    /**
     * Status of application.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Login URL
     */
    readonly url: pulumi.Input<string>;
    /**
     * A regex that further restricts URL to the specified regex
     */
    readonly urlRegex?: pulumi.Input<string>;
    /**
     * Login username field CSS selector
     */
    readonly usernameSelector: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.ThreeFieldAppUser>[]>;
}
