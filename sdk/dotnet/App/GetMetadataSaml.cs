// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App
{
    public static class GetMetadataSaml
    {
        /// <summary>
        /// Use this data source to retrieve the collaborators for a given repository.
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
        ///         var example = Output.Create(Okta.App.GetMetadataSaml.InvokeAsync(new Okta.App.GetMetadataSamlArgs
        ///         {
        ///             AppId = "&lt;app id&gt;",
        ///             KeyId = "&lt;cert key id&gt;",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMetadataSamlResult> InvokeAsync(GetMetadataSamlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMetadataSamlResult>("okta:app/getMetadataSaml:getMetadataSaml", args ?? new GetMetadataSamlArgs(), options.WithVersion());
    }


    public sealed class GetMetadataSamlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        /// <summary>
        /// Certificate Key ID.
        /// </summary>
        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        public GetMetadataSamlArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMetadataSamlResult
    {
        public readonly string AppId;
        /// <summary>
        /// public certificate from application metadata.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Entity URL for instance `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf`.
        /// </summary>
        public readonly string EntityId;
        /// <summary>
        /// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.
        /// </summary>
        public readonly string HttpPostBinding;
        /// <summary>
        /// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.
        /// </summary>
        public readonly string HttpRedirectBinding;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KeyId;
        /// <summary>
        /// raw metadata of application.
        /// </summary>
        public readonly string Metadata;
        /// <summary>
        /// Whether authn requests are signed.
        /// </summary>
        public readonly bool WantAuthnRequestsSigned;

        [OutputConstructor]
        private GetMetadataSamlResult(
            string appId,

            string certificate,

            string entityId,

            string httpPostBinding,

            string httpRedirectBinding,

            string id,

            string keyId,

            string metadata,

            bool wantAuthnRequestsSigned)
        {
            AppId = appId;
            Certificate = certificate;
            EntityId = entityId;
            HttpPostBinding = httpPostBinding;
            HttpRedirectBinding = httpRedirectBinding;
            Id = id;
            KeyId = keyId;
            Metadata = metadata;
            WantAuthnRequestsSigned = wantAuthnRequestsSigned;
        }
    }
}
