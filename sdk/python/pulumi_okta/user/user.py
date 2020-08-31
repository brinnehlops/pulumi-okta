# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['User']


class User(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_roles: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 city: Optional[pulumi.Input[str]] = None,
                 cost_center: Optional[pulumi.Input[str]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 custom_profile_attributes: Optional[pulumi.Input[str]] = None,
                 department: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 division: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 employee_number: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 group_memberships: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 honorific_prefix: Optional[pulumi.Input[str]] = None,
                 honorific_suffix: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 locale: Optional[pulumi.Input[str]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 manager: Optional[pulumi.Input[str]] = None,
                 manager_id: Optional[pulumi.Input[str]] = None,
                 middle_name: Optional[pulumi.Input[str]] = None,
                 mobile_phone: Optional[pulumi.Input[str]] = None,
                 nick_name: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 postal_address: Optional[pulumi.Input[str]] = None,
                 preferred_language: Optional[pulumi.Input[str]] = None,
                 primary_phone: Optional[pulumi.Input[str]] = None,
                 profile_url: Optional[pulumi.Input[str]] = None,
                 recovery_answer: Optional[pulumi.Input[str]] = None,
                 recovery_question: Optional[pulumi.Input[str]] = None,
                 second_email: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 street_address: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 user_type: Optional[pulumi.Input[str]] = None,
                 zip_code: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Okta User.

        This resource allows you to create and configure an Okta User.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] admin_roles: Administrator roles assigned to User.
        :param pulumi.Input[str] city: User profile property.
        :param pulumi.Input[str] cost_center: User profile property.
        :param pulumi.Input[str] country_code: User profile property.
        :param pulumi.Input[str] custom_profile_attributes: raw JSON containing all custom profile attributes.
        :param pulumi.Input[str] department: User profile property.
        :param pulumi.Input[str] display_name: User profile property.
        :param pulumi.Input[str] division: User profile property.
        :param pulumi.Input[str] email: User profile property.
        :param pulumi.Input[str] employee_number: User profile property.
        :param pulumi.Input[str] first_name: User's First Name, required by default.
        :param pulumi.Input[List[pulumi.Input[str]]] group_memberships: User profile property.
        :param pulumi.Input[str] honorific_prefix: User profile property.
        :param pulumi.Input[str] honorific_suffix: User profile property.
        :param pulumi.Input[str] last_name: User's Last Name, required by default.
        :param pulumi.Input[str] locale: User profile property.
        :param pulumi.Input[str] login: User profile property.
        :param pulumi.Input[str] manager: User profile property.
        :param pulumi.Input[str] manager_id: User profile property.
        :param pulumi.Input[str] middle_name: User profile property.
        :param pulumi.Input[str] mobile_phone: User profile property.
        :param pulumi.Input[str] nick_name: User profile property.
        :param pulumi.Input[str] organization: User profile property.
        :param pulumi.Input[str] password: User password.
        :param pulumi.Input[str] postal_address: User profile property.
        :param pulumi.Input[str] preferred_language: User profile property.
        :param pulumi.Input[str] primary_phone: User profile property.
        :param pulumi.Input[str] profile_url: User profile property.
        :param pulumi.Input[str] recovery_answer: User password recovery answer.
        :param pulumi.Input[str] recovery_question: User password recovery question.
        :param pulumi.Input[str] second_email: User profile property.
        :param pulumi.Input[str] state: User profile property.
        :param pulumi.Input[str] status: User profile property.
        :param pulumi.Input[str] street_address: User profile property.
        :param pulumi.Input[str] timezone: User profile property.
        :param pulumi.Input[str] title: User profile property.
        :param pulumi.Input[str] user_type: User profile property.
        :param pulumi.Input[str] zip_code: User profile property.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_roles'] = admin_roles
            __props__['city'] = city
            __props__['cost_center'] = cost_center
            __props__['country_code'] = country_code
            __props__['custom_profile_attributes'] = custom_profile_attributes
            __props__['department'] = department
            __props__['display_name'] = display_name
            __props__['division'] = division
            if email is None:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['employee_number'] = employee_number
            if first_name is None:
                raise TypeError("Missing required property 'first_name'")
            __props__['first_name'] = first_name
            __props__['group_memberships'] = group_memberships
            __props__['honorific_prefix'] = honorific_prefix
            __props__['honorific_suffix'] = honorific_suffix
            if last_name is None:
                raise TypeError("Missing required property 'last_name'")
            __props__['last_name'] = last_name
            __props__['locale'] = locale
            if login is None:
                raise TypeError("Missing required property 'login'")
            __props__['login'] = login
            __props__['manager'] = manager
            __props__['manager_id'] = manager_id
            __props__['middle_name'] = middle_name
            __props__['mobile_phone'] = mobile_phone
            __props__['nick_name'] = nick_name
            __props__['organization'] = organization
            __props__['password'] = password
            __props__['postal_address'] = postal_address
            __props__['preferred_language'] = preferred_language
            __props__['primary_phone'] = primary_phone
            __props__['profile_url'] = profile_url
            __props__['recovery_answer'] = recovery_answer
            __props__['recovery_question'] = recovery_question
            __props__['second_email'] = second_email
            __props__['state'] = state
            __props__['status'] = status
            __props__['street_address'] = street_address
            __props__['timezone'] = timezone
            __props__['title'] = title
            __props__['user_type'] = user_type
            __props__['zip_code'] = zip_code
            __props__['raw_status'] = None
        super(User, __self__).__init__(
            'okta:user/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_roles: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            city: Optional[pulumi.Input[str]] = None,
            cost_center: Optional[pulumi.Input[str]] = None,
            country_code: Optional[pulumi.Input[str]] = None,
            custom_profile_attributes: Optional[pulumi.Input[str]] = None,
            department: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            division: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            employee_number: Optional[pulumi.Input[str]] = None,
            first_name: Optional[pulumi.Input[str]] = None,
            group_memberships: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            honorific_prefix: Optional[pulumi.Input[str]] = None,
            honorific_suffix: Optional[pulumi.Input[str]] = None,
            last_name: Optional[pulumi.Input[str]] = None,
            locale: Optional[pulumi.Input[str]] = None,
            login: Optional[pulumi.Input[str]] = None,
            manager: Optional[pulumi.Input[str]] = None,
            manager_id: Optional[pulumi.Input[str]] = None,
            middle_name: Optional[pulumi.Input[str]] = None,
            mobile_phone: Optional[pulumi.Input[str]] = None,
            nick_name: Optional[pulumi.Input[str]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            postal_address: Optional[pulumi.Input[str]] = None,
            preferred_language: Optional[pulumi.Input[str]] = None,
            primary_phone: Optional[pulumi.Input[str]] = None,
            profile_url: Optional[pulumi.Input[str]] = None,
            raw_status: Optional[pulumi.Input[str]] = None,
            recovery_answer: Optional[pulumi.Input[str]] = None,
            recovery_question: Optional[pulumi.Input[str]] = None,
            second_email: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            street_address: Optional[pulumi.Input[str]] = None,
            timezone: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None,
            user_type: Optional[pulumi.Input[str]] = None,
            zip_code: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] admin_roles: Administrator roles assigned to User.
        :param pulumi.Input[str] city: User profile property.
        :param pulumi.Input[str] cost_center: User profile property.
        :param pulumi.Input[str] country_code: User profile property.
        :param pulumi.Input[str] custom_profile_attributes: raw JSON containing all custom profile attributes.
        :param pulumi.Input[str] department: User profile property.
        :param pulumi.Input[str] display_name: User profile property.
        :param pulumi.Input[str] division: User profile property.
        :param pulumi.Input[str] email: User profile property.
        :param pulumi.Input[str] employee_number: User profile property.
        :param pulumi.Input[str] first_name: User's First Name, required by default.
        :param pulumi.Input[List[pulumi.Input[str]]] group_memberships: User profile property.
        :param pulumi.Input[str] honorific_prefix: User profile property.
        :param pulumi.Input[str] honorific_suffix: User profile property.
        :param pulumi.Input[str] last_name: User's Last Name, required by default.
        :param pulumi.Input[str] locale: User profile property.
        :param pulumi.Input[str] login: User profile property.
        :param pulumi.Input[str] manager: User profile property.
        :param pulumi.Input[str] manager_id: User profile property.
        :param pulumi.Input[str] middle_name: User profile property.
        :param pulumi.Input[str] mobile_phone: User profile property.
        :param pulumi.Input[str] nick_name: User profile property.
        :param pulumi.Input[str] organization: User profile property.
        :param pulumi.Input[str] password: User password.
        :param pulumi.Input[str] postal_address: User profile property.
        :param pulumi.Input[str] preferred_language: User profile property.
        :param pulumi.Input[str] primary_phone: User profile property.
        :param pulumi.Input[str] profile_url: User profile property.
        :param pulumi.Input[str] raw_status: The raw status of the User in Okta - (status is mapped)
        :param pulumi.Input[str] recovery_answer: User password recovery answer.
        :param pulumi.Input[str] recovery_question: User password recovery question.
        :param pulumi.Input[str] second_email: User profile property.
        :param pulumi.Input[str] state: User profile property.
        :param pulumi.Input[str] status: User profile property.
        :param pulumi.Input[str] street_address: User profile property.
        :param pulumi.Input[str] timezone: User profile property.
        :param pulumi.Input[str] title: User profile property.
        :param pulumi.Input[str] user_type: User profile property.
        :param pulumi.Input[str] zip_code: User profile property.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_roles"] = admin_roles
        __props__["city"] = city
        __props__["cost_center"] = cost_center
        __props__["country_code"] = country_code
        __props__["custom_profile_attributes"] = custom_profile_attributes
        __props__["department"] = department
        __props__["display_name"] = display_name
        __props__["division"] = division
        __props__["email"] = email
        __props__["employee_number"] = employee_number
        __props__["first_name"] = first_name
        __props__["group_memberships"] = group_memberships
        __props__["honorific_prefix"] = honorific_prefix
        __props__["honorific_suffix"] = honorific_suffix
        __props__["last_name"] = last_name
        __props__["locale"] = locale
        __props__["login"] = login
        __props__["manager"] = manager
        __props__["manager_id"] = manager_id
        __props__["middle_name"] = middle_name
        __props__["mobile_phone"] = mobile_phone
        __props__["nick_name"] = nick_name
        __props__["organization"] = organization
        __props__["password"] = password
        __props__["postal_address"] = postal_address
        __props__["preferred_language"] = preferred_language
        __props__["primary_phone"] = primary_phone
        __props__["profile_url"] = profile_url
        __props__["raw_status"] = raw_status
        __props__["recovery_answer"] = recovery_answer
        __props__["recovery_question"] = recovery_question
        __props__["second_email"] = second_email
        __props__["state"] = state
        __props__["status"] = status
        __props__["street_address"] = street_address
        __props__["timezone"] = timezone
        __props__["title"] = title
        __props__["user_type"] = user_type
        __props__["zip_code"] = zip_code
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminRoles")
    def admin_roles(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Administrator roles assigned to User.
        """
        return pulumi.get(self, "admin_roles")

    @property
    @pulumi.getter
    def city(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="costCenter")
    def cost_center(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "cost_center")

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter(name="customProfileAttributes")
    def custom_profile_attributes(self) -> pulumi.Output[Optional[str]]:
        """
        raw JSON containing all custom profile attributes.
        """
        return pulumi.get(self, "custom_profile_attributes")

    @property
    @pulumi.getter
    def department(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "department")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def division(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "division")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        User profile property.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="employeeNumber")
    def employee_number(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "employee_number")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[str]:
        """
        User's First Name, required by default.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="groupMemberships")
    def group_memberships(self) -> pulumi.Output[Optional[List[str]]]:
        """
        User profile property.
        """
        return pulumi.get(self, "group_memberships")

    @property
    @pulumi.getter(name="honorificPrefix")
    def honorific_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "honorific_prefix")

    @property
    @pulumi.getter(name="honorificSuffix")
    def honorific_suffix(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "honorific_suffix")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[str]:
        """
        User's Last Name, required by default.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter
    def locale(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "locale")

    @property
    @pulumi.getter
    def login(self) -> pulumi.Output[str]:
        """
        User profile property.
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    def manager(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "manager")

    @property
    @pulumi.getter(name="managerId")
    def manager_id(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "manager_id")

    @property
    @pulumi.getter(name="middleName")
    def middle_name(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "middle_name")

    @property
    @pulumi.getter(name="mobilePhone")
    def mobile_phone(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "mobile_phone")

    @property
    @pulumi.getter(name="nickName")
    def nick_name(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "nick_name")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        User password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="postalAddress")
    def postal_address(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "postal_address")

    @property
    @pulumi.getter(name="preferredLanguage")
    def preferred_language(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "preferred_language")

    @property
    @pulumi.getter(name="primaryPhone")
    def primary_phone(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "primary_phone")

    @property
    @pulumi.getter(name="profileUrl")
    def profile_url(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "profile_url")

    @property
    @pulumi.getter(name="rawStatus")
    def raw_status(self) -> pulumi.Output[str]:
        """
        The raw status of the User in Okta - (status is mapped)
        """
        return pulumi.get(self, "raw_status")

    @property
    @pulumi.getter(name="recoveryAnswer")
    def recovery_answer(self) -> pulumi.Output[Optional[str]]:
        """
        User password recovery answer.
        """
        return pulumi.get(self, "recovery_answer")

    @property
    @pulumi.getter(name="recoveryQuestion")
    def recovery_question(self) -> pulumi.Output[Optional[str]]:
        """
        User password recovery question.
        """
        return pulumi.get(self, "recovery_question")

    @property
    @pulumi.getter(name="secondEmail")
    def second_email(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "second_email")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="streetAddress")
    def street_address(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "street_address")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="userType")
    def user_type(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "user_type")

    @property
    @pulumi.getter(name="zipCode")
    def zip_code(self) -> pulumi.Output[Optional[str]]:
        """
        User profile property.
        """
        return pulumi.get(self, "zip_code")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

