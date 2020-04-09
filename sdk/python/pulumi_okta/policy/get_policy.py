# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetPolicyResult:
    """
    A collection of values returned by getPolicy.
    """
    def __init__(__self__, id=None, name=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        name of policy.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        type of policy.
        """
class AwaitableGetPolicyResult(GetPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyResult(
            id=self.id,
            name=self.name,
            type=self.type)

def get_policy(name=None,type=None,opts=None):
    """
    Use this data source to retrieve a policy from Okta.



    > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/policy.html.markdown.


    :param str name: name of policy to retrieve.
    :param str type: type of policy to retrieve.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:policy/getPolicy:getPolicy', __args__, opts=opts).value

    return AwaitableGetPolicyResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        type=__ret__.get('type'))
