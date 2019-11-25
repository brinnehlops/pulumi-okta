// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Policy
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve a policy from Okta.
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/policy.html.markdown.
        /// </summary>
        public static Task<GetPolicyResult> GetPolicy(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("okta:policy/getPolicy:getPolicy", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of policy to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// type of policy to retrieve.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// name of policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// type of policy.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPolicyResult(
            string name,
            string type,
            string id)
        {
            Name = name;
            Type = type;
            Id = id;
        }
    }
}
