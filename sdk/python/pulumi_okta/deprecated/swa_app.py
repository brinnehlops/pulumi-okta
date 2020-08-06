# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class SwaApp(pulumi.CustomResource):
    accessibility_error_redirect_url: pulumi.Output[str]
    """
    Custom error page URL
    """
    accessibility_self_service: pulumi.Output[bool]
    """
    Enable self service
    """
    auto_submit_toolbar: pulumi.Output[bool]
    """
    Display auto submit toolbar
    """
    button_field: pulumi.Output[str]
    """
    Login button field
    """
    groups: pulumi.Output[list]
    """
    Groups associated with the application
    """
    hide_ios: pulumi.Output[bool]
    """
    Do not display application icon on mobile app
    """
    hide_web: pulumi.Output[bool]
    """
    Do not display application icon to users
    """
    label: pulumi.Output[str]
    """
    Pretty name of app.
    """
    name: pulumi.Output[str]
    """
    name of app.
    """
    password_field: pulumi.Output[str]
    """
    Login password field
    """
    preconfigured_app: pulumi.Output[str]
    """
    Preconfigured app name
    """
    sign_on_mode: pulumi.Output[str]
    """
    Sign on mode of application.
    """
    status: pulumi.Output[str]
    """
    Status of application.
    """
    url: pulumi.Output[str]
    """
    Login URL
    """
    url_regex: pulumi.Output[str]
    """
    A regex that further restricts URL to the specified regex
    """
    user_name_template: pulumi.Output[str]
    """
    Username template
    """
    user_name_template_type: pulumi.Output[str]
    """
    Username template type
    """
    username_field: pulumi.Output[str]
    """
    Login username field
    """
    users: pulumi.Output[list]
    """
    Users associated with the application

      * `id` (`str`)
      * `password` (`str`)
      * `scope` (`str`)
      * `username` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, accessibility_error_redirect_url=None, accessibility_self_service=None, auto_submit_toolbar=None, button_field=None, groups=None, hide_ios=None, hide_web=None, label=None, password_field=None, preconfigured_app=None, status=None, url=None, url_regex=None, username_field=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a SwaApp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessibility_error_redirect_url: Custom error page URL
        :param pulumi.Input[bool] accessibility_self_service: Enable self service
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar
        :param pulumi.Input[str] button_field: Login button field
        :param pulumi.Input[list] groups: Groups associated with the application
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app
        :param pulumi.Input[bool] hide_web: Do not display application icon to users
        :param pulumi.Input[str] label: Pretty name of app.
        :param pulumi.Input[str] password_field: Login password field
        :param pulumi.Input[str] preconfigured_app: Preconfigured app name
        :param pulumi.Input[str] status: Status of application.
        :param pulumi.Input[str] url: Login URL
        :param pulumi.Input[str] url_regex: A regex that further restricts URL to the specified regex
        :param pulumi.Input[str] username_field: Login username field
        :param pulumi.Input[list] users: Users associated with the application

        The **users** object supports the following:

          * `id` (`pulumi.Input[str]`)
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['accessibility_error_redirect_url'] = accessibility_error_redirect_url
            __props__['accessibility_self_service'] = accessibility_self_service
            __props__['auto_submit_toolbar'] = auto_submit_toolbar
            __props__['button_field'] = button_field
            __props__['groups'] = groups
            __props__['hide_ios'] = hide_ios
            __props__['hide_web'] = hide_web
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            __props__['password_field'] = password_field
            __props__['preconfigured_app'] = preconfigured_app
            __props__['status'] = status
            __props__['url'] = url
            __props__['url_regex'] = url_regex
            __props__['username_field'] = username_field
            __props__['users'] = users
            __props__['name'] = None
            __props__['sign_on_mode'] = None
            __props__['user_name_template'] = None
            __props__['user_name_template_type'] = None
        super(SwaApp, __self__).__init__(
            'okta:deprecated/swaApp:SwaApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessibility_error_redirect_url=None, accessibility_self_service=None, auto_submit_toolbar=None, button_field=None, groups=None, hide_ios=None, hide_web=None, label=None, name=None, password_field=None, preconfigured_app=None, sign_on_mode=None, status=None, url=None, url_regex=None, user_name_template=None, user_name_template_type=None, username_field=None, users=None):
        """
        Get an existing SwaApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessibility_error_redirect_url: Custom error page URL
        :param pulumi.Input[bool] accessibility_self_service: Enable self service
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar
        :param pulumi.Input[str] button_field: Login button field
        :param pulumi.Input[list] groups: Groups associated with the application
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app
        :param pulumi.Input[bool] hide_web: Do not display application icon to users
        :param pulumi.Input[str] label: Pretty name of app.
        :param pulumi.Input[str] name: name of app.
        :param pulumi.Input[str] password_field: Login password field
        :param pulumi.Input[str] preconfigured_app: Preconfigured app name
        :param pulumi.Input[str] sign_on_mode: Sign on mode of application.
        :param pulumi.Input[str] status: Status of application.
        :param pulumi.Input[str] url: Login URL
        :param pulumi.Input[str] url_regex: A regex that further restricts URL to the specified regex
        :param pulumi.Input[str] user_name_template: Username template
        :param pulumi.Input[str] user_name_template_type: Username template type
        :param pulumi.Input[str] username_field: Login username field
        :param pulumi.Input[list] users: Users associated with the application

        The **users** object supports the following:

          * `id` (`pulumi.Input[str]`)
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessibility_error_redirect_url"] = accessibility_error_redirect_url
        __props__["accessibility_self_service"] = accessibility_self_service
        __props__["auto_submit_toolbar"] = auto_submit_toolbar
        __props__["button_field"] = button_field
        __props__["groups"] = groups
        __props__["hide_ios"] = hide_ios
        __props__["hide_web"] = hide_web
        __props__["label"] = label
        __props__["name"] = name
        __props__["password_field"] = password_field
        __props__["preconfigured_app"] = preconfigured_app
        __props__["sign_on_mode"] = sign_on_mode
        __props__["status"] = status
        __props__["url"] = url
        __props__["url_regex"] = url_regex
        __props__["user_name_template"] = user_name_template
        __props__["user_name_template_type"] = user_name_template_type
        __props__["username_field"] = username_field
        __props__["users"] = users
        return SwaApp(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
