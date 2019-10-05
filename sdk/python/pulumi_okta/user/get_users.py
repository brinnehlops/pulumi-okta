# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, searches=None, users=None, id=None):
        if searches and not isinstance(searches, list):
            raise TypeError("Expected argument 'searches' to be a list")
        __self__.searches = searches
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        __self__.users = users
        """
        collection of users retrieved from Okta with the following properties.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            searches=self.searches,
            users=self.users,
            id=self.id)

def get_users(searches=None,users=None,opts=None):
    """
    Use this data source to retrieve a list of users from Okta.
    
    :param list searches: Map of search criteria to use to find users. It supports the following properties.
    
    The **searches** object supports the following:
    
      * `comparison` (`str`) - Comparison to use.
      * `name` (`str`) - Name of property to search against.
      * `value` (`str`) - Value to compare with.
    
    The **users** object supports the following:
    
      * `admin_roles` (`list`) - Administrator roles assigned to user.
      * `city` (`str`) - user profile property.
      * `cost_center` (`str`) - user profile property.
      * `country_code` (`str`) - user profile property.
      * `custom_profile_attributes` (`str`) - raw JSON containing all custom profile attributes.
      * `department` (`str`) - user profile property.
      * `display_name` (`str`) - user profile property.
      * `division` (`str`) - user profile property.
      * `email` (`str`) - user profile property.
      * `employee_number` (`str`) - user profile property.
      * `first_name` (`str`) - user profile property.
      * `group_memberships` (`list`) - user profile property.
      * `honorific_prefix` (`str`) - user profile property.
      * `honorific_suffix` (`str`) - user profile property.
      * `last_name` (`str`) - user profile property.
      * `locale` (`str`) - user profile property.
      * `login` (`str`) - user profile property.
      * `manager` (`str`) - user profile property.
      * `manager_id` (`str`) - user profile property.
      * `middle_name` (`str`) - user profile property.
      * `mobile_phone` (`str`) - user profile property.
      * `nick_name` (`str`) - user profile property.
      * `organization` (`str`) - user profile property.
      * `postal_address` (`str`) - user profile property.
      * `preferred_language` (`str`) - user profile property.
      * `primary_phone` (`str`) - user profile property.
      * `profile_url` (`str`) - user profile property.
      * `second_email` (`str`) - user profile property.
      * `state` (`str`) - user profile property.
      * `status` (`str`) - user profile property.
      * `street_address` (`str`) - user profile property.
      * `timezone` (`str`) - user profile property.
      * `title` (`str`) - user profile property.
      * `user_type` (`str`) - user profile property.
      * `zip_code` (`str`) - user profile property.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/d/users.html.markdown.
    """
    __args__ = dict()

    __args__['searches'] = searches
    __args__['users'] = users
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:user/getUsers:getUsers', __args__, opts=opts).value

    return AwaitableGetUsersResult(
        searches=__ret__.get('searches'),
        users=__ret__.get('users'),
        id=__ret__.get('id'))
