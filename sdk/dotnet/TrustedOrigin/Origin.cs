// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.TrustedOrigin
{
    /// <summary>
    /// Creates a Trusted Origin.
    /// 
    /// This resource allows you to create and configure an Trusted Origin.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Okta = Pulumi.Okta;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Okta.TrustedOrigin.Origin("example", new Okta.TrustedOrigin.OriginArgs
    ///         {
    ///             Origin = "https://example.com",
    ///             Scopes = 
    ///             {
    ///                 "CORS",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Origin : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the Trusted Origin is active or not - can only be issued post-creation.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// Name of the Trusted Origin Resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The origin to trust.
        /// </summary>
        [Output("origin")]
        public Output<string> OriginName { get; private set; } = null!;

        /// <summary>
        /// Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;


        /// <summary>
        /// Create a Origin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Origin(string name, OriginArgs args, CustomResourceOptions? options = null)
            : base("okta:trustedorigin/origin:Origin", name, args ?? new OriginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Origin(string name, Input<string> id, OriginState? state = null, CustomResourceOptions? options = null)
            : base("okta:trustedorigin/origin:Origin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Origin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Origin Get(string name, Input<string> id, OriginState? state = null, CustomResourceOptions? options = null)
        {
            return new Origin(name, id, state, options);
        }
    }

    public sealed class OriginArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the Trusted Origin is active or not - can only be issued post-creation.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Name of the Trusted Origin Resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The origin to trust.
        /// </summary>
        [Input("origin", required: true)]
        public Input<string> OriginName { get; set; } = null!;

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public OriginArgs()
        {
        }
    }

    public sealed class OriginState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the Trusted Origin is active or not - can only be issued post-creation.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Name of the Trusted Origin Resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The origin to trust.
        /// </summary>
        [Input("origin")]
        public Input<string>? OriginName { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public OriginState()
        {
        }
    }
}
