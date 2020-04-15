// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.User.Inputs
{

    public sealed class GetUserSearchArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Comparison to use.
        /// </summary>
        [Input("comparison")]
        public string? Comparison { get; set; }

        /// <summary>
        /// Name of property to search against.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Value to compare with.
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetUserSearchArgs()
        {
        }
    }
}
