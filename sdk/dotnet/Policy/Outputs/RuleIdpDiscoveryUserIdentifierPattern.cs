// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Policy.Outputs
{

    [OutputType]
    public sealed class RuleIdpDiscoveryUserIdentifierPattern
    {
        public readonly string? MatchType;
        public readonly string? Value;

        [OutputConstructor]
        private RuleIdpDiscoveryUserIdentifierPattern(
            string? matchType,

            string? value)
        {
            MatchType = matchType;
            Value = value;
        }
    }
}
