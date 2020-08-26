# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, description=None, id=None, include_users=None, name=None, users=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_users and not isinstance(include_users, bool):
            raise TypeError("Expected argument 'include_users' to be a bool")
        pulumi.set(__self__, "include_users", include_users)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        description of group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeUsers")
    def include_users(self) -> Optional[bool]:
        return pulumi.get(self, "include_users")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        name of group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def users(self) -> List[str]:
        """
        user ids that are members of this group, only included if `include_users` is set to `true`.
        """
        return pulumi.get(self, "users")


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


def get_group(include_users: Optional[bool] = None,
              name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Use this data source to retrieve a group from Okta.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_okta as okta

    example = okta.group.get_group(name="Example App")
    ```


    :param bool include_users: whether or not to retrieve all member ids.
    :param str name: name of group to retrieve.
    """
    __args__ = dict()
    __args__['includeUsers'] = include_users
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:group/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        description=__ret__.description,
        id=__ret__.id,
        include_users=__ret__.include_users,
        name=__ret__.name,
        users=__ret__.users)
