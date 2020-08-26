# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SamlIdpSigningKey']


class SamlIdpSigningKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 x5cs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a SamlIdpSigningKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] x5cs: base64-encoded X.509 certificate chain with DER encoding
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

            if x5cs is None:
                raise TypeError("Missing required property 'x5cs'")
            __props__['x5cs'] = x5cs
            __props__['created'] = None
            __props__['expires_at'] = None
            __props__['kid'] = None
            __props__['kty'] = None
            __props__['use'] = None
            __props__['x5t_s256'] = None
        super(SamlIdpSigningKey, __self__).__init__(
            'okta:deprecated/samlIdpSigningKey:SamlIdpSigningKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            kid: Optional[pulumi.Input[str]] = None,
            kty: Optional[pulumi.Input[str]] = None,
            use: Optional[pulumi.Input[str]] = None,
            x5cs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            x5t_s256: Optional[pulumi.Input[str]] = None) -> 'SamlIdpSigningKey':
        """
        Get an existing SamlIdpSigningKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] x5cs: base64-encoded X.509 certificate chain with DER encoding
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created"] = created
        __props__["expires_at"] = expires_at
        __props__["kid"] = kid
        __props__["kty"] = kty
        __props__["use"] = use
        __props__["x5cs"] = x5cs
        __props__["x5t_s256"] = x5t_s256
        return SamlIdpSigningKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> str:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> str:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def kid(self) -> str:
        return pulumi.get(self, "kid")

    @property
    @pulumi.getter
    def kty(self) -> str:
        return pulumi.get(self, "kty")

    @property
    @pulumi.getter
    def use(self) -> str:
        return pulumi.get(self, "use")

    @property
    @pulumi.getter
    def x5cs(self) -> List[str]:
        """
        base64-encoded X.509 certificate chain with DER encoding
        """
        return pulumi.get(self, "x5cs")

    @property
    @pulumi.getter(name="x5tS256")
    def x5t_s256(self) -> str:
        return pulumi.get(self, "x5t_s256")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

