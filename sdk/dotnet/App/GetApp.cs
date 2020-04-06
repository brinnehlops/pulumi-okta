// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve the collaborators for a given repository.
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/app.html.markdown.
        /// </summary>
        [Obsolete("Use GetApp.InvokeAsync() instead")]
        public static Task<GetAppResult> GetApp(GetAppArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppResult>("okta:app/getApp:getApp", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetApp
    {
        /// <summary>
        /// Use this data source to retrieve the collaborators for a given repository.
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/app.html.markdown.
        /// </summary>
        public static Task<GetAppResult> InvokeAsync(GetAppArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppResult>("okta:app/getApp:getApp", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAppArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// tells the provider to query for only `ACTIVE` applications.
        /// </summary>
        [Input("activeOnly")]
        public bool? ActiveOnly { get; set; }

        /// <summary>
        /// `id` of application to retrieve, conflicts with `label` and `label_prefix`.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The label of the app to retrieve, conflicts with `label_prefix` and `id`.
        /// </summary>
        [Input("label")]
        public string? Label { get; set; }

        /// <summary>
        /// Label prefix of the app to retrieve, conflicts with `label` and `id`. This will tell the provider to do a `starts with` query as opposed to an `equals` query.
        /// </summary>
        [Input("labelPrefix")]
        public string? LabelPrefix { get; set; }

        public GetAppArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAppResult
    {
        public readonly bool? ActiveOnly;
        /// <summary>
        /// `description` of application.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// `id` of application.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// `label` of application.
        /// </summary>
        public readonly string? Label;
        public readonly string? LabelPrefix;
        /// <summary>
        /// `name` of application.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// `status` of application.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetAppResult(
            bool? activeOnly,
            string description,
            string? id,
            string? label,
            string? labelPrefix,
            string name,
            string status)
        {
            ActiveOnly = activeOnly;
            Description = description;
            Id = id;
            Label = label;
            LabelPrefix = labelPrefix;
            Name = name;
            Status = status;
        }
    }
}
