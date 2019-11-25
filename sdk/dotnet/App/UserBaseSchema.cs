// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App
{
    /// <summary>
    /// Manages an Application User Base Schema property.
    /// 
    /// This resource allows you to configure a base app user schema property.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_user_base_schema.html.markdown.
    /// </summary>
    public partial class UserBaseSchema : Pulumi.CustomResource
    {
        /// <summary>
        /// The Application's ID the user schema property should be assigned to.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// The property name.
        /// </summary>
        [Output("index")]
        public Output<string> Index { get; private set; } = null!;

        /// <summary>
        /// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        /// </summary>
        [Output("master")]
        public Output<string?> Master { get; private set; } = null!;

        /// <summary>
        /// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        /// </summary>
        [Output("permissions")]
        public Output<string?> Permissions { get; private set; } = null!;

        /// <summary>
        /// Whether the property is required for this application's users.
        /// </summary>
        [Output("required")]
        public Output<bool?> Required { get; private set; } = null!;

        /// <summary>
        /// The property display name.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a UserBaseSchema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserBaseSchema(string name, UserBaseSchemaArgs args, CustomResourceOptions? options = null)
            : base("okta:app/userBaseSchema:UserBaseSchema", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private UserBaseSchema(string name, Input<string> id, UserBaseSchemaState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/userBaseSchema:UserBaseSchema", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserBaseSchema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserBaseSchema Get(string name, Input<string> id, UserBaseSchemaState? state = null, CustomResourceOptions? options = null)
        {
            return new UserBaseSchema(name, id, state, options);
        }
    }

    public sealed class UserBaseSchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Application's ID the user schema property should be assigned to.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// The property name.
        /// </summary>
        [Input("index", required: true)]
        public Input<string> Index { get; set; } = null!;

        /// <summary>
        /// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        /// </summary>
        [Input("master")]
        public Input<string>? Master { get; set; }

        /// <summary>
        /// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        /// </summary>
        [Input("permissions")]
        public Input<string>? Permissions { get; set; }

        /// <summary>
        /// Whether the property is required for this application's users.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The property display name.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public UserBaseSchemaArgs()
        {
        }
    }

    public sealed class UserBaseSchemaState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Application's ID the user schema property should be assigned to.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// The property name.
        /// </summary>
        [Input("index")]
        public Input<string>? Index { get; set; }

        /// <summary>
        /// Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        /// </summary>
        [Input("master")]
        public Input<string>? Master { get; set; }

        /// <summary>
        /// Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        /// </summary>
        [Input("permissions")]
        public Input<string>? Permissions { get; set; }

        /// <summary>
        /// Whether the property is required for this application's users.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The property display name.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public UserBaseSchemaState()
        {
        }
    }
}
