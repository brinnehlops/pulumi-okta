# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RulePassword(pulumi.CustomResource):
    access: pulumi.Output[str]
    authtype: pulumi.Output[str]
    enroll: pulumi.Output[str]
    mfa_lifetime: pulumi.Output[float]
    mfa_prompt: pulumi.Output[str]
    mfa_remember_device: pulumi.Output[bool]
    mfa_required: pulumi.Output[bool]
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
    password_change: pulumi.Output[str]
    """
    Allow or deny a user to change their password: `"ALLOW"` or `"DENY"`. By default it is `"ALLOW"`.
    """
    password_reset: pulumi.Output[str]
    """
    Allow or deny a user to reset their password: `"ALLOW"` or `"DENY"`. By default it is `"ALLOW"`.
    """
    password_unlock: pulumi.Output[str]
    """
    Allow or deny a user to unlock: `"ALLOW"` or `"DENY"`. By default it is `"DENY"`,
    """
    policyid: pulumi.Output[str]
    """
    Policy ID.
    """
    priority: pulumi.Output[float]
    """
    Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
    """
    session_idle: pulumi.Output[float]
    session_lifetime: pulumi.Output[float]
    session_persistent: pulumi.Output[bool]
    status: pulumi.Output[str]
    """
    Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
    """
    users_excludeds: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, access=None, authtype=None, enroll=None, mfa_lifetime=None, mfa_prompt=None, mfa_remember_device=None, mfa_required=None, name=None, network_connection=None, network_excludes=None, network_includes=None, password_change=None, password_reset=None, password_unlock=None, policyid=None, priority=None, session_idle=None, session_lifetime=None, session_persistent=None, status=None, users_excludeds=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Password Policy Rule.
        
        This resource allows you to create and configure a Password Policy Rule.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[list] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[list] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] password_change: Allow or deny a user to change their password: `"ALLOW"` or `"DENY"`. By default it is `"ALLOW"`.
        :param pulumi.Input[str] password_reset: Allow or deny a user to reset their password: `"ALLOW"` or `"DENY"`. By default it is `"ALLOW"`.
        :param pulumi.Input[str] password_unlock: Allow or deny a user to unlock: `"ALLOW"` or `"DENY"`. By default it is `"DENY"`,
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_password.html.markdown.
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

            __props__['access'] = access
            __props__['authtype'] = authtype
            __props__['enroll'] = enroll
            __props__['mfa_lifetime'] = mfa_lifetime
            __props__['mfa_prompt'] = mfa_prompt
            __props__['mfa_remember_device'] = mfa_remember_device
            __props__['mfa_required'] = mfa_required
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
            __props__['session_idle'] = session_idle
            __props__['session_lifetime'] = session_lifetime
            __props__['session_persistent'] = session_persistent
            __props__['status'] = status
            __props__['users_excludeds'] = users_excludeds
        super(RulePassword, __self__).__init__(
            'okta:policy/rulePassword:RulePassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access=None, authtype=None, enroll=None, mfa_lifetime=None, mfa_prompt=None, mfa_remember_device=None, mfa_required=None, name=None, network_connection=None, network_excludes=None, network_includes=None, password_change=None, password_reset=None, password_unlock=None, policyid=None, priority=None, session_idle=None, session_lifetime=None, session_persistent=None, status=None, users_excludeds=None):
        """
        Get an existing RulePassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[list] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[list] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] password_change: Allow or deny a user to change their password: `"ALLOW"` or `"DENY"`. By default it is `"ALLOW"`.
        :param pulumi.Input[str] password_reset: Allow or deny a user to reset their password: `"ALLOW"` or `"DENY"`. By default it is `"ALLOW"`.
        :param pulumi.Input[str] password_unlock: Allow or deny a user to unlock: `"ALLOW"` or `"DENY"`. By default it is `"DENY"`,
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_password.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access"] = access
        __props__["authtype"] = authtype
        __props__["enroll"] = enroll
        __props__["mfa_lifetime"] = mfa_lifetime
        __props__["mfa_prompt"] = mfa_prompt
        __props__["mfa_remember_device"] = mfa_remember_device
        __props__["mfa_required"] = mfa_required
        __props__["name"] = name
        __props__["network_connection"] = network_connection
        __props__["network_excludes"] = network_excludes
        __props__["network_includes"] = network_includes
        __props__["password_change"] = password_change
        __props__["password_reset"] = password_reset
        __props__["password_unlock"] = password_unlock
        __props__["policyid"] = policyid
        __props__["priority"] = priority
        __props__["session_idle"] = session_idle
        __props__["session_lifetime"] = session_lifetime
        __props__["session_persistent"] = session_persistent
        __props__["status"] = status
        __props__["users_excludeds"] = users_excludeds
        return RulePassword(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

