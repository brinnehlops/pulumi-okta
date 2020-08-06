# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class RuleMfa(pulumi.CustomResource):
    enroll: pulumi.Output[str]
    """
    When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
    """
    name: pulumi.Output[str]
    """
    Policy Rule Name.
    """
    network_connection: pulumi.Output[str]
    """
    Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
    """
    network_excludes: pulumi.Output[list]
    """
    The network zones to exclude. Conflicts with `network_includes`.
    """
    network_includes: pulumi.Output[list]
    """
    The network zones to include. Conflicts with `network_excludes`.
    """
    policyid: pulumi.Output[str]
    """
    Policy ID.
    """
    priority: pulumi.Output[float]
    """
    Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
    """
    status: pulumi.Output[str]
    """
    Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
    """
    users_excludeds: pulumi.Output[list]
    """
    Set of User IDs to Exclude
    """
    def __init__(__self__, resource_name, opts=None, enroll=None, name=None, network_connection=None, network_excludes=None, network_includes=None, policyid=None, priority=None, status=None, users_excludeds=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an MFA Policy Rule.

        This resource allows you to create and configure an MFA Policy Rule.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enroll: When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[list] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[list] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        :param pulumi.Input[list] users_excludeds: Set of User IDs to Exclude
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

            __props__['enroll'] = enroll
            __props__['name'] = name
            __props__['network_connection'] = network_connection
            __props__['network_excludes'] = network_excludes
            __props__['network_includes'] = network_includes
            if policyid is None:
                raise TypeError("Missing required property 'policyid'")
            __props__['policyid'] = policyid
            __props__['priority'] = priority
            __props__['status'] = status
            __props__['users_excludeds'] = users_excludeds
        super(RuleMfa, __self__).__init__(
            'okta:policy/ruleMfa:RuleMfa',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, enroll=None, name=None, network_connection=None, network_excludes=None, network_includes=None, policyid=None, priority=None, status=None, users_excludeds=None):
        """
        Get an existing RuleMfa resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enroll: When a user should be prompted for MFA. It can be `"CHALLENGE"`, `"LOGIN"`, or `"NEVER"`.
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[list] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[list] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        :param pulumi.Input[list] users_excludeds: Set of User IDs to Exclude
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enroll"] = enroll
        __props__["name"] = name
        __props__["network_connection"] = network_connection
        __props__["network_excludes"] = network_excludes
        __props__["network_includes"] = network_includes
        __props__["policyid"] = policyid
        __props__["priority"] = priority
        __props__["status"] = status
        __props__["users_excludeds"] = users_excludeds
        return RuleMfa(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
