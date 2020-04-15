// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App.Inputs
{

    public sealed class UserSchemaArrayOneOfArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// value mapping to member of `enum`.
        /// </summary>
        [Input("const", required: true)]
        public Input<string> Const { get; set; } = null!;

        /// <summary>
        /// display name for the enum value.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public UserSchemaArrayOneOfArgs()
        {
        }
    }
}
