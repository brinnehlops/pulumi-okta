// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Deprecated.Inputs
{

    public sealed class SamlAppAttributeStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        [Input("filterValue")]
        public Input<string>? FilterValue { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public SamlAppAttributeStatementGetArgs()
        {
        }
    }
}
