# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Signon']


class Signon(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 groups_includeds: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Sign On Policy.

        This resource allows you to create and configure a Sign On Policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.policy.Signon("example",
            description="Example",
            groups_includeds=[data["okta_group"]["everyone"]["id"]],
            status="ACTIVE")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Policy Description.
        :param pulumi.Input[List[pulumi.Input[str]]] groups_includeds: List of Group IDs to Include.
        :param pulumi.Input[str] name: Policy Name.
        :param pulumi.Input[float] priority: Priority of the policy.
        :param pulumi.Input[str] status: Policy Status: `"ACTIVE"` or `"INACTIVE"`.
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

            __props__['description'] = description
            __props__['groups_includeds'] = groups_includeds
            __props__['name'] = name
            __props__['priority'] = priority
            __props__['status'] = status
        super(Signon, __self__).__init__(
            'okta:policy/signon:Signon',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            groups_includeds: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Signon':
        """
        Get an existing Signon resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Policy Description.
        :param pulumi.Input[List[pulumi.Input[str]]] groups_includeds: List of Group IDs to Include.
        :param pulumi.Input[str] name: Policy Name.
        :param pulumi.Input[float] priority: Priority of the policy.
        :param pulumi.Input[str] status: Policy Status: `"ACTIVE"` or `"INACTIVE"`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["groups_includeds"] = groups_includeds
        __props__["name"] = name
        __props__["priority"] = priority
        __props__["status"] = status
        return Signon(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Policy Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupsIncludeds")
    def groups_includeds(self) -> Optional[List[str]]:
        """
        List of Group IDs to Include.
        """
        return pulumi.get(self, "groups_includeds")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Policy Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[float]:
        """
        Priority of the policy.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Policy Status: `"ACTIVE"` or `"INACTIVE"`.
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

