# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RuleIdpDiscovery(pulumi.CustomResource):
    app_excludes: pulumi.Output[list]
    """
    Applications to exclude in discovery rule

      * `id` (`str`) - ID of the Rule.
      * `name` (`str`) - Policy Rule Name.
      * `type` (`str`)
    """
    app_includes: pulumi.Output[list]
    """
    Applications to include in discovery rule

      * `id` (`str`) - ID of the Rule.
      * `name` (`str`) - Policy Rule Name.
      * `type` (`str`)
    """
    idp_id: pulumi.Output[str]
    idp_type: pulumi.Output[str]
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
    platform_includes: pulumi.Output[list]
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
    user_identifier_attribute: pulumi.Output[str]
    user_identifier_patterns: pulumi.Output[list]
    user_identifier_type: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, app_excludes=None, app_includes=None, idp_id=None, idp_type=None, name=None, network_connection=None, network_excludes=None, network_includes=None, platform_includes=None, policyid=None, priority=None, status=None, user_identifier_attribute=None, user_identifier_patterns=None, user_identifier_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an IdP Discovery Policy Rule.

        This resource allows you to create and configure an IdP Discovery Policy Rule.



        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_rule_idp_discovery.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_excludes: Applications to exclude in discovery rule
        :param pulumi.Input[list] app_includes: Applications to include in discovery rule
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[list] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[list] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.

        The **app_excludes** object supports the following:

          * `id` (`pulumi.Input[str]`) - ID of the Rule.
          * `name` (`pulumi.Input[str]`) - Policy Rule Name.
          * `type` (`pulumi.Input[str]`)

        The **app_includes** object supports the following:

          * `id` (`pulumi.Input[str]`) - ID of the Rule.
          * `name` (`pulumi.Input[str]`) - Policy Rule Name.
          * `type` (`pulumi.Input[str]`)

        The **platform_includes** object supports the following:

          * `osExpression` (`pulumi.Input[str]`)
          * `osType` (`pulumi.Input[str]`)
          * `type` (`pulumi.Input[str]`)

        The **user_identifier_patterns** object supports the following:

          * `match_type` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
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
    def get(resource_name, id, opts=None, app_excludes=None, app_includes=None, idp_id=None, idp_type=None, name=None, network_connection=None, network_excludes=None, network_includes=None, platform_includes=None, policyid=None, priority=None, status=None, user_identifier_attribute=None, user_identifier_patterns=None, user_identifier_type=None):
        """
        Get an existing RuleIdpDiscovery resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_excludes: Applications to exclude in discovery rule
        :param pulumi.Input[list] app_includes: Applications to include in discovery rule
        :param pulumi.Input[str] name: Policy Rule Name.
        :param pulumi.Input[str] network_connection: Network selection mode: `"ANYWHERE"`, `"ZONE"`, `"ON_NETWORK"`, or `"OFF_NETWORK"`.
        :param pulumi.Input[list] network_excludes: The network zones to exclude. Conflicts with `network_includes`.
        :param pulumi.Input[list] network_includes: The network zones to include. Conflicts with `network_excludes`.
        :param pulumi.Input[str] policyid: Policy ID.
        :param pulumi.Input[float] priority: Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Rule Status: `"ACTIVE"` or `"INACTIVE"`.

        The **app_excludes** object supports the following:

          * `id` (`pulumi.Input[str]`) - ID of the Rule.
          * `name` (`pulumi.Input[str]`) - Policy Rule Name.
          * `type` (`pulumi.Input[str]`)

        The **app_includes** object supports the following:

          * `id` (`pulumi.Input[str]`) - ID of the Rule.
          * `name` (`pulumi.Input[str]`) - Policy Rule Name.
          * `type` (`pulumi.Input[str]`)

        The **platform_includes** object supports the following:

          * `osExpression` (`pulumi.Input[str]`)
          * `osType` (`pulumi.Input[str]`)
          * `type` (`pulumi.Input[str]`)

        The **user_identifier_patterns** object supports the following:

          * `match_type` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

