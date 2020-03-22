# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServerPolicyClaim(pulumi.CustomResource):
    access_token_lifetime_minutes: pulumi.Output[float]
    """
    Lifetime of access token. Can be set to a value between 5 and 1440.
    """
    auth_server_id: pulumi.Output[str]
    """
    Auth Server ID.
    """
    grant_type_whitelists: pulumi.Output[list]
    """
    Accepted grant type values, `"authorization_code"`, `"implicit"`, `"password"`
    """
    group_blacklists: pulumi.Output[list]
    group_whitelists: pulumi.Output[list]
    inline_hook_id: pulumi.Output[str]
    """
    The ID of the inline token to trigger.
    """
    name: pulumi.Output[str]
    """
    Auth Server Policy Rule name.
    """
    policy_id: pulumi.Output[str]
    """
    Auth Server Policy ID.
    """
    priority: pulumi.Output[float]
    """
    Priority of the auth server policy rule.
    """
    refresh_token_lifetime_minutes: pulumi.Output[float]
    """
    Lifetime of refresh token.
    """
    refresh_token_window_minutes: pulumi.Output[float]
    scope_whitelists: pulumi.Output[list]
    """
    Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
    """
    status: pulumi.Output[str]
    """
    The status of the Auth Server Policy Rule.
    """
    type: pulumi.Output[str]
    """
    The type of the Auth Server Policy Rule.
    """
    user_blacklists: pulumi.Output[list]
    user_whitelists: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, access_token_lifetime_minutes=None, auth_server_id=None, grant_type_whitelists=None, group_blacklists=None, group_whitelists=None, inline_hook_id=None, name=None, policy_id=None, priority=None, refresh_token_lifetime_minutes=None, refresh_token_window_minutes=None, scope_whitelists=None, status=None, type=None, user_blacklists=None, user_whitelists=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Authorization Server Policy Rule.

        This resource allows you to create and configure an Authorization Server Policy Rule.

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/auth_server_policy_rule.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] access_token_lifetime_minutes: Lifetime of access token. Can be set to a value between 5 and 1440.
        :param pulumi.Input[str] auth_server_id: Auth Server ID.
        :param pulumi.Input[list] grant_type_whitelists: Accepted grant type values, `"authorization_code"`, `"implicit"`, `"password"`
        :param pulumi.Input[str] inline_hook_id: The ID of the inline token to trigger.
        :param pulumi.Input[str] name: Auth Server Policy Rule name.
        :param pulumi.Input[str] policy_id: Auth Server Policy ID.
        :param pulumi.Input[float] priority: Priority of the auth server policy rule.
        :param pulumi.Input[float] refresh_token_lifetime_minutes: Lifetime of refresh token.
        :param pulumi.Input[list] scope_whitelists: Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
        :param pulumi.Input[str] status: The status of the Auth Server Policy Rule.
        :param pulumi.Input[str] type: The type of the Auth Server Policy Rule.
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

            __props__['access_token_lifetime_minutes'] = access_token_lifetime_minutes
            if auth_server_id is None:
                raise TypeError("Missing required property 'auth_server_id'")
            __props__['auth_server_id'] = auth_server_id
            if grant_type_whitelists is None:
                raise TypeError("Missing required property 'grant_type_whitelists'")
            __props__['grant_type_whitelists'] = grant_type_whitelists
            __props__['group_blacklists'] = group_blacklists
            __props__['group_whitelists'] = group_whitelists
            __props__['inline_hook_id'] = inline_hook_id
            __props__['name'] = name
            if policy_id is None:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            if priority is None:
                raise TypeError("Missing required property 'priority'")
            __props__['priority'] = priority
            __props__['refresh_token_lifetime_minutes'] = refresh_token_lifetime_minutes
            __props__['refresh_token_window_minutes'] = refresh_token_window_minutes
            __props__['scope_whitelists'] = scope_whitelists
            __props__['status'] = status
            __props__['type'] = type
            __props__['user_blacklists'] = user_blacklists
            __props__['user_whitelists'] = user_whitelists
        super(ServerPolicyClaim, __self__).__init__(
            'okta:auth/serverPolicyClaim:ServerPolicyClaim',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_token_lifetime_minutes=None, auth_server_id=None, grant_type_whitelists=None, group_blacklists=None, group_whitelists=None, inline_hook_id=None, name=None, policy_id=None, priority=None, refresh_token_lifetime_minutes=None, refresh_token_window_minutes=None, scope_whitelists=None, status=None, type=None, user_blacklists=None, user_whitelists=None):
        """
        Get an existing ServerPolicyClaim resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] access_token_lifetime_minutes: Lifetime of access token. Can be set to a value between 5 and 1440.
        :param pulumi.Input[str] auth_server_id: Auth Server ID.
        :param pulumi.Input[list] grant_type_whitelists: Accepted grant type values, `"authorization_code"`, `"implicit"`, `"password"`
        :param pulumi.Input[str] inline_hook_id: The ID of the inline token to trigger.
        :param pulumi.Input[str] name: Auth Server Policy Rule name.
        :param pulumi.Input[str] policy_id: Auth Server Policy ID.
        :param pulumi.Input[float] priority: Priority of the auth server policy rule.
        :param pulumi.Input[float] refresh_token_lifetime_minutes: Lifetime of refresh token.
        :param pulumi.Input[list] scope_whitelists: Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with `"*"`.
        :param pulumi.Input[str] status: The status of the Auth Server Policy Rule.
        :param pulumi.Input[str] type: The type of the Auth Server Policy Rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_token_lifetime_minutes"] = access_token_lifetime_minutes
        __props__["auth_server_id"] = auth_server_id
        __props__["grant_type_whitelists"] = grant_type_whitelists
        __props__["group_blacklists"] = group_blacklists
        __props__["group_whitelists"] = group_whitelists
        __props__["inline_hook_id"] = inline_hook_id
        __props__["name"] = name
        __props__["policy_id"] = policy_id
        __props__["priority"] = priority
        __props__["refresh_token_lifetime_minutes"] = refresh_token_lifetime_minutes
        __props__["refresh_token_window_minutes"] = refresh_token_window_minutes
        __props__["scope_whitelists"] = scope_whitelists
        __props__["status"] = status
        __props__["type"] = type
        __props__["user_blacklists"] = user_blacklists
        __props__["user_whitelists"] = user_whitelists
        return ServerPolicyClaim(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

