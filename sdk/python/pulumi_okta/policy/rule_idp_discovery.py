# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RuleIdpDiscovery']


class RuleIdpDiscovery(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_excludes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppExcludeArgs']]]]] = None,
                 app_includes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppIncludeArgs']]]]] = None,
                 idp_id: Optional[pulumi.Input[str]] = None,
                 idp_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_connection: Optional[pulumi.Input[str]] = None,
                 network_excludes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 network_includes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 platform_includes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryPlatformIncludeArgs']]]]] = None,
                 policyid: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user_identifier_attribute: Optional[pulumi.Input[str]] = None,
                 user_identifier_patterns: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryUserIdentifierPatternArgs']]]]] = None,
                 user_identifier_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an IdP Discovery Policy Rule.

        This resource allows you to create and configure an IdP Discovery Policy Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.policy.RuleIdpDiscovery("example",
            idp_id="<idp id>",
            idp_type="SAML2",
            policyid="<policy id>",
            priority=1,
            user_identifier_attribute="company",
            user_identifier_patterns=[okta.policy.RuleIdpDiscoveryUserIdentifierPatternArgs(
                match_type="EQUALS",
                value="Articulate",
            )],
            user_identifier_type="ATTRIBUTE")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppExcludeArgs']]]] app_excludes: Applications to exclude in discovery rule
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppIncludeArgs']]]] app_includes: Applications to include in discovery rule
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[List[pulumi.Input[str]]] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[List[pulumi.Input[str]]] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
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

            __props__['app_excludes'] = app_excludes
            __props__['app_includes'] = app_includes
            __props__['idp_id'] = idp_id
            __props__['idp_type'] = idp_type
            __props__['name'] = name
            __props__['network_connection'] = network_connection
            __props__['network_excludes'] = network_excludes
            __props__['network_includes'] = network_includes
            __props__['platform_includes'] = platform_includes
            if policyid is None:
                raise TypeError("Missing required property 'policyid'")
            __props__['policyid'] = policyid
            __props__['priority'] = priority
            __props__['status'] = status
            __props__['user_identifier_attribute'] = user_identifier_attribute
            __props__['user_identifier_patterns'] = user_identifier_patterns
            __props__['user_identifier_type'] = user_identifier_type
        super(RuleIdpDiscovery, __self__).__init__(
            'okta:policy/ruleIdpDiscovery:RuleIdpDiscovery',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_excludes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppExcludeArgs']]]]] = None,
            app_includes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppIncludeArgs']]]]] = None,
            idp_id: Optional[pulumi.Input[str]] = None,
            idp_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_connection: Optional[pulumi.Input[str]] = None,
            network_excludes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            network_includes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            platform_includes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryPlatformIncludeArgs']]]]] = None,
            policyid: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            status: Optional[pulumi.Input[str]] = None,
            user_identifier_attribute: Optional[pulumi.Input[str]] = None,
            user_identifier_patterns: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryUserIdentifierPatternArgs']]]]] = None,
            user_identifier_type: Optional[pulumi.Input[str]] = None) -> 'RuleIdpDiscovery':
        """
        Get an existing RuleIdpDiscovery resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppExcludeArgs']]]] app_excludes: Applications to exclude in discovery rule
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleIdpDiscoveryAppIncludeArgs']]]] app_includes: Applications to include in discovery rule
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[List[pulumi.Input[str]]] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[List[pulumi.Input[str]]] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_excludes"] = app_excludes
        __props__["app_includes"] = app_includes
        __props__["idp_id"] = idp_id
        __props__["idp_type"] = idp_type
        __props__["name"] = name
        __props__["network_connection"] = network_connection
        __props__["network_excludes"] = network_excludes
        __props__["network_includes"] = network_includes
        __props__["platform_includes"] = platform_includes
        __props__["policyid"] = policyid
        __props__["priority"] = priority
        __props__["status"] = status
        __props__["user_identifier_attribute"] = user_identifier_attribute
        __props__["user_identifier_patterns"] = user_identifier_patterns
        __props__["user_identifier_type"] = user_identifier_type
        return RuleIdpDiscovery(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appExcludes")
    def app_excludes(self) -> pulumi.Output[Optional[List['outputs.RuleIdpDiscoveryAppExclude']]]:
        """
        Applications to exclude in discovery rule
        """
        return pulumi.get(self, "app_excludes")

    @property
    @pulumi.getter(name="appIncludes")
    def app_includes(self) -> pulumi.Output[Optional[List['outputs.RuleIdpDiscoveryAppInclude']]]:
        """
        Applications to include in discovery rule
        """
        return pulumi.get(self, "app_includes")

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "idp_id")

    @property
    @pulumi.getter(name="idpType")
    def idp_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "idp_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Policy Rule Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConnection")
    def network_connection(self) -> pulumi.Output[Optional[str]]:
        """
        Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        """
        return pulumi.get(self, "network_connection")

    @property
    @pulumi.getter(name="networkExcludes")
    def network_excludes(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The network zones to exclude. Conflicts with `network_includes`.
        """
        return pulumi.get(self, "network_excludes")

    @property
    @pulumi.getter(name="networkIncludes")
    def network_includes(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The network zones to include. Conflicts with `network_excludes`.
        """
        return pulumi.get(self, "network_includes")

    @property
    @pulumi.getter(name="platformIncludes")
    def platform_includes(self) -> pulumi.Output[Optional[List['outputs.RuleIdpDiscoveryPlatformInclude']]]:
        return pulumi.get(self, "platform_includes")

    @property
    @pulumi.getter
    def policyid(self) -> pulumi.Output[str]:
        """
        Policy ID.
        """
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[float]]:
        """
        Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userIdentifierAttribute")
    def user_identifier_attribute(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_identifier_attribute")

    @property
    @pulumi.getter(name="userIdentifierPatterns")
    def user_identifier_patterns(self) -> pulumi.Output[Optional[List['outputs.RuleIdpDiscoveryUserIdentifierPattern']]]:
        return pulumi.get(self, "user_identifier_patterns")

    @property
    @pulumi.getter(name="userIdentifierType")
    def user_identifier_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_identifier_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

