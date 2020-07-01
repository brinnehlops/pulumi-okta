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
    /// Creates an SWA Application.
    /// 
    /// This resource allows you to create and configure an SWA Application.
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
    ///         var example = new Okta.App.Swa("example", new Okta.App.SwaArgs
    ///         {
    ///             ButtonField = "btn-login",
    ///             Label = "example",
    ///             PasswordField = "txtbox-password",
    ///             Url = "https://example.com/login.html",
    ///             UsernameField = "txtbox-username",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Swa : Pulumi.CustomResource
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
        /// Login button field.
        /// </summary>
        [Output("buttonField")]
        public Output<string?> ButtonField { get; private set; } = null!;

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
        /// Login password field.
        /// </summary>
        [Output("passwordField")]
        public Output<string?> PasswordField { get; private set; } = null!;

        /// <summary>
        /// name of application from the Okta Integration Network, if not included a custom app will be created.
        /// </summary>
        [Output("preconfiguredApp")]
        public Output<string?> PreconfiguredApp { get; private set; } = null!;

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
        public Output<string?> Url { get; private set; } = null!;

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
        /// Login username field.
        /// </summary>
        [Output("usernameField")]
        public Output<string?> UsernameField { get; private set; } = null!;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.SwaUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Swa resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Swa(string name, SwaArgs args, CustomResourceOptions? options = null)
            : base("okta:app/swa:Swa", name, args ?? new SwaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Swa(string name, Input<string> id, SwaState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/swa:Swa", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Swa resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Swa Get(string name, Input<string> id, SwaState? state = null, CustomResourceOptions? options = null)
        {
            return new Swa(name, id, state, options);
        }
    }

    public sealed class SwaArgs : Pulumi.ResourceArgs
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
        /// Login button field.
        /// </summary>
        [Input("buttonField")]
        public Input<string>? ButtonField { get; set; }

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
        /// Login password field.
        /// </summary>
        [Input("passwordField")]
        public Input<string>? PasswordField { get; set; }

        /// <summary>
        /// name of application from the Okta Integration Network, if not included a custom app will be created.
        /// </summary>
        [Input("preconfiguredApp")]
        public Input<string>? PreconfiguredApp { get; set; }

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
        /// Login username field.
        /// </summary>
        [Input("usernameField")]
        public Input<string>? UsernameField { get; set; }

        [Input("users")]
        private InputList<Inputs.SwaUserArgs>? _users;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        public InputList<Inputs.SwaUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SwaUserArgs>());
            set => _users = value;
        }

        public SwaArgs()
        {
        }
    }

    public sealed class SwaState : Pulumi.ResourceArgs
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
        /// Login button field.
        /// </summary>
        [Input("buttonField")]
        public Input<string>? ButtonField { get; set; }

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
        /// Login password field.
        /// </summary>
        [Input("passwordField")]
        public Input<string>? PasswordField { get; set; }

        /// <summary>
        /// name of application from the Okta Integration Network, if not included a custom app will be created.
        /// </summary>
        [Input("preconfiguredApp")]
        public Input<string>? PreconfiguredApp { get; set; }

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
        /// Login username field.
        /// </summary>
        [Input("usernameField")]
        public Input<string>? UsernameField { get; set; }

        [Input("users")]
        private InputList<Inputs.SwaUserGetArgs>? _users;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        public InputList<Inputs.SwaUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SwaUserGetArgs>());
            set => _users = value;
        }

        public SwaState()
        {
        }
    }
}
