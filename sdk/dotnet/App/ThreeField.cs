// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App
{
    /// <summary>
    /// Creates an Three Field Application.
    /// 
    /// This resource allows you to create and configure an Three Field Application.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_three_field.html.markdown.
    /// </summary>
    public partial class ThreeField : Pulumi.CustomResource
    {
        /// <summary>
        /// Custom error page URL.
        /// </summary>
        [Output("accessibilityErrorRedirectUrl")]
        public Output<string?> AccessibilityErrorRedirectUrl { get; private set; } = null!;

        /// <summary>
        /// Enable self service. By default it is `false`.
        /// </summary>
        [Output("accessibilitySelfService")]
        public Output<bool?> AccessibilitySelfService { get; private set; } = null!;

        /// <summary>
        /// Display auto submit toolbar.
        /// </summary>
        [Output("autoSubmitToolbar")]
        public Output<bool?> AutoSubmitToolbar { get; private set; } = null!;

        /// <summary>
        /// Login button field CSS selector.
        /// </summary>
        [Output("buttonSelector")]
        public Output<string> ButtonSelector { get; private set; } = null!;

        /// <summary>
        /// Extra field CSS selector.
        /// </summary>
        [Output("extraFieldSelector")]
        public Output<string> ExtraFieldSelector { get; private set; } = null!;

        /// <summary>
        /// Value for extra form field.
        /// </summary>
        [Output("extraFieldValue")]
        public Output<string> ExtraFieldValue { get; private set; } = null!;

        /// <summary>
        /// Groups associated with the application. See `okta.app.GroupAssignment` for a more flexible approach.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// Do not display application icon on mobile app.
        /// </summary>
        [Output("hideIos")]
        public Output<bool?> HideIos { get; private set; } = null!;

        /// <summary>
        /// Do not display application icon to users.
        /// </summary>
        [Output("hideWeb")]
        public Output<bool?> HideWeb { get; private set; } = null!;

        /// <summary>
        /// The display name of the Application.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// Name assigned to the application by Okta.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Login password field CSS selector.
        /// </summary>
        [Output("passwordSelector")]
        public Output<string> PasswordSelector { get; private set; } = null!;

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Output("signOnMode")]
        public Output<string> SignOnMode { get; private set; } = null!;

        /// <summary>
        /// Status of application. By default it is `"ACTIVE"`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Login URL.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// A regex that further restricts URL to the specified regex.
        /// </summary>
        [Output("urlRegex")]
        public Output<string?> UrlRegex { get; private set; } = null!;

        /// <summary>
        /// The default username assigned to each user.
        /// </summary>
        [Output("userNameTemplate")]
        public Output<string> UserNameTemplate { get; private set; } = null!;

        /// <summary>
        /// The Username template type.
        /// </summary>
        [Output("userNameTemplateType")]
        public Output<string> UserNameTemplateType { get; private set; } = null!;

        /// <summary>
        /// Login username field CSS selector.
        /// </summary>
        [Output("usernameSelector")]
        public Output<string> UsernameSelector { get; private set; } = null!;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.ThreeFieldUsers>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a ThreeField resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThreeField(string name, ThreeFieldArgs args, CustomResourceOptions? options = null)
            : base("okta:app/threeField:ThreeField", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ThreeField(string name, Input<string> id, ThreeFieldState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/threeField:ThreeField", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ThreeField resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThreeField Get(string name, Input<string> id, ThreeFieldState? state = null, CustomResourceOptions? options = null)
        {
            return new ThreeField(name, id, state, options);
        }
    }

    public sealed class ThreeFieldArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom error page URL.
        /// </summary>
        [Input("accessibilityErrorRedirectUrl")]
        public Input<string>? AccessibilityErrorRedirectUrl { get; set; }

        /// <summary>
        /// Enable self service. By default it is `false`.
        /// </summary>
        [Input("accessibilitySelfService")]
        public Input<bool>? AccessibilitySelfService { get; set; }

        /// <summary>
        /// Display auto submit toolbar.
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        /// <summary>
        /// Login button field CSS selector.
        /// </summary>
        [Input("buttonSelector", required: true)]
        public Input<string> ButtonSelector { get; set; } = null!;

        /// <summary>
        /// Extra field CSS selector.
        /// </summary>
        [Input("extraFieldSelector", required: true)]
        public Input<string> ExtraFieldSelector { get; set; } = null!;

        /// <summary>
        /// Value for extra form field.
        /// </summary>
        [Input("extraFieldValue", required: true)]
        public Input<string> ExtraFieldValue { get; set; } = null!;

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Groups associated with the application. See `okta.app.GroupAssignment` for a more flexible approach.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Do not display application icon on mobile app.
        /// </summary>
        [Input("hideIos")]
        public Input<bool>? HideIos { get; set; }

        /// <summary>
        /// Do not display application icon to users.
        /// </summary>
        [Input("hideWeb")]
        public Input<bool>? HideWeb { get; set; }

        /// <summary>
        /// The display name of the Application.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Login password field CSS selector.
        /// </summary>
        [Input("passwordSelector", required: true)]
        public Input<string> PasswordSelector { get; set; } = null!;

        /// <summary>
        /// Status of application. By default it is `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Login URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// A regex that further restricts URL to the specified regex.
        /// </summary>
        [Input("urlRegex")]
        public Input<string>? UrlRegex { get; set; }

        /// <summary>
        /// Login username field CSS selector.
        /// </summary>
        [Input("usernameSelector", required: true)]
        public Input<string> UsernameSelector { get; set; } = null!;

        [Input("users")]
        private InputList<Inputs.ThreeFieldUsersArgs>? _users;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        public InputList<Inputs.ThreeFieldUsersArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.ThreeFieldUsersArgs>());
            set => _users = value;
        }

        public ThreeFieldArgs()
        {
        }
    }

    public sealed class ThreeFieldState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom error page URL.
        /// </summary>
        [Input("accessibilityErrorRedirectUrl")]
        public Input<string>? AccessibilityErrorRedirectUrl { get; set; }

        /// <summary>
        /// Enable self service. By default it is `false`.
        /// </summary>
        [Input("accessibilitySelfService")]
        public Input<bool>? AccessibilitySelfService { get; set; }

        /// <summary>
        /// Display auto submit toolbar.
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        /// <summary>
        /// Login button field CSS selector.
        /// </summary>
        [Input("buttonSelector")]
        public Input<string>? ButtonSelector { get; set; }

        /// <summary>
        /// Extra field CSS selector.
        /// </summary>
        [Input("extraFieldSelector")]
        public Input<string>? ExtraFieldSelector { get; set; }

        /// <summary>
        /// Value for extra form field.
        /// </summary>
        [Input("extraFieldValue")]
        public Input<string>? ExtraFieldValue { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Groups associated with the application. See `okta.app.GroupAssignment` for a more flexible approach.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Do not display application icon on mobile app.
        /// </summary>
        [Input("hideIos")]
        public Input<bool>? HideIos { get; set; }

        /// <summary>
        /// Do not display application icon to users.
        /// </summary>
        [Input("hideWeb")]
        public Input<bool>? HideWeb { get; set; }

        /// <summary>
        /// The display name of the Application.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Name assigned to the application by Okta.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Login password field CSS selector.
        /// </summary>
        [Input("passwordSelector")]
        public Input<string>? PasswordSelector { get; set; }

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Input("signOnMode")]
        public Input<string>? SignOnMode { get; set; }

        /// <summary>
        /// Status of application. By default it is `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Login URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// A regex that further restricts URL to the specified regex.
        /// </summary>
        [Input("urlRegex")]
        public Input<string>? UrlRegex { get; set; }

        /// <summary>
        /// The default username assigned to each user.
        /// </summary>
        [Input("userNameTemplate")]
        public Input<string>? UserNameTemplate { get; set; }

        /// <summary>
        /// The Username template type.
        /// </summary>
        [Input("userNameTemplateType")]
        public Input<string>? UserNameTemplateType { get; set; }

        /// <summary>
        /// Login username field CSS selector.
        /// </summary>
        [Input("usernameSelector")]
        public Input<string>? UsernameSelector { get; set; }

        [Input("users")]
        private InputList<Inputs.ThreeFieldUsersGetArgs>? _users;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        public InputList<Inputs.ThreeFieldUsersGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.ThreeFieldUsersGetArgs>());
            set => _users = value;
        }

        public ThreeFieldState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ThreeFieldUsersArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public ThreeFieldUsersArgs()
        {
        }
    }

    public sealed class ThreeFieldUsersGetArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public ThreeFieldUsersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ThreeFieldUsers
    {
        public readonly string? Id;
        public readonly string? Password;
        public readonly string Scope;
        public readonly string? Username;

        [OutputConstructor]
        private ThreeFieldUsers(
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
