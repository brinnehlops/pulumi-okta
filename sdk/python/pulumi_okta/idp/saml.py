# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Saml(pulumi.CustomResource):
    account_link_action: pulumi.Output[str]
    """
    Specifies the account linking action for an IdP user.
    """
    account_link_group_includes: pulumi.Output[list]
    """
    Group memberships to determine link candidates.
    """
    acs_binding: pulumi.Output[str]
    """
    The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
    """
    acs_type: pulumi.Output[str]
    """
    The type of ACS. It can be `"INSTANCE"` or `"ORG"`.
    """
    audience: pulumi.Output[str]
    """
    The audience restriction for the IdP.
    """
    deprovisioned_action: pulumi.Output[str]
    """
    Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
    """
    groups_action: pulumi.Output[str]
    """
    Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
    """
    groups_assignments: pulumi.Output[list]
    """
    List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
    """
    groups_attribute: pulumi.Output[str]
    """
    IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
    """
    groups_filters: pulumi.Output[list]
    """
    Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
    """
    issuer: pulumi.Output[str]
    """
    URI that identifies the issuer.
    """
    issuer_mode: pulumi.Output[str]
    """
    Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
    """
    kid: pulumi.Output[str]
    """
    The ID of the signing key.
    """
    name: pulumi.Output[str]
    """
    The Application's display name.
    """
    name_format: pulumi.Output[str]
    """
    The name identifier format to use. By default `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"`.
    """
    profile_master: pulumi.Output[bool]
    """
    Determines if the IdP should act as a source of truth for user profile attributes.
    """
    provisioning_action: pulumi.Output[str]
    """
    Provisioning action for an IdP user during authentication.
    """
    request_signature_algorithm: pulumi.Output[str]
    """
    The XML digital signature algorithm used when signing an AuthnRequest message.
    """
    request_signature_scope: pulumi.Output[str]
    """
    Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
    """
    response_signature_algorithm: pulumi.Output[str]
    """
    The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
    """
    response_signature_scope: pulumi.Output[str]
    """
    Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
    """
    sso_binding: pulumi.Output[str]
    """
    The method of making an SSO request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
    """
    sso_destination: pulumi.Output[str]
    """
    URI reference indicating the address to which the AuthnRequest message is sent.
    """
    sso_url: pulumi.Output[str]
    """
    URL of binding-specific endpoint to send an AuthnRequest message to IdP.
    """
    status: pulumi.Output[str]
    """
    Status of the IdP.
    """
    subject_filter: pulumi.Output[str]
    """
    Optional regular expression pattern used to filter untrusted IdP usernames.
    """
    subject_formats: pulumi.Output[list]
    """
    The name formate. By default `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"`.
    """
    subject_match_attribute: pulumi.Output[str]
    """
    Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
    """
    subject_match_type: pulumi.Output[str]
    """
    Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
    """
    suspended_action: pulumi.Output[str]
    """
    Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
    """
    type: pulumi.Output[str]
    """
    Type of the IdP.
    """
    username_template: pulumi.Output[str]
    """
    Okta EL Expression to generate or transform a unique username for the IdP user.
    """
    def __init__(__self__, resource_name, opts=None, account_link_action=None, account_link_group_includes=None, acs_binding=None, acs_type=None, deprovisioned_action=None, groups_action=None, groups_assignments=None, groups_attribute=None, groups_filters=None, issuer=None, issuer_mode=None, kid=None, name=None, name_format=None, profile_master=None, provisioning_action=None, request_signature_algorithm=None, request_signature_scope=None, response_signature_algorithm=None, response_signature_scope=None, sso_binding=None, sso_destination=None, sso_url=None, status=None, subject_filter=None, subject_formats=None, subject_match_attribute=None, subject_match_type=None, suspended_action=None, username_template=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a SAML Identity Provider.
        
        This resource allows you to create and configure a SAML Identity Provider.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_link_action: Specifies the account linking action for an IdP user.
        :param pulumi.Input[list] account_link_group_includes: Group memberships to determine link candidates.
        :param pulumi.Input[str] acs_binding: The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        :param pulumi.Input[str] acs_type: The type of ACS. It can be `"INSTANCE"` or `"ORG"`.
        :param pulumi.Input[str] deprovisioned_action: Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        :param pulumi.Input[str] groups_action: Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        :param pulumi.Input[list] groups_assignments: List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        :param pulumi.Input[str] groups_attribute: IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        :param pulumi.Input[list] groups_filters: Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        :param pulumi.Input[str] issuer: URI that identifies the issuer.
        :param pulumi.Input[str] issuer_mode: Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        :param pulumi.Input[str] kid: The ID of the signing key.
        :param pulumi.Input[str] name: The Application's display name.
        :param pulumi.Input[str] name_format: The name identifier format to use. By default `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"`.
        :param pulumi.Input[bool] profile_master: Determines if the IdP should act as a source of truth for user profile attributes.
        :param pulumi.Input[str] provisioning_action: Provisioning action for an IdP user during authentication.
        :param pulumi.Input[str] request_signature_algorithm: The XML digital signature algorithm used when signing an AuthnRequest message.
        :param pulumi.Input[str] request_signature_scope: Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        :param pulumi.Input[str] response_signature_algorithm: The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        :param pulumi.Input[str] response_signature_scope: Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        :param pulumi.Input[str] sso_binding: The method of making an SSO request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        :param pulumi.Input[str] sso_destination: URI reference indicating the address to which the AuthnRequest message is sent.
        :param pulumi.Input[str] sso_url: URL of binding-specific endpoint to send an AuthnRequest message to IdP.
        :param pulumi.Input[str] status: Status of the IdP.
        :param pulumi.Input[str] subject_filter: Optional regular expression pattern used to filter untrusted IdP usernames.
        :param pulumi.Input[list] subject_formats: The name formate. By default `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"`.
        :param pulumi.Input[str] subject_match_attribute: Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] subject_match_type: Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] suspended_action: Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        :param pulumi.Input[str] username_template: Okta EL Expression to generate or transform a unique username for the IdP user.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/idp_saml.html.markdown.
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

            __props__['account_link_action'] = account_link_action
            __props__['account_link_group_includes'] = account_link_group_includes
            if acs_binding is None:
                raise TypeError("Missing required property 'acs_binding'")
            __props__['acs_binding'] = acs_binding
            __props__['acs_type'] = acs_type
            __props__['deprovisioned_action'] = deprovisioned_action
            __props__['groups_action'] = groups_action
            __props__['groups_assignments'] = groups_assignments
            __props__['groups_attribute'] = groups_attribute
            __props__['groups_filters'] = groups_filters
            if issuer is None:
                raise TypeError("Missing required property 'issuer'")
            __props__['issuer'] = issuer
            __props__['issuer_mode'] = issuer_mode
            if kid is None:
                raise TypeError("Missing required property 'kid'")
            __props__['kid'] = kid
            __props__['name'] = name
            __props__['name_format'] = name_format
            __props__['profile_master'] = profile_master
            __props__['provisioning_action'] = provisioning_action
            __props__['request_signature_algorithm'] = request_signature_algorithm
            __props__['request_signature_scope'] = request_signature_scope
            __props__['response_signature_algorithm'] = response_signature_algorithm
            __props__['response_signature_scope'] = response_signature_scope
            __props__['sso_binding'] = sso_binding
            __props__['sso_destination'] = sso_destination
            if sso_url is None:
                raise TypeError("Missing required property 'sso_url'")
            __props__['sso_url'] = sso_url
            __props__['status'] = status
            __props__['subject_filter'] = subject_filter
            __props__['subject_formats'] = subject_formats
            __props__['subject_match_attribute'] = subject_match_attribute
            __props__['subject_match_type'] = subject_match_type
            __props__['suspended_action'] = suspended_action
            __props__['username_template'] = username_template
            __props__['audience'] = None
            __props__['type'] = None
        super(Saml, __self__).__init__(
            'okta:idp/saml:Saml',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_link_action=None, account_link_group_includes=None, acs_binding=None, acs_type=None, audience=None, deprovisioned_action=None, groups_action=None, groups_assignments=None, groups_attribute=None, groups_filters=None, issuer=None, issuer_mode=None, kid=None, name=None, name_format=None, profile_master=None, provisioning_action=None, request_signature_algorithm=None, request_signature_scope=None, response_signature_algorithm=None, response_signature_scope=None, sso_binding=None, sso_destination=None, sso_url=None, status=None, subject_filter=None, subject_formats=None, subject_match_attribute=None, subject_match_type=None, suspended_action=None, type=None, username_template=None):
        """
        Get an existing Saml resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_link_action: Specifies the account linking action for an IdP user.
        :param pulumi.Input[list] account_link_group_includes: Group memberships to determine link candidates.
        :param pulumi.Input[str] acs_binding: The method of making an ACS request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        :param pulumi.Input[str] acs_type: The type of ACS. It can be `"INSTANCE"` or `"ORG"`.
        :param pulumi.Input[str] audience: The audience restriction for the IdP.
        :param pulumi.Input[str] deprovisioned_action: Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        :param pulumi.Input[str] groups_action: Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        :param pulumi.Input[list] groups_assignments: List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        :param pulumi.Input[str] groups_attribute: IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        :param pulumi.Input[list] groups_filters: Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        :param pulumi.Input[str] issuer: URI that identifies the issuer.
        :param pulumi.Input[str] issuer_mode: Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        :param pulumi.Input[str] kid: The ID of the signing key.
        :param pulumi.Input[str] name: The Application's display name.
        :param pulumi.Input[str] name_format: The name identifier format to use. By default `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"`.
        :param pulumi.Input[bool] profile_master: Determines if the IdP should act as a source of truth for user profile attributes.
        :param pulumi.Input[str] provisioning_action: Provisioning action for an IdP user during authentication.
        :param pulumi.Input[str] request_signature_algorithm: The XML digital signature algorithm used when signing an AuthnRequest message.
        :param pulumi.Input[str] request_signature_scope: Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        :param pulumi.Input[str] response_signature_algorithm: The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        :param pulumi.Input[str] response_signature_scope: Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        :param pulumi.Input[str] sso_binding: The method of making an SSO request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        :param pulumi.Input[str] sso_destination: URI reference indicating the address to which the AuthnRequest message is sent.
        :param pulumi.Input[str] sso_url: URL of binding-specific endpoint to send an AuthnRequest message to IdP.
        :param pulumi.Input[str] status: Status of the IdP.
        :param pulumi.Input[str] subject_filter: Optional regular expression pattern used to filter untrusted IdP usernames.
        :param pulumi.Input[list] subject_formats: The name formate. By default `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"`.
        :param pulumi.Input[str] subject_match_attribute: Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] subject_match_type: Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] suspended_action: Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        :param pulumi.Input[str] type: Type of the IdP.
        :param pulumi.Input[str] username_template: Okta EL Expression to generate or transform a unique username for the IdP user.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/idp_saml.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["account_link_action"] = account_link_action
        __props__["account_link_group_includes"] = account_link_group_includes
        __props__["acs_binding"] = acs_binding
        __props__["acs_type"] = acs_type
        __props__["audience"] = audience
        __props__["deprovisioned_action"] = deprovisioned_action
        __props__["groups_action"] = groups_action
        __props__["groups_assignments"] = groups_assignments
        __props__["groups_attribute"] = groups_attribute
        __props__["groups_filters"] = groups_filters
        __props__["issuer"] = issuer
        __props__["issuer_mode"] = issuer_mode
        __props__["kid"] = kid
        __props__["name"] = name
        __props__["name_format"] = name_format
        __props__["profile_master"] = profile_master
        __props__["provisioning_action"] = provisioning_action
        __props__["request_signature_algorithm"] = request_signature_algorithm
        __props__["request_signature_scope"] = request_signature_scope
        __props__["response_signature_algorithm"] = response_signature_algorithm
        __props__["response_signature_scope"] = response_signature_scope
        __props__["sso_binding"] = sso_binding
        __props__["sso_destination"] = sso_destination
        __props__["sso_url"] = sso_url
        __props__["status"] = status
        __props__["subject_filter"] = subject_filter
        __props__["subject_formats"] = subject_formats
        __props__["subject_match_attribute"] = subject_match_attribute
        __props__["subject_match_type"] = subject_match_type
        __props__["suspended_action"] = suspended_action
        __props__["type"] = type
        __props__["username_template"] = username_template
        return Saml(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

