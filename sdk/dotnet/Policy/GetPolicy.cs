// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Policy
{
    public static class GetPolicy
    {
        /// <summary>
        /// Use this data source to retrieve a policy from Okta.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Okta = Pulumi.Okta;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Okta.Policy.GetPolicy.InvokeAsync(new Okta.Policy.GetPolicyArgs
        ///         {
        ///             Name = "Password Policy Example",
        ///             Type = "PASSWORD",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("okta:policy/getPolicy:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// name of policy to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// type of policy to retrieve.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// name of policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// type of policy.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyResult(
            string id,

            string name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
