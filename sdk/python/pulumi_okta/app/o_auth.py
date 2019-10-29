# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OAuth(pulumi.CustomResource):
    auto_key_rotation: pulumi.Output[bool]
    """
    Requested key rotation mode.
    """
    auto_submit_toolbar: pulumi.Output[bool]
    """
    Display auto submit toolbar.
    """
    client_basic_secret: pulumi.Output[str]
    """
    OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
    """
    client_id: pulumi.Output[str]
    """
    The client ID of the application.
    """
    client_secret: pulumi.Output[str]
    """
    The client secret of the application.
    """
    client_uri: pulumi.Output[str]
    """
    URI to a web page providing information about the client.
    """
    consent_method: pulumi.Output[str]
    """
    Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
    """
    custom_client_id: pulumi.Output[str]
    """
    This property allows you to set the application's client id.
    """
    grant_types: pulumi.Output[list]
    """
    List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
    """
    groups: pulumi.Output[list]
    """
    The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
    """
    hide_ios: pulumi.Output[bool]
    """
    Do not display application icon on mobile app.
    """
    hide_web: pulumi.Output[bool]
    """
    Do not display application icon to users.
    """
    issuer_mode: pulumi.Output[str]
    """
    Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
    """
    label: pulumi.Output[str]
    """
    The Application's display name.
    """
    login_uri: pulumi.Output[str]
    """
    URI that initiates login.
    """
    logo_uri: pulumi.Output[str]
    """
    URI that references a logo for the client.
    """
    name: pulumi.Output[str]
    """
    Name assigned to the application by Okta.
    """
    omit_secret: pulumi.Output[bool]
    """
    This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
    """
    policy_uri: pulumi.Output[str]
    """
    URI to web page providing client policy document.
    """
    post_logout_redirect_uris: pulumi.Output[list]
    """
    List of URIs for redirection after logout.
    """
    profile: pulumi.Output[str]
    """
    Custom JSON that represents an OAuth application's profile.
    """
    redirect_uris: pulumi.Output[list]
    """
    List of URIs for use in the redirect-based flow. This is required for all application types except service.
    """
    response_types: pulumi.Output[list]
    """
    List of OAuth 2.0 response type strings.
    """
    sign_on_mode: pulumi.Output[str]
    """
    Sign on mode of application.
    """
    status: pulumi.Output[str]
    """
    The status of the application, by default it is `"ACTIVE"`.
    """
    token_endpoint_auth_method: pulumi.Output[str]
    """
    Requested authentication method for the token endpoint. It can be set to `"none"`, `"client_secret_post"`, `"client_secret_basic"`, `"client_secret_jwt"`.
    """
    tos_uri: pulumi.Output[str]
    """
    URI to web page providing client tos (terms of service).
    """
    type: pulumi.Output[str]
    """
    The type of OAuth application.
    """
    users: pulumi.Output[list]
    """
    The users assigned to the application. It is recommended not to use this and instead use `app.User`.
    
      * `id` (`str`)
      * `password` (`str`)
      * `scope` (`str`)
      * `username` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, auto_key_rotation=None, auto_submit_toolbar=None, client_basic_secret=None, client_uri=None, consent_method=None, custom_client_id=None, grant_types=None, groups=None, hide_ios=None, hide_web=None, issuer_mode=None, label=None, login_uri=None, logo_uri=None, omit_secret=None, policy_uri=None, post_logout_redirect_uris=None, profile=None, redirect_uris=None, response_types=None, status=None, token_endpoint_auth_method=None, tos_uri=None, type=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an OIDC Application.
        
        This resource allows you to create and configure an OIDC Application.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_key_rotation: Requested key rotation mode.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar.
        :param pulumi.Input[str] client_basic_secret: OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
        :param pulumi.Input[str] client_uri: URI to a web page providing information about the client.
        :param pulumi.Input[str] consent_method: Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
        :param pulumi.Input[str] custom_client_id: This property allows you to set the application's client id.
        :param pulumi.Input[list] grant_types: List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
        :param pulumi.Input[list] groups: The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app.
        :param pulumi.Input[bool] hide_web: Do not display application icon to users.
        :param pulumi.Input[str] issuer_mode: Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
        :param pulumi.Input[str] label: The Application's display name.
        :param pulumi.Input[str] login_uri: URI that initiates login.
        :param pulumi.Input[str] logo_uri: URI that references a logo for the client.
        :param pulumi.Input[bool] omit_secret: This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
        :param pulumi.Input[str] policy_uri: URI to web page providing client policy document.
        :param pulumi.Input[list] post_logout_redirect_uris: List of URIs for redirection after logout.
        :param pulumi.Input[str] profile: Custom JSON that represents an OAuth application's profile.
        :param pulumi.Input[list] redirect_uris: List of URIs for use in the redirect-based flow. This is required for all application types except service.
        :param pulumi.Input[list] response_types: List of OAuth 2.0 response type strings.
        :param pulumi.Input[str] status: The status of the application, by default it is `"ACTIVE"`.
        :param pulumi.Input[str] token_endpoint_auth_method: Requested authentication method for the token endpoint. It can be set to `"none"`, `"client_secret_post"`, `"client_secret_basic"`, `"client_secret_jwt"`.
        :param pulumi.Input[str] tos_uri: URI to web page providing client tos (terms of service).
        :param pulumi.Input[str] type: The type of OAuth application.
        :param pulumi.Input[list] users: The users assigned to the application. It is recommended not to use this and instead use `app.User`.
        
        The **users** object supports the following:
        
          * `id` (`pulumi.Input[str]`)
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_oauth.html.markdown.
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

            __props__['auto_key_rotation'] = auto_key_rotation
            __props__['auto_submit_toolbar'] = auto_submit_toolbar
            __props__['client_basic_secret'] = client_basic_secret
            __props__['client_uri'] = client_uri
            __props__['consent_method'] = consent_method
            __props__['custom_client_id'] = custom_client_id
            __props__['grant_types'] = grant_types
            __props__['groups'] = groups
            __props__['hide_ios'] = hide_ios
            __props__['hide_web'] = hide_web
            __props__['issuer_mode'] = issuer_mode
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            __props__['login_uri'] = login_uri
            __props__['logo_uri'] = logo_uri
            __props__['omit_secret'] = omit_secret
            __props__['policy_uri'] = policy_uri
            __props__['post_logout_redirect_uris'] = post_logout_redirect_uris
            __props__['profile'] = profile
            __props__['redirect_uris'] = redirect_uris
            __props__['response_types'] = response_types
            __props__['status'] = status
            __props__['token_endpoint_auth_method'] = token_endpoint_auth_method
            __props__['tos_uri'] = tos_uri
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['users'] = users
            __props__['client_id'] = None
            __props__['client_secret'] = None
            __props__['name'] = None
            __props__['sign_on_mode'] = None
        super(OAuth, __self__).__init__(
            'okta:app/oAuth:OAuth',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_key_rotation=None, auto_submit_toolbar=None, client_basic_secret=None, client_id=None, client_secret=None, client_uri=None, consent_method=None, custom_client_id=None, grant_types=None, groups=None, hide_ios=None, hide_web=None, issuer_mode=None, label=None, login_uri=None, logo_uri=None, name=None, omit_secret=None, policy_uri=None, post_logout_redirect_uris=None, profile=None, redirect_uris=None, response_types=None, sign_on_mode=None, status=None, token_endpoint_auth_method=None, tos_uri=None, type=None, users=None):
        """
        Get an existing OAuth resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_key_rotation: Requested key rotation mode.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar.
        :param pulumi.Input[str] client_basic_secret: OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
        :param pulumi.Input[str] client_id: The client ID of the application.
        :param pulumi.Input[str] client_secret: The client secret of the application.
        :param pulumi.Input[str] client_uri: URI to a web page providing information about the client.
        :param pulumi.Input[str] consent_method: Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
        :param pulumi.Input[str] custom_client_id: This property allows you to set the application's client id.
        :param pulumi.Input[list] grant_types: List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
        :param pulumi.Input[list] groups: The groups assigned to the application. It is recommended not to use this and instead use `app.GroupAssignment`.
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app.
        :param pulumi.Input[bool] hide_web: Do not display application icon to users.
        :param pulumi.Input[str] issuer_mode: Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
        :param pulumi.Input[str] label: The Application's display name.
        :param pulumi.Input[str] login_uri: URI that initiates login.
        :param pulumi.Input[str] logo_uri: URI that references a logo for the client.
        :param pulumi.Input[str] name: Name assigned to the application by Okta.
        :param pulumi.Input[bool] omit_secret: This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
        :param pulumi.Input[str] policy_uri: URI to web page providing client policy document.
        :param pulumi.Input[list] post_logout_redirect_uris: List of URIs for redirection after logout.
        :param pulumi.Input[str] profile: Custom JSON that represents an OAuth application's profile.
        :param pulumi.Input[list] redirect_uris: List of URIs for use in the redirect-based flow. This is required for all application types except service.
        :param pulumi.Input[list] response_types: List of OAuth 2.0 response type strings.
        :param pulumi.Input[str] sign_on_mode: Sign on mode of application.
        :param pulumi.Input[str] status: The status of the application, by default it is `"ACTIVE"`.
        :param pulumi.Input[str] token_endpoint_auth_method: Requested authentication method for the token endpoint. It can be set to `"none"`, `"client_secret_post"`, `"client_secret_basic"`, `"client_secret_jwt"`.
        :param pulumi.Input[str] tos_uri: URI to web page providing client tos (terms of service).
        :param pulumi.Input[str] type: The type of OAuth application.
        :param pulumi.Input[list] users: The users assigned to the application. It is recommended not to use this and instead use `app.User`.
        
        The **users** object supports the following:
        
          * `id` (`pulumi.Input[str]`)
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_oauth.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["auto_key_rotation"] = auto_key_rotation
        __props__["auto_submit_toolbar"] = auto_submit_toolbar
        __props__["client_basic_secret"] = client_basic_secret
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["client_uri"] = client_uri
        __props__["consent_method"] = consent_method
        __props__["custom_client_id"] = custom_client_id
        __props__["grant_types"] = grant_types
        __props__["groups"] = groups
        __props__["hide_ios"] = hide_ios
        __props__["hide_web"] = hide_web
        __props__["issuer_mode"] = issuer_mode
        __props__["label"] = label
        __props__["login_uri"] = login_uri
        __props__["logo_uri"] = logo_uri
        __props__["name"] = name
        __props__["omit_secret"] = omit_secret
        __props__["policy_uri"] = policy_uri
        __props__["post_logout_redirect_uris"] = post_logout_redirect_uris
        __props__["profile"] = profile
        __props__["redirect_uris"] = redirect_uris
        __props__["response_types"] = response_types
        __props__["sign_on_mode"] = sign_on_mode
        __props__["status"] = status
        __props__["token_endpoint_auth_method"] = token_endpoint_auth_method
        __props__["tos_uri"] = tos_uri
        __props__["type"] = type
        __props__["users"] = users
        return OAuth(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

