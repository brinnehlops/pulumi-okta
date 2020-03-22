# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ThreeField(pulumi.CustomResource):
    accessibility_error_redirect_url: pulumi.Output[str]
    """
    Custom error page URL.
    """
    accessibility_self_service: pulumi.Output[bool]
    """
    Enable self service. By default it is `false`.
    """
    auto_submit_toolbar: pulumi.Output[bool]
    """
    Display auto submit toolbar.
    """
    button_selector: pulumi.Output[str]
    """
    Login button field CSS selector.
    """
    extra_field_selector: pulumi.Output[str]
    """
    Extra field CSS selector.
    """
    extra_field_value: pulumi.Output[str]
    """
    Value for extra form field.
    """
    groups: pulumi.Output[list]
    """
    Groups associated with the application. See `app.GroupAssignment` for a more flexible approach.
    """
    hide_ios: pulumi.Output[bool]
    """
    Do not display application icon on mobile app.
    """
    hide_web: pulumi.Output[bool]
    """
    Do not display application icon to users.
    """
    label: pulumi.Output[str]
    """
    The display name of the Application.
    """
    name: pulumi.Output[str]
    """
    Name assigned to the application by Okta.
    """
    password_selector: pulumi.Output[str]
    """
    Login password field CSS selector.
    """
    sign_on_mode: pulumi.Output[str]
    """
    Sign on mode of application.
    """
    status: pulumi.Output[str]
    """
    Status of application. By default it is `"ACTIVE"`.
    """
    url: pulumi.Output[str]
    """
    Login URL.
    """
    url_regex: pulumi.Output[str]
    """
    A regex that further restricts URL to the specified regex.
    """
    user_name_template: pulumi.Output[str]
    """
    The default username assigned to each user.
    """
    user_name_template_type: pulumi.Output[str]
    """
    The Username template type.
    """
    username_selector: pulumi.Output[str]
    """
    Login username field CSS selector.
    """
    users: pulumi.Output[list]
    """
    The users assigned to the application. See `app.User` for a more flexible approach.

      * `id` (`str`)
      * `password` (`str`)
      * `scope` (`str`)
      * `username` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, accessibility_error_redirect_url=None, accessibility_self_service=None, auto_submit_toolbar=None, button_selector=None, extra_field_selector=None, extra_field_value=None, groups=None, hide_ios=None, hide_web=None, label=None, password_selector=None, status=None, url=None, url_regex=None, username_selector=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Three Field Application.

        This resource allows you to create and configure an Three Field Application.

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_three_field.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessibility_error_redirect_url: Custom error page URL.
        :param pulumi.Input[bool] accessibility_self_service: Enable self service. By default it is `false`.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar.
        :param pulumi.Input[str] button_selector: Login button field CSS selector.
        :param pulumi.Input[str] extra_field_selector: Extra field CSS selector.
        :param pulumi.Input[str] extra_field_value: Value for extra form field.
        :param pulumi.Input[list] groups: Groups associated with the application. See `app.GroupAssignment` for a more flexible approach.
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app.
        :param pulumi.Input[bool] hide_web: Do not display application icon to users.
        :param pulumi.Input[str] label: The display name of the Application.
        :param pulumi.Input[str] password_selector: Login password field CSS selector.
        :param pulumi.Input[str] status: Status of application. By default it is `"ACTIVE"`.
        :param pulumi.Input[str] url: Login URL.
        :param pulumi.Input[str] url_regex: A regex that further restricts URL to the specified regex.
        :param pulumi.Input[str] username_selector: Login username field CSS selector.
        :param pulumi.Input[list] users: The users assigned to the application. See `app.User` for a more flexible approach.

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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['accessibility_error_redirect_url'] = accessibility_error_redirect_url
            __props__['accessibility_self_service'] = accessibility_self_service
            __props__['auto_submit_toolbar'] = auto_submit_toolbar
            if button_selector is None:
                raise TypeError("Missing required property 'button_selector'")
            __props__['button_selector'] = button_selector
            if extra_field_selector is None:
                raise TypeError("Missing required property 'extra_field_selector'")
            __props__['extra_field_selector'] = extra_field_selector
            if extra_field_value is None:
                raise TypeError("Missing required property 'extra_field_value'")
            __props__['extra_field_value'] = extra_field_value
            __props__['groups'] = groups
            __props__['hide_ios'] = hide_ios
            __props__['hide_web'] = hide_web
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            if password_selector is None:
                raise TypeError("Missing required property 'password_selector'")
            __props__['password_selector'] = password_selector
            __props__['status'] = status
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['url_regex'] = url_regex
            if username_selector is None:
                raise TypeError("Missing required property 'username_selector'")
            __props__['username_selector'] = username_selector
            __props__['users'] = users
            __props__['name'] = None
            __props__['sign_on_mode'] = None
            __props__['user_name_template'] = None
            __props__['user_name_template_type'] = None
        super(ThreeField, __self__).__init__(
            'okta:app/threeField:ThreeField',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessibility_error_redirect_url=None, accessibility_self_service=None, auto_submit_toolbar=None, button_selector=None, extra_field_selector=None, extra_field_value=None, groups=None, hide_ios=None, hide_web=None, label=None, name=None, password_selector=None, sign_on_mode=None, status=None, url=None, url_regex=None, user_name_template=None, user_name_template_type=None, username_selector=None, users=None):
        """
        Get an existing ThreeField resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessibility_error_redirect_url: Custom error page URL.
        :param pulumi.Input[bool] accessibility_self_service: Enable self service. By default it is `false`.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar.
        :param pulumi.Input[str] button_selector: Login button field CSS selector.
        :param pulumi.Input[str] extra_field_selector: Extra field CSS selector.
        :param pulumi.Input[str] extra_field_value: Value for extra form field.
        :param pulumi.Input[list] groups: Groups associated with the application. See `app.GroupAssignment` for a more flexible approach.
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app.
        :param pulumi.Input[bool] hide_web: Do not display application icon to users.
        :param pulumi.Input[str] label: The display name of the Application.
        :param pulumi.Input[str] name: Name assigned to the application by Okta.
        :param pulumi.Input[str] password_selector: Login password field CSS selector.
        :param pulumi.Input[str] sign_on_mode: Sign on mode of application.
        :param pulumi.Input[str] status: Status of application. By default it is `"ACTIVE"`.
        :param pulumi.Input[str] url: Login URL.
        :param pulumi.Input[str] url_regex: A regex that further restricts URL to the specified regex.
        :param pulumi.Input[str] user_name_template: The default username assigned to each user.
        :param pulumi.Input[str] user_name_template_type: The Username template type.
        :param pulumi.Input[str] username_selector: Login username field CSS selector.
        :param pulumi.Input[list] users: The users assigned to the application. See `app.User` for a more flexible approach.

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
        __props__["button_selector"] = button_selector
        __props__["extra_field_selector"] = extra_field_selector
        __props__["extra_field_value"] = extra_field_value
        __props__["groups"] = groups
        __props__["hide_ios"] = hide_ios
        __props__["hide_web"] = hide_web
        __props__["label"] = label
        __props__["name"] = name
        __props__["password_selector"] = password_selector
        __props__["sign_on_mode"] = sign_on_mode
        __props__["status"] = status
        __props__["url"] = url
        __props__["url_regex"] = url_regex
        __props__["user_name_template"] = user_name_template
        __props__["user_name_template_type"] = user_name_template_type
        __props__["username_selector"] = username_selector
        __props__["users"] = users
        return ThreeField(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

