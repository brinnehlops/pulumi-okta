// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Policy.Inputs
{

    public sealed class RuleIdpDiscoveryAppIncludeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Policy Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public RuleIdpDiscoveryAppIncludeArgs()
        {
        }
    }
}
