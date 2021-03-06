// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated
{
    public partial class SignonPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Policy Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of Group IDs to Include
        /// </summary>
        [Output("groupsIncludeds")]
        public Output<ImmutableArray<string>> GroupsIncludeds { get; private set; } = null!;

        /// <summary>
        /// Policy Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
        /// priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Policy Status: ACTIVE or INACTIVE.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SignonPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SignonPolicy(string name, SignonPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("okta:deprecated/signonPolicy:SignonPolicy", name, args ?? new SignonPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SignonPolicy(string name, Input<string> id, SignonPolicyState? state = null, CustomResourceOptions? options = null)
            : base("okta:deprecated/signonPolicy:SignonPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SignonPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SignonPolicy Get(string name, Input<string> id, SignonPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SignonPolicy(name, id, state, options);
        }
    }

    public sealed class SignonPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groupsIncludeds")]
        private InputList<string>? _groupsIncludeds;

        /// <summary>
        /// List of Group IDs to Include
        /// </summary>
        public InputList<string> GroupsIncludeds
        {
            get => _groupsIncludeds ?? (_groupsIncludeds = new InputList<string>());
            set => _groupsIncludeds = value;
        }

        /// <summary>
        /// Policy Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
        /// priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Policy Status: ACTIVE or INACTIVE.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SignonPolicyArgs()
        {
        }
    }

    public sealed class SignonPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groupsIncludeds")]
        private InputList<string>? _groupsIncludeds;

        /// <summary>
        /// List of Group IDs to Include
        /// </summary>
        public InputList<string> GroupsIncludeds
        {
            get => _groupsIncludeds ?? (_groupsIncludeds = new InputList<string>());
            set => _groupsIncludeds = value;
        }

        /// <summary>
        /// Policy Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
        /// priority is provided. API defaults it to the last/lowest if not there.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Policy Status: ACTIVE or INACTIVE.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SignonPolicyState()
        {
        }
    }
}
