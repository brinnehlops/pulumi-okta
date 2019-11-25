# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PasswordPolicyRule(pulumi.CustomResource):
    name: pulumi.Output[str]
    network_connection: pulumi.Output[str]
    network_excludes: pulumi.Output[list]
    network_includes: pulumi.Output[list]
    password_change: pulumi.Output[str]
    password_reset: pulumi.Output[str]
    password_unlock: pulumi.Output[str]
    policyid: pulumi.Output[str]
    priority: pulumi.Output[float]
    status: pulumi.Output[str]
    users_excludeds: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, name=None, network_connection=None, network_excludes=None, network_includes=None, password_change=None, password_reset=None, password_unlock=None, policyid=None, priority=None, status=None, users_excludeds=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a PasswordPolicyRule resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            __props__['network_connection'] = network_connection
            __props__['network_excludes'] = network_excludes
            __props__['network_includes'] = network_includes
            __props__['password_change'] = password_change
            __props__['password_reset'] = password_reset
            __props__['password_unlock'] = password_unlock
            if policyid is None:
                raise TypeError("Missing required property 'policyid'")
            __props__['policyid'] = policyid
            __props__['priority'] = priority
            __props__['status'] = status
            __props__['users_excludeds'] = users_excludeds
        super(PasswordPolicyRule, __self__).__init__(
            'okta:deprecated/passwordPolicyRule:PasswordPolicyRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, network_connection=None, network_excludes=None, network_includes=None, password_change=None, password_reset=None, password_unlock=None, policyid=None, priority=None, status=None, users_excludeds=None):
        """
        Get an existing PasswordPolicyRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["name"] = name
        __props__["network_connection"] = network_connection
        __props__["network_excludes"] = network_excludes
        __props__["network_includes"] = network_includes
        __props__["password_change"] = password_change
        __props__["password_reset"] = password_reset
        __props__["password_unlock"] = password_unlock
        __props__["policyid"] = policyid
        __props__["priority"] = priority
        __props__["status"] = status
        __props__["users_excludeds"] = users_excludeds
        return PasswordPolicyRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

