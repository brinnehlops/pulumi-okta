// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class BookmarkApp extends pulumi.CustomResource {
    /**
     * Get an existing BookmarkApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BookmarkAppState, opts?: pulumi.CustomResourceOptions): BookmarkApp {
        return new BookmarkApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:deprecated/bookmarkApp:BookmarkApp';

    /**
     * Returns true if the given object is an instance of BookmarkApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BookmarkApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BookmarkApp.__pulumiType;
    }

    /**
     * Display auto submit toolbar
     */
    public readonly autoSubmitToolbar!: pulumi.Output<boolean | undefined>;
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
    public readonly requestIntegration!: pulumi.Output<boolean | undefined>;
    /**
     * Sign on mode of application.
     */
    public /*out*/ readonly signOnMode!: pulumi.Output<string>;
    /**
     * Status of application.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly url!: pulumi.Output<string>;
    /**
     * Users associated with the application
     */
    public readonly users!: pulumi.Output<outputs.deprecated.BookmarkAppUser[] | undefined>;

    /**
     * Create a BookmarkApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BookmarkAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BookmarkAppArgs | BookmarkAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BookmarkAppState | undefined;
            inputs["autoSubmitToolbar"] = state ? state.autoSubmitToolbar : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["hideIos"] = state ? state.hideIos : undefined;
            inputs["hideWeb"] = state ? state.hideWeb : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["requestIntegration"] = state ? state.requestIntegration : undefined;
            inputs["signOnMode"] = state ? state.signOnMode : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as BookmarkAppArgs | undefined;
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["autoSubmitToolbar"] = args ? args.autoSubmitToolbar : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["hideIos"] = args ? args.hideIos : undefined;
            inputs["hideWeb"] = args ? args.hideWeb : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["requestIntegration"] = args ? args.requestIntegration : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["signOnMode"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BookmarkApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BookmarkApp resources.
 */
export interface BookmarkAppState {
    /**
     * Display auto submit toolbar
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
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
    readonly requestIntegration?: pulumi.Input<boolean>;
    /**
     * Sign on mode of application.
     */
    readonly signOnMode?: pulumi.Input<string>;
    /**
     * Status of application.
     */
    readonly status?: pulumi.Input<string>;
    readonly url?: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.BookmarkAppUser>[]>;
}

/**
 * The set of arguments for constructing a BookmarkApp resource.
 */
export interface BookmarkAppArgs {
    /**
     * Display auto submit toolbar
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
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
    readonly requestIntegration?: pulumi.Input<boolean>;
    /**
     * Status of application.
     */
    readonly status?: pulumi.Input<string>;
    readonly url: pulumi.Input<string>;
    /**
     * Users associated with the application
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.deprecated.BookmarkAppUser>[]>;
}
