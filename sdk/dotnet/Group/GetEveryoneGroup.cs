// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Group
{
    public static class GetEveryoneGroup
    {
        /// <summary>
        /// Use this data source to retrieve the Everyone group from Okta. The same can be achieved with the `okta.group.Group` data source with `name = "Everyone"`. This is simply a shortcut.
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
        ///         var example = Output.Create(Okta.Group.GetEveryoneGroup.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEveryoneGroupResult> InvokeAsync(GetEveryoneGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEveryoneGroupResult>("okta:group/getEveryoneGroup:getEveryoneGroup", args ?? new GetEveryoneGroupArgs(), options.WithVersion());
    }


    public sealed class GetEveryoneGroupArgs : Pulumi.InvokeArgs
    {
        [Input("includeUsers")]
        public bool? IncludeUsers { get; set; }

        public GetEveryoneGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEveryoneGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeUsers;

        [OutputConstructor]
        private GetEveryoneGroupResult(
            string id,

            bool? includeUsers)
        {
            Id = id;
            IncludeUsers = includeUsers;
        }
    }
}
