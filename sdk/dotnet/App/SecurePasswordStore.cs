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
    /// Creates a Secure Password Store Application.
    /// 
    /// This resource allows you to create and configure a Secure Password Store Application.
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
    ///         var example = new Okta.App.SecurePasswordStore("example", new Okta.App.SecurePasswordStoreArgs
    ///         {
    ///             CredentialsScheme = "ADMIN_SETS_CREDENTIALS",
    ///             Label = "example",
    ///             PasswordField = "pass",
    ///             Url = "http://test.com",
    ///             UsernameField = "user",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SecurePasswordStore : Pulumi.CustomResource
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
        /// Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.
        /// </summary>
        [Output("credentialsScheme")]
        public Output<string?> CredentialsScheme { get; private set; } = null!;

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
        /// Name of optional param in the login form.
        /// </summary>
        [Output("optionalField1")]
        public Output<string?> OptionalField1 { get; private set; } = null!;

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Output("optionalField1Value")]
        public Output<string?> OptionalField1Value { get; private set; } = null!;

        /// <summary>
        /// Name of optional param in the login form.
        /// </summary>
        [Output("optionalField2")]
        public Output<string?> OptionalField2 { get; private set; } = null!;

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Output("optionalField2Value")]
        public Output<string?> OptionalField2Value { get; private set; } = null!;

        /// <summary>
        /// Name of optional param in the login form.
        /// </summary>
        [Output("optionalField3")]
        public Output<string?> OptionalField3 { get; private set; } = null!;

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Output("optionalField3Value")]
        public Output<string?> OptionalField3Value { get; private set; } = null!;

        /// <summary>
        /// Login password field.
        /// </summary>
        [Output("passwordField")]
        public Output<string> PasswordField { get; private set; } = null!;

        /// <summary>
        /// Allow user to reveal password.
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
        public Output<string> UsernameField { get; private set; } = null!;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.SecurePasswordStoreUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a SecurePasswordStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurePasswordStore(string name, SecurePasswordStoreArgs args, CustomResourceOptions? options = null)
            : base("okta:app/securePasswordStore:SecurePasswordStore", name, args ?? new SecurePasswordStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurePasswordStore(string name, Input<string> id, SecurePasswordStoreState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/securePasswordStore:SecurePasswordStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurePasswordStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurePasswordStore Get(string name, Input<string> id, SecurePasswordStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurePasswordStore(name, id, state, options);
        }
    }

    public sealed class SecurePasswordStoreArgs : Pulumi.ResourceArgs
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
        /// Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.
        /// </summary>
        [Input("credentialsScheme")]
        public Input<string>? CredentialsScheme { get; set; }

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
        /// Name of optional param in the login form.
        /// </summary>
        [Input("optionalField1")]
        public Input<string>? OptionalField1 { get; set; }

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Input("optionalField1Value")]
        public Input<string>? OptionalField1Value { get; set; }

        /// <summary>
        /// Name of optional param in the login form.
        /// </summary>
        [Input("optionalField2")]
        public Input<string>? OptionalField2 { get; set; }

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Input("optionalField2Value")]
        public Input<string>? OptionalField2Value { get; set; }

        /// <summary>
        /// Name of optional param in the login form.
        /// </summary>
        [Input("optionalField3")]
        public Input<string>? OptionalField3 { get; set; }

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Input("optionalField3Value")]
        public Input<string>? OptionalField3Value { get; set; }

        /// <summary>
        /// Login password field.
        /// </summary>
        [Input("passwordField", required: true)]
        public Input<string> PasswordField { get; set; } = null!;

        /// <summary>
        /// Allow user to reveal password.
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
        /// Login username field.
        /// </summary>
        [Input("usernameField", required: true)]
        public Input<string> UsernameField { get; set; } = null!;

        [Input("users")]
        private InputList<Inputs.SecurePasswordStoreUserArgs>? _users;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        public InputList<Inputs.SecurePasswordStoreUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SecurePasswordStoreUserArgs>());
            set => _users = value;
        }

        public SecurePasswordStoreArgs()
        {
        }
    }

    public sealed class SecurePasswordStoreState : Pulumi.ResourceArgs
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
        /// Application credentials scheme. Can be set to `"EDIT_USERNAME_AND_PASSWORD"`, `"ADMIN_SETS_CREDENTIALS"`, `"EDIT_PASSWORD_ONLY"`, `"EXTERNAL_PASSWORD_SYNC"`, or `"SHARED_USERNAME_AND_PASSWORD"`.
        /// </summary>
        [Input("credentialsScheme")]
        public Input<string>? CredentialsScheme { get; set; }

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
        /// Name of optional param in the login form.
        /// </summary>
        [Input("optionalField1")]
        public Input<string>? OptionalField1 { get; set; }

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Input("optionalField1Value")]
        public Input<string>? OptionalField1Value { get; set; }

        /// <summary>
        /// Name of optional param in the login form.
        /// </summary>
        [Input("optionalField2")]
        public Input<string>? OptionalField2 { get; set; }

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Input("optionalField2Value")]
        public Input<string>? OptionalField2Value { get; set; }

        /// <summary>
        /// Name of optional param in the login form.
        /// </summary>
        [Input("optionalField3")]
        public Input<string>? OptionalField3 { get; set; }

        /// <summary>
        /// Name of optional value in the login form.
        /// </summary>
        [Input("optionalField3Value")]
        public Input<string>? OptionalField3Value { get; set; }

        /// <summary>
        /// Login password field.
        /// </summary>
        [Input("passwordField")]
        public Input<string>? PasswordField { get; set; }

        /// <summary>
        /// Allow user to reveal password.
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
        private InputList<Inputs.SecurePasswordStoreUserGetArgs>? _users;

        /// <summary>
        /// The users assigned to the application. See `okta.app.User` for a more flexible approach.
        /// </summary>
        public InputList<Inputs.SecurePasswordStoreUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.SecurePasswordStoreUserGetArgs>());
            set => _users = value;
        }

        public SecurePasswordStoreState()
        {
        }
    }
}
