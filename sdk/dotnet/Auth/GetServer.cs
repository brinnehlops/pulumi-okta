// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Auth
{
    public static class GetServer
    {
        /// <summary>
        /// Use this data source to retrieve an auth server from Okta.
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
        ///         var example = Output.Create(Okta.Auth.GetServer.InvokeAsync(new Okta.Auth.GetServerArgs
        ///         {
        ///             Name = "Example Auth",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("okta:auth/getServer:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the auth server to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// array of audiences,
        /// </summary>
        public readonly ImmutableArray<string> Audiences;
        /// <summary>
        /// last time credentials were rotated.
        /// </summary>
        public readonly string CredentialsLastRotated;
        /// <summary>
        /// next time credentials will be rotated
        /// </summary>
        public readonly string CredentialsNextRotation;
        /// <summary>
        /// mode of credential rotation, auto or manual.
        /// </summary>
        public readonly string CredentialsRotationMode;
        /// <summary>
        /// description of Authorization server.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// auth server key id.
        /// </summary>
        public readonly string Kid;
        /// <summary>
        /// The name of the auth server.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// the activation status of the authorization server.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetServerResult(
            ImmutableArray<string> audiences,

            string credentialsLastRotated,

            string credentialsNextRotation,

            string credentialsRotationMode,

            string description,

            string id,

            string kid,

            string name,

            string status)
        {
            Audiences = audiences;
            CredentialsLastRotated = credentialsLastRotated;
            CredentialsNextRotation = credentialsNextRotation;
            CredentialsRotationMode = credentialsRotationMode;
            Description = description;
            Id = id;
            Kid = kid;
            Name = name;
            Status = status;
        }
    }
}
