# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Social']


class Social(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_link_action: Optional[pulumi.Input[str]] = None,
                 account_link_group_includes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 deprovisioned_action: Optional[pulumi.Input[str]] = None,
                 groups_action: Optional[pulumi.Input[str]] = None,
                 groups_assignments: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 groups_attribute: Optional[pulumi.Input[str]] = None,
                 groups_filters: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 issuer_mode: Optional[pulumi.Input[str]] = None,
                 match_attribute: Optional[pulumi.Input[str]] = None,
                 match_type: Optional[pulumi.Input[str]] = None,
                 max_clock_skew: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile_master: Optional[pulumi.Input[bool]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 provisioning_action: Optional[pulumi.Input[str]] = None,
                 request_signature_algorithm: Optional[pulumi.Input[str]] = None,
                 request_signature_scope: Optional[pulumi.Input[str]] = None,
                 response_signature_algorithm: Optional[pulumi.Input[str]] = None,
                 response_signature_scope: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 subject_match_attribute: Optional[pulumi.Input[str]] = None,
                 subject_match_type: Optional[pulumi.Input[str]] = None,
                 suspended_action: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username_template: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Social Identity Provider.

        This resource allows you to create and configure an Social Identity Provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.idp.Social("example",
            client_id="abcd123",
            client_secret="abcd123",
            match_attribute="customfieldId",
            match_type="CUSTOM_ATTRIBUTE",
            protocol_type="OAUTH2",
            scopes=[
                "public_profile",
                "email",
            ],
            type="FACEBOOK",
            username_template="idpuser.email")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_link_action: Specifies the account linking action for an IdP user.
        :param pulumi.Input[List[pulumi.Input[str]]] account_link_group_includes: Group memberships to determine link candidates.
        :param pulumi.Input[str] client_id: Unique identifier issued by AS for the Okta IdP instance.
        :param pulumi.Input[str] client_secret: Client secret issued by AS for the Okta IdP instance.
        :param pulumi.Input[str] deprovisioned_action: Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        :param pulumi.Input[str] groups_action: Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        :param pulumi.Input[List[pulumi.Input[str]]] groups_assignments: List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        :param pulumi.Input[str] groups_attribute: IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        :param pulumi.Input[List[pulumi.Input[str]]] groups_filters: Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        :param pulumi.Input[str] issuer_mode: Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        :param pulumi.Input[float] max_clock_skew: Maximum allowable clock-skew when processing messages from the IdP.
        :param pulumi.Input[str] name: The Application's display name.
        :param pulumi.Input[bool] profile_master: Determines if the IdP should act as a source of truth for user profile attributes.
        :param pulumi.Input[str] protocol_type: The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
        :param pulumi.Input[str] provisioning_action: Provisioning action for an IdP user during authentication.
        :param pulumi.Input[str] request_signature_algorithm: The XML digital signature algorithm used when signing an AuthnRequest message.
        :param pulumi.Input[str] request_signature_scope: Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        :param pulumi.Input[str] response_signature_algorithm: The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        :param pulumi.Input[str] response_signature_scope: Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        :param pulumi.Input[List[pulumi.Input[str]]] scopes: The scopes of the IdP.
        :param pulumi.Input[str] status: Status of the IdP.
        :param pulumi.Input[str] subject_match_attribute: Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] subject_match_type: Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] suspended_action: Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        :param pulumi.Input[str] type: The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
        :param pulumi.Input[str] username_template: Okta EL Expression to generate or transform a unique username for the IdP user.
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
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['deprovisioned_action'] = deprovisioned_action
            __props__['groups_action'] = groups_action
            __props__['groups_assignments'] = groups_assignments
            __props__['groups_attribute'] = groups_attribute
            __props__['groups_filters'] = groups_filters
            __props__['issuer_mode'] = issuer_mode
            if match_attribute is not None:
                warnings.warn("This property was incorrectly added to this resource, you should use \"subject_match_attribute\"", DeprecationWarning)
                pulumi.log.warn("match_attribute is deprecated: This property was incorrectly added to this resource, you should use \"subject_match_attribute\"")
            __props__['match_attribute'] = match_attribute
            if match_type is not None:
                warnings.warn("This property was incorrectly added to this resource, you should use \"subject_match_type\"", DeprecationWarning)
                pulumi.log.warn("match_type is deprecated: This property was incorrectly added to this resource, you should use \"subject_match_type\"")
            __props__['match_type'] = match_type
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
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['username_template'] = username_template
            __props__['authorization_binding'] = None
            __props__['authorization_url'] = None
            __props__['token_binding'] = None
            __props__['token_url'] = None
        super(Social, __self__).__init__(
            'okta:idp/social:Social',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_link_action: Optional[pulumi.Input[str]] = None,
            account_link_group_includes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            authorization_binding: Optional[pulumi.Input[str]] = None,
            authorization_url: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            deprovisioned_action: Optional[pulumi.Input[str]] = None,
            groups_action: Optional[pulumi.Input[str]] = None,
            groups_assignments: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            groups_attribute: Optional[pulumi.Input[str]] = None,
            groups_filters: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            issuer_mode: Optional[pulumi.Input[str]] = None,
            match_attribute: Optional[pulumi.Input[str]] = None,
            match_type: Optional[pulumi.Input[str]] = None,
            max_clock_skew: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            profile_master: Optional[pulumi.Input[bool]] = None,
            protocol_type: Optional[pulumi.Input[str]] = None,
            provisioning_action: Optional[pulumi.Input[str]] = None,
            request_signature_algorithm: Optional[pulumi.Input[str]] = None,
            request_signature_scope: Optional[pulumi.Input[str]] = None,
            response_signature_algorithm: Optional[pulumi.Input[str]] = None,
            response_signature_scope: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subject_match_attribute: Optional[pulumi.Input[str]] = None,
            subject_match_type: Optional[pulumi.Input[str]] = None,
            suspended_action: Optional[pulumi.Input[str]] = None,
            token_binding: Optional[pulumi.Input[str]] = None,
            token_url: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            username_template: Optional[pulumi.Input[str]] = None) -> 'Social':
        """
        Get an existing Social resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_link_action: Specifies the account linking action for an IdP user.
        :param pulumi.Input[List[pulumi.Input[str]]] account_link_group_includes: Group memberships to determine link candidates.
        :param pulumi.Input[str] authorization_binding: The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        :param pulumi.Input[str] authorization_url: IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
        :param pulumi.Input[str] client_id: Unique identifier issued by AS for the Okta IdP instance.
        :param pulumi.Input[str] client_secret: Client secret issued by AS for the Okta IdP instance.
        :param pulumi.Input[str] deprovisioned_action: Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        :param pulumi.Input[str] groups_action: Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        :param pulumi.Input[List[pulumi.Input[str]]] groups_assignments: List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        :param pulumi.Input[str] groups_attribute: IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        :param pulumi.Input[List[pulumi.Input[str]]] groups_filters: Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        :param pulumi.Input[str] issuer_mode: Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        :param pulumi.Input[float] max_clock_skew: Maximum allowable clock-skew when processing messages from the IdP.
        :param pulumi.Input[str] name: The Application's display name.
        :param pulumi.Input[bool] profile_master: Determines if the IdP should act as a source of truth for user profile attributes.
        :param pulumi.Input[str] protocol_type: The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
        :param pulumi.Input[str] provisioning_action: Provisioning action for an IdP user during authentication.
        :param pulumi.Input[str] request_signature_algorithm: The XML digital signature algorithm used when signing an AuthnRequest message.
        :param pulumi.Input[str] request_signature_scope: Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        :param pulumi.Input[str] response_signature_algorithm: The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        :param pulumi.Input[str] response_signature_scope: Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        :param pulumi.Input[List[pulumi.Input[str]]] scopes: The scopes of the IdP.
        :param pulumi.Input[str] status: Status of the IdP.
        :param pulumi.Input[str] subject_match_attribute: Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] subject_match_type: Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        :param pulumi.Input[str] suspended_action: Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        :param pulumi.Input[str] token_binding: The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        :param pulumi.Input[str] token_url: IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
        :param pulumi.Input[str] type: The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
        :param pulumi.Input[str] username_template: Okta EL Expression to generate or transform a unique username for the IdP user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_link_action"] = account_link_action
        __props__["account_link_group_includes"] = account_link_group_includes
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
        __props__["match_attribute"] = match_attribute
        __props__["match_type"] = match_type
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
        __props__["username_template"] = username_template
        return Social(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountLinkAction")
    def account_link_action(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the account linking action for an IdP user.
        """
        return pulumi.get(self, "account_link_action")

    @property
    @pulumi.getter(name="accountLinkGroupIncludes")
    def account_link_group_includes(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Group memberships to determine link candidates.
        """
        return pulumi.get(self, "account_link_group_includes")

    @property
    @pulumi.getter(name="authorizationBinding")
    def authorization_binding(self) -> pulumi.Output[str]:
        """
        The method of making an authorization request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        """
        return pulumi.get(self, "authorization_binding")

    @property
    @pulumi.getter(name="authorizationUrl")
    def authorization_url(self) -> pulumi.Output[str]:
        """
        IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
        """
        return pulumi.get(self, "authorization_url")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        Unique identifier issued by AS for the Okta IdP instance.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        Client secret issued by AS for the Okta IdP instance.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="deprovisionedAction")
    def deprovisioned_action(self) -> pulumi.Output[Optional[str]]:
        """
        Action for a previously deprovisioned IdP user during authentication. Can be `"NONE"` or `"REACTIVATE"`.
        """
        return pulumi.get(self, "deprovisioned_action")

    @property
    @pulumi.getter(name="groupsAction")
    def groups_action(self) -> pulumi.Output[Optional[str]]:
        """
        Provisioning action for IdP user's group memberships. It can be `"NONE"`, `"SYNC"`, `"APPEND"`, or `"ASSIGN"`.
        """
        return pulumi.get(self, "groups_action")

    @property
    @pulumi.getter(name="groupsAssignments")
    def groups_assignments(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of Okta Group IDs to add an IdP user as a member with the `"ASSIGN"` `groups_action`.
        """
        return pulumi.get(self, "groups_assignments")

    @property
    @pulumi.getter(name="groupsAttribute")
    def groups_attribute(self) -> pulumi.Output[Optional[str]]:
        """
        IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
        """
        return pulumi.get(self, "groups_attribute")

    @property
    @pulumi.getter(name="groupsFilters")
    def groups_filters(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Whitelist of Okta Group identifiers that are allowed for the `"APPEND"` or `"SYNC"` `groups_action`.
        """
        return pulumi.get(self, "groups_filters")

    @property
    @pulumi.getter(name="issuerMode")
    def issuer_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be `"ORG_URL"` or `"CUSTOM_URL"`.
        """
        return pulumi.get(self, "issuer_mode")

    @property
    @pulumi.getter(name="matchAttribute")
    def match_attribute(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "match_attribute")

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "match_type")

    @property
    @pulumi.getter(name="maxClockSkew")
    def max_clock_skew(self) -> pulumi.Output[Optional[float]]:
        """
        Maximum allowable clock-skew when processing messages from the IdP.
        """
        return pulumi.get(self, "max_clock_skew")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Application's display name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileMaster")
    def profile_master(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the IdP should act as a source of truth for user profile attributes.
        """
        return pulumi.get(self, "profile_master")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of protocol to use. It can be `"OIDC"` or `"OAUTH2"`.
        """
        return pulumi.get(self, "protocol_type")

    @property
    @pulumi.getter(name="provisioningAction")
    def provisioning_action(self) -> pulumi.Output[Optional[str]]:
        """
        Provisioning action for an IdP user during authentication.
        """
        return pulumi.get(self, "provisioning_action")

    @property
    @pulumi.getter(name="requestSignatureAlgorithm")
    def request_signature_algorithm(self) -> pulumi.Output[Optional[str]]:
        """
        The XML digital signature algorithm used when signing an AuthnRequest message.
        """
        return pulumi.get(self, "request_signature_algorithm")

    @property
    @pulumi.getter(name="requestSignatureScope")
    def request_signature_scope(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be `"REQUEST"` or `"NONE"`.
        """
        return pulumi.get(self, "request_signature_scope")

    @property
    @pulumi.getter(name="responseSignatureAlgorithm")
    def response_signature_algorithm(self) -> pulumi.Output[Optional[str]]:
        """
        The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.
        """
        return pulumi.get(self, "response_signature_algorithm")

    @property
    @pulumi.getter(name="responseSignatureScope")
    def response_signature_scope(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be `"RESPONSE"`, `"ASSERTION"`, or `"ANY"`.
        """
        return pulumi.get(self, "response_signature_scope")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[List[str]]:
        """
        The scopes of the IdP.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Status of the IdP.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subjectMatchAttribute")
    def subject_match_attribute(self) -> pulumi.Output[Optional[str]]:
        """
        Okta user profile attribute for matching transformed IdP username. Only for matchType `"CUSTOM_ATTRIBUTE"`.
        """
        return pulumi.get(self, "subject_match_attribute")

    @property
    @pulumi.getter(name="subjectMatchType")
    def subject_match_type(self) -> pulumi.Output[Optional[str]]:
        """
        Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to `"USERNAME"`. It can be set to `"USERNAME"`, `"EMAIL"`, `"USERNAME_OR_EMAIL"` or `"CUSTOM_ATTRIBUTE"`.
        """
        return pulumi.get(self, "subject_match_type")

    @property
    @pulumi.getter(name="suspendedAction")
    def suspended_action(self) -> pulumi.Output[Optional[str]]:
        """
        Action for a previously suspended IdP user during authentication. Can be set to `"NONE"` or `"UNSUSPEND"`
        """
        return pulumi.get(self, "suspended_action")

    @property
    @pulumi.getter(name="tokenBinding")
    def token_binding(self) -> pulumi.Output[str]:
        """
        The method of making a token request. It can be set to `"HTTP-POST"` or `"HTTP-REDIRECT"`.
        """
        return pulumi.get(self, "token_binding")

    @property
    @pulumi.getter(name="tokenUrl")
    def token_url(self) -> pulumi.Output[str]:
        """
        IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
        """
        return pulumi.get(self, "token_url")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of Social IdP. It can be `"FACEBOOK"`, `"LINKEDIN"`, `"MICROSOFT"`, or `"GOOGLE"`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> pulumi.Output[Optional[str]]:
        """
        Okta EL Expression to generate or transform a unique username for the IdP user.
        """
        return pulumi.get(self, "username_template")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

