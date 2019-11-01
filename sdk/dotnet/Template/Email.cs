// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Template
{
    /// <summary>
    /// Creates an Okta Email Template.
    /// 
    /// This resource allows you to create and configure an Okta Email Template.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/template_email.html.markdown.
    /// </summary>
    public partial class Email : Pulumi.CustomResource
    {
        /// <summary>
        /// The default language, by default is set to `"en"`.
        /// </summary>
        [Output("defaultLanguage")]
        public Output<string?> DefaultLanguage { get; private set; } = null!;

        /// <summary>
        /// Set of translations for particular template.
        /// </summary>
        [Output("translations")]
        public Output<ImmutableArray<Outputs.EmailTranslations>> Translations { get; private set; } = null!;

        /// <summary>
        /// Email template type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Email resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Email(string name, EmailArgs args, CustomResourceOptions? options = null)
            : base("okta:template/email:Email", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Email(string name, Input<string> id, EmailState? state = null, CustomResourceOptions? options = null)
            : base("okta:template/email:Email", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Email resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Email Get(string name, Input<string> id, EmailState? state = null, CustomResourceOptions? options = null)
        {
            return new Email(name, id, state, options);
        }
    }

    public sealed class EmailArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default language, by default is set to `"en"`.
        /// </summary>
        [Input("defaultLanguage")]
        public Input<string>? DefaultLanguage { get; set; }

        [Input("translations", required: true)]
        private InputList<Inputs.EmailTranslationsArgs>? _translations;

        /// <summary>
        /// Set of translations for particular template.
        /// </summary>
        public InputList<Inputs.EmailTranslationsArgs> Translations
        {
            get => _translations ?? (_translations = new InputList<Inputs.EmailTranslationsArgs>());
            set => _translations = value;
        }

        /// <summary>
        /// Email template type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public EmailArgs()
        {
        }
    }

    public sealed class EmailState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default language, by default is set to `"en"`.
        /// </summary>
        [Input("defaultLanguage")]
        public Input<string>? DefaultLanguage { get; set; }

        [Input("translations")]
        private InputList<Inputs.EmailTranslationsGetArgs>? _translations;

        /// <summary>
        /// Set of translations for particular template.
        /// </summary>
        public InputList<Inputs.EmailTranslationsGetArgs> Translations
        {
            get => _translations ?? (_translations = new InputList<Inputs.EmailTranslationsGetArgs>());
            set => _translations = value;
        }

        /// <summary>
        /// Email template type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EmailState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EmailTranslationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The language to map tthe template to.
        /// </summary>
        [Input("language", required: true)]
        public Input<string> Language { get; set; } = null!;

        /// <summary>
        /// The email subject line.
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        /// <summary>
        /// The email body.
        /// </summary>
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public EmailTranslationsArgs()
        {
        }
    }

    public sealed class EmailTranslationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The language to map tthe template to.
        /// </summary>
        [Input("language", required: true)]
        public Input<string> Language { get; set; } = null!;

        /// <summary>
        /// The email subject line.
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        /// <summary>
        /// The email body.
        /// </summary>
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public EmailTranslationsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EmailTranslations
    {
        /// <summary>
        /// The language to map tthe template to.
        /// </summary>
        public readonly string Language;
        /// <summary>
        /// The email subject line.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// The email body.
        /// </summary>
        public readonly string Template;

        [OutputConstructor]
        private EmailTranslations(
            string language,
            string subject,
            string template)
        {
            Language = language;
            Subject = subject;
            Template = template;
        }
    }
    }
}