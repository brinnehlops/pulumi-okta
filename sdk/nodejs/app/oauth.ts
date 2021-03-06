// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an OIDC Application.
 *
 * This resource allows you to create and configure an OIDC Application.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.app.OAuth("example", {
 *     grantTypes: ["authorization_code"],
 *     label: "example",
 *     redirectUris: ["https://example.com/"],
 *     responseTypes: ["code"],
 *     type: "web",
 * });
 * ```
 */
export class OAuth extends pulumi.CustomResource {
    /**
     * Get an existing OAuth resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OAuthState, opts?: pulumi.CustomResourceOptions): OAuth {
        return new OAuth(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:app/oAuth:OAuth';

    /**
     * Returns true if the given object is an instance of OAuth.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OAuth {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OAuth.__pulumiType;
    }

    /**
     * Requested key rotation mode.
     */
    public readonly autoKeyRotation!: pulumi.Output<boolean | undefined>;
    /**
     * Display auto submit toolbar.
     */
    public readonly autoSubmitToolbar!: pulumi.Output<boolean | undefined>;
    /**
     * OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
     */
    public readonly clientBasicSecret!: pulumi.Output<string | undefined>;
    /**
     * The client ID of the application.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * The client secret of the application.
     */
    public /*out*/ readonly clientSecret!: pulumi.Output<string>;
    /**
     * URI to a web page providing information about the client.
     */
    public readonly clientUri!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
     */
    public readonly consentMethod!: pulumi.Output<string | undefined>;
    /**
     * This property allows you to set the application's client id.
     */
    public readonly customClientId!: pulumi.Output<string | undefined>;
    /**
     * List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
     */
    public readonly grantTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The groups assigned to the application. It is recommended not to use this and instead use `okta.app.GroupAssignment`.
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
     * Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
     */
    public readonly issuerMode!: pulumi.Output<string | undefined>;
    /**
     * The Application's display name.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * URI that initiates login.
     */
    public readonly loginUri!: pulumi.Output<string | undefined>;
    /**
     * URI that references a logo for the client.
     */
    public readonly logoUri!: pulumi.Output<string | undefined>;
    /**
     * Name assigned to the application by Okta.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
     */
    public readonly omitSecret!: pulumi.Output<boolean | undefined>;
    /**
     * URI to web page providing client policy document.
     */
    public readonly policyUri!: pulumi.Output<string | undefined>;
    /**
     * List of URIs for redirection after logout.
     */
    public readonly postLogoutRedirectUris!: pulumi.Output<string[] | undefined>;
    /**
     * Custom JSON that represents an OAuth application's profile.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * List of URIs for use in the redirect-based flow. This is required for all application types except service.
     */
    public readonly redirectUris!: pulumi.Output<string[] | undefined>;
    /**
     * List of OAuth 2.0 response type strings.
     */
    public readonly responseTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Sign on mode of application.
     */
    public /*out*/ readonly signOnMode!: pulumi.Output<string>;
    /**
     * The status of the application, by default it is `"ACTIVE"`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
     */
    public readonly tokenEndpointAuthMethod!: pulumi.Output<string | undefined>;
    /**
     * URI to web page providing client tos (terms of service).
     */
    public readonly tosUri!: pulumi.Output<string | undefined>;
    /**
     * The type of OAuth application.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The users assigned to the application. It is recommended not to use this and instead use `okta.app.User`.
     */
    public readonly users!: pulumi.Output<outputs.app.OAuthUser[] | undefined>;

    /**
     * Create a OAuth resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OAuthArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OAuthArgs | OAuthState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OAuthState | undefined;
            inputs["autoKeyRotation"] = state ? state.autoKeyRotation : undefined;
            inputs["autoSubmitToolbar"] = state ? state.autoSubmitToolbar : undefined;
            inputs["clientBasicSecret"] = state ? state.clientBasicSecret : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["clientUri"] = state ? state.clientUri : undefined;
            inputs["consentMethod"] = state ? state.consentMethod : undefined;
            inputs["customClientId"] = state ? state.customClientId : undefined;
            inputs["grantTypes"] = state ? state.grantTypes : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["hideIos"] = state ? state.hideIos : undefined;
            inputs["hideWeb"] = state ? state.hideWeb : undefined;
            inputs["issuerMode"] = state ? state.issuerMode : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["loginUri"] = state ? state.loginUri : undefined;
            inputs["logoUri"] = state ? state.logoUri : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["omitSecret"] = state ? state.omitSecret : undefined;
            inputs["policyUri"] = state ? state.policyUri : undefined;
            inputs["postLogoutRedirectUris"] = state ? state.postLogoutRedirectUris : undefined;
            inputs["profile"] = state ? state.profile : undefined;
            inputs["redirectUris"] = state ? state.redirectUris : undefined;
            inputs["responseTypes"] = state ? state.responseTypes : undefined;
            inputs["signOnMode"] = state ? state.signOnMode : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tokenEndpointAuthMethod"] = state ? state.tokenEndpointAuthMethod : undefined;
            inputs["tosUri"] = state ? state.tosUri : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as OAuthArgs | undefined;
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["autoKeyRotation"] = args ? args.autoKeyRotation : undefined;
            inputs["autoSubmitToolbar"] = args ? args.autoSubmitToolbar : undefined;
            inputs["clientBasicSecret"] = args ? args.clientBasicSecret : undefined;
            inputs["clientUri"] = args ? args.clientUri : undefined;
            inputs["consentMethod"] = args ? args.consentMethod : undefined;
            inputs["customClientId"] = args ? args.customClientId : undefined;
            inputs["grantTypes"] = args ? args.grantTypes : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["hideIos"] = args ? args.hideIos : undefined;
            inputs["hideWeb"] = args ? args.hideWeb : undefined;
            inputs["issuerMode"] = args ? args.issuerMode : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["loginUri"] = args ? args.loginUri : undefined;
            inputs["logoUri"] = args ? args.logoUri : undefined;
            inputs["omitSecret"] = args ? args.omitSecret : undefined;
            inputs["policyUri"] = args ? args.policyUri : undefined;
            inputs["postLogoutRedirectUris"] = args ? args.postLogoutRedirectUris : undefined;
            inputs["profile"] = args ? args.profile : undefined;
            inputs["redirectUris"] = args ? args.redirectUris : undefined;
            inputs["responseTypes"] = args ? args.responseTypes : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tokenEndpointAuthMethod"] = args ? args.tokenEndpointAuthMethod : undefined;
            inputs["tosUri"] = args ? args.tosUri : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["clientId"] = undefined /*out*/;
            inputs["clientSecret"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["signOnMode"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OAuth.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OAuth resources.
 */
export interface OAuthState {
    /**
     * Requested key rotation mode.
     */
    readonly autoKeyRotation?: pulumi.Input<boolean>;
    /**
     * Display auto submit toolbar.
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
    /**
     * OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
     */
    readonly clientBasicSecret?: pulumi.Input<string>;
    /**
     * The client ID of the application.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client secret of the application.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * URI to a web page providing information about the client.
     */
    readonly clientUri?: pulumi.Input<string>;
    /**
     * Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
     */
    readonly consentMethod?: pulumi.Input<string>;
    /**
     * This property allows you to set the application's client id.
     */
    readonly customClientId?: pulumi.Input<string>;
    /**
     * List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
     */
    readonly grantTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The groups assigned to the application. It is recommended not to use this and instead use `okta.app.GroupAssignment`.
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
     * Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
     */
    readonly issuerMode?: pulumi.Input<string>;
    /**
     * The Application's display name.
     */
    readonly label?: pulumi.Input<string>;
    /**
     * URI that initiates login.
     */
    readonly loginUri?: pulumi.Input<string>;
    /**
     * URI that references a logo for the client.
     */
    readonly logoUri?: pulumi.Input<string>;
    /**
     * Name assigned to the application by Okta.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
     */
    readonly omitSecret?: pulumi.Input<boolean>;
    /**
     * URI to web page providing client policy document.
     */
    readonly policyUri?: pulumi.Input<string>;
    /**
     * List of URIs for redirection after logout.
     */
    readonly postLogoutRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom JSON that represents an OAuth application's profile.
     */
    readonly profile?: pulumi.Input<string>;
    /**
     * List of URIs for use in the redirect-based flow. This is required for all application types except service.
     */
    readonly redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of OAuth 2.0 response type strings.
     */
    readonly responseTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sign on mode of application.
     */
    readonly signOnMode?: pulumi.Input<string>;
    /**
     * The status of the application, by default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
     */
    readonly tokenEndpointAuthMethod?: pulumi.Input<string>;
    /**
     * URI to web page providing client tos (terms of service).
     */
    readonly tosUri?: pulumi.Input<string>;
    /**
     * The type of OAuth application.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The users assigned to the application. It is recommended not to use this and instead use `okta.app.User`.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.OAuthUser>[]>;
}

/**
 * The set of arguments for constructing a OAuth resource.
 */
export interface OAuthArgs {
    /**
     * Requested key rotation mode.
     */
    readonly autoKeyRotation?: pulumi.Input<boolean>;
    /**
     * Display auto submit toolbar.
     */
    readonly autoSubmitToolbar?: pulumi.Input<boolean>;
    /**
     * OAuth client secret key, this can be set when tokenEndpointAuthMethod is client_secret_basic.
     */
    readonly clientBasicSecret?: pulumi.Input<string>;
    /**
     * URI to a web page providing information about the client.
     */
    readonly clientUri?: pulumi.Input<string>;
    /**
     * Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
     */
    readonly consentMethod?: pulumi.Input<string>;
    /**
     * This property allows you to set the application's client id.
     */
    readonly customClientId?: pulumi.Input<string>;
    /**
     * List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
     */
    readonly grantTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The groups assigned to the application. It is recommended not to use this and instead use `okta.app.GroupAssignment`.
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
     * Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
     */
    readonly issuerMode?: pulumi.Input<string>;
    /**
     * The Application's display name.
     */
    readonly label: pulumi.Input<string>;
    /**
     * URI that initiates login.
     */
    readonly loginUri?: pulumi.Input<string>;
    /**
     * URI that references a logo for the client.
     */
    readonly logoUri?: pulumi.Input<string>;
    /**
     * This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
     */
    readonly omitSecret?: pulumi.Input<boolean>;
    /**
     * URI to web page providing client policy document.
     */
    readonly policyUri?: pulumi.Input<string>;
    /**
     * List of URIs for redirection after logout.
     */
    readonly postLogoutRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom JSON that represents an OAuth application's profile.
     */
    readonly profile?: pulumi.Input<string>;
    /**
     * List of URIs for use in the redirect-based flow. This is required for all application types except service.
     */
    readonly redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of OAuth 2.0 response type strings.
     */
    readonly responseTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the application, by default it is `"ACTIVE"`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Requested authentication method for the token endpoint. It can be set to `"none"`, `"clientSecretPost"`, `"clientSecretBasic"`, `"clientSecretJwt"`.
     */
    readonly tokenEndpointAuthMethod?: pulumi.Input<string>;
    /**
     * URI to web page providing client tos (terms of service).
     */
    readonly tosUri?: pulumi.Input<string>;
    /**
     * The type of OAuth application.
     */
    readonly type: pulumi.Input<string>;
    /**
     * The users assigned to the application. It is recommended not to use this and instead use `okta.app.User`.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.app.OAuthUser>[]>;
}
