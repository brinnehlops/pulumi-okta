// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.User
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve a users from Okta.
        /// 
        /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/user.html.markdown.
        /// </summary>
        public static Task<GetUserResult> GetUser(GetUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("okta:user/getUser:getUser", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetUserArgs : Pulumi.InvokeArgs
    {
        [Input("searches", required: true)]
        private List<Inputs.GetUserSearchesArgs>? _searches;

        /// <summary>
        /// Map of search criteria. It supports the following properties.
        /// </summary>
        public List<Inputs.GetUserSearchesArgs> Searches
        {
            get => _searches ?? (_searches = new List<Inputs.GetUserSearchesArgs>());
            set => _searches = value;
        }

        public GetUserArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetUserResult
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
        public readonly ImmutableArray<Outputs.GetUserSearchesResult> Searches;
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
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetUserResult(
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
            ImmutableArray<Outputs.GetUserSearchesResult> searches,
            string secondEmail,
            string state,
            string status,
            string streetAddress,
            string timezone,
            string title,
            string userType,
            string zipCode,
            string id)
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
            Searches = searches;
            SecondEmail = secondEmail;
            State = state;
            Status = status;
            StreetAddress = streetAddress;
            Timezone = timezone;
            Title = title;
            UserType = userType;
            ZipCode = zipCode;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetUserSearchesArgs : Pulumi.InvokeArgs
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

        public GetUserSearchesArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetUserSearchesResult
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
        private GetUserSearchesResult(
            string? comparison,
            string name,
            string value)
        {
            Comparison = comparison;
            Name = name;
            Value = value;
        }
    }
    }
}
