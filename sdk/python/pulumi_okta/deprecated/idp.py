# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Idp(pulumi.CustomResource):
    account_link_action: pulumi.Output[str]
    account_link_group_includes: pulumi.Output[list]
    acs_binding: pulumi.Output[str]
    acs_type: pulumi.Output[str]
    authorization_binding: pulumi.Output[str]
    authorization_url: pulumi.Output[str]
    client_id: pulumi.Output[str]
    client_secret: pulumi.Output[str]
    deprovisioned_action: pulumi.Output[str]
    groups_action: pulumi.Output[str]
    groups_assignments: pulumi.Output[list]
    groups_attribute: pulumi.Output[str]
    groups_filters: pulumi.Output[list]
    issuer_mode: pulumi.Output[str]
    """
    Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
    """
    issuer_url: pulumi.Output[str]
    jwks_binding: pulumi.Output[str]
    jwks_url: pulumi.Output[str]
    max_clock_skew: pulumi.Output[float]
    name: pulumi.Output[str]
    """
    name of idp
    """
    profile_master: pulumi.Output[bool]
    protocol_type: pulumi.Output[str]
    provisioning_action: pulumi.Output[str]
    request_signature_algorithm: pulumi.Output[str]
    """
    algorithm to use to sign requests
    """
    request_signature_scope: pulumi.Output[str]
    """
    algorithm to use to sign response
    """
    response_signature_algorithm: pulumi.Output[str]
    """
    algorithm to use to sign requests
    """
    response_signature_scope: pulumi.Output[str]
    """
    algorithm to use to sign response
    """
    scopes: pulumi.Output[list]
    status: pulumi.Output[str]
    subject_match_attribute: pulumi.Output[str]
    subject_match_type: pulumi.Output[str]
    suspended_action: pulumi.Output[str]
    token_binding: pulumi.Output[str]
    token_url: pulumi.Output[str]
    type: pulumi.Output[str]
    user_info_binding: pulumi.Output[str]
    user_info_url: pulumi.Output[str]
    username_template: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, account_link_action=None, account_link_group_includes=None, acs_binding=None, acs_type=None, authorization_binding=None, authorization_url=None, client_id=None, client_secret=None, deprovisioned_action=None, groups_action=None, groups_assignments=None, groups_attribute=None, groups_filters=None, issuer_mode=None, issuer_url=None, jwks_binding=None, jwks_url=None, max_clock_skew=None, name=None, profile_master=None, protocol_type=None, provisioning_action=None, request_signature_algorithm=None, request_signature_scope=None, response_signature_algorithm=None, response_signature_scope=None, scopes=None, status=None, subject_match_attribute=None, subject_match_type=None, suspended_action=None, token_binding=None, token_url=None, user_info_binding=None, user_info_url=None, username_template=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Idp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] issuer_mode: Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
        :param pulumi.Input[str] name: name of idp
        :param pulumi.Input[str] request_signature_algorithm: algorithm to use to sign requests
        :param pulumi.Input[str] request_signature_scope: algorithm to use to sign response
        :param pulumi.Input[str] response_signature_algorithm: algorithm to use to sign requests
        :param pulumi.Input[str] response_signature_scope: algorithm to use to sign response
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

            __props__['account_link_action'] = account_link_action
            __props__['account_link_group_includes'] = account_link_group_includes
            if acs_binding is None:
                raise TypeError("Missing required property 'acs_binding'")
            __props__['acs_binding'] = acs_binding
            __props__['acs_type'] = acs_type
            if authorization_binding is None:
                raise TypeError("Missing required property 'authorization_binding'")
            __props__['authorization_binding'] = authorization_binding
            if authorization_url is None:
                raise TypeError("Missing required property 'authorization_url'")
            __props__['authorization_url'] = authorization_url
            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['deprovisioned_action'] = deprovisioned_action
            __props__['groups_action'] = groups_action
            __props__['groups_assignments'] = groups_assignments
            __props__['groups_attribute'] = groups_attribute
            __props__['groups_filters'] = groups_filters
            __props__['issuer_mode'] = issuer_mode
            if issuer_url is None:
                raise TypeError("Missing required property 'issuer_url'")
            __props__['issuer_url'] = issuer_url
            if jwks_binding is None:
                raise TypeError("Missing required property 'jwks_binding'")
            __props__['jwks_binding'] = jwks_binding
            if jwks_url is None:
                raise TypeError("Missing required property 'jwks_url'")
            __props__['jwks_url'] = jwks_url
            __props__['max_clock_skew'] = max_clock_skew
            __props__['name'] = name
            __props__['profile_master'] = profile_master
            __props__['protocol_type'] = protocol_type
            __props__['provisioning_action'] = provisioning_action
            __props__['request_signature_algorithm'] = request_signature_algorithm
            __props__['request_signature_scope'] = request_signature_scope
            __props__['response_signature_algorithm'] = response_signature_algorithm
            __props__['response_signature_scope'] = response_signature_scope
            if scopes is None:
                raise TypeError("Missing required property 'scopes'")
            __props__['scopes'] = scopes
            __props__['status'] = status
            __props__['subject_match_attribute'] = subject_match_attribute
            __props__['subject_match_type'] = subject_match_type
            __props__['suspended_action'] = suspended_action
            if token_binding is None:
                raise TypeError("Missing required property 'token_binding'")
            __props__['token_binding'] = token_binding
            if token_url is None:
                raise TypeError("Missing required property 'token_url'")
            __props__['token_url'] = token_url
            __props__['user_info_binding'] = user_info_binding
            __props__['user_info_url'] = user_info_url
            __props__['username_template'] = username_template
            __props__['type'] = None
        super(Idp, __self__).__init__(
            'okta:deprecated/idp:Idp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_link_action=None, account_link_group_includes=None, acs_binding=None, acs_type=None, authorization_binding=None, authorization_url=None, client_id=None, client_secret=None, deprovisioned_action=None, groups_action=None, groups_assignments=None, groups_attribute=None, groups_filters=None, issuer_mode=None, issuer_url=None, jwks_binding=None, jwks_url=None, max_clock_skew=None, name=None, profile_master=None, protocol_type=None, provisioning_action=None, request_signature_algorithm=None, request_signature_scope=None, response_signature_algorithm=None, response_signature_scope=None, scopes=None, status=None, subject_match_attribute=None, subject_match_type=None, suspended_action=None, token_binding=None, token_url=None, type=None, user_info_binding=None, user_info_url=None, username_template=None):
        """
        Get an existing Idp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] issuer_mode: Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL
        :param pulumi.Input[str] name: name of idp
        :param pulumi.Input[str] request_signature_algorithm: algorithm to use to sign requests
        :param pulumi.Input[str] request_signature_scope: algorithm to use to sign response
        :param pulumi.Input[str] response_signature_algorithm: algorithm to use to sign requests
        :param pulumi.Input[str] response_signature_scope: algorithm to use to sign response
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_link_action"] = account_link_action
        __props__["account_link_group_includes"] = account_link_group_includes
        __props__["acs_binding"] = acs_binding
        __props__["acs_type"] = acs_type
        __props__["authorization_binding"] = authorization_binding
        __props__["authorization_url"] = authorization_url
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["deprovisioned_action"] = deprovisioned_action
        __props__["groups_action"] = groups_action
        __props__["groups_assignments"] = groups_assignments
        __props__["groups_attribute"] = groups_attribute
        __props__["groups_filters"] = groups_filters
        __props__["issuer_mode"] = issuer_mode
        __props__["issuer_url"] = issuer_url
        __props__["jwks_binding"] = jwks_binding
        __props__["jwks_url"] = jwks_url
        __props__["max_clock_skew"] = max_clock_skew
        __props__["name"] = name
        __props__["profile_master"] = profile_master
        __props__["protocol_type"] = protocol_type
        __props__["provisioning_action"] = provisioning_action
        __props__["request_signature_algorithm"] = request_signature_algorithm
        __props__["request_signature_scope"] = request_signature_scope
        __props__["response_signature_algorithm"] = response_signature_algorithm
        __props__["response_signature_scope"] = response_signature_scope
        __props__["scopes"] = scopes
        __props__["status"] = status
        __props__["subject_match_attribute"] = subject_match_attribute
        __props__["subject_match_type"] = subject_match_type
        __props__["suspended_action"] = suspended_action
        __props__["token_binding"] = token_binding
        __props__["token_url"] = token_url
        __props__["type"] = type
        __props__["user_info_binding"] = user_info_binding
        __props__["user_info_url"] = user_info_url
        __props__["username_template"] = username_template
        return Idp(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
