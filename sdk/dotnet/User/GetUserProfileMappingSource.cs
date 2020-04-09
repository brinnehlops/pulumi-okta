// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.User
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve the base user Profile Mapping source or target from Okta.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/user_profile_mapping_source.html.markdown.
        /// </summary>
        [Obsolete("Use GetUserProfileMappingSource.InvokeAsync() instead")]
        public static Task<GetUserProfileMappingSourceResult> GetUserProfileMappingSource(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserProfileMappingSourceResult>("okta:user/getUserProfileMappingSource:getUserProfileMappingSource", InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetUserProfileMappingSource
    {
        /// <summary>
        /// Use this data source to retrieve the base user Profile Mapping source or target from Okta.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/user_profile_mapping_source.html.markdown.
        /// </summary>
        public static Task<GetUserProfileMappingSourceResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserProfileMappingSourceResult>("okta:user/getUserProfileMappingSource:getUserProfileMappingSource", InvokeArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetUserProfileMappingSourceResult
    {
        /// <summary>
        /// name of source.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// type of source.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetUserProfileMappingSourceResult(
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
