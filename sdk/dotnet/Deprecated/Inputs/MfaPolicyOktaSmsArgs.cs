// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated.Inputs
{

    public sealed class MfaPolicyOktaSmsArgs : Pulumi.ResourceArgs
    {
        [Input("consentType")]
        public Input<string>? ConsentType { get; set; }

        [Input("enroll")]
        public Input<string>? Enroll { get; set; }

        public MfaPolicyOktaSmsArgs()
        {
        }
    }
}
