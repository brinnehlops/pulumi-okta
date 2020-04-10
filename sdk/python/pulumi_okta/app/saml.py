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
    accessibility_error_redirect_url: pulumi.Output[str]
    """
    Custom error page URL.
    """
    accessibility_login_redirect_url: pulumi.Output[str]
    """
    Custom login page URL.
    """
    accessibility_self_service: pulumi.Output[bool]
    """
    Enable self service.
    """
    app_settings_json: pulumi.Output[str]
    """
    Application settings in JSON format.
    """
    assertion_signed: pulumi.Output[bool]
    """
    Determines whether the SAML assertion is digitally signed.
    """
    attribute_statements: pulumi.Output[list]
    """
    List of SAML Attribute statements.

      * `filterType` (`str`) - Type of group attribute filter.
      * `filterValue` (`str`) - Filter value to use.
      * `name` (`str`) - The name of the attribute statement.
      * `namespace` (`str`) - The attribute namespace. It can be set to `"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"`, `"urn:oasis:names:tc:SAML:2.0:attrname-format:uri"`, or `"urn:oasis:names:tc:SAML:2.0:attrname-format:basic"`.
      * `type` (`str`) - The type of attribute statement value. Can be `"EXPRESSION"` or `"GROUP"`.
      * `values` (`list`) - Array of values to use.
    """
    audience: pulumi.Output[str]
    """
    Audience restriction.
    """
    authn_context_class_ref: pulumi.Output[str]
    """
    Identifies the SAML authentication context class for the assertion’s authentication statement.
    """
    auto_submit_toolbar: pulumi.Output[bool]
    """
    Display auto submit toolbar.
    """
    certificate: pulumi.Output[str]
    """
    The raw signing certificate.
    """
    default_relay_state: pulumi.Output[str]
    """
    Identifies a specific application resource in an IDP initiated SSO scenario.
    """
    destination: pulumi.Output[str]
    """
    Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
    """
    digest_algorithm: pulumi.Output[str]
    """
    Determines the digest algorithm used to digitally sign the SAML assertion and response.
    """
    entity_key: pulumi.Output[str]
    """
    Entity ID, the ID portion of the `entity_url`.
    """
    entity_url: pulumi.Output[str]
    """
    Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8.
    """
    features: pulumi.Output[list]
    """
    features enabled.
    """
    groups: pulumi.Output[list]
    """
    Groups associated with the application
    """
    hide_ios: pulumi.Output[bool]
    """
    Do not display application icon on mobile app.
    """
    hide_web: pulumi.Output[bool]
    """
    Do not display application icon to users
    """
    honor_force_authn: pulumi.Output[bool]
    """
    Prompt user to re-authenticate if SP asks for it.
    """
    http_post_binding: pulumi.Output[str]
    """
    `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` location from the SAML metadata.
    """
    http_redirect_binding: pulumi.Output[str]
    """
    `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` location from the SAML metadata.
    """
    idp_issuer: pulumi.Output[str]
    """
    SAML issuer ID.
    """
    key_id: pulumi.Output[str]
    """
    Certificate key ID.
    """
    key_name: pulumi.Output[str]
    """
    Certificate name. This modulates the rotation of keys. New name == new key.
    """
    key_years_valid: pulumi.Output[float]
    """
    Number of years the certificate is valid.
    """
    label: pulumi.Output[str]
    """
    label of application.
    """
    metadata: pulumi.Output[str]
    """
    The raw SAML metadata in XML.
    """
    name: pulumi.Output[str]
    """
    The name of the attribute statement.
    """
    preconfigured_app: pulumi.Output[str]
    """
    name of application from the Okta Integration Network, if not included a custom app will be created.
    """
    recipient: pulumi.Output[str]
    """
    The location where the app may present the SAML assertion.
    """
    request_compressed: pulumi.Output[bool]
    """
    Denotes whether the request is compressed or not.
    """
    response_signed: pulumi.Output[bool]
    """
    Determines whether the SAML auth response message is digitally signed.
    """
    sign_on_mode: pulumi.Output[str]
    """
    Sign on mode of application.
    """
    signature_algorithm: pulumi.Output[str]
    """
    Signature algorithm used ot digitally sign the assertion and response.
    """
    sp_issuer: pulumi.Output[str]
    """
    SAML service provider issuer.
    """
    sso_url: pulumi.Output[str]
    """
    Single Sign on Url.
    """
    status: pulumi.Output[str]
    """
    status of application.
    """
    subject_name_id_format: pulumi.Output[str]
    """
    Identifies the SAML processing rules.
    """
    subject_name_id_template: pulumi.Output[str]
    """
    Template for app user's username when a user is assigned to the app.
    """
    user_name_template: pulumi.Output[str]
    """
    Username template.
    """
    user_name_template_suffix: pulumi.Output[str]
    """
    Username template suffix.
    """
    user_name_template_type: pulumi.Output[str]
    """
    Username template type.
    """
    users: pulumi.Output[list]
    """
    Users associated with the application

      * `id` (`str`) - id of application.
      * `password` (`str`)
      * `scope` (`str`)
      * `username` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, accessibility_error_redirect_url=None, accessibility_login_redirect_url=None, accessibility_self_service=None, app_settings_json=None, assertion_signed=None, attribute_statements=None, audience=None, authn_context_class_ref=None, auto_submit_toolbar=None, default_relay_state=None, destination=None, digest_algorithm=None, features=None, groups=None, hide_ios=None, hide_web=None, honor_force_authn=None, idp_issuer=None, key_name=None, key_years_valid=None, label=None, preconfigured_app=None, recipient=None, request_compressed=None, response_signed=None, signature_algorithm=None, sp_issuer=None, sso_url=None, status=None, subject_name_id_format=None, subject_name_id_template=None, user_name_template=None, user_name_template_suffix=None, user_name_template_type=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an SAML Application.

        This resource allows you to create and configure an SAML Application.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessibility_error_redirect_url: Custom error page URL.
        :param pulumi.Input[str] accessibility_login_redirect_url: Custom login page URL.
        :param pulumi.Input[bool] accessibility_self_service: Enable self service.
        :param pulumi.Input[str] app_settings_json: Application settings in JSON format.
        :param pulumi.Input[bool] assertion_signed: Determines whether the SAML assertion is digitally signed.
        :param pulumi.Input[list] attribute_statements: List of SAML Attribute statements.
        :param pulumi.Input[str] audience: Audience restriction.
        :param pulumi.Input[str] authn_context_class_ref: Identifies the SAML authentication context class for the assertion’s authentication statement.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar.
        :param pulumi.Input[str] default_relay_state: Identifies a specific application resource in an IDP initiated SSO scenario.
        :param pulumi.Input[str] destination: Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
        :param pulumi.Input[str] digest_algorithm: Determines the digest algorithm used to digitally sign the SAML assertion and response.
        :param pulumi.Input[list] features: features enabled.
        :param pulumi.Input[list] groups: Groups associated with the application
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app.
        :param pulumi.Input[bool] hide_web: Do not display application icon to users
        :param pulumi.Input[bool] honor_force_authn: Prompt user to re-authenticate if SP asks for it.
        :param pulumi.Input[str] idp_issuer: SAML issuer ID.
        :param pulumi.Input[str] key_name: Certificate name. This modulates the rotation of keys. New name == new key.
        :param pulumi.Input[float] key_years_valid: Number of years the certificate is valid.
        :param pulumi.Input[str] label: label of application.
        :param pulumi.Input[str] preconfigured_app: name of application from the Okta Integration Network, if not included a custom app will be created.
        :param pulumi.Input[str] recipient: The location where the app may present the SAML assertion.
        :param pulumi.Input[bool] request_compressed: Denotes whether the request is compressed or not.
        :param pulumi.Input[bool] response_signed: Determines whether the SAML auth response message is digitally signed.
        :param pulumi.Input[str] signature_algorithm: Signature algorithm used ot digitally sign the assertion and response.
        :param pulumi.Input[str] sp_issuer: SAML service provider issuer.
        :param pulumi.Input[str] sso_url: Single Sign on Url.
        :param pulumi.Input[str] status: status of application.
        :param pulumi.Input[str] subject_name_id_format: Identifies the SAML processing rules.
        :param pulumi.Input[str] subject_name_id_template: Template for app user's username when a user is assigned to the app.
        :param pulumi.Input[str] user_name_template: Username template.
        :param pulumi.Input[str] user_name_template_suffix: Username template suffix.
        :param pulumi.Input[str] user_name_template_type: Username template type.
        :param pulumi.Input[list] users: Users associated with the application

        The **attribute_statements** object supports the following:

          * `filterType` (`pulumi.Input[str]`) - Type of group attribute filter.
          * `filterValue` (`pulumi.Input[str]`) - Filter value to use.
          * `name` (`pulumi.Input[str]`) - The name of the attribute statement.
          * `namespace` (`pulumi.Input[str]`) - The attribute namespace. It can be set to `"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"`, `"urn:oasis:names:tc:SAML:2.0:attrname-format:uri"`, or `"urn:oasis:names:tc:SAML:2.0:attrname-format:basic"`.
          * `type` (`pulumi.Input[str]`) - The type of attribute statement value. Can be `"EXPRESSION"` or `"GROUP"`.
          * `values` (`pulumi.Input[list]`) - Array of values to use.

        The **users** object supports the following:

          * `id` (`pulumi.Input[str]`) - id of application.
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)
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

            __props__['accessibility_error_redirect_url'] = accessibility_error_redirect_url
            __props__['accessibility_login_redirect_url'] = accessibility_login_redirect_url
            __props__['accessibility_self_service'] = accessibility_self_service
            __props__['app_settings_json'] = app_settings_json
            __props__['assertion_signed'] = assertion_signed
            __props__['attribute_statements'] = attribute_statements
            __props__['audience'] = audience
            __props__['authn_context_class_ref'] = authn_context_class_ref
            __props__['auto_submit_toolbar'] = auto_submit_toolbar
            __props__['default_relay_state'] = default_relay_state
            __props__['destination'] = destination
            __props__['digest_algorithm'] = digest_algorithm
            __props__['features'] = features
            __props__['groups'] = groups
            __props__['hide_ios'] = hide_ios
            __props__['hide_web'] = hide_web
            __props__['honor_force_authn'] = honor_force_authn
            __props__['idp_issuer'] = idp_issuer
            __props__['key_name'] = key_name
            __props__['key_years_valid'] = key_years_valid
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            __props__['preconfigured_app'] = preconfigured_app
            __props__['recipient'] = recipient
            __props__['request_compressed'] = request_compressed
            __props__['response_signed'] = response_signed
            __props__['signature_algorithm'] = signature_algorithm
            __props__['sp_issuer'] = sp_issuer
            __props__['sso_url'] = sso_url
            __props__['status'] = status
            __props__['subject_name_id_format'] = subject_name_id_format
            __props__['subject_name_id_template'] = subject_name_id_template
            __props__['user_name_template'] = user_name_template
            __props__['user_name_template_suffix'] = user_name_template_suffix
            __props__['user_name_template_type'] = user_name_template_type
            __props__['users'] = users
            __props__['certificate'] = None
            __props__['entity_key'] = None
            __props__['entity_url'] = None
            __props__['http_post_binding'] = None
            __props__['http_redirect_binding'] = None
            __props__['key_id'] = None
            __props__['metadata'] = None
            __props__['name'] = None
            __props__['sign_on_mode'] = None
        super(Saml, __self__).__init__(
            'okta:app/saml:Saml',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessibility_error_redirect_url=None, accessibility_login_redirect_url=None, accessibility_self_service=None, app_settings_json=None, assertion_signed=None, attribute_statements=None, audience=None, authn_context_class_ref=None, auto_submit_toolbar=None, certificate=None, default_relay_state=None, destination=None, digest_algorithm=None, entity_key=None, entity_url=None, features=None, groups=None, hide_ios=None, hide_web=None, honor_force_authn=None, http_post_binding=None, http_redirect_binding=None, idp_issuer=None, key_id=None, key_name=None, key_years_valid=None, label=None, metadata=None, name=None, preconfigured_app=None, recipient=None, request_compressed=None, response_signed=None, sign_on_mode=None, signature_algorithm=None, sp_issuer=None, sso_url=None, status=None, subject_name_id_format=None, subject_name_id_template=None, user_name_template=None, user_name_template_suffix=None, user_name_template_type=None, users=None):
        """
        Get an existing Saml resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessibility_error_redirect_url: Custom error page URL.
        :param pulumi.Input[str] accessibility_login_redirect_url: Custom login page URL.
        :param pulumi.Input[bool] accessibility_self_service: Enable self service.
        :param pulumi.Input[str] app_settings_json: Application settings in JSON format.
        :param pulumi.Input[bool] assertion_signed: Determines whether the SAML assertion is digitally signed.
        :param pulumi.Input[list] attribute_statements: List of SAML Attribute statements.
        :param pulumi.Input[str] audience: Audience restriction.
        :param pulumi.Input[str] authn_context_class_ref: Identifies the SAML authentication context class for the assertion’s authentication statement.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar.
        :param pulumi.Input[str] certificate: The raw signing certificate.
        :param pulumi.Input[str] default_relay_state: Identifies a specific application resource in an IDP initiated SSO scenario.
        :param pulumi.Input[str] destination: Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
        :param pulumi.Input[str] digest_algorithm: Determines the digest algorithm used to digitally sign the SAML assertion and response.
        :param pulumi.Input[str] entity_key: Entity ID, the ID portion of the `entity_url`.
        :param pulumi.Input[str] entity_url: Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8.
        :param pulumi.Input[list] features: features enabled.
        :param pulumi.Input[list] groups: Groups associated with the application
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app.
        :param pulumi.Input[bool] hide_web: Do not display application icon to users
        :param pulumi.Input[bool] honor_force_authn: Prompt user to re-authenticate if SP asks for it.
        :param pulumi.Input[str] http_post_binding: `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` location from the SAML metadata.
        :param pulumi.Input[str] http_redirect_binding: `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` location from the SAML metadata.
        :param pulumi.Input[str] idp_issuer: SAML issuer ID.
        :param pulumi.Input[str] key_id: Certificate key ID.
        :param pulumi.Input[str] key_name: Certificate name. This modulates the rotation of keys. New name == new key.
        :param pulumi.Input[float] key_years_valid: Number of years the certificate is valid.
        :param pulumi.Input[str] label: label of application.
        :param pulumi.Input[str] metadata: The raw SAML metadata in XML.
        :param pulumi.Input[str] name: The name of the attribute statement.
        :param pulumi.Input[str] preconfigured_app: name of application from the Okta Integration Network, if not included a custom app will be created.
        :param pulumi.Input[str] recipient: The location where the app may present the SAML assertion.
        :param pulumi.Input[bool] request_compressed: Denotes whether the request is compressed or not.
        :param pulumi.Input[bool] response_signed: Determines whether the SAML auth response message is digitally signed.
        :param pulumi.Input[str] sign_on_mode: Sign on mode of application.
        :param pulumi.Input[str] signature_algorithm: Signature algorithm used ot digitally sign the assertion and response.
        :param pulumi.Input[str] sp_issuer: SAML service provider issuer.
        :param pulumi.Input[str] sso_url: Single Sign on Url.
        :param pulumi.Input[str] status: status of application.
        :param pulumi.Input[str] subject_name_id_format: Identifies the SAML processing rules.
        :param pulumi.Input[str] subject_name_id_template: Template for app user's username when a user is assigned to the app.
        :param pulumi.Input[str] user_name_template: Username template.
        :param pulumi.Input[str] user_name_template_suffix: Username template suffix.
        :param pulumi.Input[str] user_name_template_type: Username template type.
        :param pulumi.Input[list] users: Users associated with the application

        The **attribute_statements** object supports the following:

          * `filterType` (`pulumi.Input[str]`) - Type of group attribute filter.
          * `filterValue` (`pulumi.Input[str]`) - Filter value to use.
          * `name` (`pulumi.Input[str]`) - The name of the attribute statement.
          * `namespace` (`pulumi.Input[str]`) - The attribute namespace. It can be set to `"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"`, `"urn:oasis:names:tc:SAML:2.0:attrname-format:uri"`, or `"urn:oasis:names:tc:SAML:2.0:attrname-format:basic"`.
          * `type` (`pulumi.Input[str]`) - The type of attribute statement value. Can be `"EXPRESSION"` or `"GROUP"`.
          * `values` (`pulumi.Input[list]`) - Array of values to use.

        The **users** object supports the following:

          * `id` (`pulumi.Input[str]`) - id of application.
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessibility_error_redirect_url"] = accessibility_error_redirect_url
        __props__["accessibility_login_redirect_url"] = accessibility_login_redirect_url
        __props__["accessibility_self_service"] = accessibility_self_service
        __props__["app_settings_json"] = app_settings_json
        __props__["assertion_signed"] = assertion_signed
        __props__["attribute_statements"] = attribute_statements
        __props__["audience"] = audience
        __props__["authn_context_class_ref"] = authn_context_class_ref
        __props__["auto_submit_toolbar"] = auto_submit_toolbar
        __props__["certificate"] = certificate
        __props__["default_relay_state"] = default_relay_state
        __props__["destination"] = destination
        __props__["digest_algorithm"] = digest_algorithm
        __props__["entity_key"] = entity_key
        __props__["entity_url"] = entity_url
        __props__["features"] = features
        __props__["groups"] = groups
        __props__["hide_ios"] = hide_ios
        __props__["hide_web"] = hide_web
        __props__["honor_force_authn"] = honor_force_authn
        __props__["http_post_binding"] = http_post_binding
        __props__["http_redirect_binding"] = http_redirect_binding
        __props__["idp_issuer"] = idp_issuer
        __props__["key_id"] = key_id
        __props__["key_name"] = key_name
        __props__["key_years_valid"] = key_years_valid
        __props__["label"] = label
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["preconfigured_app"] = preconfigured_app
        __props__["recipient"] = recipient
        __props__["request_compressed"] = request_compressed
        __props__["response_signed"] = response_signed
        __props__["sign_on_mode"] = sign_on_mode
        __props__["signature_algorithm"] = signature_algorithm
        __props__["sp_issuer"] = sp_issuer
        __props__["sso_url"] = sso_url
        __props__["status"] = status
        __props__["subject_name_id_format"] = subject_name_id_format
        __props__["subject_name_id_template"] = subject_name_id_template
        __props__["user_name_template"] = user_name_template
        __props__["user_name_template_suffix"] = user_name_template_suffix
        __props__["user_name_template_type"] = user_name_template_type
        __props__["users"] = users
        return Saml(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

