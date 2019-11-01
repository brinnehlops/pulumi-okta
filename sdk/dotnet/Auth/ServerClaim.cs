// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Auth
{
    /// <summary>
    /// Creates an Authorization Server Claim.
    /// 
    /// This resource allows you to create and configure an Authorization Server Claim.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/auth_server_claim.html.markdown.
    /// </summary>
    public partial class ServerClaim : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to include claims in token, by default is is set to `true`.
        /// </summary>
        [Output("alwaysIncludeInToken")]
        public Output<bool?> AlwaysIncludeInToken { get; private set; } = null!;

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Output("authServerId")]
        public Output<string> AuthServerId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the claim is for an access token `"RESOURCE"` or ID token `"IDENTITY"`.
        /// </summary>
        [Output("claimType")]
        public Output<string> ClaimType { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of group filter if `value_type` is `"GROUPS"`. Can be set to one of the following `"STARTS_WITH"`, `"EQUALS"`, `"CONTAINS"`, `"REGEX"`.
        /// </summary>
        [Output("groupFilterType")]
        public Output<string?> GroupFilterType { get; private set; } = null!;

        /// <summary>
        /// The name of the claim.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of scopes the auth server claim is tied to.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// The status of the application. It defaults to `"ACTIVE"`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The value of the claim.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// The type of value of the claim. It can be set to `"EXPRESSION"` or `"GROUPS"`. It defaults to `"EXPRESSION"`.
        /// </summary>
        [Output("valueType")]
        public Output<string?> ValueType { get; private set; } = null!;


        /// <summary>
        /// Create a ServerClaim resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerClaim(string name, ServerClaimArgs args, CustomResourceOptions? options = null)
            : base("okta:auth/serverClaim:ServerClaim", name, args, MakeResourceOptions(options, ""))
        {
        }

        private ServerClaim(string name, Input<string> id, ServerClaimState? state = null, CustomResourceOptions? options = null)
            : base("okta:auth/serverClaim:ServerClaim", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerClaim resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerClaim Get(string name, Input<string> id, ServerClaimState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerClaim(name, id, state, options);
        }
    }

    public sealed class ServerClaimArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to include claims in token, by default is is set to `true`.
        /// </summary>
        [Input("alwaysIncludeInToken")]
        public Input<bool>? AlwaysIncludeInToken { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("authServerId", required: true)]
        public Input<string> AuthServerId { get; set; } = null!;

        /// <summary>
        /// Specifies whether the claim is for an access token `"RESOURCE"` or ID token `"IDENTITY"`.
        /// </summary>
        [Input("claimType", required: true)]
        public Input<string> ClaimType { get; set; } = null!;

        /// <summary>
        /// Specifies the type of group filter if `value_type` is `"GROUPS"`. Can be set to one of the following `"STARTS_WITH"`, `"EQUALS"`, `"CONTAINS"`, `"REGEX"`.
        /// </summary>
        [Input("groupFilterType")]
        public Input<string>? GroupFilterType { get; set; }

        /// <summary>
        /// The name of the claim.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The list of scopes the auth server claim is tied to.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The status of the application. It defaults to `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The value of the claim.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The type of value of the claim. It can be set to `"EXPRESSION"` or `"GROUPS"`. It defaults to `"EXPRESSION"`.
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public ServerClaimArgs()
        {
        }
    }

    public sealed class ServerClaimState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to include claims in token, by default is is set to `true`.
        /// </summary>
        [Input("alwaysIncludeInToken")]
        public Input<bool>? AlwaysIncludeInToken { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("authServerId")]
        public Input<string>? AuthServerId { get; set; }

        /// <summary>
        /// Specifies whether the claim is for an access token `"RESOURCE"` or ID token `"IDENTITY"`.
        /// </summary>
        [Input("claimType")]
        public Input<string>? ClaimType { get; set; }

        /// <summary>
        /// Specifies the type of group filter if `value_type` is `"GROUPS"`. Can be set to one of the following `"STARTS_WITH"`, `"EQUALS"`, `"CONTAINS"`, `"REGEX"`.
        /// </summary>
        [Input("groupFilterType")]
        public Input<string>? GroupFilterType { get; set; }

        /// <summary>
        /// The name of the claim.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The list of scopes the auth server claim is tied to.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The status of the application. It defaults to `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The value of the claim.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The type of value of the claim. It can be set to `"EXPRESSION"` or `"GROUPS"`. It defaults to `"EXPRESSION"`.
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public ServerClaimState()
        {
        }
    }
}