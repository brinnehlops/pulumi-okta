# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetEveryoneGroupResult:
    """
    A collection of values returned by getEveryoneGroup.
    """
    def __init__(__self__, id=None, include_users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if include_users and not isinstance(include_users, bool):
            raise TypeError("Expected argument 'include_users' to be a bool")
        __self__.include_users = include_users
class AwaitableGetEveryoneGroupResult(GetEveryoneGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEveryoneGroupResult(
            id=self.id,
            include_users=self.include_users)

def get_everyone_group(include_users=None,opts=None):
    """
    Use this data source to retrieve the Everyone group from Okta. The same can be achieved with the `group.Group` data source with `name = "Everyone"`. This is simply a shortcut.
    """
    __args__ = dict()


    __args__['includeUsers'] = include_users
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:group/getEveryoneGroup:getEveryoneGroup', __args__, opts=opts).value

    return AwaitableGetEveryoneGroupResult(
        id=__ret__.get('id'),
        include_users=__ret__.get('includeUsers'))
