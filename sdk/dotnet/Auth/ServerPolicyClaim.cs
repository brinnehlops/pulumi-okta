// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Auth
{
    /// <summary>
    /// Creates an Authorization Server Policy Rule.
    /// 
    /// This resource allows you to create and configure an Authorization Server Policy Rule.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Okta = Pulumi.Okta;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Okta.Auth.ServerPolicyClaim("example", new Okta.Auth.ServerPolicyClaimArgs
    ///         {
    ///             AuthServerId = "&lt;auth server id&gt;",
    ///             GrantTypeWhitelists = 
    ///             {
    ///                 "implicit",
    ///             },
    ///             GroupWhitelists = 
    ///             {
    ///                 "&lt;group ids&gt;",
    ///             },
    ///             PolicyId = "&lt;auth server policy id&gt;",
    ///             Priority = 1,
    ///             Status = "ACTIVE",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ServerPolicyClaim : Pulumi.CustomResource
    {
        /// <summary>
        /// Lifetime of access token. Can be set to a value between 5 and 1440.
        /// </summary>
        [Output("accessTokenLifetimeMinutes")]
        public Output<int?> AccessTokenLifetimeMinutes { get; private set; } = null!;

        /// <summary>
        /// Auth Server ID.
        /// </summary>
        [Output("authServerId")]
        public Output<string> AuthServerId { get; private set; } = null!;

        /// <summary>
        /// Accepted grant type values, `"authorization_code"`, `"implicit"`, `"password"`
        /// </summary>
        [Output("grantTypeWhitelists")]
        public Output<ImmutableArray<string>> GrantTypeWhitelists { get; private set; } = null!;

        [Output("groupBlacklists")]
        public Output<ImmutableArray<string>> GroupBlacklists { get; private set; } = null!;

        [Output("groupWhitelists")]
        public Output<ImmutableArray<string>> GroupWhitelists { get; private set; } = null!;

        /// <summary>
        /// The ID of the inline token to trigger.
        /// </summary>
        [Output("inlineHookId")]
        public Output<string?> InlineHookId { get; private set; } = null!;

        /// <summary>
        /// Auth Server Policy Rule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Auth Server Policy ID.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Priority of the auth server policy rule.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Lifetime of refresh token.
        /// </summary>
        [Output("refreshTokenLifetimeMinutes")]
        public Output<int?> RefreshTokenLifetimeMinutes { get; private set; } = null!;

        [Output("refreshTokenWindowMinutes")]
        public Output<int?> RefreshTokenWindowMinutes { get; private set; } = null!;

        /// <summary>
        /// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
        /// </summary>
        [Output("scopeWhitelists")]
        public Output<ImmutableArray<string>> ScopeWhitelists { get; private set; } = null!;

        /// <summary>
        /// The status of the Auth Server Policy Rule.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The type of the Auth Server Policy Rule.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("userBlacklists")]
        public Output<ImmutableArray<string>> UserBlacklists { get; private set; } = null!;

        [Output("userWhitelists")]
        public Output<ImmutableArray<string>> UserWhitelists { get; private set; } = null!;


        /// <summary>
        /// Create a ServerPolicyClaim resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerPolicyClaim(string name, ServerPolicyClaimArgs args, CustomResourceOptions? options = null)
            : base("okta:auth/serverPolicyClaim:ServerPolicyClaim", name, args ?? new ServerPolicyClaimArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerPolicyClaim(string name, Input<string> id, ServerPolicyClaimState? state = null, CustomResourceOptions? options = null)
            : base("okta:auth/serverPolicyClaim:ServerPolicyClaim", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerPolicyClaim resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerPolicyClaim Get(string name, Input<string> id, ServerPolicyClaimState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerPolicyClaim(name, id, state, options);
        }
    }

    public sealed class ServerPolicyClaimArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Lifetime of access token. Can be set to a value between 5 and 1440.
        /// </summary>
        [Input("accessTokenLifetimeMinutes")]
        public Input<int>? AccessTokenLifetimeMinutes { get; set; }

        /// <summary>
        /// Auth Server ID.
        /// </summary>
        [Input("authServerId", required: true)]
        public Input<string> AuthServerId { get; set; } = null!;

        [Input("grantTypeWhitelists", required: true)]
        private InputList<string>? _grantTypeWhitelists;

        /// <summary>
        /// Accepted grant type values, `"authorization_code"`, `"implicit"`, `"password"`
        /// </summary>
        public InputList<string> GrantTypeWhitelists
        {
            get => _grantTypeWhitelists ?? (_grantTypeWhitelists = new InputList<string>());
            set => _grantTypeWhitelists = value;
        }

        [Input("groupBlacklists")]
        private InputList<string>? _groupBlacklists;
        public InputList<string> GroupBlacklists
        {
            get => _groupBlacklists ?? (_groupBlacklists = new InputList<string>());
            set => _groupBlacklists = value;
        }

        [Input("groupWhitelists")]
        private InputList<string>? _groupWhitelists;
        public InputList<string> GroupWhitelists
        {
            get => _groupWhitelists ?? (_groupWhitelists = new InputList<string>());
            set => _groupWhitelists = value;
        }

        /// <summary>
        /// The ID of the inline token to trigger.
        /// </summary>
        [Input("inlineHookId")]
        public Input<string>? InlineHookId { get; set; }

        /// <summary>
        /// Auth Server Policy Rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Auth Server Policy ID.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// Priority of the auth server policy rule.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// Lifetime of refresh token.
        /// </summary>
        [Input("refreshTokenLifetimeMinutes")]
        public Input<int>? RefreshTokenLifetimeMinutes { get; set; }

        [Input("refreshTokenWindowMinutes")]
        public Input<int>? RefreshTokenWindowMinutes { get; set; }

        [Input("scopeWhitelists")]
        private InputList<string>? _scopeWhitelists;

        /// <summary>
        /// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
        /// </summary>
        public InputList<string> ScopeWhitelists
        {
            get => _scopeWhitelists ?? (_scopeWhitelists = new InputList<string>());
            set => _scopeWhitelists = value;
        }

        /// <summary>
        /// The status of the Auth Server Policy Rule.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of the Auth Server Policy Rule.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userBlacklists")]
        private InputList<string>? _userBlacklists;
        public InputList<string> UserBlacklists
        {
            get => _userBlacklists ?? (_userBlacklists = new InputList<string>());
            set => _userBlacklists = value;
        }

        [Input("userWhitelists")]
        private InputList<string>? _userWhitelists;
        public InputList<string> UserWhitelists
        {
            get => _userWhitelists ?? (_userWhitelists = new InputList<string>());
            set => _userWhitelists = value;
        }

        public ServerPolicyClaimArgs()
        {
        }
    }

    public sealed class ServerPolicyClaimState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Lifetime of access token. Can be set to a value between 5 and 1440.
        /// </summary>
        [Input("accessTokenLifetimeMinutes")]
        public Input<int>? AccessTokenLifetimeMinutes { get; set; }

        /// <summary>
        /// Auth Server ID.
        /// </summary>
        [Input("authServerId")]
        public Input<string>? AuthServerId { get; set; }

        [Input("grantTypeWhitelists")]
        private InputList<string>? _grantTypeWhitelists;

        /// <summary>
        /// Accepted grant type values, `"authorization_code"`, `"implicit"`, `"password"`
        /// </summary>
        public InputList<string> GrantTypeWhitelists
        {
            get => _grantTypeWhitelists ?? (_grantTypeWhitelists = new InputList<string>());
            set => _grantTypeWhitelists = value;
        }

        [Input("groupBlacklists")]
        private InputList<string>? _groupBlacklists;
        public InputList<string> GroupBlacklists
        {
            get => _groupBlacklists ?? (_groupBlacklists = new InputList<string>());
            set => _groupBlacklists = value;
        }

        [Input("groupWhitelists")]
        private InputList<string>? _groupWhitelists;
        public InputList<string> GroupWhitelists
        {
            get => _groupWhitelists ?? (_groupWhitelists = new InputList<string>());
            set => _groupWhitelists = value;
        }

        /// <summary>
        /// The ID of the inline token to trigger.
        /// </summary>
        [Input("inlineHookId")]
        public Input<string>? InlineHookId { get; set; }

        /// <summary>
        /// Auth Server Policy Rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Auth Server Policy ID.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Priority of the auth server policy rule.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Lifetime of refresh token.
        /// </summary>
        [Input("refreshTokenLifetimeMinutes")]
        public Input<int>? RefreshTokenLifetimeMinutes { get; set; }

        [Input("refreshTokenWindowMinutes")]
        public Input<int>? RefreshTokenWindowMinutes { get; set; }

        [Input("scopeWhitelists")]
        private InputList<string>? _scopeWhitelists;

        /// <summary>
        /// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
        /// </summary>
        public InputList<string> ScopeWhitelists
        {
            get => _scopeWhitelists ?? (_scopeWhitelists = new InputList<string>());
            set => _scopeWhitelists = value;
        }

        /// <summary>
        /// The status of the Auth Server Policy Rule.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of the Auth Server Policy Rule.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userBlacklists")]
        private InputList<string>? _userBlacklists;
        public InputList<string> UserBlacklists
        {
            get => _userBlacklists ?? (_userBlacklists = new InputList<string>());
            set => _userBlacklists = value;
        }

        [Input("userWhitelists")]
        private InputList<string>? _userWhitelists;
        public InputList<string> UserWhitelists
        {
            get => _userWhitelists ?? (_userWhitelists = new InputList<string>());
            set => _userWhitelists = value;
        }

        public ServerPolicyClaimState()
        {
        }
    }
}
