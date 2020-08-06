# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, admin_roles=None, city=None, cost_center=None, country_code=None, custom_profile_attributes=None, department=None, display_name=None, division=None, email=None, employee_number=None, first_name=None, group_memberships=None, honorific_prefix=None, honorific_suffix=None, id=None, last_name=None, locale=None, login=None, manager=None, manager_id=None, middle_name=None, mobile_phone=None, nick_name=None, organization=None, postal_address=None, preferred_language=None, primary_phone=None, profile_url=None, searches=None, second_email=None, state=None, status=None, street_address=None, timezone=None, title=None, user_type=None, zip_code=None):
        if admin_roles and not isinstance(admin_roles, list):
            raise TypeError("Expected argument 'admin_roles' to be a list")
        __self__.admin_roles = admin_roles
        """
        Administrator roles assigned to user.
        """
        if city and not isinstance(city, str):
            raise TypeError("Expected argument 'city' to be a str")
        __self__.city = city
        """
        user profile property.
        """
        if cost_center and not isinstance(cost_center, str):
            raise TypeError("Expected argument 'cost_center' to be a str")
        __self__.cost_center = cost_center
        """
        user profile property.
        """
        if country_code and not isinstance(country_code, str):
            raise TypeError("Expected argument 'country_code' to be a str")
        __self__.country_code = country_code
        """
        user profile property.
        """
        if custom_profile_attributes and not isinstance(custom_profile_attributes, str):
            raise TypeError("Expected argument 'custom_profile_attributes' to be a str")
        __self__.custom_profile_attributes = custom_profile_attributes
        """
        raw JSON containing all custom profile attributes.
        """
        if department and not isinstance(department, str):
            raise TypeError("Expected argument 'department' to be a str")
        __self__.department = department
        """
        user profile property.
        """
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        user profile property.
        """
        if division and not isinstance(division, str):
            raise TypeError("Expected argument 'division' to be a str")
        __self__.division = division
        """
        user profile property.
        """
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        """
        user profile property.
        """
        if employee_number and not isinstance(employee_number, str):
            raise TypeError("Expected argument 'employee_number' to be a str")
        __self__.employee_number = employee_number
        """
        user profile property.
        """
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        __self__.first_name = first_name
        """
        user profile property.
        """
        if group_memberships and not isinstance(group_memberships, list):
            raise TypeError("Expected argument 'group_memberships' to be a list")
        __self__.group_memberships = group_memberships
        """
        user profile property.
        """
        if honorific_prefix and not isinstance(honorific_prefix, str):
            raise TypeError("Expected argument 'honorific_prefix' to be a str")
        __self__.honorific_prefix = honorific_prefix
        """
        user profile property.
        """
        if honorific_suffix and not isinstance(honorific_suffix, str):
            raise TypeError("Expected argument 'honorific_suffix' to be a str")
        __self__.honorific_suffix = honorific_suffix
        """
        user profile property.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        __self__.last_name = last_name
        """
        user profile property.
        """
        if locale and not isinstance(locale, str):
            raise TypeError("Expected argument 'locale' to be a str")
        __self__.locale = locale
        """
        user profile property.
        """
        if login and not isinstance(login, str):
            raise TypeError("Expected argument 'login' to be a str")
        __self__.login = login
        """
        user profile property.
        """
        if manager and not isinstance(manager, str):
            raise TypeError("Expected argument 'manager' to be a str")
        __self__.manager = manager
        """
        user profile property.
        """
        if manager_id and not isinstance(manager_id, str):
            raise TypeError("Expected argument 'manager_id' to be a str")
        __self__.manager_id = manager_id
        """
        user profile property.
        """
        if middle_name and not isinstance(middle_name, str):
            raise TypeError("Expected argument 'middle_name' to be a str")
        __self__.middle_name = middle_name
        """
        user profile property.
        """
        if mobile_phone and not isinstance(mobile_phone, str):
            raise TypeError("Expected argument 'mobile_phone' to be a str")
        __self__.mobile_phone = mobile_phone
        """
        user profile property.
        """
        if nick_name and not isinstance(nick_name, str):
            raise TypeError("Expected argument 'nick_name' to be a str")
        __self__.nick_name = nick_name
        """
        user profile property.
        """
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        __self__.organization = organization
        """
        user profile property.
        """
        if postal_address and not isinstance(postal_address, str):
            raise TypeError("Expected argument 'postal_address' to be a str")
        __self__.postal_address = postal_address
        """
        user profile property.
        """
        if preferred_language and not isinstance(preferred_language, str):
            raise TypeError("Expected argument 'preferred_language' to be a str")
        __self__.preferred_language = preferred_language
        """
        user profile property.
        """
        if primary_phone and not isinstance(primary_phone, str):
            raise TypeError("Expected argument 'primary_phone' to be a str")
        __self__.primary_phone = primary_phone
        """
        user profile property.
        """
        if profile_url and not isinstance(profile_url, str):
            raise TypeError("Expected argument 'profile_url' to be a str")
        __self__.profile_url = profile_url
        """
        user profile property.
        """
        if searches and not isinstance(searches, list):
            raise TypeError("Expected argument 'searches' to be a list")
        __self__.searches = searches
        if second_email and not isinstance(second_email, str):
            raise TypeError("Expected argument 'second_email' to be a str")
        __self__.second_email = second_email
        """
        user profile property.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        user profile property.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        user profile property.
        """
        if street_address and not isinstance(street_address, str):
            raise TypeError("Expected argument 'street_address' to be a str")
        __self__.street_address = street_address
        """
        user profile property.
        """
        if timezone and not isinstance(timezone, str):
            raise TypeError("Expected argument 'timezone' to be a str")
        __self__.timezone = timezone
        """
        user profile property.
        """
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        __self__.title = title
        """
        user profile property.
        """
        if user_type and not isinstance(user_type, str):
            raise TypeError("Expected argument 'user_type' to be a str")
        __self__.user_type = user_type
        """
        user profile property.
        """
        if zip_code and not isinstance(zip_code, str):
            raise TypeError("Expected argument 'zip_code' to be a str")
        __self__.zip_code = zip_code
        """
        user profile property.
        """


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            admin_roles=self.admin_roles,
            city=self.city,
            cost_center=self.cost_center,
            country_code=self.country_code,
            custom_profile_attributes=self.custom_profile_attributes,
            department=self.department,
            display_name=self.display_name,
            division=self.division,
            email=self.email,
            employee_number=self.employee_number,
            first_name=self.first_name,
            group_memberships=self.group_memberships,
            honorific_prefix=self.honorific_prefix,
            honorific_suffix=self.honorific_suffix,
            id=self.id,
            last_name=self.last_name,
            locale=self.locale,
            login=self.login,
            manager=self.manager,
            manager_id=self.manager_id,
            middle_name=self.middle_name,
            mobile_phone=self.mobile_phone,
            nick_name=self.nick_name,
            organization=self.organization,
            postal_address=self.postal_address,
            preferred_language=self.preferred_language,
            primary_phone=self.primary_phone,
            profile_url=self.profile_url,
            searches=self.searches,
            second_email=self.second_email,
            state=self.state,
            status=self.status,
            street_address=self.street_address,
            timezone=self.timezone,
            title=self.title,
            user_type=self.user_type,
            zip_code=self.zip_code)


