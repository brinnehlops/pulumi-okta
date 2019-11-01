// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Group
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve a group from Okta.
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/group.html.markdown.
        /// </summary>
        public static Task<GetGroupResult> GetGroup(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("okta:group/getGroup:getGroup", args, options.WithVersion());
    }

    public sealed class GetGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether or not to retrieve all member ids.
        /// </summary>
        [Input("includeUsers")]
        public Input<bool>? IncludeUsers { get; set; }

        /// <summary>
        /// name of group to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// description of group.
        /// </summary>
        public readonly string Description;
        public readonly bool? IncludeUsers;
        /// <summary>
        /// name of group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// user ids that are members of this group, only included if `include_users` is set to `true`.
        /// </summary>
        public readonly ImmutableArray<string> Users;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupResult(
            string description,
            bool? includeUsers,
            string name,
            ImmutableArray<string> users,
            string id)
        {
            Description = description;
            IncludeUsers = includeUsers;
            Name = name;
            Users = users;
            Id = id;
        }
    }
}
