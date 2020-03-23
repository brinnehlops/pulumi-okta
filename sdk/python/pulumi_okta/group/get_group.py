# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, description=None, id=None, include_users=None, name=None, users=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        description of group.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if include_users and not isinstance(include_users, bool):
            raise TypeError("Expected argument 'include_users' to be a bool")
        __self__.include_users = include_users
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        name of group.
        """
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        __self__.users = users
        """
        user ids that are members of this group, only included if `include_users` is set to `true`.
        """
class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            description=self.description,
            id=self.id,
            include_users=self.include_users,
            name=self.name,
            users=self.users)

def get_group(include_users=None,name=None,opts=None):
    """
    Use this data source to retrieve a group from Okta.

    > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/group.html.markdown.


    :param bool include_users: whether or not to retrieve all member ids.
    :param str name: name of group to retrieve.
    """
    __args__ = dict()


    __args__['includeUsers'] = include_users
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:group/getGroup:getGroup', __args__, opts=opts).value

    return AwaitableGetGroupResult(
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        include_users=__ret__.get('includeUsers'),
        name=__ret__.get('name'),
        users=__ret__.get('users'))