def get_user(searches=None, opts=None):
    """
    Use this data source to retrieve a users from Okta.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_okta as okta

    example = okta.user.get_user(searches=[
        {
            "name": "profile.firstName",
            "value": "John",
        },
        {
            "name": "profile.lastName",
            "value": "Doe",
        },
    ])
    ```


    :param list searches: Map of search criteria. It supports the following properties.

    The **searches** object supports the following:

      * `comparison` (`str`) - Comparison to use.
      * `name` (`str`) - Name of property to search against.
      * `value` (`str`) - Value to compare with.
    """
    __args__ = dict()
    __args__['searches'] = searches
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:user/getUser:getUser', __args__, opts=opts).value

    return AwaitableGetUserResult(
        admin_roles=__ret__.get('adminRoles'),
        city=__ret__.get('city'),
        cost_center=__ret__.get('costCenter'),
        country_code=__ret__.get('countryCode'),
        custom_profile_attributes=__ret__.get('customProfileAttributes'),
        department=__ret__.get('department'),
        display_name=__ret__.get('displayName'),
        division=__ret__.get('division'),
        email=__ret__.get('email'),
        employee_number=__ret__.get('employeeNumber'),
        first_name=__ret__.get('firstName'),
        group_memberships=__ret__.get('groupMemberships'),
        honorific_prefix=__ret__.get('honorificPrefix'),
        honorific_suffix=__ret__.get('honorificSuffix'),
        id=__ret__.get('id'),
        last_name=__ret__.get('lastName'),
        locale=__ret__.get('locale'),
        login=__ret__.get('login'),
        manager=__ret__.get('manager'),
        manager_id=__ret__.get('managerId'),
        middle_name=__ret__.get('middleName'),
        mobile_phone=__ret__.get('mobilePhone'),
        nick_name=__ret__.get('nickName'),
        organization=__ret__.get('organization'),
        postal_address=__ret__.get('postalAddress'),
        preferred_language=__ret__.get('preferredLanguage'),
        primary_phone=__ret__.get('primaryPhone'),
        profile_url=__ret__.get('profileUrl'),
        searches=__ret__.get('searches'),
        second_email=__ret__.get('secondEmail'),
        state=__ret__.get('state'),
        status=__ret__.get('status'),
        street_address=__ret__.get('streetAddress'),
        timezone=__ret__.get('timezone'),
        title=__ret__.get('title'),
        user_type=__ret__.get('userType'),
        zip_code=__ret__.get('zipCode'))
