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
    /// Creates an Auto Login Okta Application.
    /// 
    /// This resource allows you to create and configure an Auto Login Okta Application.
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
    ///         var example = new Okta.App.AutoLogin("example", new Okta.App.AutoLoginArgs
    ///         {
    ///             CredentialsScheme = "EDIT_USERNAME_AND_PASSWORD",
    ///             Label = "Example App",
    ///             RevealPassword = true,
    ///             SignOnRedirectUrl = "https://example.com",
    ///             SignOnUrl = "https://example.com/login.html",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AutoLogin : Pulumi.CustomResource
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
        /// Application credentials scheme
        /// </summary>
        [Output("credentialsScheme")]
        public Output<string?> CredentialsScheme { get; private set; } = null!;

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
        /// The Application's display name.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// Name assigned to the application by Okta.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
        /// </summary>
        [Output("preconfiguredApp")]
        public Output<string?> PreconfiguredApp { get; private set; } = null!;

        /// <summary>
        /// Allow user to reveal password
        /// </summary>
        [Output("revealPassword")]
        public Output<bool?> RevealPassword { get; private set; } = null!;

        /// <summary>
        /// Shared password, required for certain schemes.
        /// </summary>
        [Output("sharedPassword")]
        public Output<string?> SharedPassword { get; private set; } = null!;

        /// <summary>
        /// Shared username, required for certain schemes.
        /// </summary>
        [Output("sharedUsername")]
        public Output<string?> SharedUsername { get; private set; } = null!;

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Output("signOnMode")]
        public Output<string> SignOnMode { get; private set; } = null!;

        /// <summary>
        /// Post login redirect URL
        /// </summary>
        [Output("signOnRedirectUrl")]
        public Output<string?> SignOnRedirectUrl { get; private set; } = null!;

        /// <summary>
        /// Login URL
        /// </summary>
        [Output("signOnUrl")]
        public Output<string?> SignOnUrl { get; private set; } = null!;

        /// <summary>
        /// The status of the application, by default it is `"ACTIVE"`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

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
        /// Users associated with the application
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.AutoLoginUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a AutoLogin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoLogin(string name, AutoLoginArgs args, CustomResourceOptions? options = null)
            : base("okta:app/autoLogin:AutoLogin", name, args ?? new AutoLoginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoLogin(string name, Input<string> id, AutoLoginState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/autoLogin:AutoLogin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoLogin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoLogin Get(string name, Input<string> id, AutoLoginState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoLogin(name, id, state, options);
        }
    }

    public sealed class AutoLoginArgs : Pulumi.ResourceArgs
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
        /// Application credentials scheme
        /// </summary>
        [Input("credentialsScheme")]
        public Input<string>? CredentialsScheme { get; set; }

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
        /// The Application's display name.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
        /// </summary>
        [Input("preconfiguredApp")]
        public Input<string>? PreconfiguredApp { get; set; }

        /// <summary>
        /// Allow user to reveal password
        /// </summary>
        [Input("revealPassword")]
        public Input<bool>? RevealPassword { get; set; }

        /// <summary>
        /// Shared password, required for certain schemes.
        /// </summary>
        [Input("sharedPassword")]
        public Input<string>? SharedPassword { get; set; }

        /// <summary>
        /// Shared username, required for certain schemes.
        /// </summary>
        [Input("sharedUsername")]
        public Input<string>? SharedUsername { get; set; }

        /// <summary>
        /// Post login redirect URL
        /// </summary>
        [Input("signOnRedirectUrl")]
        public Input<string>? SignOnRedirectUrl { get; set; }

        /// <summary>
        /// Login URL
        /// </summary>
        [Input("signOnUrl")]
        public Input<string>? SignOnUrl { get; set; }

        /// <summary>
        /// The status of the application, by default it is `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("users")]
        private InputList<Inputs.AutoLoginUserArgs>? _users;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        public InputList<Inputs.AutoLoginUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.AutoLoginUserArgs>());
            set => _users = value;
        }

        public AutoLoginArgs()
        {
        }
    }

    public sealed class AutoLoginState : Pulumi.ResourceArgs
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
        /// Application credentials scheme
        /// </summary>
        [Input("credentialsScheme")]
        public Input<string>? CredentialsScheme { get; set; }

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
        /// The Application's display name.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Name assigned to the application by Okta.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Tells Okta to use an existing application in their application catalog, as opposed to a custom application.
        /// </summary>
        [Input("preconfiguredApp")]
        public Input<string>? PreconfiguredApp { get; set; }

        /// <summary>
        /// Allow user to reveal password
        /// </summary>
        [Input("revealPassword")]
        public Input<bool>? RevealPassword { get; set; }

        /// <summary>
        /// Shared password, required for certain schemes.
        /// </summary>
        [Input("sharedPassword")]
        public Input<string>? SharedPassword { get; set; }

        /// <summary>
        /// Shared username, required for certain schemes.
        /// </summary>
        [Input("sharedUsername")]
        public Input<string>? SharedUsername { get; set; }

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Input("signOnMode")]
        public Input<string>? SignOnMode { get; set; }

        /// <summary>
        /// Post login redirect URL
        /// </summary>
        [Input("signOnRedirectUrl")]
        public Input<string>? SignOnRedirectUrl { get; set; }

        /// <summary>
        /// Login URL
        /// </summary>
        [Input("signOnUrl")]
        public Input<string>? SignOnUrl { get; set; }

        /// <summary>
        /// The status of the application, by default it is `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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

        [Input("users")]
        private InputList<Inputs.AutoLoginUserGetArgs>? _users;

        /// <summary>
        /// Users associated with the application
        /// </summary>
        public InputList<Inputs.AutoLoginUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.AutoLoginUserGetArgs>());
            set => _users = value;
        }

        public AutoLoginState()
        {
        }
    }
}
