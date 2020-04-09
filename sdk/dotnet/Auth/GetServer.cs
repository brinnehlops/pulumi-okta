// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Auth
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve an auth server from Okta.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/auth_server.html.markdown.
        /// </summary>
        [Obsolete("Use GetServer.InvokeAsync() instead")]
        public static Task<GetServerResult> GetServer(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("okta:auth/getServer:getServer", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetServer
    {
        /// <summary>
        /// Use this data source to retrieve an auth server from Okta.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/auth_server.html.markdown.
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("okta:auth/getServer:getServer", args ?? InvokeArgs.Empty, options.WithVersion());
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
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetServerResult(
            ImmutableArray<string> audiences,
            string credentialsLastRotated,
            string credentialsNextRotation,
            string credentialsRotationMode,
            string description,
            string kid,
            string name,
            string status,
            string id)
        {
            Audiences = audiences;
            CredentialsLastRotated = credentialsLastRotated;
            CredentialsNextRotation = credentialsNextRotation;
            CredentialsRotationMode = credentialsRotationMode;
            Description = description;
            Kid = kid;
            Name = name;
            Status = status;
            Id = id;
        }
    }
}
