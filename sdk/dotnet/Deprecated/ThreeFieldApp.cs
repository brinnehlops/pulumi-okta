// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated
{
    public partial class ThreeFieldApp : Pulumi.CustomResource
    {
        /// <summary>
        /// Custom error page URL
        /// </summary>
        [Output("accessibilityErrorRedirectUrl")]
        public Output<string?> AccessibilityErrorRedirectUrl { get; private set; } = null!;

        /// <summary>
        /// Enable self service
        /// </summary>
        [Output("accessibilitySelfService")]
        public Output<bool?> AccessibilitySelfService { get; private set; } = null!;

        /// <summary>
        /// Display auto submit toolbar
        /// </summary>
        [Output("autoSubmitToolbar")]
        public Output<bool?> AutoSubmitToolbar { get; private set; } = null!;

        /// <summary>
        /// Login button field CSS selector
        /// </summary>
        [Output("buttonSelector")]
        public Output<string> ButtonSelector { get; private set; } = null!;

        /// <summary>
        /// Extra field CSS selector
        /// </summary>
        [Output("extraFieldSelector")]
        public Output<string> ExtraFieldSelector { get; private set; } = null!;

        /// <summary>
        /// Value for extra form field
        /// </summary>
        [Output("extraFieldValue")]
        public Output<string> ExtraFieldValue { get; private set; } = null!;

        /// <summary>
        /// Groups associated with the application
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// Do not display application icon on mobile app
        /// </summary>
        [Output("hideIos")]
        public Output<bool?> HideIos { get; private set; } = null!;

        /// <summary>
        /// Do not display application icon to users
        /// </summary>
        [Output("hideWeb")]
        public Output<bool?> HideWeb { get; private set; } = null!;

        /// <summary>
        /// Pretty name of app.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// name of app.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Login password field CSS selector
        /// </summary>
        [Output("passwordSelector")]
        public Output<string> PasswordSelector { get; private set; } = null!;

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Output("signOnMode")]
        public Output<string> SignOnMode { get; private set; } = null!;

        /// <summary>
        /// Status of application.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Login URL
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// A regex that further restricts URL to the specified regex
        /// </summary>
        [Output("urlRegex")]
        public Output<string?> UrlRegex { get; private set; } = null!;

        /// <summary>
        /// Username template
        /// </summary>
        [Output("userNameTemplate")]
        public Output<string> UserNameTemplate { get; private set; } = null!;

        /// <summary>
        /// Username template type
        /// </summary>
        [Output("userNameTemplateType")]
        public Output<string> UserNameTemplateType { get; private set; } = null!;

        /// <summary>
        /// Login username field CSS selector
        /// </summary>
        [Output("usernameSelector")]
        public Output<string> UsernameSelector { get; private set; } = null!;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.ThreeFieldAppUsers>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a ThreeFieldApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThreeFieldApp(string name, ThreeFieldAppArgs args, CustomResourceOptions? options = null)
            : base("okta:deprecated/threeFieldApp:ThreeFieldApp", name, args, MakeResourceOptions(options, ""))
        {
        }

        private ThreeFieldApp(string name, Input<string> id, ThreeFieldAppState? state = null, CustomResourceOptions? options = null)
            : base("okta:deprecated/threeFieldApp:ThreeFieldApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ThreeFieldApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThreeFieldApp Get(string name, Input<string> id, ThreeFieldAppState? state = null, CustomResourceOptions? options = null)
        {
            return new ThreeFieldApp(name, id, state, options);
        }
    }

    public sealed class ThreeFieldAppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom error page URL
        /// </summary>
        [Input("accessibilityErrorRedirectUrl")]
        public Input<string>? AccessibilityErrorRedirectUrl { get; set; }

        /// <summary>
        /// Enable self service
        /// </summary>
        [Input("accessibilitySelfService")]
        public Input<bool>? AccessibilitySelfService { get; set; }

        /// <summary>
        /// Display auto submit toolbar
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        /// <summary>
        /// Login button field CSS selector
        /// </summary>
        [Input("buttonSelector", required: true)]
        public Input<string> ButtonSelector { get; set; } = null!;

        /// <summary>
        /// Extra field CSS selector
        /// </summary>
        [Input("extraFieldSelector", required: true)]
        public Input<string> ExtraFieldSelector { get; set; } = null!;

        /// <summary>
        /// Value for extra form field
        /// </summary>
        [Input("extraFieldValue", required: true)]
        public Input<string> ExtraFieldValue { get; set; } = null!;

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Groups associated with the application
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Do not display application icon on mobile app
        /// </summary>
        [Input("hideIos")]
        public Input<bool>? HideIos { get; set; }

        /// <summary>
        /// Do not display application icon to users
        /// </summary>
        [Input("hideWeb")]
        public Input<bool>? HideWeb { get; set; }

        /// <summary>
        /// Pretty name of app.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Login password field CSS selector
        /// </summary>
        [Input("passwordSelector", required: true)]
        public Input<string> PasswordSelector { get; set; } = null!;

        /// <summary>
        /// Status of application.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Login URL
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// A regex that further restricts URL to the specified regex
        /// </summary>
        [Input("urlRegex")]
        public Input<string>? UrlRegex { get; set; }

        /// <summary>
        /// Login username field CSS selector
        /// </summary>
        [Input("usernameSelector", required: true)]
        public Input<string> UsernameSelector { get; set; } = null!;

        [Input("users")]
        private InputList<Inputs.ThreeFieldAppUsersArgs>? _users;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        public InputList<Inputs.ThreeFieldAppUsersArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.ThreeFieldAppUsersArgs>());
            set => _users = value;
        }

        public ThreeFieldAppArgs()
        {
        }
    }

    public sealed class ThreeFieldAppState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom error page URL
        /// </summary>
        [Input("accessibilityErrorRedirectUrl")]
        public Input<string>? AccessibilityErrorRedirectUrl { get; set; }

        /// <summary>
        /// Enable self service
        /// </summary>
        [Input("accessibilitySelfService")]
        public Input<bool>? AccessibilitySelfService { get; set; }

        /// <summary>
        /// Display auto submit toolbar
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        /// <summary>
        /// Login button field CSS selector
        /// </summary>
        [Input("buttonSelector")]
        public Input<string>? ButtonSelector { get; set; }

        /// <summary>
        /// Extra field CSS selector
        /// </summary>
        [Input("extraFieldSelector")]
        public Input<string>? ExtraFieldSelector { get; set; }

        /// <summary>
        /// Value for extra form field
        /// </summary>
        [Input("extraFieldValue")]
        public Input<string>? ExtraFieldValue { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Groups associated with the application
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Do not display application icon on mobile app
        /// </summary>
        [Input("hideIos")]
        public Input<bool>? HideIos { get; set; }

        /// <summary>
        /// Do not display application icon to users
        /// </summary>
        [Input("hideWeb")]
        public Input<bool>? HideWeb { get; set; }

        /// <summary>
        /// Pretty name of app.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// name of app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Login password field CSS selector
        /// </summary>
        [Input("passwordSelector")]
        public Input<string>? PasswordSelector { get; set; }

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Input("signOnMode")]
        public Input<string>? SignOnMode { get; set; }

        /// <summary>
        /// Status of application.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Login URL
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// A regex that further restricts URL to the specified regex
        /// </summary>
        [Input("urlRegex")]
        public Input<string>? UrlRegex { get; set; }

        /// <summary>
        /// Username template
        /// </summary>
        [Input("userNameTemplate")]
        public Input<string>? UserNameTemplate { get; set; }

        /// <summary>
        /// Username template type
        /// </summary>
        [Input("userNameTemplateType")]
        public Input<string>? UserNameTemplateType { get; set; }

        /// <summary>
        /// Login username field CSS selector
        /// </summary>
        [Input("usernameSelector")]
        public Input<string>? UsernameSelector { get; set; }

        [Input("users")]
        private InputList<Inputs.ThreeFieldAppUsersGetArgs>? _users;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        public InputList<Inputs.ThreeFieldAppUsersGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.ThreeFieldAppUsersGetArgs>());
            set => _users = value;
        }

        public ThreeFieldAppState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ThreeFieldAppUsersArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public ThreeFieldAppUsersArgs()
        {
        }
    }

    public sealed class ThreeFieldAppUsersGetArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public ThreeFieldAppUsersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ThreeFieldAppUsers
    {
        public readonly string? Id;
        public readonly string? Password;
        public readonly string Scope;
        public readonly string? Username;

        [OutputConstructor]
        private ThreeFieldAppUsers(
            string? id,
            string? password,
            string scope,
            string? username)
        {
            Id = id;
            Password = password;
            Scope = scope;
            Username = username;
        }
    }
    }
}