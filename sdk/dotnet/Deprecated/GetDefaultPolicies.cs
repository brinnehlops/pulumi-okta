// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated
{
    public static class GetDefaultPolicies
    {
        public static Task<GetDefaultPoliciesResult> InvokeAsync(GetDefaultPoliciesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultPoliciesResult>("okta:deprecated/getDefaultPolicies:getDefaultPolicies", args ?? new GetDefaultPoliciesArgs(), options.WithVersion());
    }


    public sealed class GetDefaultPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetDefaultPoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDefaultPoliciesResult
    {
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Type;

        [OutputConstructor]
        private GetDefaultPoliciesResult(
            string id,

            string type)
        {
            Id = id;
            Type = type;
        }
    }
}
