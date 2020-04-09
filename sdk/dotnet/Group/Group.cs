// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Group
{
    /// <summary>
    /// Creates an Okta Group.
    /// 
    /// This resource allows you to create and configure an Okta Group.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/group.html.markdown.
    /// </summary>
    public partial class Group : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Okta Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the Okta Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Users associated with the group. This can also be done per user.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("okta:group/group:Group", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("okta:group/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Okta Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Okta Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// Users associated with the group. This can also be done per user.
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Okta Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Okta Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// Users associated with the group. This can also be done per user.
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public GroupState()
        {
        }
    }
}
