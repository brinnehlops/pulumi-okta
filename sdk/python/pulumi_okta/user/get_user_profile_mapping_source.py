# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetUserProfileMappingSourceResult:
    """
    A collection of values returned by getUserProfileMappingSource.
    """
    def __init__(__self__, id=None, name=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        name of source.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        type of source.
        """
class AwaitableGetUserProfileMappingSourceResult(GetUserProfileMappingSourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserProfileMappingSourceResult(
            id=self.id,
            name=self.name,
            type=self.type)

def get_user_profile_mapping_source(opts=None):
    """
    Use this data source to retrieve the base user Profile Mapping source or target from Okta.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_okta as okta

    example = okta.user.get_user_profile_mapping_source()
    ```
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:user/getUserProfileMappingSource:getUserProfileMappingSource', __args__, opts=opts).value

    return AwaitableGetUserProfileMappingSourceResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        type=__ret__.get('type'))
