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
        /// Use this data source to retrieve a list of users from Okta.
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/users.html.markdown.
        /// </summary>
        public static Task<GetUsersResult> GetUsers(GetUsersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("okta:user/getUsers:getUsers", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetUsersArgs : Pulumi.InvokeArgs
    {
        [Input("searches", required: true)]
        private List<Inputs.GetUsersSearchesArgs>? _searches;

        /// <summary>
        /// Map of search criteria to use to find users. It supports the following properties.
        /// </summary>
        public List<Inputs.GetUsersSearchesArgs> Searches
        {
            get => _searches ?? (_searches = new List<Inputs.GetUsersSearchesArgs>());
            set => _searches = value;
        }

        [Input("users")]
        private List<Inputs.GetUsersUsersArgs>? _users;
        public List<Inputs.GetUsersUsersArgs> Users
        {
            get => _users ?? (_users = new List<Inputs.GetUsersUsersArgs>());
            set => _users = value;
        }

        public GetUsersArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetUsersResult
    {
        public readonly ImmutableArray<Outputs.GetUsersSearchesResult> Searches;
        /// <summary>
        /// collection of users retrieved from Okta with the following properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersUsersResult> Users;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetUsersResult(
            ImmutableArray<Outputs.GetUsersSearchesResult> searches,
            ImmutableArray<Outputs.GetUsersUsersResult> users,
            string id)
        {
            Searches = searches;
            Users = users;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetUsersSearchesArgs : Pulumi.InvokeArgs
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

        public GetUsersSearchesArgs()
        {
        }
    }

    public sealed class GetUsersUsersArgs : Pulumi.InvokeArgs
    {
        [Input("adminRoles")]
        private List<string>? _adminRoles;

        /// <summary>
        /// Administrator roles assigned to user.
        /// </summary>
        public List<string> AdminRoles
        {
            get => _adminRoles ?? (_adminRoles = new List<string>());
            set => _adminRoles = value;
        }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("city")]
        public string? City { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("costCenter")]
        public string? CostCenter { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("countryCode")]
        public string? CountryCode { get; set; }

        /// <summary>
        /// raw JSON containing all custom profile attributes.
        /// </summary>
        [Input("customProfileAttributes")]
        public string? CustomProfileAttributes { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("department")]
        public string? Department { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("division")]
        public string? Division { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("employeeNumber")]
        public string? EmployeeNumber { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("firstName")]
        public string? FirstName { get; set; }

        [Input("groupMemberships")]
        private List<string>? _groupMemberships;

        /// <summary>
        /// user profile property.
        /// </summary>
        public List<string> GroupMemberships
        {
            get => _groupMemberships ?? (_groupMemberships = new List<string>());
            set => _groupMemberships = value;
        }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("honorificPrefix")]
        public string? HonorificPrefix { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("honorificSuffix")]
        public string? HonorificSuffix { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("lastName")]
        public string? LastName { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("locale")]
        public string? Locale { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("login")]
        public string? Login { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("manager")]
        public string? Manager { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("managerId")]
        public string? ManagerId { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("middleName")]
        public string? MiddleName { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("mobilePhone")]
        public string? MobilePhone { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("nickName")]
        public string? NickName { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("organization")]
        public string? Organization { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("postalAddress")]
        public string? PostalAddress { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("preferredLanguage")]
        public string? PreferredLanguage { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("primaryPhone")]
        public string? PrimaryPhone { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("profileUrl")]
        public string? ProfileUrl { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("secondEmail")]
        public string? SecondEmail { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("streetAddress")]
        public string? StreetAddress { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("timezone")]
        public string? Timezone { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("title")]
        public string? Title { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("userType")]
        public string? UserType { get; set; }

        /// <summary>
        /// user profile property.
        /// </summary>
        [Input("zipCode")]
        public string? ZipCode { get; set; }

        public GetUsersUsersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetUsersSearchesResult
    {
        /// <summary>
        /// Comparison to use.
        /// </summary>
        public readonly string? Comparison;
        /// <summary>
        /// Name of property to search against.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Value to compare with.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetUsersSearchesResult(
            string? comparison,
            string name,
            string value)
        {
            Comparison = comparison;
            Name = name;
            Value = value;
        }
    }

    [OutputType]
    public sealed class GetUsersUsersResult
    {
        /// <summary>
        /// Administrator roles assigned to user.
        /// </summary>
        public readonly ImmutableArray<string> AdminRoles;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string CostCenter;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string CountryCode;
        /// <summary>
        /// raw JSON containing all custom profile attributes.
        /// </summary>
        public readonly string CustomProfileAttributes;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Department;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Division;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string EmployeeNumber;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly ImmutableArray<string> GroupMemberships;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string HonorificPrefix;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string HonorificSuffix;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Locale;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Login;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Manager;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string ManagerId;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string MiddleName;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string MobilePhone;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string NickName;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string PostalAddress;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string PreferredLanguage;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string PrimaryPhone;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string ProfileUrl;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string SecondEmail;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string StreetAddress;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Timezone;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string UserType;
        /// <summary>
        /// user profile property.
        /// </summary>
        public readonly string ZipCode;

        [OutputConstructor]
        private GetUsersUsersResult(
            ImmutableArray<string> adminRoles,
            string city,
            string costCenter,
            string countryCode,
            string customProfileAttributes,
            string department,
            string displayName,
            string division,
            string email,
            string employeeNumber,
            string firstName,
            ImmutableArray<string> groupMemberships,
            string honorificPrefix,
            string honorificSuffix,
            string lastName,
            string locale,
            string login,
            string manager,
            string managerId,
            string middleName,
            string mobilePhone,
            string nickName,
            string organization,
            string postalAddress,
            string preferredLanguage,
            string primaryPhone,
            string profileUrl,
            string secondEmail,
            string state,
            string status,
            string streetAddress,
            string timezone,
            string title,
            string userType,
            string zipCode)
        {
            AdminRoles = adminRoles;
            City = city;
            CostCenter = costCenter;
            CountryCode = countryCode;
            CustomProfileAttributes = customProfileAttributes;
            Department = department;
            DisplayName = displayName;
            Division = division;
            Email = email;
            EmployeeNumber = employeeNumber;
            FirstName = firstName;
            GroupMemberships = groupMemberships;
            HonorificPrefix = honorificPrefix;
            HonorificSuffix = honorificSuffix;
            LastName = lastName;
            Locale = locale;
            Login = login;
            Manager = manager;
            ManagerId = managerId;
            MiddleName = middleName;
            MobilePhone = mobilePhone;
            NickName = nickName;
            Organization = organization;
            PostalAddress = postalAddress;
            PreferredLanguage = preferredLanguage;
            PrimaryPhone = primaryPhone;
            ProfileUrl = profileUrl;
            SecondEmail = secondEmail;
            State = state;
            Status = status;
            StreetAddress = streetAddress;
            Timezone = timezone;
            Title = title;
            UserType = userType;
            ZipCode = zipCode;
        }
    }
    }
}
