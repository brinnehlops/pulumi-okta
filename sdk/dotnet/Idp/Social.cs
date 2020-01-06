// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Idp
{
    /// <summary>
    /// Creates an Social Identity Provider.
    /// 
    /// This resource allows you to create and configure an Social Identity Provider.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/idp_social.html.markdown.
    /// </summary>
    public partial class Social : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the account linking action for an IdP user.
        /// </summary>
        [Output("accountLinkAction")]
        public Output<string?> AccountLinkAction { get; private set; } = null!;

        /// <summary>
        /// Group memberships to determine link candidates.
        /// </summary>
        [Output("accountLinkGroupIncludes")]
        public Output<ImmutableArray<string>> AccountLinkGroupIncludes { get; private set; } = null!;

        /// <summary>
        /// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        /// </summary>
        [Output("authorizationBinding")]
        public Output<string> AuthorizationBinding { get; private set; } = null!;

        /// <summary>
        /// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
        /// </summary>
        [Output("authorizationUrl")]
        public Output<string> AuthorizationUrl { get; private set; } = null!;

        /// <summary>
        /// Unique identifier issued by AS for the Okta IdP instance.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret issued by AS for the Okta IdP instance.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        /// </summary>
        [Output("deprovisionedAction")]
        public Output<string?> DeprovisionedAction { get; private set; } = null!;

        /// <summary>
        /// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        /// </summary>
        [Output("groupsAction")]
        public Output<string?> GroupsAction { get; private set; } = null!;

        /// <summary>
        /// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        /// </summary>
        [Output("groupsAssignments")]
        public Output<ImmutableArray<string>> GroupsAssignments { get; private set; } = null!;

        /// <summary>
        /// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        /// </summary>
        [Output("groupsAttribute")]
        public Output<string?> GroupsAttribute { get; private set; } = null!;

        /// <summary>
        /// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        /// </summary>
        [Output("groupsFilters")]
        public Output<ImmutableArray<string>> GroupsFilters { get; private set; } = null!;

        /// <summary>
        /// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        /// </summary>
        [Output("issuerMode")]
        public Output<string?> IssuerMode { get; private set; } = null!;

        [Output("matchAttribute")]
        public Output<string?> MatchAttribute { get; private set; } = null!;

        [Output("matchType")]
        public Output<string?> MatchType { get; private set; } = null!;

        /// <summary>
        /// Maximum allowable clock-skew when processing messages from the IdP.
        /// </summary>
        [Output("maxClockSkew")]
        public Output<int?> MaxClockSkew { get; private set; } = null!;

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Determines if the IdP should act as a source of truth for user profile attributes.
        /// </summary>
        [Output("profileMaster")]
        public Output<bool?> ProfileMaster { get; private set; } = null!;

        /// <summary>
        /// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
        /// </summary>
        [Output("protocolType")]
        public Output<string?> ProtocolType { get; private set; } = null!;

        /// <summary>
        /// Provisioning action for an IdP user during authentication.
        /// </summary>
        [Output("provisioningAction")]
        public Output<string?> ProvisioningAction { get; private set; } = null!;

        /// <summary>
        /// The XML digital signature algorithm used when signing an AuthnRequest message.
        /// </summary>
        [Output("requestSignatureAlgorithm")]
        public Output<string?> RequestSignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        /// </summary>
        [Output("requestSignatureScope")]
        public Output<string?> RequestSignatureScope { get; private set; } = null!;

        /// <summary>
        /// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        /// </summary>
        [Output("responseSignatureAlgorithm")]
        public Output<string?> ResponseSignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        /// </summary>
        [Output("responseSignatureScope")]
        public Output<string?> ResponseSignatureScope { get; private set; } = null!;

        /// <summary>
        /// The scopes of the IdP.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Status of the IdP.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        /// </summary>
        [Output("subjectMatchAttribute")]
        public Output<string?> SubjectMatchAttribute { get; private set; } = null!;

        /// <summary>
        /// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        /// </summary>
        [Output("subjectMatchType")]
        public Output<string?> SubjectMatchType { get; private set; } = null!;

        /// <summary>
        /// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        /// </summary>
        [Output("suspendedAction")]
        public Output<string?> SuspendedAction { get; private set; } = null!;

        /// <summary>
        /// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        /// </summary>
        [Output("tokenBinding")]
        public Output<string> TokenBinding { get; private set; } = null!;

        /// <summary>
        /// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
        /// </summary>
        [Output("tokenUrl")]
        public Output<string> TokenUrl { get; private set; } = null!;

        /// <summary>
        /// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Okta EL Expression to generate or transform a unique username for the IdP user.
        /// </summary>
        [Output("usernameTemplate")]
        public Output<string?> UsernameTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a Social resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Social(string name, SocialArgs args, CustomResourceOptions? options = null)
            : base("okta:idp/social:Social", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Social(string name, Input<string> id, SocialState? state = null, CustomResourceOptions? options = null)
            : base("okta:idp/social:Social", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Social resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Social Get(string name, Input<string> id, SocialState? state = null, CustomResourceOptions? options = null)
        {
            return new Social(name, id, state, options);
        }
    }

    public sealed class SocialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the account linking action for an IdP user.
        /// </summary>
        [Input("accountLinkAction")]
        public Input<string>? AccountLinkAction { get; set; }

        [Input("accountLinkGroupIncludes")]
        private InputList<string>? _accountLinkGroupIncludes;

        /// <summary>
        /// Group memberships to determine link candidates.
        /// </summary>
        public InputList<string> AccountLinkGroupIncludes
        {
            get => _accountLinkGroupIncludes ?? (_accountLinkGroupIncludes = new InputList<string>());
            set => _accountLinkGroupIncludes = value;
        }

        /// <summary>
        /// Unique identifier issued by AS for the Okta IdP instance.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client secret issued by AS for the Okta IdP instance.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        /// </summary>
        [Input("deprovisionedAction")]
        public Input<string>? DeprovisionedAction { get; set; }

        /// <summary>
        /// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        /// </summary>
        [Input("groupsAction")]
        public Input<string>? GroupsAction { get; set; }

        [Input("groupsAssignments")]
        private InputList<string>? _groupsAssignments;

        /// <summary>
        /// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        /// </summary>
        public InputList<string> GroupsAssignments
        {
            get => _groupsAssignments ?? (_groupsAssignments = new InputList<string>());
            set => _groupsAssignments = value;
        }

        /// <summary>
        /// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        /// </summary>
        [Input("groupsAttribute")]
        public Input<string>? GroupsAttribute { get; set; }

        [Input("groupsFilters")]
        private InputList<string>? _groupsFilters;

        /// <summary>
        /// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        /// </summary>
        public InputList<string> GroupsFilters
        {
            get => _groupsFilters ?? (_groupsFilters = new InputList<string>());
            set => _groupsFilters = value;
        }

        /// <summary>
        /// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        [Input("matchAttribute")]
        public Input<string>? MatchAttribute { get; set; }

        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Maximum allowable clock-skew when processing messages from the IdP.
        /// </summary>
        [Input("maxClockSkew")]
        public Input<int>? MaxClockSkew { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Determines if the IdP should act as a source of truth for user profile attributes.
        /// </summary>
        [Input("profileMaster")]
        public Input<bool>? ProfileMaster { get; set; }

        /// <summary>
        /// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        /// <summary>
        /// Provisioning action for an IdP user during authentication.
        /// </summary>
        [Input("provisioningAction")]
        public Input<string>? ProvisioningAction { get; set; }

        /// <summary>
        /// The XML digital signature algorithm used when signing an AuthnRequest message.
        /// </summary>
        [Input("requestSignatureAlgorithm")]
        public Input<string>? RequestSignatureAlgorithm { get; set; }

        /// <summary>
        /// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        /// </summary>
        [Input("requestSignatureScope")]
        public Input<string>? RequestSignatureScope { get; set; }

        /// <summary>
        /// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        /// </summary>
        [Input("responseSignatureAlgorithm")]
        public Input<string>? ResponseSignatureAlgorithm { get; set; }

        /// <summary>
        /// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        /// </summary>
        [Input("responseSignatureScope")]
        public Input<string>? ResponseSignatureScope { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of the IdP.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Status of the IdP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        /// </summary>
        [Input("subjectMatchAttribute")]
        public Input<string>? SubjectMatchAttribute { get; set; }

        /// <summary>
        /// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        /// </summary>
        [Input("subjectMatchType")]
        public Input<string>? SubjectMatchType { get; set; }

        /// <summary>
        /// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        /// </summary>
        [Input("suspendedAction")]
        public Input<string>? SuspendedAction { get; set; }

        /// <summary>
        /// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Okta EL Expression to generate or transform a unique username for the IdP user.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SocialArgs()
        {
        }
    }

    public sealed class SocialState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the account linking action for an IdP user.
        /// </summary>
        [Input("accountLinkAction")]
        public Input<string>? AccountLinkAction { get; set; }

        [Input("accountLinkGroupIncludes")]
        private InputList<string>? _accountLinkGroupIncludes;

        /// <summary>
        /// Group memberships to determine link candidates.
        /// </summary>
        public InputList<string> AccountLinkGroupIncludes
        {
            get => _accountLinkGroupIncludes ?? (_accountLinkGroupIncludes = new InputList<string>());
            set => _accountLinkGroupIncludes = value;
        }

        /// <summary>
        /// The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        /// </summary>
        [Input("authorizationBinding")]
        public Input<string>? AuthorizationBinding { get; set; }

        /// <summary>
        /// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
        /// </summary>
        [Input("authorizationUrl")]
        public Input<string>? AuthorizationUrl { get; set; }

        /// <summary>
        /// Unique identifier issued by AS for the Okta IdP instance.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client secret issued by AS for the Okta IdP instance.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        /// </summary>
        [Input("deprovisionedAction")]
        public Input<string>? DeprovisionedAction { get; set; }

        /// <summary>
        /// Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        /// </summary>
        [Input("groupsAction")]
        public Input<string>? GroupsAction { get; set; }

        [Input("groupsAssignments")]
        private InputList<string>? _groupsAssignments;

        /// <summary>
        /// List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        /// </summary>
        public InputList<string> GroupsAssignments
        {
            get => _groupsAssignments ?? (_groupsAssignments = new InputList<string>());
            set => _groupsAssignments = value;
        }

        /// <summary>
        /// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        /// </summary>
        [Input("groupsAttribute")]
        public Input<string>? GroupsAttribute { get; set; }

        [Input("groupsFilters")]
        private InputList<string>? _groupsFilters;

        /// <summary>
        /// Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        /// </summary>
        public InputList<string> GroupsFilters
        {
            get => _groupsFilters ?? (_groupsFilters = new InputList<string>());
            set => _groupsFilters = value;
        }

        /// <summary>
        /// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        [Input("matchAttribute")]
        public Input<string>? MatchAttribute { get; set; }

        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Maximum allowable clock-skew when processing messages from the IdP.
        /// </summary>
        [Input("maxClockSkew")]
        public Input<int>? MaxClockSkew { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Determines if the IdP should act as a source of truth for user profile attributes.
        /// </summary>
        [Input("profileMaster")]
        public Input<bool>? ProfileMaster { get; set; }

        /// <summary>
        /// The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        /// <summary>
        /// Provisioning action for an IdP user during authentication.
        /// </summary>
        [Input("provisioningAction")]
        public Input<string>? ProvisioningAction { get; set; }

        /// <summary>
        /// The XML digital signature algorithm used when signing an AuthnRequest message.
        /// </summary>
        [Input("requestSignatureAlgorithm")]
        public Input<string>? RequestSignatureAlgorithm { get; set; }

        /// <summary>
        /// Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        /// </summary>
        [Input("requestSignatureScope")]
        public Input<string>? RequestSignatureScope { get; set; }

        /// <summary>
        /// The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        /// </summary>
        [Input("responseSignatureAlgorithm")]
        public Input<string>? ResponseSignatureAlgorithm { get; set; }

        /// <summary>
        /// Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        /// </summary>
        [Input("responseSignatureScope")]
        public Input<string>? ResponseSignatureScope { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of the IdP.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Status of the IdP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        /// </summary>
        [Input("subjectMatchAttribute")]
        public Input<string>? SubjectMatchAttribute { get; set; }

        /// <summary>
        /// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        /// </summary>
        [Input("subjectMatchType")]
        public Input<string>? SubjectMatchType { get; set; }

        /// <summary>
        /// Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        /// </summary>
        [Input("suspendedAction")]
        public Input<string>? SuspendedAction { get; set; }

        /// <summary>
        /// The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        /// </summary>
        [Input("tokenBinding")]
        public Input<string>? TokenBinding { get; set; }

        /// <summary>
        /// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
        /// </summary>
        [Input("tokenUrl")]
        public Input<string>? TokenUrl { get; set; }

        /// <summary>
        /// The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Okta EL Expression to generate or transform a unique username for the IdP user.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SocialState()
        {
        }
    }
}
