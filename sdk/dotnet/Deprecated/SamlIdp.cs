// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated
{
    public partial class SamlIdp : Pulumi.CustomResource
    {
        [Output("accountLinkAction")]
        public Output<string?> AccountLinkAction { get; private set; } = null!;

        [Output("accountLinkGroupIncludes")]
        public Output<ImmutableArray<string>> AccountLinkGroupIncludes { get; private set; } = null!;

        [Output("acsBinding")]
        public Output<string> AcsBinding { get; private set; } = null!;

        [Output("acsType")]
        public Output<string?> AcsType { get; private set; } = null!;

        [Output("audience")]
        public Output<string> Audience { get; private set; } = null!;

        [Output("deprovisionedAction")]
        public Output<string?> DeprovisionedAction { get; private set; } = null!;

        [Output("groupsAction")]
        public Output<string?> GroupsAction { get; private set; } = null!;

        [Output("groupsAssignments")]
        public Output<ImmutableArray<string>> GroupsAssignments { get; private set; } = null!;

        [Output("groupsAttribute")]
        public Output<string?> GroupsAttribute { get; private set; } = null!;

        [Output("groupsFilters")]
        public Output<ImmutableArray<string>> GroupsFilters { get; private set; } = null!;

        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
        /// </summary>
        [Output("issuerMode")]
        public Output<string?> IssuerMode { get; private set; } = null!;

        [Output("kid")]
        public Output<string> Kid { get; private set; } = null!;

        /// <summary>
        /// name of idp
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nameFormat")]
        public Output<string?> NameFormat { get; private set; } = null!;

        [Output("profileMaster")]
        public Output<bool?> ProfileMaster { get; private set; } = null!;

        [Output("provisioningAction")]
        public Output<string?> ProvisioningAction { get; private set; } = null!;

        /// <summary>
        /// algorithm to use to sign requests
        /// </summary>
        [Output("requestSignatureAlgorithm")]
        public Output<string?> RequestSignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// algorithm to use to sign response
        /// </summary>
        [Output("requestSignatureScope")]
        public Output<string?> RequestSignatureScope { get; private set; } = null!;

        /// <summary>
        /// algorithm to use to sign requests
        /// </summary>
        [Output("responseSignatureAlgorithm")]
        public Output<string?> ResponseSignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// algorithm to use to sign response
        /// </summary>
        [Output("responseSignatureScope")]
        public Output<string?> ResponseSignatureScope { get; private set; } = null!;

        [Output("ssoBinding")]
        public Output<string?> SsoBinding { get; private set; } = null!;

        [Output("ssoDestination")]
        public Output<string?> SsoDestination { get; private set; } = null!;

        [Output("ssoUrl")]
        public Output<string> SsoUrl { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("subjectFilter")]
        public Output<string?> SubjectFilter { get; private set; } = null!;

        [Output("subjectFormats")]
        public Output<ImmutableArray<string>> SubjectFormats { get; private set; } = null!;

        [Output("subjectMatchAttribute")]
        public Output<string?> SubjectMatchAttribute { get; private set; } = null!;

        [Output("subjectMatchType")]
        public Output<string?> SubjectMatchType { get; private set; } = null!;

        [Output("suspendedAction")]
        public Output<string?> SuspendedAction { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("usernameTemplate")]
        public Output<string?> UsernameTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a SamlIdp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamlIdp(string name, SamlIdpArgs args, CustomResourceOptions? options = null)
            : base("okta:deprecated/samlIdp:SamlIdp", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SamlIdp(string name, Input<string> id, SamlIdpState? state = null, CustomResourceOptions? options = null)
            : base("okta:deprecated/samlIdp:SamlIdp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SamlIdp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamlIdp Get(string name, Input<string> id, SamlIdpState? state = null, CustomResourceOptions? options = null)
        {
            return new SamlIdp(name, id, state, options);
        }
    }

    public sealed class SamlIdpArgs : Pulumi.ResourceArgs
    {
        [Input("accountLinkAction")]
        public Input<string>? AccountLinkAction { get; set; }

        [Input("accountLinkGroupIncludes")]
        private InputList<string>? _accountLinkGroupIncludes;
        public InputList<string> AccountLinkGroupIncludes
        {
            get => _accountLinkGroupIncludes ?? (_accountLinkGroupIncludes = new InputList<string>());
            set => _accountLinkGroupIncludes = value;
        }

        [Input("acsBinding", required: true)]
        public Input<string> AcsBinding { get; set; } = null!;

        [Input("acsType")]
        public Input<string>? AcsType { get; set; }

        [Input("deprovisionedAction")]
        public Input<string>? DeprovisionedAction { get; set; }

        [Input("groupsAction")]
        public Input<string>? GroupsAction { get; set; }

        [Input("groupsAssignments")]
        private InputList<string>? _groupsAssignments;
        public InputList<string> GroupsAssignments
        {
            get => _groupsAssignments ?? (_groupsAssignments = new InputList<string>());
            set => _groupsAssignments = value;
        }

        [Input("groupsAttribute")]
        public Input<string>? GroupsAttribute { get; set; }

        [Input("groupsFilters")]
        private InputList<string>? _groupsFilters;
        public InputList<string> GroupsFilters
        {
            get => _groupsFilters ?? (_groupsFilters = new InputList<string>());
            set => _groupsFilters = value;
        }

        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        [Input("kid", required: true)]
        public Input<string> Kid { get; set; } = null!;

        /// <summary>
        /// name of idp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameFormat")]
        public Input<string>? NameFormat { get; set; }

        [Input("profileMaster")]
        public Input<bool>? ProfileMaster { get; set; }

        [Input("provisioningAction")]
        public Input<string>? ProvisioningAction { get; set; }

        /// <summary>
        /// algorithm to use to sign requests
        /// </summary>
        [Input("requestSignatureAlgorithm")]
        public Input<string>? RequestSignatureAlgorithm { get; set; }

        /// <summary>
        /// algorithm to use to sign response
        /// </summary>
        [Input("requestSignatureScope")]
        public Input<string>? RequestSignatureScope { get; set; }

        /// <summary>
        /// algorithm to use to sign requests
        /// </summary>
        [Input("responseSignatureAlgorithm")]
        public Input<string>? ResponseSignatureAlgorithm { get; set; }

        /// <summary>
        /// algorithm to use to sign response
        /// </summary>
        [Input("responseSignatureScope")]
        public Input<string>? ResponseSignatureScope { get; set; }

        [Input("ssoBinding")]
        public Input<string>? SsoBinding { get; set; }

        [Input("ssoDestination")]
        public Input<string>? SsoDestination { get; set; }

        [Input("ssoUrl", required: true)]
        public Input<string> SsoUrl { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subjectFilter")]
        public Input<string>? SubjectFilter { get; set; }

        [Input("subjectFormats")]
        private InputList<string>? _subjectFormats;
        public InputList<string> SubjectFormats
        {
            get => _subjectFormats ?? (_subjectFormats = new InputList<string>());
            set => _subjectFormats = value;
        }

        [Input("subjectMatchAttribute")]
        public Input<string>? SubjectMatchAttribute { get; set; }

        [Input("subjectMatchType")]
        public Input<string>? SubjectMatchType { get; set; }

        [Input("suspendedAction")]
        public Input<string>? SuspendedAction { get; set; }

        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SamlIdpArgs()
        {
        }
    }

    public sealed class SamlIdpState : Pulumi.ResourceArgs
    {
        [Input("accountLinkAction")]
        public Input<string>? AccountLinkAction { get; set; }

        [Input("accountLinkGroupIncludes")]
        private InputList<string>? _accountLinkGroupIncludes;
        public InputList<string> AccountLinkGroupIncludes
        {
            get => _accountLinkGroupIncludes ?? (_accountLinkGroupIncludes = new InputList<string>());
            set => _accountLinkGroupIncludes = value;
        }

        [Input("acsBinding")]
        public Input<string>? AcsBinding { get; set; }

        [Input("acsType")]
        public Input<string>? AcsType { get; set; }

        [Input("audience")]
        public Input<string>? Audience { get; set; }

        [Input("deprovisionedAction")]
        public Input<string>? DeprovisionedAction { get; set; }

        [Input("groupsAction")]
        public Input<string>? GroupsAction { get; set; }

        [Input("groupsAssignments")]
        private InputList<string>? _groupsAssignments;
        public InputList<string> GroupsAssignments
        {
            get => _groupsAssignments ?? (_groupsAssignments = new InputList<string>());
            set => _groupsAssignments = value;
        }

        [Input("groupsAttribute")]
        public Input<string>? GroupsAttribute { get; set; }

        [Input("groupsFilters")]
        private InputList<string>? _groupsFilters;
        public InputList<string> GroupsFilters
        {
            get => _groupsFilters ?? (_groupsFilters = new InputList<string>());
            set => _groupsFilters = value;
        }

        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        [Input("kid")]
        public Input<string>? Kid { get; set; }

        /// <summary>
        /// name of idp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameFormat")]
        public Input<string>? NameFormat { get; set; }

        [Input("profileMaster")]
        public Input<bool>? ProfileMaster { get; set; }

        [Input("provisioningAction")]
        public Input<string>? ProvisioningAction { get; set; }

        /// <summary>
        /// algorithm to use to sign requests
        /// </summary>
        [Input("requestSignatureAlgorithm")]
        public Input<string>? RequestSignatureAlgorithm { get; set; }

        /// <summary>
        /// algorithm to use to sign response
        /// </summary>
        [Input("requestSignatureScope")]
        public Input<string>? RequestSignatureScope { get; set; }

        /// <summary>
        /// algorithm to use to sign requests
        /// </summary>
        [Input("responseSignatureAlgorithm")]
        public Input<string>? ResponseSignatureAlgorithm { get; set; }

        /// <summary>
        /// algorithm to use to sign response
        /// </summary>
        [Input("responseSignatureScope")]
        public Input<string>? ResponseSignatureScope { get; set; }

        [Input("ssoBinding")]
        public Input<string>? SsoBinding { get; set; }

        [Input("ssoDestination")]
        public Input<string>? SsoDestination { get; set; }

        [Input("ssoUrl")]
        public Input<string>? SsoUrl { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subjectFilter")]
        public Input<string>? SubjectFilter { get; set; }

        [Input("subjectFormats")]
        private InputList<string>? _subjectFormats;
        public InputList<string> SubjectFormats
        {
            get => _subjectFormats ?? (_subjectFormats = new InputList<string>());
            set => _subjectFormats = value;
        }

        [Input("subjectMatchAttribute")]
        public Input<string>? SubjectMatchAttribute { get; set; }

        [Input("subjectMatchType")]
        public Input<string>? SubjectMatchType { get; set; }

        [Input("suspendedAction")]
        public Input<string>? SuspendedAction { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SamlIdpState()
        {
        }
    }
}
