// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Idp
{
    /// <summary>
    /// Creates a SAML Identity Provider Signing Key.
    /// 
    /// This resource allows you to create and configure a SAML Identity Provider Signing Key.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/idp_saml_signing_key.html.markdown.
    /// </summary>
    public partial class SamlKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Date created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Date the cert expires.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// Key ID.
        /// </summary>
        [Output("kid")]
        public Output<string> Kid { get; private set; } = null!;

        /// <summary>
        /// Identifies the cryptographic algorithm family used with the key.
        /// </summary>
        [Output("kty")]
        public Output<string> Kty { get; private set; } = null!;

        /// <summary>
        /// Intended use of the public key.
        /// </summary>
        [Output("use")]
        public Output<string> Use { get; private set; } = null!;

        /// <summary>
        /// base64-encoded X.509 certificate chain with DER encoding.
        /// </summary>
        [Output("x5cs")]
        public Output<ImmutableArray<string>> X5cs { get; private set; } = null!;

        /// <summary>
        /// base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate.
        /// </summary>
        [Output("x5tS256")]
        public Output<string> X5tS256 { get; private set; } = null!;


        /// <summary>
        /// Create a SamlKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamlKey(string name, SamlKeyArgs args, CustomResourceOptions? options = null)
            : base("okta:idp/samlKey:SamlKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SamlKey(string name, Input<string> id, SamlKeyState? state = null, CustomResourceOptions? options = null)
            : base("okta:idp/samlKey:SamlKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SamlKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamlKey Get(string name, Input<string> id, SamlKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SamlKey(name, id, state, options);
        }
    }

    public sealed class SamlKeyArgs : Pulumi.ResourceArgs
    {
        [Input("x5cs", required: true)]
        private InputList<string>? _x5cs;

        /// <summary>
        /// base64-encoded X.509 certificate chain with DER encoding.
        /// </summary>
        public InputList<string> X5cs
        {
            get => _x5cs ?? (_x5cs = new InputList<string>());
            set => _x5cs = value;
        }

        public SamlKeyArgs()
        {
        }
    }

    public sealed class SamlKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Date the cert expires.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// Key ID.
        /// </summary>
        [Input("kid")]
        public Input<string>? Kid { get; set; }

        /// <summary>
        /// Identifies the cryptographic algorithm family used with the key.
        /// </summary>
        [Input("kty")]
        public Input<string>? Kty { get; set; }

        /// <summary>
        /// Intended use of the public key.
        /// </summary>
        [Input("use")]
        public Input<string>? Use { get; set; }

        [Input("x5cs")]
        private InputList<string>? _x5cs;

        /// <summary>
        /// base64-encoded X.509 certificate chain with DER encoding.
        /// </summary>
        public InputList<string> X5cs
        {
            get => _x5cs ?? (_x5cs = new InputList<string>());
            set => _x5cs = value;
        }

        /// <summary>
        /// base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate.
        /// </summary>
        [Input("x5tS256")]
        public Input<string>? X5tS256 { get; set; }

        public SamlKeyState()
        {
        }
    }
}
