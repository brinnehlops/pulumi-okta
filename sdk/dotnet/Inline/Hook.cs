// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Inline
{
    /// <summary>
    /// Creates an inline hook.
    /// 
    /// This resource allows you to create and configure an inline hook.
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
    ///         var example = new Okta.Inline.Hook("example", new Okta.Inline.HookArgs
    ///         {
    ///             Auth = new Okta.Inline.Inputs.HookAuthArgs
    ///             {
    ///                 Key = "Authorization",
    ///                 Type = "HEADER",
    ///                 Value = "secret",
    ///             },
    ///             Channel = new Okta.Inline.Inputs.HookChannelArgs
    ///             {
    ///                 Method = "POST",
    ///                 Uri = "https://example.com/test",
    ///                 Version = "1.0.0",
    ///             },
    ///             Type = "com.okta.oauth2.tokens.transform",
    ///             Version = "1.0.1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Hook : Pulumi.CustomResource
    {
        /// <summary>
        /// Authentication required for inline hook request.
        /// </summary>
        [Output("auth")]
        public Output<Outputs.HookAuth?> Auth { get; private set; } = null!;

        /// <summary>
        /// Details of the endpoint the inline hook will hit.
        /// </summary>
        [Output("channel")]
        public Output<Outputs.HookChannel?> Channel { get; private set; } = null!;

        /// <summary>
        /// Map of headers to send along in inline hook request.
        /// </summary>
        [Output("headers")]
        public Output<ImmutableArray<Outputs.HookHeader>> Headers { get; private set; } = null!;

        /// <summary>
        /// The inline hook display name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The type of hook to trigger. Currently only `"HTTP"` is supported.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The version of the endpoint.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Hook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hook(string name, HookArgs args, CustomResourceOptions? options = null)
            : base("okta:inline/hook:Hook", name, args ?? new HookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hook(string name, Input<string> id, HookState? state = null, CustomResourceOptions? options = null)
            : base("okta:inline/hook:Hook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Hook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hook Get(string name, Input<string> id, HookState? state = null, CustomResourceOptions? options = null)
        {
            return new Hook(name, id, state, options);
        }
    }

    public sealed class HookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication required for inline hook request.
        /// </summary>
        [Input("auth")]
        public Input<Inputs.HookAuthArgs>? Auth { get; set; }

        /// <summary>
        /// Details of the endpoint the inline hook will hit.
        /// </summary>
        [Input("channel")]
        public Input<Inputs.HookChannelArgs>? Channel { get; set; }

        [Input("headers")]
        private InputList<Inputs.HookHeaderArgs>? _headers;

        /// <summary>
        /// Map of headers to send along in inline hook request.
        /// </summary>
        public InputList<Inputs.HookHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.HookHeaderArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// The inline hook display name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of hook to trigger. Currently only `"HTTP"` is supported.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The version of the endpoint.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public HookArgs()
        {
        }
    }

    public sealed class HookState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication required for inline hook request.
        /// </summary>
        [Input("auth")]
        public Input<Inputs.HookAuthGetArgs>? Auth { get; set; }

        /// <summary>
        /// Details of the endpoint the inline hook will hit.
        /// </summary>
        [Input("channel")]
        public Input<Inputs.HookChannelGetArgs>? Channel { get; set; }

        [Input("headers")]
        private InputList<Inputs.HookHeaderGetArgs>? _headers;

        /// <summary>
        /// Map of headers to send along in inline hook request.
        /// </summary>
        public InputList<Inputs.HookHeaderGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.HookHeaderGetArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// The inline hook display name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of hook to trigger. Currently only `"HTTP"` is supported.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The version of the endpoint.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public HookState()
        {
        }
    }
}
