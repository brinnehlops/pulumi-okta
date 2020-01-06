// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App
{
    /// <summary>
    /// Assigns a group to an application.
    /// 
    /// This resource allows you to create an App Group assignment.
    /// 
    /// __When using this resource, make sure to add the following `lifefycle` argument to the application resource you are assigning to:__
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_group_assignment.html.markdown.
    /// </summary>
    public partial class GroupAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the application to assign a group to.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// The ID of the group to assign the app to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;


        /// <summary>
        /// Create a GroupAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupAssignment(string name, GroupAssignmentArgs args, CustomResourceOptions? options = null)
            : base("okta:app/groupAssignment:GroupAssignment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GroupAssignment(string name, Input<string> id, GroupAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("okta:app/groupAssignment:GroupAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupAssignment Get(string name, Input<string> id, GroupAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupAssignment(name, id, state, options);
        }
    }

    public sealed class GroupAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application to assign a group to.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// The ID of the group to assign the app to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        public GroupAssignmentArgs()
        {
        }
    }

    public sealed class GroupAssignmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application to assign a group to.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// The ID of the group to assign the app to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        public GroupAssignmentState()
        {
        }
    }
}
