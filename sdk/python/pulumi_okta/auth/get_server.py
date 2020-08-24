# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetServerResult',
    'AwaitableGetServerResult',
    'get_server',
]

@pulumi.output_type
class GetServerResult:
    """
    A collection of values returned by getServer.
    """
    def __init__(__self__, audiences=None, credentials_last_rotated=None, credentials_next_rotation=None, credentials_rotation_mode=None, description=None, id=None, kid=None, name=None, status=None):
        if audiences and not isinstance(audiences, list):
            raise TypeError("Expected argument 'audiences' to be a list")
        pulumi.set(__self__, "audiences", audiences)
        if credentials_last_rotated and not isinstance(credentials_last_rotated, str):
            raise TypeError("Expected argument 'credentials_last_rotated' to be a str")
        pulumi.set(__self__, "credentials_last_rotated", credentials_last_rotated)
        if credentials_next_rotation and not isinstance(credentials_next_rotation, str):
            raise TypeError("Expected argument 'credentials_next_rotation' to be a str")
        pulumi.set(__self__, "credentials_next_rotation", credentials_next_rotation)
        if credentials_rotation_mode and not isinstance(credentials_rotation_mode, str):
            raise TypeError("Expected argument 'credentials_rotation_mode' to be a str")
        pulumi.set(__self__, "credentials_rotation_mode", credentials_rotation_mode)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kid and not isinstance(kid, str):
            raise TypeError("Expected argument 'kid' to be a str")
        pulumi.set(__self__, "kid", kid)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def audiences(self) -> List[str]:
        """
        array of audiences,
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter(name="credentialsLastRotated")
    def credentials_last_rotated(self) -> str:
        """
        last time credentials were rotated.
        """
        return pulumi.get(self, "credentials_last_rotated")

    @property
    @pulumi.getter(name="credentialsNextRotation")
    def credentials_next_rotation(self) -> str:
        """
        next time credentials will be rotated
        """
        return pulumi.get(self, "credentials_next_rotation")

    @property
    @pulumi.getter(name="credentialsRotationMode")
    def credentials_rotation_mode(self) -> str:
        """
        mode of credential rotation, auto or manual.
        """
        return pulumi.get(self, "credentials_rotation_mode")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        description of Authorization server.
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
    @pulumi.getter
    def kid(self) -> str:
        """
        auth server key id.
        """
        return pulumi.get(self, "kid")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the auth server.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        the activation status of the authorization server.
        """
        return pulumi.get(self, "status")


class AwaitableGetServerResult(GetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerResult(
            audiences=self.audiences,
            credentials_last_rotated=self.credentials_last_rotated,
            credentials_next_rotation=self.credentials_next_rotation,
            credentials_rotation_mode=self.credentials_rotation_mode,
            description=self.description,
            id=self.id,
            kid=self.kid,
            name=self.name,
            status=self.status)


def get_server(name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerResult:
    """
    Use this data source to retrieve an auth server from Okta.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_okta as okta

    example = okta.auth.get_server(name="Example Auth")
    ```


    :param str name: The name of the auth server to retrieve.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:auth/getServer:getServer', __args__, opts=opts, typ=GetServerResult).value

    return AwaitableGetServerResult(
        audiences=__ret__.audiences,
        credentials_last_rotated=__ret__.credentials_last_rotated,
        credentials_next_rotation=__ret__.credentials_next_rotation,
        credentials_rotation_mode=__ret__.credentials_rotation_mode,
        description=__ret__.description,
        id=__ret__.id,
        kid=__ret__.kid,
        name=__ret__.name,
        status=__ret__.status)
