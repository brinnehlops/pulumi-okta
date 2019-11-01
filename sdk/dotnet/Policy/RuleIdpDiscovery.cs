// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Policy
{
    /// <summary>
    /// Creates an IdP Discovery Policy Rule.
    /// 
    /// This resource allows you to create and configure an IdP Discovery Policy Rule.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_idp_discovery.html.markdown.
    /// </summary>
    public partial class RuleIdpDiscovery : Pulumi.CustomResource
    {
        /// <summary>
        /// Applications to exclude in discovery rule
        /// </summary>
        [Output("appExcludes")]
        public Output<ImmutableArray<Outputs.RuleIdpDiscoveryAppExcludes>> AppExcludes { get; private set; } = null!;

        /// <summary>
        /// Applications to include in discovery rule
        /// </summary>
        [Output("appIncludes")]
        public Output<ImmutableArray<Outputs.RuleIdpDiscoveryAppIncludes>> AppIncludes { get; private set; } = null!;

        [Output("idpId")]
        public Output<string?> IdpId { get; private set; } = null!;

        [Output("idpType")]
        public Output<string?> IdpType { get; private set; } = null!;

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        /// </summary>
        [Output("networkConnection")]
        public Output<string?> NetworkConnection { get; private set; } = null!;

        /// <summary>
        /// The network zones to exclude. Conflicts with `network_includes`.
        /// </summary>
        [Output("networkExcludes")]
        public Output<ImmutableArray<string>> NetworkExcludes { get; private set; } = null!;

        /// <summary>
        /// The network zones to include. Conflicts with `network_excludes`.
        /// </summary>
        [Output("networkIncludes")]
        public Output<ImmutableArray<string>> NetworkIncludes { get; private set; } = null!;

        [Output("platformIncludes")]
        public Output<ImmutableArray<Outputs.RuleIdpDiscoveryPlatformIncludes>> PlatformIncludes { get; private set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Output("policyid")]
        public Output<string> Policyid { get; private set; } = null!;

        /// <summary>
        /// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("userIdentifierAttribute")]
        public Output<string?> UserIdentifierAttribute { get; private set; } = null!;

        [Output("userIdentifierPatterns")]
        public Output<ImmutableArray<Outputs.RuleIdpDiscoveryUserIdentifierPatterns>> UserIdentifierPatterns { get; private set; } = null!;

        [Output("userIdentifierType")]
        public Output<string?> UserIdentifierType { get; private set; } = null!;


        /// <summary>
        /// Create a RuleIdpDiscovery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleIdpDiscovery(string name, RuleIdpDiscoveryArgs args, CustomResourceOptions? options = null)
            : base("okta:policy/ruleIdpDiscovery:RuleIdpDiscovery", name, args, MakeResourceOptions(options, ""))
        {
        }

        private RuleIdpDiscovery(string name, Input<string> id, RuleIdpDiscoveryState? state = null, CustomResourceOptions? options = null)
            : base("okta:policy/ruleIdpDiscovery:RuleIdpDiscovery", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RuleIdpDiscovery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleIdpDiscovery Get(string name, Input<string> id, RuleIdpDiscoveryState? state = null, CustomResourceOptions? options = null)
        {
            return new RuleIdpDiscovery(name, id, state, options);
        }
    }

    public sealed class RuleIdpDiscoveryArgs : Pulumi.ResourceArgs
    {
        [Input("appExcludes")]
        private InputList<Inputs.RuleIdpDiscoveryAppExcludesArgs>? _appExcludes;

        /// <summary>
        /// Applications to exclude in discovery rule
        /// </summary>
        public InputList<Inputs.RuleIdpDiscoveryAppExcludesArgs> AppExcludes
        {
            get => _appExcludes ?? (_appExcludes = new InputList<Inputs.RuleIdpDiscoveryAppExcludesArgs>());
            set => _appExcludes = value;
        }

        [Input("appIncludes")]
        private InputList<Inputs.RuleIdpDiscoveryAppIncludesArgs>? _appIncludes;

        /// <summary>
        /// Applications to include in discovery rule
        /// </summary>
        public InputList<Inputs.RuleIdpDiscoveryAppIncludesArgs> AppIncludes
        {
            get => _appIncludes ?? (_appIncludes = new InputList<Inputs.RuleIdpDiscoveryAppIncludesArgs>());
            set => _appIncludes = value;
        }

        [Input("idpId")]
        public Input<string>? IdpId { get; set; }

        [Input("idpType")]
        public Input<string>? IdpType { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        /// </summary>
        [Input("networkConnection")]
        public Input<string>? NetworkConnection { get; set; }

        [Input("networkExcludes")]
        private InputList<string>? _networkExcludes;

        /// <summary>
        /// The network zones to exclude. Conflicts with `network_includes`.
        /// </summary>
        public InputList<string> NetworkExcludes
        {
            get => _networkExcludes ?? (_networkExcludes = new InputList<string>());
            set => _networkExcludes = value;
        }

        [Input("networkIncludes")]
        private InputList<string>? _networkIncludes;

        /// <summary>
        /// The network zones to include. Conflicts with `network_excludes`.
        /// </summary>
        public InputList<string> NetworkIncludes
        {
            get => _networkIncludes ?? (_networkIncludes = new InputList<string>());
            set => _networkIncludes = value;
        }

        [Input("platformIncludes")]
        private InputList<Inputs.RuleIdpDiscoveryPlatformIncludesArgs>? _platformIncludes;
        public InputList<Inputs.RuleIdpDiscoveryPlatformIncludesArgs> PlatformIncludes
        {
            get => _platformIncludes ?? (_platformIncludes = new InputList<Inputs.RuleIdpDiscoveryPlatformIncludesArgs>());
            set => _platformIncludes = value;
        }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid", required: true)]
        public Input<string> Policyid { get; set; } = null!;

        /// <summary>
        /// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("userIdentifierAttribute")]
        public Input<string>? UserIdentifierAttribute { get; set; }

        [Input("userIdentifierPatterns")]
        private InputList<Inputs.RuleIdpDiscoveryUserIdentifierPatternsArgs>? _userIdentifierPatterns;
        public InputList<Inputs.RuleIdpDiscoveryUserIdentifierPatternsArgs> UserIdentifierPatterns
        {
            get => _userIdentifierPatterns ?? (_userIdentifierPatterns = new InputList<Inputs.RuleIdpDiscoveryUserIdentifierPatternsArgs>());
            set => _userIdentifierPatterns = value;
        }

        [Input("userIdentifierType")]
        public Input<string>? UserIdentifierType { get; set; }

        public RuleIdpDiscoveryArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryState : Pulumi.ResourceArgs
    {
        [Input("appExcludes")]
        private InputList<Inputs.RuleIdpDiscoveryAppExcludesGetArgs>? _appExcludes;

        /// <summary>
        /// Applications to exclude in discovery rule
        /// </summary>
        public InputList<Inputs.RuleIdpDiscoveryAppExcludesGetArgs> AppExcludes
        {
            get => _appExcludes ?? (_appExcludes = new InputList<Inputs.RuleIdpDiscoveryAppExcludesGetArgs>());
            set => _appExcludes = value;
        }

        [Input("appIncludes")]
        private InputList<Inputs.RuleIdpDiscoveryAppIncludesGetArgs>? _appIncludes;

        /// <summary>
        /// Applications to include in discovery rule
        /// </summary>
        public InputList<Inputs.RuleIdpDiscoveryAppIncludesGetArgs> AppIncludes
        {
            get => _appIncludes ?? (_appIncludes = new InputList<Inputs.RuleIdpDiscoveryAppIncludesGetArgs>());
            set => _appIncludes = value;
        }

        [Input("idpId")]
        public Input<string>? IdpId { get; set; }

        [Input("idpType")]
        public Input<string>? IdpType { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        /// </summary>
        [Input("networkConnection")]
        public Input<string>? NetworkConnection { get; set; }

        [Input("networkExcludes")]
        private InputList<string>? _networkExcludes;

        /// <summary>
        /// The network zones to exclude. Conflicts with `network_includes`.
        /// </summary>
        public InputList<string> NetworkExcludes
        {
            get => _networkExcludes ?? (_networkExcludes = new InputList<string>());
            set => _networkExcludes = value;
        }

        [Input("networkIncludes")]
        private InputList<string>? _networkIncludes;

        /// <summary>
        /// The network zones to include. Conflicts with `network_excludes`.
        /// </summary>
        public InputList<string> NetworkIncludes
        {
            get => _networkIncludes ?? (_networkIncludes = new InputList<string>());
            set => _networkIncludes = value;
        }

        [Input("platformIncludes")]
        private InputList<Inputs.RuleIdpDiscoveryPlatformIncludesGetArgs>? _platformIncludes;
        public InputList<Inputs.RuleIdpDiscoveryPlatformIncludesGetArgs> PlatformIncludes
        {
            get => _platformIncludes ?? (_platformIncludes = new InputList<Inputs.RuleIdpDiscoveryPlatformIncludesGetArgs>());
            set => _platformIncludes = value;
        }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyid")]
        public Input<string>? Policyid { get; set; }

        /// <summary>
        /// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("userIdentifierAttribute")]
        public Input<string>? UserIdentifierAttribute { get; set; }

        [Input("userIdentifierPatterns")]
        private InputList<Inputs.RuleIdpDiscoveryUserIdentifierPatternsGetArgs>? _userIdentifierPatterns;
        public InputList<Inputs.RuleIdpDiscoveryUserIdentifierPatternsGetArgs> UserIdentifierPatterns
        {
            get => _userIdentifierPatterns ?? (_userIdentifierPatterns = new InputList<Inputs.RuleIdpDiscoveryUserIdentifierPatternsGetArgs>());
            set => _userIdentifierPatterns = value;
        }

        [Input("userIdentifierType")]
        public Input<string>? UserIdentifierType { get; set; }

        public RuleIdpDiscoveryState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RuleIdpDiscoveryAppExcludesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryAppExcludesArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryAppExcludesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryAppExcludesGetArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryAppIncludesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryAppIncludesArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryAppIncludesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryAppIncludesGetArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryPlatformIncludesArgs : Pulumi.ResourceArgs
    {
        [Input("osExpression")]
        public Input<string>? OsExpression { get; set; }

        [Input("osType")]
        public Input<string>? OsType { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryPlatformIncludesArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryPlatformIncludesGetArgs : Pulumi.ResourceArgs
    {
        [Input("osExpression")]
        public Input<string>? OsExpression { get; set; }

        [Input("osType")]
        public Input<string>? OsType { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryPlatformIncludesGetArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryUserIdentifierPatternsArgs : Pulumi.ResourceArgs
    {
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public RuleIdpDiscoveryUserIdentifierPatternsArgs()
        {
        }
    }

    public sealed class RuleIdpDiscoveryUserIdentifierPatternsGetArgs : Pulumi.ResourceArgs
    {
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public RuleIdpDiscoveryUserIdentifierPatternsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RuleIdpDiscoveryAppExcludes
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        public readonly string? Name;
        public readonly string? Type;

        [OutputConstructor]
        private RuleIdpDiscoveryAppExcludes(
            string? id,
            string? name,
            string? type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }

    [OutputType]
    public sealed class RuleIdpDiscoveryAppIncludes
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        public readonly string? Name;
        public readonly string? Type;

        [OutputConstructor]
        private RuleIdpDiscoveryAppIncludes(
            string? id,
            string? name,
            string? type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }

    [OutputType]
    public sealed class RuleIdpDiscoveryPlatformIncludes
    {
        public readonly string? OsExpression;
        public readonly string? OsType;
        public readonly string? Type;

        [OutputConstructor]
        private RuleIdpDiscoveryPlatformIncludes(
            string? osExpression,
            string? osType,
            string? type)
        {
            OsExpression = osExpression;
            OsType = osType;
            Type = type;
        }
    }

    [OutputType]
    public sealed class RuleIdpDiscoveryUserIdentifierPatterns
    {
        public readonly string? MatchType;
        public readonly string? Value;

        [OutputConstructor]
        private RuleIdpDiscoveryUserIdentifierPatterns(
            string? matchType,
            string? value)
        {
            MatchType = matchType;
            Value = value;
        }
    }
    }
}