// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated
{
    public partial class MfaPolicyRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Should the user be enrolled the first time they LOGIN, the next time they are CHALLENGEd, or NEVER?
        /// </summary>
        [Output("enroll")]
        public Output<string?> Enroll { get; private set; } = null!;

        /// <summary>
        /// Policy Rule Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
        /// </summary>
        [Output("networkConnection")]
        public Output<string?> NetworkConnection { get; private set; } = null!;

        /// <summary>
        /// The zones to exclude
        /// </summary>
        [Output("networkExcludes")]
        public Output<ImmutableArray<string>> NetworkExcludes { get; private set; } = null!;

        /// <summary>
        /// The zones to include
        /// </summary>
        [Output("networkIncludes")]
        public Output<ImmutableArray<string>> NetworkIncludes { get; private set; } = null!;

        /// <summary>
        /// Policy ID of the Rule
        /// </summary>
        [Output("policyid")]
        public Output<string> Policyid { get; private set; } = null!;

        /// <summary>
        /// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we
        /// error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Policy Rule Status: ACTIVE or INACTIVE.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Set of User IDs to Exclude
        /// </summary>
        [Output("usersExcludeds")]
        public Output<ImmutableArray<string>> UsersExcludeds { get; private set; } = null!;


        /// <summary>
        /// Create a MfaPolicyRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaPolicyRule(string name, MfaPolicyRuleArgs args, CustomResourceOptions? options = null)
            : base("okta:deprecated/mfaPolicyRule:MfaPolicyRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private MfaPolicyRule(string name, Input<string> id, MfaPolicyRuleState? state = null, CustomResourceOptions? options = null)
            : base("okta:deprecated/mfaPolicyRule:MfaPolicyRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MfaPolicyRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MfaPolicyRule Get(string name, Input<string> id, MfaPolicyRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new MfaPolicyRule(name, id, state, options);
        }
    }

    public sealed class MfaPolicyRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the user be enrolled the first time they LOGIN, the next time they are CHALLENGEd, or NEVER?
        /// </summary>
        [Input("enroll")]
        public Input<string>? Enroll { get; set; }

        /// <summary>
        /// Policy Rule Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
        /// </summary>
        [Input("networkConnection")]
        public Input<string>? NetworkConnection { get; set; }

        [Input("networkExcludes")]
        private InputList<string>? _networkExcludes;

        /// <summary>
        /// The zones to exclude
        /// </summary>
        public InputList<string> NetworkExcludes
        {
            get => _networkExcludes ?? (_networkExcludes = new InputList<string>());
            set => _networkExcludes = value;
        }

        [Input("networkIncludes")]
        private InputList<string>? _networkIncludes;

        /// <summary>
        /// The zones to include
        /// </summary>
        public InputList<string> NetworkIncludes
        {
            get => _networkIncludes ?? (_networkIncludes = new InputList<string>());
            set => _networkIncludes = value;
        }

        /// <summary>
        /// Policy ID of the Rule
        /// </summary>
        [Input("policyid", required: true)]
        public Input<string> Policyid { get; set; } = null!;

        /// <summary>
        /// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we
        /// error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Policy Rule Status: ACTIVE or INACTIVE.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("usersExcludeds")]
        private InputList<string>? _usersExcludeds;

        /// <summary>
        /// Set of User IDs to Exclude
        /// </summary>
        public InputList<string> UsersExcludeds
        {
            get => _usersExcludeds ?? (_usersExcludeds = new InputList<string>());
            set => _usersExcludeds = value;
        }

        public MfaPolicyRuleArgs()
        {
        }
    }

    public sealed class MfaPolicyRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the user be enrolled the first time they LOGIN, the next time they are CHALLENGEd, or NEVER?
        /// </summary>
        [Input("enroll")]
        public Input<string>? Enroll { get; set; }

        /// <summary>
        /// Policy Rule Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
        /// </summary>
        [Input("networkConnection")]
        public Input<string>? NetworkConnection { get; set; }

        [Input("networkExcludes")]
        private InputList<string>? _networkExcludes;

        /// <summary>
        /// The zones to exclude
        /// </summary>
        public InputList<string> NetworkExcludes
        {
            get => _networkExcludes ?? (_networkExcludes = new InputList<string>());
            set => _networkExcludes = value;
        }

        [Input("networkIncludes")]
        private InputList<string>? _networkIncludes;

        /// <summary>
        /// The zones to include
        /// </summary>
        public InputList<string> NetworkIncludes
        {
            get => _networkIncludes ?? (_networkIncludes = new InputList<string>());
            set => _networkIncludes = value;
        }

        /// <summary>
        /// Policy ID of the Rule
        /// </summary>
        [Input("policyid")]
        public Input<string>? Policyid { get; set; }

        /// <summary>
        /// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we
        /// error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Policy Rule Status: ACTIVE or INACTIVE.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("usersExcludeds")]
        private InputList<string>? _usersExcludeds;

        /// <summary>
        /// Set of User IDs to Exclude
        /// </summary>
        public InputList<string> UsersExcludeds
        {
            get => _usersExcludeds ?? (_usersExcludeds = new InputList<string>());
            set => _usersExcludeds = value;
        }

        public MfaPolicyRuleState()
        {
        }
    }
}
