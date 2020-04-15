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
    /// Creates an OIDC Application.
    /// 
    /// This resource allows you to create and configure an OIDC Application.
    /// </summary>
    public partial class OAuth : Pulumi.CustomResource
    {
        /// <summary>
        /// Requested key rotation mode.
        /// </summary>
        [Output("autoKeyRotation")]
        public Output<bool?> AutoKeyRotation { get; private set; } = null!;

        /// <summary>
        /// Display auto submit toolbar.
        /// </summary>
        [Output("autoSubmitToolbar")]
        public Output<bool?> AutoSubmitToolbar { get; private set; } = null!;

        /// <summary>
        /// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
        /// </summary>
        [Output("clientBasicSecret")]
        public Output<string?> ClientBasicSecret { get; private set; } = null!;

        /// <summary>
        /// The client ID of the application.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client secret of the application.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// URI to a web page providing information about the client.
        /// </summary>
        [Output("clientUri")]
        public Output<string?> ClientUri { get; private set; } = null!;

        /// <summary>
        /// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
        /// </summary>
        [Output("consentMethod")]
        public Output<string?> ConsentMethod { get; private set; } = null!;

        /// <summary>
        /// This property allows you to set the application's client id.
        /// </summary>
        [Output("customClientId")]
        public Output<string?> CustomClientId { get; private set; } = null!;

        /// <summary>
        /// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
        /// </summary>
        [Output("grantTypes")]
        public Output<ImmutableArray<string>> GrantTypes { get; private set; } = null!;

        /// <summary>
        /// The groups assigned to the application. It is recommended not to use this and instead use `okta.app.GroupAssignment`.
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
        /// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
        /// </summary>
        [Output("issuerMode")]
        public Output<string?> IssuerMode { get; private set; } = null!;

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// URI that initiates login.
        /// </summary>
        [Output("loginUri")]
        public Output<string?> LoginUri { get; private set; } = null!;

        /// <summary>
        /// URI that references a logo for the client.
        /// </summary>
        [Output("logoUri")]
        public Output<string?> LogoUri { get; private set; } = null!;

        /// <summary>
        /// Name assigned to the application by Okta.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This tells the provider not to persist the application's secret to state. If this is ever changes from true =&gt; false your app will be recreated.
        /// </summary>
        [Output("omitSecret")]
        public Output<bool?> OmitSecret { get; private set; } = null!;

        /// <summary>
        /// URI to web page providing client policy document.
        /// </summary>
        [Output("policyUri")]
        public Output<string?> PolicyUri { get; private set; } = null!;

        /// <summary>
        /// List of URIs for redirection after logout.
        /// </summary>
        [Output("postLogoutRedirectUris")]
        public Output<ImmutableArray<string>> PostLogoutRedirectUris { get; private set; } = null!;

        /// <summary>
        /// Custom JSON that represents an OAuth application's profile.
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// List of URIs for use in the redirect-based flow. This is required for all application types except service.
        /// </summary>
        [Output("redirectUris")]
        public Output<ImmutableArray<string>> RedirectUris { get; private set; } = null!;

        /// <summary>
        /// List of OAuth 2.0 response type strings.
        /// </summary>
        [Output("responseTypes")]
        public Output<ImmutableArray<string>> ResponseTypes { get; private set; } = null!;

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Output("signOnMode")]
        public Output<string> SignOnMode { get; private set; } = null!;

        /// <summary>
        /// The status of the application, by default it is `"ACTIVE"`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Requested authentication method for the token endpoint. It can be set to `"none"`, `"client_secret_post"`, `"client_secret_basic"`, `"client_secret_jwt"`.
        /// </summary>
        [Output("tokenEndpointAuthMethod")]
        public Output<string?> TokenEndpointAuthMethod { get; private set; } = null!;

        /// <summary>
        /// URI to web page providing client tos (terms of service).
        /// </summary>
        [Output("tosUri")]
        public Output<string?> TosUri { get; private set; } = null!;

        /// <summary>
        /// The type of OAuth application.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The users assigned to the application. It is recommended not to use this and instead use `okta.app.User`.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.OAuthUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a OAuth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OAuth(string name, OAuthArgs args, CustomResourceOptions? options = null)
            : base("okta:app/oAuth:OAuth", name, args ?? new OAuthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OAuth(string name, Input<string> id, OAuthState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/oAuth:OAuth", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OAuth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OAuth Get(string name, Input<string> id, OAuthState? state = null, CustomResourceOptions? options = null)
        {
            return new OAuth(name, id, state, options);
        }
    }

    public sealed class OAuthArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Requested key rotation mode.
        /// </summary>
        [Input("autoKeyRotation")]
        public Input<bool>? AutoKeyRotation { get; set; }

        /// <summary>
        /// Display auto submit toolbar.
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        /// <summary>
        /// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
        /// </summary>
        [Input("clientBasicSecret")]
        public Input<string>? ClientBasicSecret { get; set; }

        /// <summary>
        /// URI to a web page providing information about the client.
        /// </summary>
        [Input("clientUri")]
        public Input<string>? ClientUri { get; set; }

        /// <summary>
        /// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
        /// </summary>
        [Input("consentMethod")]
        public Input<string>? ConsentMethod { get; set; }

        /// <summary>
        /// This property allows you to set the application's client id.
        /// </summary>
        [Input("customClientId")]
        public Input<string>? CustomClientId { get; set; }

        [Input("grantTypes")]
        private InputList<string>? _grantTypes;

        /// <summary>
        /// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
        /// </summary>
        public InputList<string> GrantTypes
        {
            get => _grantTypes ?? (_grantTypes = new InputList<string>());
            set => _grantTypes = value;
        }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// The groups assigned to the application. It is recommended not to use this and instead use `okta.app.GroupAssignment`.
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
        /// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// URI that initiates login.
        /// </summary>
        [Input("loginUri")]
        public Input<string>? LoginUri { get; set; }

        /// <summary>
        /// URI that references a logo for the client.
        /// </summary>
        [Input("logoUri")]
        public Input<string>? LogoUri { get; set; }

        /// <summary>
        /// This tells the provider not to persist the application's secret to state. If this is ever changes from true =&gt; false your app will be recreated.
        /// </summary>
        [Input("omitSecret")]
        public Input<bool>? OmitSecret { get; set; }

        /// <summary>
        /// URI to web page providing client policy document.
        /// </summary>
        [Input("policyUri")]
        public Input<string>? PolicyUri { get; set; }

        [Input("postLogoutRedirectUris")]
        private InputList<string>? _postLogoutRedirectUris;

        /// <summary>
        /// List of URIs for redirection after logout.
        /// </summary>
        public InputList<string> PostLogoutRedirectUris
        {
            get => _postLogoutRedirectUris ?? (_postLogoutRedirectUris = new InputList<string>());
            set => _postLogoutRedirectUris = value;
        }

        /// <summary>
        /// Custom JSON that represents an OAuth application's profile.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// List of URIs for use in the redirect-based flow. This is required for all application types except service.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        [Input("responseTypes")]
        private InputList<string>? _responseTypes;

        /// <summary>
        /// List of OAuth 2.0 response type strings.
        /// </summary>
        public InputList<string> ResponseTypes
        {
            get => _responseTypes ?? (_responseTypes = new InputList<string>());
            set => _responseTypes = value;
        }

        /// <summary>
        /// The status of the application, by default it is `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Requested authentication method for the token endpoint. It can be set to `"none"`, `"client_secret_post"`, `"client_secret_basic"`, `"client_secret_jwt"`.
        /// </summary>
        [Input("tokenEndpointAuthMethod")]
        public Input<string>? TokenEndpointAuthMethod { get; set; }

        /// <summary>
        /// URI to web page providing client tos (terms of service).
        /// </summary>
        [Input("tosUri")]
        public Input<string>? TosUri { get; set; }

        /// <summary>
        /// The type of OAuth application.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("users")]
        private InputList<Inputs.OAuthUserArgs>? _users;

        /// <summary>
        /// The users assigned to the application. It is recommended not to use this and instead use `okta.app.User`.
        /// </summary>
        public InputList<Inputs.OAuthUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.OAuthUserArgs>());
            set => _users = value;
        }

        public OAuthArgs()
        {
        }
    }

    public sealed class OAuthState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Requested key rotation mode.
        /// </summary>
        [Input("autoKeyRotation")]
        public Input<bool>? AutoKeyRotation { get; set; }

        /// <summary>
        /// Display auto submit toolbar.
        /// </summary>
        [Input("autoSubmitToolbar")]
        public Input<bool>? AutoSubmitToolbar { get; set; }

        /// <summary>
        /// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
        /// </summary>
        [Input("clientBasicSecret")]
        public Input<string>? ClientBasicSecret { get; set; }

        /// <summary>
        /// The client ID of the application.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret of the application.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// URI to a web page providing information about the client.
        /// </summary>
        [Input("clientUri")]
        public Input<string>? ClientUri { get; set; }

        /// <summary>
        /// Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
        /// </summary>
        [Input("consentMethod")]
        public Input<string>? ConsentMethod { get; set; }

        /// <summary>
        /// This property allows you to set the application's client id.
        /// </summary>
        [Input("customClientId")]
        public Input<string>? CustomClientId { get; set; }

        [Input("grantTypes")]
        private InputList<string>? _grantTypes;

        /// <summary>
        /// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
        /// </summary>
        public InputList<string> GrantTypes
        {
            get => _grantTypes ?? (_grantTypes = new InputList<string>());
            set => _grantTypes = value;
        }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// The groups assigned to the application. It is recommended not to use this and instead use `okta.app.GroupAssignment`.
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
        /// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        /// <summary>
        /// The Application's display name.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// URI that initiates login.
        /// </summary>
        [Input("loginUri")]
        public Input<string>? LoginUri { get; set; }

        /// <summary>
        /// URI that references a logo for the client.
        /// </summary>
        [Input("logoUri")]
        public Input<string>? LogoUri { get; set; }

        /// <summary>
        /// Name assigned to the application by Okta.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This tells the provider not to persist the application's secret to state. If this is ever changes from true =&gt; false your app will be recreated.
        /// </summary>
        [Input("omitSecret")]
        public Input<bool>? OmitSecret { get; set; }

        /// <summary>
        /// URI to web page providing client policy document.
        /// </summary>
        [Input("policyUri")]
        public Input<string>? PolicyUri { get; set; }

        [Input("postLogoutRedirectUris")]
        private InputList<string>? _postLogoutRedirectUris;

        /// <summary>
        /// List of URIs for redirection after logout.
        /// </summary>
        public InputList<string> PostLogoutRedirectUris
        {
            get => _postLogoutRedirectUris ?? (_postLogoutRedirectUris = new InputList<string>());
            set => _postLogoutRedirectUris = value;
        }

        /// <summary>
        /// Custom JSON that represents an OAuth application's profile.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// List of URIs for use in the redirect-based flow. This is required for all application types except service.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        [Input("responseTypes")]
        private InputList<string>? _responseTypes;

        /// <summary>
        /// List of OAuth 2.0 response type strings.
        /// </summary>
        public InputList<string> ResponseTypes
        {
            get => _responseTypes ?? (_responseTypes = new InputList<string>());
            set => _responseTypes = value;
        }

        /// <summary>
        /// Sign on mode of application.
        /// </summary>
        [Input("signOnMode")]
        public Input<string>? SignOnMode { get; set; }

        /// <summary>
        /// The status of the application, by default it is `"ACTIVE"`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Requested authentication method for the token endpoint. It can be set to `"none"`, `"client_secret_post"`, `"client_secret_basic"`, `"client_secret_jwt"`.
        /// </summary>
        [Input("tokenEndpointAuthMethod")]
        public Input<string>? TokenEndpointAuthMethod { get; set; }

        /// <summary>
        /// URI to web page providing client tos (terms of service).
        /// </summary>
        [Input("tosUri")]
        public Input<string>? TosUri { get; set; }

        /// <summary>
        /// The type of OAuth application.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("users")]
        private InputList<Inputs.OAuthUserGetArgs>? _users;

        /// <summary>
        /// The users assigned to the application. It is recommended not to use this and instead use `okta.app.User`.
        /// </summary>
        public InputList<Inputs.OAuthUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.OAuthUserGetArgs>());
            set => _users = value;
        }

        public OAuthState()
        {
        }
    }
}
