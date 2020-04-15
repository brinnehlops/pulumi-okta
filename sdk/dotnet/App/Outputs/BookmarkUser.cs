// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.App.Outputs
{

    [OutputType]
    public sealed class BookmarkUser
    {
        /// <summary>
        /// ID of the Application.
        /// </summary>
        public readonly string? Id;
        public readonly string? Password;
        public readonly string? Scope;
        public readonly string? Username;

        [OutputConstructor]
        private BookmarkUser(
            string? id,

            string? password,

            string? scope,

            string? username)
        {
            Id = id;
            Password = password;
            Scope = scope;
            Username = username;
        }
    }
}