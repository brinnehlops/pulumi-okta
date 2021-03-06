# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetDefaultPoliciesResult',
    'AwaitableGetDefaultPoliciesResult',
    'get_default_policies',
]

@pulumi.output_type
class GetDefaultPoliciesResult:
    """
    A collection of values returned by getDefaultPolicies.
    """
    def __init__(__self__, id=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetDefaultPoliciesResult(GetDefaultPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultPoliciesResult(
            id=self.id,
            type=self.type)


def get_default_policies(type: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultPoliciesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:deprecated/getDefaultPolicies:getDefaultPolicies', __args__, opts=opts, typ=GetDefaultPoliciesResult).value

    return AwaitableGetDefaultPoliciesResult(
        id=__ret__.id,
        type=__ret__.type)
