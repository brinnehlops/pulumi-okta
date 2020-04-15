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
    public sealed class MfaYubikeyToken
    {
        /// <summary>
        /// User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
        /// </summary>
        public readonly string? ConsentType;
        /// <summary>
        /// Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        /// </summary>
        public readonly string? Enroll;

        [OutputConstructor]
        private MfaYubikeyToken(
            string? consentType,

            string? enroll)
        {
            ConsentType = consentType;
            Enroll = enroll;
        }
    }
}
