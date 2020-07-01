// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Social Identity Provider.
 *
 * This resource allows you to create and configure an Social Identity Provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as okta from "@pulumi/okta";
 *
 * const example = new okta.idp.Social("example", {
 *     clientId: "abcd123",
 *     clientSecret: "abcd123",
 *     matchAttribute: "customfieldId",
 *     matchType: "CUSTOM_ATTRIBUTE",
 *     protocolType: "OAUTH2",
 *     scopes: [
 *         "public_profile",
 *         "email",
 *     ],
 *     type: "FACEBOOK",
 *     usernameTemplate: "idpuser.email",
 * });
 * ```
 */
export class Social extends pulumi.CustomResource {
    /**
     * Get an existing Social resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SocialState, opts?: pulumi.CustomResourceOptions): Social {
        return new Social(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'okta:idp/social:Social';

    /**
     * Returns true if the given object is an instance of Social.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Social {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Social.__pulumiType;
    }

    /**
     * Specifies the account linking action for an IdP user.
     */
    public readonly accountLinkAction!: pulumi.Output<string | undefined>;
    /**
     * Group memberships to determine link candidates.
     */
    public readonly accountLinkGroupIncludes!: pulumi.Output<string[] | undefined>;
    /**
     * The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
     */
    public /*out*/ readonly authorizationBinding!: pulumi.Output<string>;
    /**
     * IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
     */
    public /*out*/ readonly authorizationUrl!: pulumi.Output<string>;
    /**
     * Unique identifier issued by AS for the Okta IdP instance.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * Client secret issued by AS for the Okta IdP instance.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
     */
    public readonly deprovisionedAction!: pulumi.Output<string | undefined>;
    /**
     * Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
     */
    public readonly groupsAction!: pulumi.Output<string | undefined>;
    /**
     * List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
     */
    public readonly groupsAssignments!: pulumi.Output<string[] | undefined>;
    /**
     * IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
     */
    public readonly groupsAttribute!: pulumi.Output<string | undefined>;
    /**
     * Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
     */
    public readonly groupsFilters!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
     */
    public readonly issuerMode!: pulumi.Output<string | undefined>;
    /**
     * @deprecated This property was incorrectly added to this resource, you should use "subject_match_attribute"
     */
    public readonly matchAttribute!: pulumi.Output<string | undefined>;
    /**
     * @deprecated This property was incorrectly added to this resource, you should use "subject_match_type"
     */
    public readonly matchType!: pulumi.Output<string | undefined>;
    /**
     * Maximum allowable clock-skew when processing messages from the IdP.
     */
    public readonly maxClockSkew!: pulumi.Output<number | undefined>;
    /**
     * The Application's display name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Determines if the IdP should act as a source of truth for user profile attributes.
     */
    public readonly profileMaster!: pulumi.Output<boolean | undefined>;
    /**
     * The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
     */
    public readonly protocolType!: pulumi.Output<string | undefined>;
    /**
     * Provisioning action for an IdP user during authentication.
     */
    public readonly provisioningAction!: pulumi.Output<string | undefined>;
    /**
     * The XML digital signature algorithm used when signing an AuthnRequest message.
     */
    public readonly requestSignatureAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
     */
    public readonly requestSignatureScope!: pulumi.Output<string | undefined>;
    /**
     * The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
     */
    public readonly responseSignatureAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
     */
    public readonly responseSignatureScope!: pulumi.Output<string | undefined>;
    /**
     * The scopes of the IdP.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * Status of the IdP.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
     */
    public readonly subjectMatchAttribute!: pulumi.Output<string | undefined>;
    /**
     * Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
     */
    public readonly subjectMatchType!: pulumi.Output<string | undefined>;
    /**
     * Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
     */
    public readonly suspendedAction!: pulumi.Output<string | undefined>;
    /**
     * The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
     */
    public /*out*/ readonly tokenBinding!: pulumi.Output<string>;
    /**
     * IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
     */
    public /*out*/ readonly tokenUrl!: pulumi.Output<string>;
    /**
     * The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Okta EL Expression to generate or transform a unique username for the IdP user.
     */
    public readonly usernameTemplate!: pulumi.Output<string | undefined>;

    /**
     * Create a Social resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SocialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SocialArgs | SocialState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SocialState | undefined;
            inputs["accountLinkAction"] = state ? state.accountLinkAction : undefined;
            inputs["accountLinkGroupIncludes"] = state ? state.accountLinkGroupIncludes : undefined;
            inputs["authorizationBinding"] = state ? state.authorizationBinding : undefined;
            inputs["authorizationUrl"] = state ? state.authorizationUrl : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["deprovisionedAction"] = state ? state.deprovisionedAction : undefined;
            inputs["groupsAction"] = state ? state.groupsAction : undefined;
            inputs["groupsAssignments"] = state ? state.groupsAssignments : undefined;
            inputs["groupsAttribute"] = state ? state.groupsAttribute : undefined;
            inputs["groupsFilters"] = state ? state.groupsFilters : undefined;
            inputs["issuerMode"] = state ? state.issuerMode : undefined;
            inputs["matchAttribute"] = state ? state.matchAttribute : undefined;
            inputs["matchType"] = state ? state.matchType : undefined;
            inputs["maxClockSkew"] = state ? state.maxClockSkew : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["profileMaster"] = state ? state.profileMaster : undefined;
            inputs["protocolType"] = state ? state.protocolType : undefined;
            inputs["provisioningAction"] = state ? state.provisioningAction : undefined;
            inputs["requestSignatureAlgorithm"] = state ? state.requestSignatureAlgorithm : undefined;
            inputs["requestSignatureScope"] = state ? state.requestSignatureScope : undefined;
            inputs["responseSignatureAlgorithm"] = state ? state.responseSignatureAlgorithm : undefined;
            inputs["responseSignatureScope"] = state ? state.responseSignatureScope : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subjectMatchAttribute"] = state ? state.subjectMatchAttribute : undefined;
            inputs["subjectMatchType"] = state ? state.subjectMatchType : undefined;
            inputs["suspendedAction"] = state ? state.suspendedAction : undefined;
            inputs["tokenBinding"] = state ? state.tokenBinding : undefined;
            inputs["tokenUrl"] = state ? state.tokenUrl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["usernameTemplate"] = state ? state.usernameTemplate : undefined;
        } else {
            const args = argsOrState as SocialArgs | undefined;
            if (!args || args.scopes === undefined) {
                throw new Error("Missing required property 'scopes'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["accountLinkAction"] = args ? args.accountLinkAction : undefined;
            inputs["accountLinkGroupIncludes"] = args ? args.accountLinkGroupIncludes : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["deprovisionedAction"] = args ? args.deprovisionedAction : undefined;
            inputs["groupsAction"] = args ? args.groupsAction : undefined;
            inputs["groupsAssignments"] = args ? args.groupsAssignments : undefined;
            inputs["groupsAttribute"] = args ? args.groupsAttribute : undefined;
            inputs["groupsFilters"] = args ? args.groupsFilters : undefined;
            inputs["issuerMode"] = args ? args.issuerMode : undefined;
            inputs["matchAttribute"] = args ? args.matchAttribute : undefined;
            inputs["matchType"] = args ? args.matchType : undefined;
            inputs["maxClockSkew"] = args ? args.maxClockSkew : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["profileMaster"] = args ? args.profileMaster : undefined;
            inputs["protocolType"] = args ? args.protocolType : undefined;
            inputs["provisioningAction"] = args ? args.provisioningAction : undefined;
            inputs["requestSignatureAlgorithm"] = args ? args.requestSignatureAlgorithm : undefined;
            inputs["requestSignatureScope"] = args ? args.requestSignatureScope : undefined;
            inputs["responseSignatureAlgorithm"] = args ? args.responseSignatureAlgorithm : undefined;
            inputs["responseSignatureScope"] = args ? args.responseSignatureScope : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["subjectMatchAttribute"] = args ? args.subjectMatchAttribute : undefined;
            inputs["subjectMatchType"] = args ? args.subjectMatchType : undefined;
            inputs["suspendedAction"] = args ? args.suspendedAction : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["usernameTemplate"] = args ? args.usernameTemplate : undefined;
            inputs["authorizationBinding"] = undefined /*out*/;
            inputs["authorizationUrl"] = undefined /*out*/;
            inputs["tokenBinding"] = undefined /*out*/;
            inputs["tokenUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Social.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Social resources.
 */
export interface SocialState {
    /**
     * Specifies the account linking action for an IdP user.
     */
    readonly accountLinkAction?: pulumi.Input<string>;
    /**
     * Group memberships to determine link candidates.
     */
    readonly accountLinkGroupIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
     */
    readonly authorizationBinding?: pulumi.Input<string>;
    /**
     * IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
     */
    readonly authorizationUrl?: pulumi.Input<string>;
    /**
     * Unique identifier issued by AS for the Okta IdP instance.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client secret issued by AS for the Okta IdP instance.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
     */
    readonly deprovisionedAction?: pulumi.Input<string>;
    /**
     * Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
     */
    readonly groupsAction?: pulumi.Input<string>;
    /**
     * List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
     */
    readonly groupsAssignments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
     */
    readonly groupsAttribute?: pulumi.Input<string>;
    /**
     * Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
     */
    readonly groupsFilters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
     */
    readonly issuerMode?: pulumi.Input<string>;
    /**
     * @deprecated This property was incorrectly added to this resource, you should use "subject_match_attribute"
     */
    readonly matchAttribute?: pulumi.Input<string>;
    /**
     * @deprecated This property was incorrectly added to this resource, you should use "subject_match_type"
     */
    readonly matchType?: pulumi.Input<string>;
    /**
     * Maximum allowable clock-skew when processing messages from the IdP.
     */
    readonly maxClockSkew?: pulumi.Input<number>;
    /**
     * The Application's display name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines if the IdP should act as a source of truth for user profile attributes.
     */
    readonly profileMaster?: pulumi.Input<boolean>;
    /**
     * The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
     */
    readonly protocolType?: pulumi.Input<string>;
    /**
     * Provisioning action for an IdP user during authentication.
     */
    readonly provisioningAction?: pulumi.Input<string>;
    /**
     * The XML digital signature algorithm used when signing an AuthnRequest message.
     */
    readonly requestSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
     */
    readonly requestSignatureScope?: pulumi.Input<string>;
    /**
     * The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
     */
    readonly responseSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
     */
    readonly responseSignatureScope?: pulumi.Input<string>;
    /**
     * The scopes of the IdP.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the IdP.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
     */
    readonly subjectMatchAttribute?: pulumi.Input<string>;
    /**
     * Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
     */
    readonly subjectMatchType?: pulumi.Input<string>;
    /**
     * Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
     */
    readonly suspendedAction?: pulumi.Input<string>;
    /**
     * The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
     */
    readonly tokenBinding?: pulumi.Input<string>;
    /**
     * IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
     */
    readonly tokenUrl?: pulumi.Input<string>;
    /**
     * The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Okta EL Expression to generate or transform a unique username for the IdP user.
     */
    readonly usernameTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Social resource.
 */
export interface SocialArgs {
    /**
     * Specifies the account linking action for an IdP user.
     */
    readonly accountLinkAction?: pulumi.Input<string>;
    /**
     * Group memberships to determine link candidates.
     */
    readonly accountLinkGroupIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unique identifier issued by AS for the Okta IdP instance.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client secret issued by AS for the Okta IdP instance.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
     */
    readonly deprovisionedAction?: pulumi.Input<string>;
    /**
     * Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
     */
    readonly groupsAction?: pulumi.Input<string>;
    /**
     * List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groupsAction`.
     */
    readonly groupsAssignments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
     */
    readonly groupsAttribute?: pulumi.Input<string>;
    /**
     * Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groupsAction`.
     */
    readonly groupsFilters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
     */
    readonly issuerMode?: pulumi.Input<string>;
    /**
     * @deprecated This property was incorrectly added to this resource, you should use "subject_match_attribute"
     */
    readonly matchAttribute?: pulumi.Input<string>;
    /**
     * @deprecated This property was incorrectly added to this resource, you should use "subject_match_type"
     */
    readonly matchType?: pulumi.Input<string>;
    /**
     * Maximum allowable clock-skew when processing messages from the IdP.
     */
    readonly maxClockSkew?: pulumi.Input<number>;
    /**
     * The Application's display name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines if the IdP should act as a source of truth for user profile attributes.
     */
    readonly profileMaster?: pulumi.Input<boolean>;
    /**
     * The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
     */
    readonly protocolType?: pulumi.Input<string>;
    /**
     * Provisioning action for an IdP user during authentication.
     */
    readonly provisioningAction?: pulumi.Input<string>;
    /**
     * The XML digital signature algorithm used when signing an AuthnRequest message.
     */
    readonly requestSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
     */
    readonly requestSignatureScope?: pulumi.Input<string>;
    /**
     * The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
     */
    readonly responseSignatureAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
     */
    readonly responseSignatureScope?: pulumi.Input<string>;
    /**
     * The scopes of the IdP.
     */
    readonly scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the IdP.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
     */
    readonly subjectMatchAttribute?: pulumi.Input<string>;
    /**
     * Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
     */
    readonly subjectMatchType?: pulumi.Input<string>;
    /**
     * Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
     */
    readonly suspendedAction?: pulumi.Input<string>;
    /**
     * The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
     */
    readonly type: pulumi.Input<string>;
    /**
     * Okta EL Expression to generate or transform a unique username for the IdP user.
     */
    readonly usernameTemplate?: pulumi.Input<string>;
}
